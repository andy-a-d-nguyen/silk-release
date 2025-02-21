package vtep_test

import (
	"errors"
	"net"

	"code.cloudfoundry.org/lager/v3/lagertest"
	mcn "code.cloudfoundry.org/lib/multiple-cidr-network"
	"code.cloudfoundry.org/silk/daemon/vtep"
	"code.cloudfoundry.org/silk/daemon/vtep/fakes"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
	"github.com/vishvananda/netlink"
)

var _ = Describe("Factory", func() {
	var (
		fakeNetlinkAdapter *fakes.NetlinkAdapter
		factory            *vtep.Factory
		vtepConfig         *vtep.Config
		fakeLogger         *lagertest.TestLogger
		overlayMAC         net.HardwareAddr
	)

	BeforeEach(func() {
		fakeNetlinkAdapter = &fakes.NetlinkAdapter{}
		fakeLogger = lagertest.NewTestLogger("test")
		factory = &vtep.Factory{
			NetlinkAdapter: fakeNetlinkAdapter,
			Logger:         fakeLogger,
		}

		overlayMAC = net.HardwareAddr{0xee, 0xee, 0x0a, 0xff, 0x20, 0x00}

		underlayInterface := net.Interface{
			Index:        4,
			MTU:          1450,
			Name:         "eth4",
			HardwareAddr: net.HardwareAddr{0xbb, 0xbb, 0x00, 0x00, 0x12, 0x34},
			Flags:        net.FlagUp | net.FlagMulticast,
		}

		overlayNetworks, err := mcn.NewMultipleCIDRNetwork([]string{"10.2.0.0/24", "10.240.0.0/16", "10.255.0.0/16"})
		Expect(err).NotTo(HaveOccurred())

		vtepConfig = &vtep.Config{
			VTEPName:            "some-device",
			UnderlayInterface:   underlayInterface,
			UnderlayIP:          net.IP{172, 255, 0, 0},
			LeaseIP:             net.IP{10, 255, 32, 0},
			OverlayNetworks:     overlayNetworks,
			OverlayHardwareAddr: overlayMAC,
			VNI:                 99,
			VTEPPort:            4913,
		}
	})

	Describe("CreateVTEP", func() {
		var vtepLinkIndex int
		BeforeEach(func() {
			vtepLinkIndex = 2
			fakeNetlinkAdapter.LinkByNameReturns(&netlink.Vxlan{LinkAttrs: netlink.LinkAttrs{Index: vtepLinkIndex}}, nil)
		})
		It("creates the link, with the HW address", func() {
			err := factory.CreateVTEP(vtepConfig)
			Expect(err).NotTo(HaveOccurred())

			expectedLink := &netlink.Vxlan{
				LinkAttrs: netlink.LinkAttrs{
					Name:         "some-device",
					HardwareAddr: overlayMAC,
				},
				VxlanId:      99,
				SrcAddr:      net.IP{172, 255, 0, 0},
				GBP:          true,
				Port:         4913,
				VtepDevIndex: 4,
			}

			Expect(fakeNetlinkAdapter.LinkAddCallCount()).To(Equal(1))
			Expect(fakeNetlinkAdapter.LinkAddArgsForCall(0)).To(Equal(expectedLink))

			Expect(fakeNetlinkAdapter.LinkSetUpCallCount()).To(Equal(1))
			Expect(fakeNetlinkAdapter.LinkSetUpArgsForCall(0)).To(Equal(expectedLink))

			Expect(fakeNetlinkAdapter.LinkSetHardwareAddrCallCount()).To(Equal(0))

			// for network: 10.2.0.0/24
			Expect(fakeNetlinkAdapter.AddrAddScopeLinkCallCount()).To(Equal(vtepConfig.OverlayNetworks.Length()))
			link, addr := fakeNetlinkAdapter.AddrAddScopeLinkArgsForCall(0)
			Expect(link).To(Equal(expectedLink))
			Expect(addr).To(Equal(&netlink.Addr{
				IPNet: &net.IPNet{
					IP:   net.IP{10, 2, 0, 0},
					Mask: net.IPMask{0xff, 0xff, 0xff, 0x00},
				},
			}))

			// for network: 10.240.0.0/16
			link, addr = fakeNetlinkAdapter.AddrAddScopeLinkArgsForCall(1)
			Expect(link).To(Equal(expectedLink))
			Expect(addr).To(Equal(&netlink.Addr{
				IPNet: &net.IPNet{
					IP:   net.IP{10, 240, 0, 0},
					Mask: net.IPMask{0xff, 0xff, 0x00, 0x00},
				},
			}))

			// for network: 10.255.0.0/16
			link, addr = fakeNetlinkAdapter.AddrAddScopeLinkArgsForCall(2)
			Expect(link).To(Equal(expectedLink))
			Expect(addr).To(Equal(&netlink.Addr{
				IPNet: &net.IPNet{
					// the lease for this daemon is in this subnet, so we use the lease IP instead of the first IP
					IP:   net.IP{10, 255, 32, 0},
					Mask: net.IPMask{0xff, 0xff, 0x00, 0x00},
				},
			}))
		})

		Context("when adding the link fails", func() {
			BeforeEach(func() {
				fakeNetlinkAdapter.LinkAddReturns(errors.New("potato"))
			})
			It("wraps and returns the error", func() {
				err := factory.CreateVTEP(vtepConfig)
				Expect(err).To(MatchError("create link some-device: potato"))
			})
		})

		Context("when setting the link up fails", func() {
			BeforeEach(func() {
				fakeNetlinkAdapter.LinkSetUpReturns(errors.New("potato"))
			})
			It("wraps and returns the error", func() {
				err := factory.CreateVTEP(vtepConfig)
				Expect(err).To(MatchError("up link: potato"))
			})
		})

		Context("when adding the overlay address", func() {
			BeforeEach(func() {
				fakeNetlinkAdapter.AddrAddScopeLinkReturns(errors.New("potato"))
			})
			It("wraps and returns the error", func() {
				err := factory.CreateVTEP(vtepConfig)
				Expect(err).To(MatchError("add address: potato"))
			})
		})

		Context("when the leaseIP is not in any of the overlay networks", func() {
			It("returns an error", func() {
				vtepConfig.LeaseIP = net.IP{10, 30, 0, 0}
				err := factory.CreateVTEP(vtepConfig)
				Expect(err).To(MatchError(ContainSubstring("lease IP '10.30.0.0' is not in any of the overlay networks")))
			})
		})
	})

	Describe("GetVTEPState", func() {
		BeforeEach(func() {
			fakeNetlinkAdapter.LinkByNameReturns(&netlink.Vxlan{
				LinkAttrs: netlink.LinkAttrs{
					Name:         "some-device",
					MTU:          1400,
					HardwareAddr: net.HardwareAddr{0xee, 0xee, 0x0a, 0xff, 0x42, 0x00},
				},
			}, nil)
			fakeNetlinkAdapter.AddrListReturns([]netlink.Addr{{
				IPNet: &net.IPNet{
					IP:   net.IP{10, 255, 32, 0},
					Mask: net.IPMask{0xff, 0xff, 0xff, 0xff},
				},
			}}, nil)
		})

		It("returns the overlay address, hardware addr, and MTU", func() {
			hwAddr, ip, mtu, err := factory.GetVTEPState(vtepConfig.VTEPName)
			Expect(err).NotTo(HaveOccurred())
			Expect(ip).To(Equal(net.IP{10, 255, 32, 0}))
			Expect(hwAddr).To(Equal(net.HardwareAddr{0xee, 0xee, 0x0a, 0xff, 0x42, 0x00}))
			Expect(mtu).To(Equal(1400))

			Expect(fakeNetlinkAdapter.LinkByNameCallCount()).To(Equal(1))
			Expect(fakeNetlinkAdapter.LinkByNameArgsForCall(0)).To(Equal(vtepConfig.VTEPName))

			Expect(fakeNetlinkAdapter.AddrListCallCount()).To(Equal(1))
			link, family := (fakeNetlinkAdapter.AddrListArgsForCall(0))
			Expect(link).To(Equal(&netlink.Vxlan{
				LinkAttrs: netlink.LinkAttrs{
					Name:         "some-device",
					HardwareAddr: net.HardwareAddr{0xee, 0xee, 0x0a, 0xff, 0x42, 0x00},
					MTU:          1400,
				},
			}))
			Expect(family).To(Equal(netlink.FAMILY_V4))
		})

		Context("when finding the link errors", func() {
			BeforeEach(func() {
				fakeNetlinkAdapter.LinkByNameReturns(nil, errors.New("potato"))
			})
			It("returns an error", func() {
				_, _, _, err := factory.GetVTEPState(vtepConfig.VTEPName)
				Expect(err).To(MatchError("find link: potato"))
			})
		})

		Context("when listing the addresses fails", func() {
			BeforeEach(func() {
				fakeNetlinkAdapter.AddrListReturns(nil, errors.New("potato"))
			})
			It("returns an error", func() {
				_, _, _, err := factory.GetVTEPState(vtepConfig.VTEPName)
				Expect(err).To(MatchError("list addresses: potato"))
			})
		})

		Context("when there are no addresses", func() {
			BeforeEach(func() {
				fakeNetlinkAdapter.AddrListReturns(nil, nil)
			})
			It("returns an error", func() {
				_, _, _, err := factory.GetVTEPState(vtepConfig.VTEPName)
				Expect(err).To(MatchError("no addresses"))
			})
		})
	})

	Describe("DeleteVTEP", func() {
		BeforeEach(func() {
			fakeNetlinkAdapter.LinkByNameReturns(&netlink.Vxlan{
				LinkAttrs: netlink.LinkAttrs{
					Name:         "some-device",
					HardwareAddr: net.HardwareAddr{0xee, 0xee, 0x0a, 0xff, 0x42, 0x00},
				},
			}, nil)
		})

		It("deletes the vtep", func() {
			err := factory.DeleteVTEP(vtepConfig.VTEPName)
			Expect(err).NotTo(HaveOccurred())

			Expect(fakeNetlinkAdapter.LinkByNameCallCount()).To(Equal(1))
			Expect(fakeNetlinkAdapter.LinkByNameArgsForCall(0)).To(Equal(vtepConfig.VTEPName))

			Expect(fakeNetlinkAdapter.LinkDelCallCount()).To(Equal(1))
			Expect(fakeNetlinkAdapter.LinkDelArgsForCall(0)).To(Equal(&netlink.Vxlan{
				LinkAttrs: netlink.LinkAttrs{
					Name:         "some-device",
					HardwareAddr: net.HardwareAddr{0xee, 0xee, 0x0a, 0xff, 0x42, 0x00},
				},
			}))
		})

		Context("when the link cannot be found", func() {
			BeforeEach(func() {
				fakeNetlinkAdapter.LinkByNameReturns(nil, errors.New("banana"))
			})
			It("returns an error", func() {
				err := factory.DeleteVTEP(vtepConfig.VTEPName)
				Expect(err).To(MatchError("find link some-device: banana"))
			})
		})

		Context("when the link cannot be deleted", func() {
			BeforeEach(func() {
				fakeNetlinkAdapter.LinkDelReturns(errors.New("banana"))
			})
			It("returns an error", func() {
				err := factory.DeleteVTEP(vtepConfig.VTEPName)
				Expect(err).To(MatchError("delete link some-device: banana"))
			})
		})
	})
})

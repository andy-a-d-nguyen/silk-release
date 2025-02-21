package vtep

import (
	"fmt"
	"net"

	mcn "code.cloudfoundry.org/lib/multiple-cidr-network"
	clientConfig "code.cloudfoundry.org/silk/client/config"
	"code.cloudfoundry.org/silk/controller"
)

//go:generate counterfeiter -o fakes/netAdapter.go --fake-name NetAdapter . netAdapter
type netAdapter interface {
	Interfaces() ([]net.Interface, error)
	InterfaceAddrs(net.Interface) ([]net.Addr, error)
	InterfaceByName(name string) (*net.Interface, error)
}

type ConfigCreator struct {
	NetAdapter netAdapter
}

type Config struct {
	VTEPName            string
	UnderlayInterface   net.Interface
	UnderlayIP          net.IP
	LeaseIP             net.IP
	OverlayHardwareAddr net.HardwareAddr
	VNI                 int
	VTEPPort            int
	OverlayNetworks     mcn.MultipleCIDRNetwork
}

func (c *ConfigCreator) Create(clientConf clientConfig.Config, lease controller.Lease) (*Config, error) {
	if clientConf.VTEPName == "" {
		return nil, fmt.Errorf("empty vtep name")
	}

	if clientConf.VTEPPort < 1 {
		return nil, fmt.Errorf("vtep port must be greater than 0")
	}

	underlayIP := net.ParseIP(clientConf.UnderlayIP)
	if underlayIP == nil {
		return nil, fmt.Errorf("parse underlay ip: %s", clientConf.UnderlayIP)
	}

	var underlayInterface net.Interface
	var err error

	if clientConf.VxlanInterfaceName != "" {
		underlayInterfacePointer, err := c.NetAdapter.InterfaceByName(clientConf.VxlanInterfaceName)
		if err != nil {
			return nil, fmt.Errorf("find device from name %s: %s", clientConf.VxlanInterfaceName, err)
		}
		underlayInterface = *underlayInterfacePointer
	} else {
		underlayInterface, err = c.locateInterface(underlayIP)
		if err != nil {
			return nil, fmt.Errorf("find device from ip %s: %s", underlayIP, err)
		}
	}

	leaseIP, _, err := net.ParseCIDR(lease.OverlaySubnet)
	if err != nil {
		return nil, fmt.Errorf("determine vtep overlay ip: %s", err)
	}

	overlayHardwareAddr, err := net.ParseMAC(lease.OverlayHardwareAddr)
	if err != nil {
		return nil, fmt.Errorf("parsing hardware address: %s", err)
	}

	if len(clientConf.OverlayNetworks) == 0 {
		return nil, fmt.Errorf("no overlay networks specified")
	}

	overlayNetworks, err := mcn.NewMultipleCIDRNetwork(clientConf.OverlayNetworks)
	if err != nil {
		return nil, fmt.Errorf("creating multiple CIDR Network: %s", err)
	}

	if overlayNetworks.SmallestMask >= clientConf.SubnetPrefixLength {
		return nil, fmt.Errorf("overlay prefix %d must be smaller than subnet prefix %d",
			overlayNetworks.SmallestMask, clientConf.SubnetPrefixLength)
	}

	return &Config{
		VTEPName:            clientConf.VTEPName,
		UnderlayInterface:   underlayInterface,
		UnderlayIP:          underlayIP,
		LeaseIP:             leaseIP,
		OverlayNetworks:     overlayNetworks,
		OverlayHardwareAddr: overlayHardwareAddr,
		VNI:                 clientConf.VNI,
		VTEPPort:            clientConf.VTEPPort,
	}, nil
}

func (c *ConfigCreator) locateInterface(toFind net.IP) (net.Interface, error) {
	ifaces, err := c.NetAdapter.Interfaces()
	if err != nil {
		return net.Interface{}, fmt.Errorf("find interfaces: %s", err)
	}
	for _, iface := range ifaces {
		addrs, err := c.NetAdapter.InterfaceAddrs(iface)
		if err != nil {
			return net.Interface{}, fmt.Errorf("get addresses: %s", err)
		}

		for _, addr := range addrs {
			ip, _, err := net.ParseCIDR(addr.String())
			if err != nil {
				return net.Interface{}, fmt.Errorf("parse address: %s", err)
			}
			if ip.String() == toFind.String() {
				return iface, nil
			}
		}
	}

	return net.Interface{}, fmt.Errorf("no interface with address %s", toFind.String())
}

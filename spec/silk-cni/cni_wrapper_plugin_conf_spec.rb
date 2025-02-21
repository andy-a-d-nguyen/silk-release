require 'rspec'
require 'bosh/template/test'
require 'yaml'
require 'json'


module Bosh::Template::Test
  describe 'silk-cni job' do
    let(:release_path) {File.join(File.dirname(__FILE__), '../..')}
    let(:release) {ReleaseDir.new(release_path)}
    let(:links) {[
      Link.new(
        name: 'vpa',
        properties: {
          'force_policy_poll_cycle_port' => 5555
        }
      )
    ]}
    let(:merged_manifest_properties) do
      {
        'mtu' => mtu,
        'silk_daemon' => {
          'listen_port' => 8080
        },
        'iptables_logging' => true,
        'no_masquerade_cidr_range' => '222.22.0.0/16',
        'dns_servers' => ['8.8.8.8'],
        'rate' => 100,
        'burst' => 200,
        'iptables_denied_logs_per_sec' => 2,
        'iptables_accepted_udp_logs_per_sec' => 3,
        'host_tcp_services' => ['169.254.0.2:9001', '169.254.0.2:9002'],
        'host_udp_services' => ['169.254.0.2:9003', '169.254.0.2:9004'],
        'deny_networks' => {
          'always' => ['1.1.1.1/32'],
          'running' => ['2.2.2.2/32'],
          'staging' => ['3.3.3.3/32'],
        },
        'outbound_connections' => {
          'limit' => true,
        }
      }
    end
    let(:job) {release.job('silk-cni')}
    let(:mtu) {0}
    let(:disable) {false}
    let(:networks) {{'fake-network' => {'fake-network-settings' => {}, 'ip' => '192.74.65.4'}}}
    let(:spec) {InstanceSpec.new(networks: networks, ip: '111.11.11.1')}


    describe 'cni-wrapper-plugin.conflist' do
      let(:template) {job.template('config/cni/cni-wrapper-plugin.conflist')}

      it 'creates a config/cni/cni-wrapper-plugin.conflist from properties' do
        clientConfig = JSON.parse(template.render(merged_manifest_properties, spec: spec, consumes: links))
        expect(clientConfig).to eq({
          'name' => 'cni-wrapper',
          'cniVersion' => '1.0.0',
          'disableCheck' => true,
          'plugins' => [{
            'type' => 'cni-wrapper-plugin',
            'datastore' => '/var/vcap/data/container-metadata/store.json',
            'datastore_file_owner' => 'vcap',
            'datastore_file_group' => 'vcap',
            'iptables_lock_file' => '/var/vcap/data/garden-cni/iptables.lock',
            'instance_address' => '111.11.11.1',
            'no_masquerade_cidr_range' => '222.22.0.0/16',
            'temporary_underlay_interface_names' => [],
            'underlay_ips' => ['192.74.65.4'],
            'iptables_asg_logging' => true,
            'iptables_c2c_logging' => true,
            'iptables_denied_logs_per_sec' => 2,
            'iptables_accepted_udp_logs_per_sec' => 3,
            'ingress_tag' => 'ffff0000',
            'vtep_name' => 'silk-vtep',
            'dns_servers' => ['8.8.8.8'],
            'policy_agent_force_poll_address' => '127.0.0.1:5555',
            'host_tcp_services' => ['169.254.0.2:9001', '169.254.0.2:9002'],
            'host_udp_services' => ['169.254.0.2:9003', '169.254.0.2:9004'],
            'deny_networks' => {
              'always' => ['1.1.1.1/32'],
              'running' => ['2.2.2.2/32'],
              'staging' => ['3.3.3.3/32'],
            },
            'delegate' => {
              'cniVersion' => '1.0.0',
              'name' => 'silk',
              'type' => 'silk-cni',
              'daemonPort' => 8080,
              'dataDir' => '/var/vcap/data/host-local',
              'datastore' => '/var/vcap/data/silk/store.json',
              'mtu' => 0
            },
            'outbound_connections' => {
              'limit' => true,
              'logging' => true,
              'burst' => 1000,
              'rate_per_sec' => 100,
              'dry_run' => false,
            }
          }, {
            'name' => 'bandwidth-limit',
            'type' => 'bandwidth',
            'ingressRate' => 100 * 1024,
            'ingressBurst' => 200 * 1024,
            'egressRate' => 100 * 1024,
            'egressBurst' => 200 * 1024
          }]
        })
      end

      context 'when ips have leading 0s' do
        it 'no_masquerade_cidr_range fails with a nice message' do
          merged_manifest_properties['no_masquerade_cidr_range'] = '222.022.0.2/16'
          expect {
            template.render(merged_manifest_properties, spec: spec, consumes: links)
          }.to raise_error (/Invalid no_masquerade_cidr_range/)
        end

        it 'dns_servers fails with a nice message' do
          merged_manifest_properties['dns_servers'] = ['1.2.3.4', '8.08.08.08']
          expect {
            template.render(merged_manifest_properties, spec: spec, consumes: links)
          }.to raise_error (/Invalid dns_servers '8.08.08.08':/)
        end

        it 'host_tcp_services fails with a nice message' do
          merged_manifest_properties['host_tcp_services'] = ['1.2.3.4:33333', '8.08.08.08:8888']
          expect {
            template.render(merged_manifest_properties, spec: spec, consumes: links)
          }.to raise_error (/Invalid host_tcp_services '8.08.08.08':/)
        end

        it 'host_udp_services fails with a nice message' do
          merged_manifest_properties['host_udp_services'] = ['1.2.3.4:33333', '8.08.08.08:8888']
          expect {
            template.render(merged_manifest_properties, spec: spec, consumes: links)
          }.to raise_error (/Invalid host_udp_services '8.08.08.08':/)
        end

        it 'deny_networks.running fails with a nice message' do
          merged_manifest_properties['deny_networks']['running'] = ['1.2.3.4/12', '8.08.08.08/13']
          expect {
            template.render(merged_manifest_properties, spec: spec, consumes: links)
          }.to raise_error (/Invalid deny_networks.running entry 8.08.08.08\/13/)
        end

        it 'deny_networks.staging fails with a nice message' do
          merged_manifest_properties['deny_networks']['staging'] = ['1.2.3.4/12', '8.08.08.08/13']
          expect {
            template.render(merged_manifest_properties, spec: spec, consumes: links)
          }.to raise_error (/Invalid deny_networks.staging entry 8.08.08.08\/13/)
        end

        it 'deny_networks.always fails with a nice message' do
          merged_manifest_properties['deny_networks']['always'] = ['1.2.3.4/12', '8.08.08.08/13']
          expect {
            template.render(merged_manifest_properties, spec: spec, consumes: links)
          }.to raise_error (/Invalid deny_networks.always entry 8.08.08.08\/13/)
        end
      end

      context 'when no_masquerade_cidr_range is not provided' do
        let(:merged_manifest_properties) {}
        it 'does not set the no_masquerade_cidr_range' do
          clientConfig = JSON.parse(template.render(merged_manifest_properties, spec: spec, consumes: links))
          expect(clientConfig['plugins'][0]['no_masquerade_cidr_range']).to eq('')
        end
      end

      context 'when mtu is greater than 0' do
        let(:mtu) {100}
        it 'subtracts VXLAN_OVERHEAD from the mtu value' do
          clientConfig = JSON.parse(template.render(merged_manifest_properties, spec: spec, consumes: links))
          expect(clientConfig['plugins'][0]['delegate']['mtu']).to eq(50)
        end
      end

      context 'when deny_networks are provided' do
        context 'when a destination is IPv6' do
          it 'raises a descriptive error' do
            contents = merged_manifest_properties.merge(
              'deny_networks' => {
                'always' => ['2001:db8:0:1:1:1:1:1', '1.1.0.0/16']
              }
            )

            expect {
              template.render(contents, spec: spec, consumes: links)
            }.to raise_error /Invalid deny_networks.always entry 2001:db8:0:1:1:1:1:1 not an IPv4 address/
          end
        end

        context 'when a destination is invalid' do
          it 'raises a descriptive error' do
            contents = merged_manifest_properties.merge(
              'deny_networks' => {
                'running' => ['invalid-network']
              }
            )

            expect {
              template.render(contents, spec: spec, consumes: links)
            }.to raise_error /Invalid deny_networks.running entry invalid-network invalid address/
          end
        end
      end

      context 'when deny_networks are not provided' do
        it 'does not raise an error' do
          contents = merged_manifest_properties.clone.delete('deny_networks')

          expect {
            template.render(contents, spec: spec, consumes: links)
          }.not_to raise_error
        end
      end
    end
  end
end

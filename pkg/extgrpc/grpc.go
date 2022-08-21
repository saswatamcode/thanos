package extgrpc

import (
	"github.com/thanos-io/thanos/pkg/httpconfig"
)

// EndpointsConfig configures a cluster of gRPC endpoints from static addresses and
// file service discovery. Similar to httpconfig.EndpointConfig but for gRPC.
type EndpointsConfig struct {
	// List of addresses with DNS prefixes.
	Addresses []string `yaml:"addresses"`
	// List of file configurations (our FileSD supports different DNS lookups).
	FileSDConfigs []httpconfig.FileSDConfig `yaml:"file_sd_configs"`
}

package query

import (
	"github.com/pkg/errors"
	"github.com/prometheus/prometheus/discovery/file"
	"github.com/thanos-io/thanos/pkg/extgrpc"
	"gopkg.in/yaml.v2"
)

type EndpointConfig struct {
	extgrpc.Config `yaml:",inline"`
	Endpoints      []string
	EndpointsSD    []file.SDConfig
}

// LoadConfig returns list of per-endpoint TLS config.
func LoadConfig(confYAML []byte, endpointAddrs []string, globalFileSDConfig *file.SDConfig) ([]EndpointConfig, error) {
	var endpointConfig []EndpointConfig

	if len(confYAML) > 0 {
		if err := yaml.UnmarshalStrict(confYAML, &endpointConfig); err != nil {
			return nil, err
		}

	}

	// Adding --store, rule, metadata, target, exemplar and --store.sd-files, if provided.
	// Global TLS config applies until deprecated.
	if len(endpointAddrs) > 0 || globalFileSDConfig != nil {
		cfg := EndpointConfig{}
		cfg.Endpoints = endpointAddrs
		if globalFileSDConfig != nil {
			cfg.EndpointsSD = []file.SDConfig{
				{
					Files:           globalFileSDConfig.Files,
					RefreshInterval: globalFileSDConfig.RefreshInterval,
				},
			}
		}
		endpointConfig = append(endpointConfig, cfg)
	}

	// Checking for duplicates.
	// NOTE: This does not check dynamic endpoints of course.
	allEndpoints := make(map[string]struct{})
	for _, config := range endpointConfig {
		for _, addr := range config.EndpointsConfig.Addresses {
			if _, exists := allEndpoints[addr]; exists {
				return nil, errors.Errorf("%s endpoint provided more than once", addr)
			}
			allEndpoints[addr] = struct{}{}
		}
	}
	return endpointConfig, nil
}

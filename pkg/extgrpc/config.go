package extgrpc

import "github.com/thanos-io/thanos/pkg/httpconfig"

// Config is a structure that allows pointing to various gRPC endpoint, e.g Querier connecting to StoreAPI.
type Config struct {
	EndpointsConfig EndpointsConfig `yaml:",inline"`
}

func DefaultConfig() Config {
	return Config{
		EndpointsConfig: EndpointsConfig{
			Addresses:     []string{},
			FileSDConfigs: []httpconfig.FileSDConfig{},
		},
	}
}

// UnmarshalYAML implements the yaml.Unmarshaler interface.
func (c *Config) UnmarshalYAML(unmarshal func(interface{}) error) error {
	*c = DefaultConfig()
	type plain Config
	return unmarshal((*plain)(c))
}

package di

import (
	"github.com/ITegs/trippify/apiserver"
)

// APIServerConfig is the configuration for the API server.
type APIServerConfig struct {
	ListenPort int64 `yaml:"listen_port" env:"APISERVER_LISTEN_PORT"`
}

// APIServerDefaultConfig is the default configuration for the API server.
var APIServerDefaultConfig = APIServerConfig{
	ListenPort: 80,
}

// APIServerBuilder is a builder (DI container) for the API server.
type APIServerBuilder struct {
	*baseContainer
}

// LoadAPIServerConfig loads the videoservice config from the given path.
func LoadAPIServerConfig(path string) (*APIServerConfig, error) {
	c := APIServerDefaultConfig
	err := loadConfig(path, &c)
	return &c, err
}

// NewAPIServerBuilder creates a new builder for the API server.
func NewAPIServerBuilder(c *APIServerConfig, baseConfig *BaseConfig) *APIServerBuilder {
	ct := APIServerBuilder{
		baseContainer: NewBaseContainer(baseConfig),
	}

	// Make the configuration available to the container
	ct.SingletonLazy(func() *APIServerConfig { return c })

	// Build the API server
	ct.TransientLazy(apiserver.NewAPIServer)

	return &ct
}

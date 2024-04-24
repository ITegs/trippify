package di

import (
	"context"
	"net/url"
	"strings"

	"github.com/golobby/container/v3"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type BaseConfig struct {
	Endpoints endpointsConfig `yaml:"endpoints"`
	BaseURL   string          `yaml:"base_url" env:"BASE_URL"`
}

type endpointsConfig struct {
	MongoDB url.URL `yaml:"mongodb" env:"MONGODB_URL"`
}

var baseDefaultConfig = BaseConfig{
	Endpoints: endpointsConfig{
		MongoDB: url.URL{
			Scheme: "mongodb",
			Host:   "localhost:27017",
			Path:   "trippify",
		},
	},
}

// LoadBaseConfig loads the base config from the given path.
func LoadBaseConfig(path string) (*BaseConfig, error) {
	c := baseDefaultConfig
	err := loadConfig(path, &c)
	return &c, err
}

// baseContainer is a common base container holding the basic services required by all components.
type baseContainer struct {
	container.Container
}

func NewBaseContainer(config *BaseConfig) *baseContainer {
	ct := baseContainer{Container: container.New()}

	// Make the provider config available to the baseContainer
	ct.SingletonLazy(func() *BaseConfig { return config })

	// Register external services
	ct.TransientLazy(ct.buildMongoDB)

	return &ct
}

// buildMongoDB builds a mongodb client
func (ct *baseContainer) buildMongoDB(config *BaseConfig) (*mongo.Database, error) {
	dbName := strings.TrimLeft(config.Endpoints.MongoDB.Path, "/")
	db, err := mongo.Connect(context.TODO(), options.Client().ApplyURI(config.Endpoints.MongoDB.String()))
	if err != nil {
		return nil, err
	}
	return db.Database(dbName), nil
}

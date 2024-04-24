package di

import (
	"errors"
	"os"

	"github.com/caarlos0/env/v10"
	"gopkg.in/yaml.v3"
)

// loadConfig loads the config from given yaml file and from the envinronment.
//
// Values are first loaded from the yaml file, then from the environment. This
// means that values from the environment will overwrite values from the yaml.
//
// ENV variables are mapped to config values by using the `env` struct tag.
// This function uses [github.com/caarlos0/env] for that purpose. The `env`
// tag should not use the `default` option, as this would cause the values
// loaded from the yaml file to be overwritten.
func loadConfig(path string, config any) error {
	// First, load the config from the yaml file, skipping if the file does not exist.
	cf, err := os.ReadFile(string(path))
	if err == nil {
		err = yaml.Unmarshal(cf, config)
		if err != nil {
			return err
		}
	} else if !errors.Is(err, os.ErrNotExist) {
		return err
	}

	// Then, load the config from the environment.
	return env.Parse(config)
}

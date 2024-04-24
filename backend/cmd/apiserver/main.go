package main

import (
	"fmt"
	"net/http"
	"os"

	"github.com/ITegs/trippify/apiserver"
	"github.com/ITegs/trippify/di"
	"github.com/gorilla/handlers"
)

// defaultConfigPath is the default path to the config file of the apiserver
const defaultConfigPath = "config/apiserver.yaml"

func main() {
	fmt.Printf("Starting API server...\n")
	a := app{}

	if err := a.serve(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

// app is the main application struct
type app struct {
	builder *di.APIServerBuilder
}

// load loads the config and sets up the DI container
func (a *app) load() error {
	// First, load the cfg from given file path
	bcfg, err := di.LoadBaseConfig("") // Empty string -> use default config
	if err != nil {
		return err
	}
	cfg, err := di.LoadAPIServerConfig(defaultConfigPath)
	if err != nil {
		return err
	}

	// Retrieve our DI container, passing it the config
	a.builder = di.NewAPIServerBuilder(cfg, bcfg)

	return nil
}

// serve runs the HTTP API server
func (a *app) serve() error {
	if err := a.load(); err != nil {
		return err
	}

	var config *di.APIServerConfig
	if err := a.builder.Resolve(&config); err != nil {
		return err
	}

	var apiserver apiserver.APIServer
	if err := a.builder.Resolve(&apiserver); err != nil {
		return err
	}

	cors := handlers.CORS(
		handlers.AllowedOrigins([]string{"*"}),
		handlers.AllowedMethods([]string{http.MethodOptions, http.MethodGet, http.MethodPost, http.MethodPut, http.MethodDelete}),
		handlers.AllowedHeaders([]string{"Content-Type"}),
	)

	server := http.Server{
		Addr:    fmt.Sprintf(":%d", config.ListenPort),
		Handler: cors(apiserver),
	}

	fmt.Printf("API server listening on port %d\n", config.ListenPort)
	server.ListenAndServe()

	return nil
}

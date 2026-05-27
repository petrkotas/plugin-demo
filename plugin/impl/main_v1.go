//go:build v1

package main

import (
	"github.com/hashicorp/go-hclog"
	"github.com/hashicorp/go-plugin"

	"github.com/petrkotas/plugin-demo/plugin/shared"
)

func main() {
	logger := hclog.New(&hclog.LoggerOptions{
		Name:  "V1 Plugin",
		Level: hclog.Debug,
	})

	srelib := &SRELibraryPlugin{logger: logger}

	plugin.Serve(&plugin.ServeConfig{
		Logger:          logger,
		HandshakeConfig: shared.HandshakeConfig,
		Plugins: map[string]plugin.Plugin{
			"sre_library": &shared.SRELibraryPluginV1{Impl: srelib},
		},
		VersionedPlugins: map[int]plugin.PluginSet{
			1: {
				"sre_library": &shared.SRELibraryPluginV1{Impl: srelib},
			},
		},
	})
}

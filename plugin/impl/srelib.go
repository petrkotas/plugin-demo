package main

import (
	"github.com/hashicorp/go-hclog"
	"github.com/hashicorp/go-plugin"

	"github.com/petrkotas/plugin-demo/plugin/shared"
)

// =================================================
// Plugin implementation
// =================================================

type SRELibraryPlugin struct {
	logger hclog.Logger
	state  int
}

func (p *SRELibraryPlugin) GetClusterInfo() (string, error) {
	p.logger.Info("GetClusterInfo called")
	return "Cluster Info: This is a sample SRE Library plugin.", nil
}

func (p *SRELibraryPlugin) GetManagementCluster() (string, error) {
	p.logger.Info("GetManagementCluster called", "state", p.state)
	p.state++

	return "Management Cluster: This is a sample management cluster info.", nil
}

// =================================================
// Main function to serve the plugin
// =================================================

func main() {
	logger := hclog.New(&hclog.LoggerOptions{
		Level:      hclog.Debug,
		JSONFormat: true,
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
			2: {
				"sre_library": &shared.SRELibraryPluginV2{Impl: srelib},
			},
		},
	})
}

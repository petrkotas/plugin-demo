//go:build v2

package main

import (
	"os"
	"os/exec"

	"github.com/hashicorp/go-hclog"
	"github.com/hashicorp/go-plugin"

	"github.com/petrkotas/plugin-demo/plugin/shared"
)

func main() {
	logger := hclog.New(&hclog.LoggerOptions{
		Name:   "V2 Client",
		Output: os.Stdout,
		Level:  hclog.Debug,
	})

	client := plugin.NewClient(&plugin.ClientConfig{
		HandshakeConfig: shared.HandshakeConfig,
		VersionedPlugins: map[int]plugin.PluginSet{
			1: {
				"sre_library": &shared.SRELibraryPluginV1{},
			},
			2: {
				"sre_library": &shared.SRELibraryPluginV2{},
			},
		},
		Cmd:    exec.Command("./srelib"),
		Logger: logger,
	})
	defer client.Kill()

	rpcClient, err := client.Client()
	if err != nil {
		logger.Error("Error creating RPC client", "error", err)
		return
	}

	raw, err := rpcClient.Dispense("sre_library")
	if err != nil {
		logger.Error("Error dispensing plugin", "error", err)
		return
	}

	sreLibrary := raw.(shared.SRELibraryV1)
	info, err := sreLibrary.GetClusterInfo()
	if err != nil {
		logger.Error("Error getting cluster info", "error", err)
		return
	}
	logger.Info("Cluster Info V1", "info", info)

	if client.NegotiatedVersion() == 2 {
		logger.Info("Using SRE Library V2")

		sreLibraryV2 := raw.(shared.SRELibraryV2)

		mgmtInfo, err := sreLibraryV2.GetManagementCluster()
		if err != nil {
			logger.Error("Error getting management cluster info", "error", err)
			return
		}
		logger.Info("Management Cluster Info V2", "info", mgmtInfo)
	}
}

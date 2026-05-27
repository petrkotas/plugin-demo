package main

import (
	"github.com/hashicorp/go-hclog"
)

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

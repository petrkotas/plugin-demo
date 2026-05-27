//go:build v2

package shared

import (
	"net/rpc"

	"github.com/hashicorp/go-plugin"
)

// The interface that gets communicated as an API
type SRELibraryV2 interface {
	GetClusterInfo() (string, error)
	GetManagementCluster() (string, error)
}

// ================================================
// Client implementation of SRELibrary
// =================================================

type SRELibraryRPCClientV2 struct {
	Client *rpc.Client
}

func (m *SRELibraryRPCClientV2) GetClusterInfo() (string, error) {
	var resp string
	err := m.Client.Call("Plugin.GetClusterInfo", new(interface{}), &resp)
	if err != nil {
		return "", err
	}
	return resp, nil
}

func (m *SRELibraryRPCClientV2) GetManagementCluster() (string, error) {
	var resp string
	err := m.Client.Call("Plugin.GetManagementCluster", new(interface{}), &resp)
	if err != nil {
		return "", err
	}
	return resp, nil
}

// =================================================
// Server implementation of SRELibrary
// =================================================

type SRELibraryRPCServerV2 struct {
	Impl SRELibraryV2
}

func (s *SRELibraryRPCServerV2) GetClusterInfo(args interface{}, resp *string) error {
	val, err := s.Impl.GetClusterInfo()
	if err != nil {
		return err
	}
	*resp = val
	return nil
}

func (s *SRELibraryRPCServerV2) GetManagementCluster(args interface{}, resp *string) error {
	val, err := s.Impl.GetManagementCluster()
	if err != nil {
		return err
	}
	*resp = val
	return nil
}

// =================================================
// Plugin implementation
// =================================================

type SRELibraryPluginV2 struct {
	Impl SRELibraryV2
}

func (p *SRELibraryPluginV2) Server(*plugin.MuxBroker) (interface{}, error) {
	return &SRELibraryRPCServerV2{Impl: p.Impl}, nil
}

func (p *SRELibraryPluginV2) Client(_ *plugin.MuxBroker, rpcClient *rpc.Client) (interface{}, error) {
	return &SRELibraryRPCClientV2{Client: rpcClient}, nil
}

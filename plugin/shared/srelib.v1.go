package shared

import (
	"net/rpc"

	"github.com/hashicorp/go-plugin"
)

// The interface that gets communicated as an API
type SRELibraryV1 interface {
	GetClusterInfo() (string, error)
}

// HandshakeConfigs are used to just do a basic handshake between a plugin and host.
// If the handshake fails, a user friendly error is shown.
// This prevents users from executing bad plugins or executing a plugin in the wrong process.
// This is not a security measure, but it is a good sanity check.
var HandshakeConfig = plugin.HandshakeConfig{
	MagicCookieKey:   "SRE_LIBRARY_PLUGIN",
	MagicCookieValue: "magic_cookie_value",
}

// =================================================
// Client implementation of SRELibrary
// =================================================

type SRELibraryRPCClient struct {
	Client *rpc.Client
}

func (m *SRELibraryRPCClient) GetClusterInfo() (string, error) {
	var resp string
	err := m.Client.Call("Plugin.GetClusterInfo", new(interface{}), &resp)
	if err != nil {
		return "", err
	}
	return resp, nil
}

// =================================================
// Server implementation of SRELibrary
// =================================================

type SRELibraryRPCServerV1 struct {
	Impl SRELibraryV1
}

func (s *SRELibraryRPCServerV1) GetClusterInfo(args interface{}, resp *string) error {
	val, err := s.Impl.GetClusterInfo()
	if err != nil {
		return err
	}
	*resp = val
	return nil
}

// =================================================
// Plugin implementation
// =================================================

type SRELibraryPluginV1 struct {
	Impl SRELibraryV1
}

func (p *SRELibraryPluginV1) Server(*plugin.MuxBroker) (interface{}, error) {
	return &SRELibraryRPCServerV1{Impl: p.Impl}, nil
}

func (p *SRELibraryPluginV1) Client(_ *plugin.MuxBroker, rpcClient *rpc.Client) (interface{}, error) {
	return &SRELibraryRPCClient{Client: rpcClient}, nil
}

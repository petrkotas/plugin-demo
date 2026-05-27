# go-plugin Versioned Plugin Demo

Demonstrates **versioned plugin negotiation** using [HashiCorp go-plugin](https://github.com/hashicorp/go-plugin). A host application loads an SRE library as a separate process over RPC, and the two sides negotiate the highest mutually supported protocol version.

## Components

| Component | Path | Role |
|-----------|------|------|
| **Shared interfaces** | `plugin/shared/` | Defines the `SRELibraryV1` and `SRELibraryV2` interfaces, their RPC client/server stubs, and the plugin wrappers. Both host and plugin import this package. |
| **Plugin (server)** | `plugin/impl/` | Implements the SRE library logic (`GetClusterInfo`, `GetManagementCluster`) and serves it via `plugin.Serve`. Built with either the `v1` or `v2` build tag to register different version sets. |
| **Host (client)** | `main_v1.go` / `main_v2.go` | Launches the plugin binary as a subprocess, negotiates a version, and calls the library over RPC. The V2 host registers both V1 and V2 and checks `NegotiatedVersion()` to use V2-only methods. |

## Version Negotiation

- **V1 interface** &mdash; `GetClusterInfo()`
- **V2 interface** &mdash; `GetClusterInfo()` + `GetManagementCluster()`

The host and plugin each advertise which versions they support. `go-plugin` picks the highest version both sides share. This means a V2 host talking to a V1 plugin gracefully falls back to V1.

## Build & Run

```sh
make all   # builds all 4 combinations into V1_V1/, V1_V2/, V2_V1/, V2_V2/
make run   # builds and runs all 4
```

| Combo | Plugin | Host | Result |
|-------|--------|------|--------|
| `V1_V1` | V1 | V1 | Negotiates V1 &mdash; `GetClusterInfo` only |
| `V1_V2` | V1 | V2 | Host wants V2, plugin only has V1 &mdash; falls back to V1 |
| `V2_V1` | V2 | V1 | Plugin has V2, host only knows V1 &mdash; negotiates V1 |
| `V2_V2` | V2 | V2 | Both support V2 &mdash; negotiates V2, calls `GetManagementCluster` |

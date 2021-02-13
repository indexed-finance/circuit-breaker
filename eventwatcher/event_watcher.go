package eventwatcher

import (
	"github.com/bonedaddy/go-indexed/bclient"
	poolbindings "github.com/bonedaddy/go-indexed/bindings/pool"
	"github.com/ethereum/go-ethereum/common"
)

// EventWatcher wraps one or more interfaces to the ethereum blockchain to watch events
type EventWatcher struct {
	bc *bclient.Client
	// TODO(bonedaddy): add highly available lookups
}

// New returns a new EventWatcher implementation
func New(bc *bclient.Client) *EventWatcher {
	return &EventWatcher{bc: bc}
}

// NewBindings constructs poolbindings for various indices, the input addresses argument is a map of pool name -> pool address
func (ew *EventWatcher) NewBindings(addresses map[string]string) (map[string]*poolbindings.Poolbindings, error) {
	out := make(map[string]*poolbindings.Poolbindings)
	for name, addr := range addresses {
		binding, err := poolbindings.NewPoolbindings(common.HexToAddress(addr), ew.bc.EthClient())
		if err != nil {
			return nil, err
		}
		out[name] = binding
	}
	return out, nil
}

// BC returns the underlying blockchain client
func (ew *EventWatcher) BC() *bclient.Client {
	return ew.bc
}

// Close is used to terminate the event watcher
func (ew *EventWatcher) Close() error {
	ew.bc.Close()
	return nil
}

package eventwatcher

import (
	"github.com/bonedaddy/go-defi/utils"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/indexed-finance/circuit-breaker/bindings/sigmacore"
)

// EventWatcher wraps one or more interfaces to the ethereum blockchain to watch events
type EventWatcher struct {
	bc utils.Blockchain
	// bc *bclient.Client
	// TODO(bonedaddy): add highly available lookups
}

// New returns a new EventWatcher implementation
func New(bc utils.Blockchain) *EventWatcher {
	return &EventWatcher{bc: bc}
}

// NewBindings constructs poolbindings for various indices, the input addresses argument is a map of pool name -> pool address
func (ew *EventWatcher) NewBindings(addresses map[string]string) (map[string]*sigmacore.Sigmacore, error) {
	out := make(map[string]*sigmacore.Sigmacore)
	for name, addr := range addresses {
		binding, err := sigmacore.NewSigmacore(common.HexToAddress(addr), ew.bc)
		if err != nil {
			return nil, err
		}
		out[name] = binding
	}
	return out, nil
}

// BC returns the underlying blockchain client
func (ew *EventWatcher) BC() utils.Blockchain {
	return ew.bc
}

// Close is used to terminate the event watcher
func (ew *EventWatcher) Close() error {
	if ibc, ok := ew.bc.(*ethclient.Client); ok {
		ibc.Close()
	}
	return nil
}

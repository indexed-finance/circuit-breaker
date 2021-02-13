package eventwatcher

import (
	"os"
	"testing"

	"github.com/bonedaddy/go-indexed/bclient"
	"github.com/stretchr/testify/require"
)

var (
	pools = map[string]string{
		"defi5": "0xfa6de2697d59e88ed7fc4dfe5a33dac43565ea41",
		"cc10":  "0x17ac188e09a7890a1844e5e65471fe8b0ccfadf3",
	}
)

func doSetup(t *testing.T) *EventWatcher {
	infuraAPIKey := os.Getenv("INFURA_API_KEY")
	if infuraAPIKey == "" {
		t.Fatal("INFURA_API_KEY env var is empty")
	}
	client, err := bclient.NewInfuraClient(infuraAPIKey, true)
	require.NoError(t, err)
	return New(client)
}

func TestEventWatcher(t *testing.T) {
	client := doSetup(t)
	t.Cleanup(func() {
		client.Close()
	})
	_, err := client.NewBindings(pools)
	require.NoError(t, err)
}

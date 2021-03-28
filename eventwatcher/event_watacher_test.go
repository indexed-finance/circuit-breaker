package eventwatcher

import (
	"context"
	"testing"

	"github.com/bonedaddy/go-defi/testenv"
	"github.com/stretchr/testify/require"
)

var (
	pools = map[string]string{
		"defi5": "0xfa6de2697d59e88ed7fc4dfe5a33dac43565ea41",
		"cc10":  "0x17ac188e09a7890a1844e5e65471fe8b0ccfadf3",
	}
)

func TestEventWatcher(t *testing.T) {
	tenv, err := testenv.NewBlockchain(context.Background())
	require.NoError(t, err)
	client := New(tenv)
	_, err = client.NewBindings(pools)
	require.NoError(t, err)
}

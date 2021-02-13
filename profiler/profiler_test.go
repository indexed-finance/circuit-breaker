package profiler

import (
	"testing"
	"time"

	"github.com/stretchr/testify/require"
)

func TestProfiler(t *testing.T) {
	profiler := New()
	profiler.Start("localhost:6060", true)
	time.Sleep(time.Second * 2)
	require.NoError(t, profiler.Close())
}

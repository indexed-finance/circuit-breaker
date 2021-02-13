package utils

import (
	"math/big"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestGetChangePercentageBig(t *testing.T) {
	type args struct {
		final *big.Int
		start *big.Int
	}
	tests := []struct {
		name       string
		args       args
		wantChange float64
	}{
		{
			"0",
			args{
				ToWei("300.0", 18),
				ToWei("30.0", 18),
			},
			9.0,
		},
		{
			"1",
			args{
				ToWei("10000000.0", 18), // 10 mil
				ToWei("3000000.0", 18),  // 3 mil
			},
			2.3333333333333335,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ret := GetChangePercentageBig(tt.args.final, tt.args.start)
			t.Log("ret: ", ret, " want: ", tt.wantChange)
			require.Equal(t, ret, tt.wantChange)
		})
	}
}

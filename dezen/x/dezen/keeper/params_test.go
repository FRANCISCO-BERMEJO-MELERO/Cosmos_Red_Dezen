package keeper_test

import (
	"testing"

	"github.com/stretchr/testify/require"

	keepertest "dezen/testutil/keeper"
	"dezen/x/dezen/types"
)

func TestGetParams(t *testing.T) {
	k, ctx := keepertest.DezenKeeper(t)
	params := types.DefaultParams()

	require.NoError(t, k.SetParams(ctx, params))
	require.EqualValues(t, params, k.GetParams(ctx))
}

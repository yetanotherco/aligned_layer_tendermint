package keeper_test

import (
	"testing"

	"github.com/stretchr/testify/require"

	keepertest "alignedlayer/testutil/keeper"
	"alignedlayer/x/verify/types"
)

func TestGetParams(t *testing.T) {
	k, ctx := keepertest.VerifyKeeper(t)
	params := types.DefaultParams()

	require.NoError(t, k.SetParams(ctx, params))
	require.EqualValues(t, params, k.GetParams(ctx))
}

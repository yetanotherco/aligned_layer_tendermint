package verify_test

import (
	"testing"

	keepertest "alignedlayer/testutil/keeper"
	"alignedlayer/testutil/nullify"
	verify "alignedlayer/x/verify/module"
	"alignedlayer/x/verify/types"

	"github.com/stretchr/testify/require"
)

func TestGenesis(t *testing.T) {
	genesisState := types.GenesisState{
		Params: types.DefaultParams(),

		// this line is used by starport scaffolding # genesis/test/state
	}

	k, ctx := keepertest.VerifyKeeper(t)
	verify.InitGenesis(ctx, k, genesisState)
	got := verify.ExportGenesis(ctx, k)
	require.NotNil(t, got)

	nullify.Fill(&genesisState)
	nullify.Fill(got)

	// this line is used by starport scaffolding # genesis/test/assert
}

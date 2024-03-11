package verification_test

import (
	"testing"

	keepertest "alignedlayer/testutil/keeper"
	"alignedlayer/testutil/nullify"
	verification "alignedlayer/x/verification/module"
	"alignedlayer/x/verification/types"

	"github.com/stretchr/testify/require"
)

func TestGenesis(t *testing.T) {
	genesisState := types.GenesisState{
		Params: types.DefaultParams(),

		// this line is used by starport scaffolding # genesis/test/state
	}

	k, ctx := keepertest.VerificationKeeper(t)
	verification.InitGenesis(ctx, k, genesisState)
	got := verification.ExportGenesis(ctx, k)
	require.NotNil(t, got)

	nullify.Fill(&genesisState)
	nullify.Fill(got)

	// this line is used by starport scaffolding # genesis/test/assert
}

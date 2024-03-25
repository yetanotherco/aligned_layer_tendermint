package verify

import (
	"math/rand"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/module"
	simtypes "github.com/cosmos/cosmos-sdk/types/simulation"
	"github.com/cosmos/cosmos-sdk/x/simulation"

	"alignedlayer/testutil/sample"
	verifysimulation "alignedlayer/x/verify/simulation"
	"alignedlayer/x/verify/types"
)

// avoid unused import issue
var (
	_ = verifysimulation.FindAccount
	_ = rand.Rand{}
	_ = sample.AccAddress
	_ = sdk.AccAddress{}
	_ = simulation.MsgEntryKind
)

const (
	opWeightMsgGnarkPlonk = "op_weight_msg_gnark_plonk"
	// TODO: Determine the simulation weight value
	defaultWeightMsgGnarkPlonk int = 100

	opWeightMsgCairoPlatinum = "op_weight_msg_cairo_platinum"
	// TODO: Determine the simulation weight value
	defaultWeightMsgCairoPlatinum int = 100

	// this line is used by starport scaffolding # simapp/module/const
)

// GenerateGenesisState creates a randomized GenState of the module.
func (AppModule) GenerateGenesisState(simState *module.SimulationState) {
	accs := make([]string, len(simState.Accounts))
	for i, acc := range simState.Accounts {
		accs[i] = acc.Address.String()
	}
	verifyGenesis := types.GenesisState{
		Params: types.DefaultParams(),
		// this line is used by starport scaffolding # simapp/module/genesisState
	}
	simState.GenState[types.ModuleName] = simState.Cdc.MustMarshalJSON(&verifyGenesis)
}

// RegisterStoreDecoder registers a decoder.
func (am AppModule) RegisterStoreDecoder(_ simtypes.StoreDecoderRegistry) {}

// ProposalContents doesn't return any content functions for governance proposals.
func (AppModule) ProposalContents(_ module.SimulationState) []simtypes.WeightedProposalContent {
	return nil
}

// WeightedOperations returns the all the gov module operations with their respective weights.
func (am AppModule) WeightedOperations(simState module.SimulationState) []simtypes.WeightedOperation {
	operations := make([]simtypes.WeightedOperation, 0)

	var weightMsgGnarkPlonk int
	simState.AppParams.GetOrGenerate(opWeightMsgGnarkPlonk, &weightMsgGnarkPlonk, nil,
		func(_ *rand.Rand) {
			weightMsgGnarkPlonk = defaultWeightMsgGnarkPlonk
		},
	)
	operations = append(operations, simulation.NewWeightedOperation(
		weightMsgGnarkPlonk,
		verifysimulation.SimulateMsgGnarkPlonk(am.accountKeeper, am.bankKeeper, am.keeper),
	))

	var weightMsgCairoPlatinum int
	simState.AppParams.GetOrGenerate(opWeightMsgCairoPlatinum, &weightMsgCairoPlatinum, nil,
		func(_ *rand.Rand) {
			weightMsgCairoPlatinum = defaultWeightMsgCairoPlatinum
		},
	)
	operations = append(operations, simulation.NewWeightedOperation(
		weightMsgCairoPlatinum,
		verifysimulation.SimulateMsgCairoPlatinum(am.accountKeeper, am.bankKeeper, am.keeper),
	))

	// this line is used by starport scaffolding # simapp/module/operation

	return operations
}

// ProposalMsgs returns msgs used for governance proposals for simulations.
func (am AppModule) ProposalMsgs(simState module.SimulationState) []simtypes.WeightedProposalMsg {
	return []simtypes.WeightedProposalMsg{
		simulation.NewWeightedProposalMsg(
			opWeightMsgGnarkPlonk,
			defaultWeightMsgGnarkPlonk,
			func(r *rand.Rand, ctx sdk.Context, accs []simtypes.Account) sdk.Msg {
				verifysimulation.SimulateMsgGnarkPlonk(am.accountKeeper, am.bankKeeper, am.keeper)
				return nil
			},
		),
		simulation.NewWeightedProposalMsg(
			opWeightMsgCairoPlatinum,
			defaultWeightMsgCairoPlatinum,
			func(r *rand.Rand, ctx sdk.Context, accs []simtypes.Account) sdk.Msg {
				verifysimulation.SimulateMsgCairoPlatinum(am.accountKeeper, am.bankKeeper, am.keeper)
				return nil
			},
		),
		// this line is used by starport scaffolding # simapp/module/OpMsg
	}
}

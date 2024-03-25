package verification

import (
	"math/rand"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/module"
	simtypes "github.com/cosmos/cosmos-sdk/types/simulation"
	"github.com/cosmos/cosmos-sdk/x/simulation"

	"alignedlayer/testutil/sample"
	verificationsimulation "alignedlayer/x/verification/simulation"
	"alignedlayer/x/verification/types"
)

// avoid unused import issue
var (
	_ = verificationsimulation.FindAccount
	_ = rand.Rand{}
	_ = sample.AccAddress
	_ = sdk.AccAddress{}
	_ = simulation.MsgEntryKind
)

const (
	opWeightMsgVerifyPlonk = "op_weight_msg_verify_plonk"
	// TODO: Determine the simulation weight value
	defaultWeightMsgVerifyPlonk int = 100

	opWeightMsgVerifyCairo = "op_weight_msg_verify_cairo"
	// TODO: Determine the simulation weight value
	defaultWeightMsgVerifyCairo int = 100

	// this line is used by starport scaffolding # simapp/module/const
)

// GenerateGenesisState creates a randomized GenState of the module.
func (AppModule) GenerateGenesisState(simState *module.SimulationState) {
	accs := make([]string, len(simState.Accounts))
	for i, acc := range simState.Accounts {
		accs[i] = acc.Address.String()
	}
	verificationGenesis := types.GenesisState{
		Params: types.DefaultParams(),
		// this line is used by starport scaffolding # simapp/module/genesisState
	}
	simState.GenState[types.ModuleName] = simState.Cdc.MustMarshalJSON(&verificationGenesis)
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

	var weightMsgVerifyPlonk int
	simState.AppParams.GetOrGenerate(opWeightMsgVerifyPlonk, &weightMsgVerifyPlonk, nil,
		func(_ *rand.Rand) {
			weightMsgVerifyPlonk = defaultWeightMsgVerifyPlonk
		},
	)
	operations = append(operations, simulation.NewWeightedOperation(
		weightMsgVerifyPlonk,
		verificationsimulation.SimulateMsgVerifyPlonk(am.accountKeeper, am.bankKeeper, am.keeper),
	))

	var weightMsgVerifyCairo int
	simState.AppParams.GetOrGenerate(opWeightMsgVerifyCairo, &weightMsgVerifyCairo, nil,
		func(_ *rand.Rand) {
			weightMsgVerifyCairo = defaultWeightMsgVerifyCairo
		},
	)
	operations = append(operations, simulation.NewWeightedOperation(
		weightMsgVerifyCairo,
		verificationsimulation.SimulateMsgVerifyCairo(am.accountKeeper, am.bankKeeper, am.keeper),
	))

	// this line is used by starport scaffolding # simapp/module/operation

	return operations
}

// ProposalMsgs returns msgs used for governance proposals for simulations.
func (am AppModule) ProposalMsgs(simState module.SimulationState) []simtypes.WeightedProposalMsg {
	return []simtypes.WeightedProposalMsg{
		simulation.NewWeightedProposalMsg(
			opWeightMsgVerifyPlonk,
			defaultWeightMsgVerifyPlonk,
			func(r *rand.Rand, ctx sdk.Context, accs []simtypes.Account) sdk.Msg {
				verificationsimulation.SimulateMsgVerifyPlonk(am.accountKeeper, am.bankKeeper, am.keeper)
				return nil
			},
		),
		simulation.NewWeightedProposalMsg(
			opWeightMsgVerifyCairo,
			defaultWeightMsgVerifyCairo,
			func(r *rand.Rand, ctx sdk.Context, accs []simtypes.Account) sdk.Msg {
				verificationsimulation.SimulateMsgVerifyCairo(am.accountKeeper, am.bankKeeper, am.keeper)
				return nil
			},
		),
		// this line is used by starport scaffolding # simapp/module/OpMsg
	}
}

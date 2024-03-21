package simulation

import (
	"math/rand"

	"alignedlayer/x/verification/keeper"
	"alignedlayer/x/verification/types"

	"github.com/cosmos/cosmos-sdk/baseapp"
	sdk "github.com/cosmos/cosmos-sdk/types"
	simtypes "github.com/cosmos/cosmos-sdk/types/simulation"
)

func SimulateMsgVerifyPlonk(
	ak types.AccountKeeper,
	bk types.BankKeeper,
	k keeper.Keeper,
) simtypes.Operation {
	return func(r *rand.Rand, app *baseapp.BaseApp, ctx sdk.Context, accs []simtypes.Account, chainID string,
	) (simtypes.OperationMsg, []simtypes.FutureOperation, error) {
		simAccount, _ := simtypes.RandomAcc(r, accs)
		msg := &types.MsgVerifyPlonk{
			Creator: simAccount.Address.String(),
		}

		// TODO: Handling the VerifyPlonk simulation

		return simtypes.NoOpMsg(types.ModuleName, sdk.MsgTypeURL(msg), "VerifyPlonk simulation not implemented"), nil, nil
	}
}

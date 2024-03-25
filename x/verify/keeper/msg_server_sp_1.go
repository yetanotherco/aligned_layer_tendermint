package keeper

import (
	"context"

	"alignedlayer/x/verify/types"

	sdk "github.com/cosmos/cosmos-sdk/types"
)

func (k msgServer) Sp1(goCtx context.Context, msg *types.MsgSp1) (*types.MsgSp1Response, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	// TODO: Handling the message
	_ = ctx

	return &types.MsgSp1Response{}, nil
}

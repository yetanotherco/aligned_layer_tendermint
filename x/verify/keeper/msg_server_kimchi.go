package keeper

import (
	"context"

	"alignedlayer/x/verify/types"

	sdk "github.com/cosmos/cosmos-sdk/types"
)

func (k msgServer) Kimchi(goCtx context.Context, msg *types.MsgKimchi) (*types.MsgKimchiResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	// TODO: Handling the message
	_ = ctx

	return &types.MsgKimchiResponse{}, nil
}

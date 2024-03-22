package keeper

import (
	"context"

	"alignedlayer/x/verify/types"

	sdk "github.com/cosmos/cosmos-sdk/types"
)

func (k msgServer) GnarkPlonk(goCtx context.Context, msg *types.MsgGnarkPlonk) (*types.MsgGnarkPlonkResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	// TODO: Handling the message
	_ = ctx

	return &types.MsgGnarkPlonkResponse{}, nil
}

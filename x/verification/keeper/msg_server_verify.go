package keeper

import (
	"context"
	"strings"

	"alignedlayer/x/verification/types"

	sdk "github.com/cosmos/cosmos-sdk/types"
)

func (k msgServer) Verify(goCtx context.Context, msg *types.MsgVerify) (*types.MsgVerifyResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)
	_ = ctx

	if !verify(msg.Proof) {
		return nil, types.ErrSample
	} else {
		return &types.MsgVerifyResponse{}, nil
	}
}

func verify(proof string) bool {
	return !strings.Contains(proof, "invalid")
}

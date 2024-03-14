package keeper

import (
	"context"
	"strings"

	"alignedlayer/x/verification/types"

	sdk "github.com/cosmos/cosmos-sdk/types"
)

func (k msgServer) VerifySp1(goCtx context.Context, msg *types.MsgVerifySp1) (*types.MsgVerifySp1Response, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)
	_ = ctx

	if !verifySP1(msg.Proof) {
		return nil, types.ErrSample
	} else {
		return &types.MsgVerifySp1Response{}, nil
	}
}

func verifySP1(proof string) bool {
	return !strings.Contains(proof, "invalid")
}

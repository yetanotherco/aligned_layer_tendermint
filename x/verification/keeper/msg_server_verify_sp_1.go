package keeper

import (
	"context"

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
	//proofBytes := []byte(proof)

	//var proofArray [sp1.MAX_PROOF_SIZE]byte
	//copy(proofArray[:], proofBytes)

	//return sp1.VerifySp1Proof(([sp1.MAX_PROOF_SIZE]byte)(proofBytes), uint(len(proofBytes)))
	return true
}

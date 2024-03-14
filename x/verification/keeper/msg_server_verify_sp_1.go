package keeper

import (
	"alignedlayer/x/verification/types"
	"context"
	"strconv"

	sp1 "alignedlayer/operators/sp1"

	sdk "github.com/cosmos/cosmos-sdk/types"
)

func (k msgServer) VerifySp1(goCtx context.Context, msg *types.MsgVerifySp1) (*types.MsgVerifySp1Response, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)
	_ = ctx

	result := verifySP1(msg.Proof)
	event := sdk.NewEvent("SP1verification_finished",
		sdk.NewAttribute("SP1proof_verifies", strconv.FormatBool(result)))

	ctx.EventManager().EmitEvent(event)

	return &types.MsgVerifySp1Response{}, nil
}

func verifySP1(proof string) bool {
	proofBytes := []byte(proof)

	var proofArray [sp1.MAX_PROOF_SIZE]byte
	copy(proofArray[:], proofBytes)

	return sp1.VerifySp1Proof(([sp1.MAX_PROOF_SIZE]byte)(proofBytes), uint(len(proofBytes)))
}

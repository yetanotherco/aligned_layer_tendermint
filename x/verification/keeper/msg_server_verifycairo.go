package keeper

import (
	"context"
	"strconv"

	"alignedlayer/x/verification/types"

	cp "alignedlayer/operators/cairo_platinum"

	sdk "github.com/cosmos/cosmos-sdk/types"
)

func (k msgServer) Verifycairo(goCtx context.Context, msg *types.MsgVerifycairo) (*types.MsgVerifycairoResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)
	_ = ctx

	result := verifyCairo(msg.Proof)
	event := sdk.NewEvent("SP1verification_finished",
		sdk.NewAttribute("SP1proof_verifies", strconv.FormatBool(result)))

	ctx.EventManager().EmitEvent(event)

	return &types.MsgVerifycairoResponse{}, nil
}

func verifyCairo(proof string) bool {
	proofBytes := []byte(proof)

	var proofArray [cp.MAX_PROOF_SIZE]byte
	copy(proofArray[:], proofBytes)

	return !cp.VerifyCairoProof100Bits(([cp.MAX_PROOF_SIZE]byte)(proofBytes), uint(len(proofBytes)))
}

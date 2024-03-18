package keeper

import (
	"context"
	"encoding/base64"
	"strconv"

	"alignedlayer/x/verification/types"

	cp "alignedlayer/operators/cairo_platinum"

	sdk "github.com/cosmos/cosmos-sdk/types"
)

func (k msgServer) Verifycairo(goCtx context.Context, msg *types.MsgVerifycairo) (*types.MsgVerifycairoResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)
	_ = ctx

	result := verifyCairo(msg.Proof)
	event := sdk.NewEvent("CAIROverification_finished",
		sdk.NewAttribute("CAIROproof_verifies", strconv.FormatBool(result)))

	ctx.EventManager().EmitEvent(event)

	return &types.MsgVerifycairoResponse{}, nil
}

func verifyCairo(proof string) bool {
	decodedBytes := make([]byte, cp.MAX_PROOF_SIZE)
	nDecoded, err := base64.StdEncoding.Decode(decodedBytes, []byte(proof))
	if err != nil {
		return false
	}

	return !cp.VerifyCairoProof100Bits(([cp.MAX_PROOF_SIZE]byte)(decodedBytes), uint(nDecoded))
}

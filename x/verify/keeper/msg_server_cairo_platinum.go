package keeper

import (
	"context"
	"encoding/base64"
	"strconv"

	"alignedlayer/x/verify/types"

	cp "alignedlayer/verifiers/cairo_platinum"

	sdk "github.com/cosmos/cosmos-sdk/types"
)

func (k msgServer) CairoPlatinum(goCtx context.Context, msg *types.MsgCairoPlatinum) (*types.MsgCairoPlatinumResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	result := verifyCairoPlatinum(msg.Proof)
	event := sdk.NewEvent("verification_finished",
		sdk.NewAttribute("proof_verifies", strconv.FormatBool(result)),
		sdk.NewAttribute("prover", "CAIRO"))

	ctx.EventManager().EmitEvent(event)

	return &types.MsgCairoPlatinumResponse{}, nil
}

func verifyCairoPlatinum(proof string) bool {
	if len(proof)%3 != 0 {
		return false
	}
	decodedBytes := make([]byte, cp.MAX_PROOF_SIZE)
	nDecoded, err := base64.StdEncoding.Decode(decodedBytes, []byte(proof))
	if err != nil {
		return false
	}

	return cp.VerifyCairoProof100Bits(([cp.MAX_PROOF_SIZE]byte)(decodedBytes), uint(nDecoded))
}

package keeper

import (
	"context"
	"encoding/base64"
	"fmt"
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
	decodedBytes, err := base64.StdEncoding.DecodeString(proof)
	if err != nil {
		fmt.Println("Error decoding base64 string:", err)
		return false
	}

	return !cp.VerifyCairoProof100Bits(([cp.MAX_PROOF_SIZE]byte)(decodedBytes), uint(len(decodedBytes)))
}

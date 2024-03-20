package keeper

import (
	"alignedlayer/x/verification/types"
	"context"
	"encoding/base64"
	"strconv"

	sp1 "alignedlayer/operators/sp1"

	sdk "github.com/cosmos/cosmos-sdk/types"
)

func (k msgServer) VerifySp1(goCtx context.Context, msg *types.MsgVerifySp1) (*types.MsgVerifySp1Response, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	result := verifySP1(msg.Proof)
	event := sdk.NewEvent("SP1verification_finished",
		sdk.NewAttribute("SP1proof_verifies", strconv.FormatBool(result)))

	ctx.EventManager().EmitEvent(event)

	return &types.MsgVerifySp1Response{}, nil
}

func verifySP1(proof string) bool {
	if len(proof)%3 != 0 {
		return false
	}
	decodedBytes := make([]byte, sp1.MAX_PROOF_SIZE)
	nDecoded, err := base64.StdEncoding.Decode(decodedBytes, []byte(proof))
	if err != nil {
		return false
	}

	return sp1.VerifySp1Proof(([sp1.MAX_PROOF_SIZE]byte)(decodedBytes), uint(nDecoded))
}

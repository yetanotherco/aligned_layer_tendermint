package keeper

import (
	"context"
	"encoding/base64"
	"strconv"

	sp1 "alignedlayer/verifiers/sp1"
	"alignedlayer/x/verify/types"

	sdk "github.com/cosmos/cosmos-sdk/types"
)

func (k msgServer) Sp1(goCtx context.Context, msg *types.MsgSp1) (*types.MsgSp1Response, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	result := verifySp1(msg.Proof)
	event := sdk.NewEvent("verification_finished",
		sdk.NewAttribute("proof_verifies", strconv.FormatBool(result)),
		sdk.NewAttribute("prover", "SP1"))

	ctx.EventManager().EmitEvent(event)

	return &types.MsgSp1Response{}, nil
}

func verifySp1(proof string) bool {
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

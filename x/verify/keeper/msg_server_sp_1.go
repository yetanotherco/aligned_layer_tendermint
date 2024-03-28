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

	result := verifySp1(msg.Proof, msg.Elf)
	event := sdk.NewEvent("verification_finished",
		sdk.NewAttribute("proof_verifies", strconv.FormatBool(result)),
		sdk.NewAttribute("prover", "SP1"))

	ctx.EventManager().EmitEvent(event)

	return &types.MsgSp1Response{}, nil
}

func verifySp1(proof string, elf string) bool {
	decodedProof, err := base64.StdEncoding.DecodeString(proof)
	if err != nil {
		return false
	}

	decodedElf, err := base64.StdEncoding.DecodeString(elf)
	if err != nil {
		return false
	}

	return sp1.VerifySp1ProofElf(decodedProof, decodedElf)
}

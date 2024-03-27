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
	if len(proof)%3 != 0 {
		return false
	}

	decodedProof := make([]byte, sp1.MAX_PROOF_SIZE)
	nDecodedProof, err := base64.StdEncoding.Decode(decodedProof, []byte(proof))
	if err != nil {
		return false
	}

	decodedElf := make([]byte, sp1.MAX_PROOF_SIZE)
	nDecodedElf, err := base64.StdEncoding.Decode(decodedElf, []byte(elf))
	if err != nil {
		return false
	}

	return sp1.VerifySp1ProofElf(([sp1.MAX_PROOF_SIZE]byte)(decodedProof), ([sp1.MAX_PROOF_SIZE]byte)(decodedElf), uint(nDecodedProof), uint(nDecodedElf))
}

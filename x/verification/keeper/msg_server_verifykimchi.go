package keeper

import (
	"context"
	"encoding/base64"
	"os"
	"strconv"

	kim "alignedlayer/operators/kimchi"
	"alignedlayer/x/verification/types"

	sdk "github.com/cosmos/cosmos-sdk/types"
)

func (k msgServer) Verifykimchi(goCtx context.Context, msg *types.MsgVerifykimchi) (*types.MsgVerifykimchiResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	result := verifykim(msg.Proof)
	event := sdk.NewEvent("KIMCHIverification_finished",
		sdk.NewAttribute("KIMCHIproof_verifies", strconv.FormatBool(result)))

	ctx.EventManager().EmitEvent(event)

	return &types.MsgVerifykimchiResponse{}, nil
}

func verifykim(proof string) bool {
	decodedBytes := make([]byte, kim.MAX_PROOF_SIZE)
	nDecoded, err := base64.StdEncoding.Decode(decodedBytes, []byte(proof))

	if err != nil {
		return false
	}
	pubInputFile, err := os.ReadFile("../../../operators/kimchi/lib/verifier_base64")
	if err != nil {
		return false
	}
	pubInputBuffer := make([]byte, kim.MAX_PUB_INPUT_SIZE)
	pubInputLen, err := base64.StdEncoding.Decode(pubInputBuffer, []byte(pubInputFile))
	if err != nil {
		return false
	}

	return kim.VerifyKimchiProof(([kim.MAX_PROOF_SIZE]byte)(decodedBytes), uint(nDecoded), ([kim.MAX_PUB_INPUT_SIZE]byte)(pubInputBuffer), uint(pubInputLen))
}

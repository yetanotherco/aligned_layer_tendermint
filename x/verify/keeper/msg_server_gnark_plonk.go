package keeper

import (
	"bytes"
	"context"
	"encoding/base64"
	"io"
	"strconv"

	"alignedlayer/x/verify/types"

	"github.com/consensys/gnark-crypto/ecc"
	"github.com/consensys/gnark/backend/plonk"
	"github.com/consensys/gnark/backend/witness"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

func (k msgServer) GnarkPlonk(goCtx context.Context, msg *types.MsgGnarkPlonk) (*types.MsgGnarkPlonkResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	result := verifyGnarkPlonk(msg)
	event := sdk.NewEvent("verification_finished",
		sdk.NewAttribute("proof_verifies", strconv.FormatBool(result)),
		sdk.NewAttribute("prover", "PLONK"))

	ctx.EventManager().EmitEvent(event)

	return &types.MsgGnarkPlonkResponse{}, nil
}

func verifyGnarkPlonk(msg *types.MsgGnarkPlonk) bool {
	proof := plonk.NewProof(ecc.BN254)
	deserialize(proof, msg.Proof)

	public_input, _ := witness.New(ecc.BN254.ScalarField())
	deserialize(public_input, msg.PublicInputs)

	verifying_key := plonk.NewVerifyingKey(ecc.BN254)
	deserialize(verifying_key, msg.VerifyingKey)

	err := plonk.Verify(proof, verifying_key, public_input)

	return err == nil
}

func deserialize[r io.ReaderFrom](dst r, encoded string) error {
	bytes_buffer, err := base64.StdEncoding.DecodeString(encoded)
	if err != nil {
		return err
	}
	reader := bytes.NewReader(bytes_buffer)
	_, err = dst.ReadFrom(reader)

	return err
}

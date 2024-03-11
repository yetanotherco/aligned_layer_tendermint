package keeper

import (
	"bytes"
	"context"
	"encoding/base64"

	"lambchain/x/lambchain/types"

	"github.com/consensys/gnark-crypto/ecc"
	"github.com/consensys/gnark/backend/plonk"
	"github.com/consensys/gnark/backend/witness"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

func (k msgServer) Verify(goCtx context.Context, msg *types.MsgVerify) (*types.MsgVerifyResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)
	_ = ctx
	if !verify(msg.Proof, msg.Pw) {
		return nil, types.ErrSample
	} else {
		return &types.MsgVerifyResponse{}, nil
	}
}

func verify(proof string, pw string) bool {
	// Deserialize Proof
	proof_des := plonk.NewProof(ecc.BN254)
	proof_base64 := []byte(proof)
	proof_bytes := make([]byte, base64.StdEncoding.DecodedLen(len(proof_base64)))
	_, err := base64.StdEncoding.Decode(proof_bytes, proof_base64)
	if err != nil {
		return false
	}
	proof_reader := bytes.NewReader(proof_bytes)
	proof_des.ReadFrom(proof_reader)

	// Deserialize VK
	vk_base64 := []byte("AAAAAAAAEAAwYUgt+gOND7W0wLImGUBHomFlCfUx1Po6zbd0lsEAAQkx1ZbeL9EPAd3Qc/1akKl28WnHbwObuRxHdXIAQtQ6AAAAAAAAAAIAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAABY0AAYjsYnxFxsZkSu01UJCzT1TY4wYKO83x85HuGpqkzSzR+y1tHogeTh/+BcPLL/2AeRJKWvh4eVXYfAmQIy/D7eif1IruCwSxWxwZTmtNYy/363lzydub3QyU5wvIfdyNz9ebgVHFSJk0W0bbfce9QInLceebJI0g1EOb/FWqiWd0pz3ECLSbEO3s9B6fud9SXnoAiDU/vWOFyTcMwtDmXLhhsW9q+2Dc31zs5cjt7xb6XH0pJtRgta56sagSyYpq3fUBO/lz6ywSPm8dQcrjV+ySL18Cx0kMqR5UGAw8luJ/u3gnVlJlcXmammrOj7e80CJ3s7uOfYk3t9QAqzwAAAAAgAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAGZjpOTkg1IOnJgv7cx+10l8apJMzWp5xKX5IW3rvMSwhgA3u8SHx52QmoAZl5cRHlnQyLU917a3UbevVzZkvbt5dt2YLAKXdBOzOUs6dbrwKWzdDbNCPOVHnlXonvXKRkacomr1SkbI3KZEOkEZK7iZBUCXU8tSzV7m2o/9P9hmgAAAAA=")
	vk_des := plonk.NewVerifyingKey(ecc.BN254)
	vk_bytes := make([]byte, base64.StdEncoding.DecodedLen(len(vk_base64)))
	base64.StdEncoding.Decode(vk_bytes, vk_base64)
	vk_reader := bytes.NewReader(vk_bytes)
	vk_des.ReadFrom(vk_reader)

	// Deserialize PublicWitness
	pw_des, err := witness.New(ecc.BN254.ScalarField())
	if err != nil {
		return false
	}
	pw_base64 := []byte(pw)
	pw_bytes := make([]byte, base64.StdEncoding.DecodedLen(len(pw_base64)))
	base64.StdEncoding.Decode(pw_bytes, pw_base64)
	pw_des.UnmarshalBinary(pw_bytes)

	err = plonk.Verify(proof_des, vk_des, pw_des)

	return err == nil
}

package main

import (
	"bytes"
	"encoding/base64"
	"fmt"
	"os"

	"github.com/consensys/gnark-crypto/ecc"
	"github.com/consensys/gnark/backend/plonk"
	"github.com/consensys/gnark/backend/witness"
	cs "github.com/consensys/gnark/constraint/bn254"
	"github.com/consensys/gnark/frontend/cs/scs"
	"github.com/consensys/gnark/test"

	"github.com/consensys/gnark/frontend"
)

// gnark is a zk-SNARK library written in Go. Circuits are regular structs.
// The inputs must be of type frontend.Variable and make up the witness.
// The witness has a
//   - secret part --> known to the prover only
//   - public part --> known to the prover and the verifier
type Circuit struct {
	X frontend.Variable `gnark:",public"` // x  --> public visibility
	Y frontend.Variable `gnark:",public"` // Y  --> public visibility

	E frontend.Variable
}

// Define declares the circuit logic. The compiler then produces a list of constraints
// y == x**e
func (circuit *Circuit) Define(api frontend.API) error {
	// number of bits of exponent
	const bitSize = 4000

	// specify constraints
	output := frontend.Variable(1)
	bits := api.ToBinary(circuit.E, bitSize)

	for i := 0; i < len(bits); i++ {
		// api.Println(fmt.Sprintf("e[%d]", i), bits[i]) // we may print a variable for testing and / or debugging purposes

		if i != 0 {
			output = api.Mul(output, output)
		}
		multiply := api.Mul(output, circuit.X)
		output = api.Select(bits[len(bits)-1-i], multiply, output)

	}

	api.AssertIsEqual(circuit.Y, output)

	return nil
}

func main() {
	var myCircuit Circuit
	ccs, _ := frontend.Compile(ecc.BN254.ScalarField(), scs.NewBuilder, &myCircuit)

	scs := ccs.(*cs.SparseR1CS)
	kzgsrs, _ := test.NewKZGSRS(scs)

	// Witnesses instantiation. Witness is known only by the prover,
	// while public w is a public data known by the verifier.
	var w Circuit
	w.X = 2
	w.E = 2
	w.Y = 4

	witnessFull, _ := frontend.NewWitness(&w, ecc.BN254.ScalarField())
	witnessPublic, _ := frontend.NewWitness(&w, ecc.BN254.ScalarField(), frontend.PublicOnly())

	// public data consists of the polynomials describing the constants involved
	// in the constraints, the polynomial describing the permutation ("grand
	// product argument"), and the FFT domains.

	pk, vk, _ := plonk.Setup(ccs, kzgsrs)

	proof, _ := plonk.Prove(ccs, pk, witnessFull)

	// Serialize PK
	var pk_buf bytes.Buffer
	pk.WriteTo(&pk_buf)
	pk_buf_base64 := make([]byte, base64.StdEncoding.EncodedLen(pk_buf.Len()))
	base64.StdEncoding.Encode(pk_buf_base64, pk_buf.Bytes())
	os.WriteFile("pk.bin", pk_buf_base64, 0644)

	// Serialize VK
	var vk_buf bytes.Buffer
	vk.WriteTo(&vk_buf)
	vk_buf_base64 := make([]byte, base64.StdEncoding.EncodedLen(vk_buf.Len()))
	base64.StdEncoding.Encode(vk_buf_base64, vk_buf.Bytes())
	os.WriteFile("vk.bin", vk_buf_base64, 0644)

	// Serialize Proof
	var proof_buf bytes.Buffer
	proof.WriteTo(&proof_buf)
	proof_buf_base64 := make([]byte, base64.StdEncoding.EncodedLen(proof_buf.Len()))
	base64.StdEncoding.Encode(proof_buf_base64, proof_buf.Bytes())
	os.WriteFile("proof.bin", proof_buf_base64, 0644)

	// Serialize PublicWitness
	pw_buf, _ := witnessPublic.MarshalBinary()
	pw_buf_base64 := make([]byte, base64.StdEncoding.EncodedLen(len(pw_buf)))
	base64.StdEncoding.Encode(pw_buf_base64, pw_buf)
	os.WriteFile("pw.bin", pw_buf_base64, 0644)

	// Deserialize Proof
	proof_des := plonk.NewProof(ecc.BN254)
	proof_base64, _ := os.ReadFile("proof.bin")
	proof_bytes := make([]byte, base64.StdEncoding.DecodedLen(len(proof_base64)))
	base64.StdEncoding.Decode(proof_bytes, proof_base64)
	proof_reader := bytes.NewReader(proof_bytes)
	proof_des.ReadFrom(proof_reader)

	// Deserialize VK
	vk_des := plonk.NewVerifyingKey(ecc.BN254)
	vk_base64, _ := os.ReadFile("vk.bin")
	vk_bytes := make([]byte, base64.StdEncoding.DecodedLen(len(vk_base64)))
	base64.StdEncoding.Decode(vk_bytes, vk_base64)
	vk_reader := bytes.NewReader(vk_bytes)
	vk_des.ReadFrom(vk_reader)

	// Deserialize PublicWitness
	pw_des, err := witness.New(ecc.BN254.ScalarField())
	if err != nil {
		panic(err)
	}
	pw_base64, _ := os.ReadFile("pw.bin")
	pw_bytes := make([]byte, base64.StdEncoding.DecodedLen(len(pw_base64)))
	base64.StdEncoding.Decode(pw_bytes, pw_base64)
	pw_des.UnmarshalBinary(pw_bytes)

	err = plonk.Verify(proof_des, vk_des, pw_des)

	if err != nil {
		fmt.Print("Invalid Proof")
		return
	}
	fmt.Print("ValidProof")

}

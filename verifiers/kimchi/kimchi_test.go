package kimchi

import (
	"fmt"
	"os"
	"testing"
)

func TestEcAddKimchiProofVerifies(t *testing.T) {
	fmt.Println(os.Getwd())
	proofFile, err := os.Open("../../prover_examples/kimchi/example/kimchi_ec_add.proof")
	if err != nil {
		t.Errorf("could not open kimchi proof file")
	}

	proofBuffer := make([]byte, MAX_PROOF_SIZE)
	proofLen, err := proofFile.Read(proofBuffer)
	if err != nil {
		t.Errorf("could not read bytes from kimchi proof file")
	}

	pubInputFile, err := os.Open("../../prover_examples/kimchi/example/kimchi_verifier_index_example.bin")
	if err != nil {
		t.Errorf("could not open kimchi aggregated public input file")
	}
	pubInputBuffer := make([]byte, MAX_PUB_INPUT_SIZE)
	pubInputLen, err := pubInputFile.Read(pubInputBuffer)
	if err != nil {
		t.Errorf("could not read bytes from kimchi aggregated public input")
	}

	if !VerifyKimchiProof(([MAX_PROOF_SIZE]byte)(proofBuffer), uint(proofLen), ([MAX_PUB_INPUT_SIZE]byte)(pubInputBuffer), uint(pubInputLen)) {
		t.Errorf("proof did not verify")
	}
}

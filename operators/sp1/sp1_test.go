package sp1

import (
	"fmt"
	"os"
	"testing"
)

func TestFibonacciSp1ProofVerifies(t *testing.T) {
	fmt.Println(os.Getwd())
	f, err := os.Open("../../tests/testing_data/sp1_fibonacci.proof")
	if err != nil {
		t.Errorf("could not open proof file")
	}

	proofBytes := make([]byte, MAX_PROOF_SIZE)
	nReadBytes, err := f.Read(proofBytes)
	if err != nil {
		t.Errorf("could not read bytes from file")
	}

	if !VerifySp1Proof(([MAX_PROOF_SIZE]byte)(proofBytes), uint(nReadBytes)) {
		t.Errorf("proof did not verify")
	}
}

package cairo_platinum

import (
	"fmt"
	"os"
	"testing"
)

func TestFibonacci5ProofVerifies(t *testing.T) {
	fmt.Println(os.Getwd())
	f, err := os.Open("../../prover_examples/cairo_platinum/example/fibonacci_5.proof")
	if err != nil {
		t.Errorf("could not open proof file")
	}

	proofBytes := make([]byte, MAX_PROOF_SIZE)
	nReadBytes, err := f.Read(proofBytes)
	if err != nil {
		t.Errorf("could not read bytes from file")
	}

	if !VerifyCairoProof100Bits(([MAX_PROOF_SIZE]byte)(proofBytes), uint(nReadBytes)) {
		t.Errorf("proof did not verify")
	}
}

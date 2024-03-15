package cairo_platinum

import (
	"encoding/base64"
	"fmt"
	"os"
	"testing"
)

func TestFibonacci5ProofVerifies(t *testing.T) {
	fmt.Println(os.Getwd())
	f, err := os.Open("../../tests/testing_data/fibonacci_10.b64")
	if err != nil {
		t.Errorf("could not open proof file")
	}

	proofBytes := make([]byte, MAX_PROOF_SIZE)
	nReadBytes, err := f.Read(proofBytes)
	if err != nil {
		t.Errorf("could not read bytes from file")
	}

	// Decode base64 string back to bytes
	fmt.Printf(string(proofBytes[379056]))
	decodedBytes, err := base64.StdEncoding.DecodeString(string(proofBytes))
	if err != nil {
		fmt.Println("Error decoding base64 string:", err)
		return
	}

	if !VerifyCairoProof100Bits(([MAX_PROOF_SIZE]byte)(decodedBytes), uint(nReadBytes)) {
		t.Errorf("proof did not verify")
	}
}

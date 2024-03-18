package cairo_platinum

import (
	"encoding/base64"
	"fmt"
	"log"
	"os"
	"testing"
)

func TestFibonacci5ProofVerifies(t *testing.T) {
	fmt.Println(os.Getwd())
	f, err := os.Open("../../tests/testing_data/fibo_5.base64")
	if err != nil {
		t.Errorf("could not open proof file")
	}

	proofBytes := make([]byte, MAX_PROOF_SIZE)
	nReadBytes, err := f.Read(proofBytes)
	if err != nil {
		t.Errorf("could not read bytes from file")
	}

	// Pass contents of proof to string
	// to simulate the blockchain's workflow
	proofString := string(proofBytes[:nReadBytes])

	// Decode base64 string back to bytes
	decodedBytes := make([]byte, MAX_PROOF_SIZE)
	nDecoded, err := base64.StdEncoding.Decode(decodedBytes, []byte(proofString))
	if err != nil {
		log.Fatalf("could not decode base64 string: %s\n", err)
	}

	if !VerifyCairoProof100Bits(([MAX_PROOF_SIZE]byte)(decodedBytes), uint(nDecoded)) {
		t.Errorf("proof did not verify")
	}
}

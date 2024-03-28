package sp1_test

import (
	"alignedlayer/verifiers/sp1"
	"fmt"
	"os"
	"testing"
)

func TestFibonacciSp1ProofVerifies(t *testing.T) {
	fmt.Println(os.Getwd())

	proof, err := os.ReadFile("lib/src/fibonacci.proof")
	if err != nil {
		t.Errorf("could not open proof file")
	}

	elf, err := os.ReadFile("lib/src/fibonacci.elf")
	if err != nil {
		t.Errorf("could not open elf file")
	}

	if !sp1.VerifySp1ProofElf(proof, elf) {
		t.Errorf("proof did not verify")
	}
}

package sp1_test

import (
	"alignedlayer/verifiers/sp1"
	"fmt"
	"os"
	"testing"
)

func TestFibonacciSp1ProofVerifies(t *testing.T) {
	fmt.Println(os.Getwd())

	proofFile, err := os.Open("../../prover_examples/sp1/example/fibonacci.proof")
	if err != nil {
		t.Errorf("could not open proof file")
	}

	proof := make([]byte, sp1.MAX_PROOF_SIZE)
	proofSize, err := proofFile.Read(proof)
	if err != nil {
		t.Errorf("could not read bytes from file")
	}

	elfFile, err := os.Open("lib/elf/riscv32im-succinct-zkvm-elf")
	if err != nil {
		t.Errorf("could not open elf file")
	}

	elf := make([]byte, sp1.MAX_ELF_SIZE)
	elfSize, err := elfFile.Read(elf)
	if err != nil {
		t.Errorf("could not read bytes from file")
	}

	if !sp1.VerifySp1ProofElf(([sp1.MAX_PROOF_SIZE]byte)(proof), ([sp1.MAX_ELF_SIZE]byte)(elf), uint(proofSize), uint(elfSize)) {
		t.Errorf("proof did not verify")
	}
}

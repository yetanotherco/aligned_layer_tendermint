package sp1

/*
#cgo darwin LDFLAGS: -L./lib -lsp1_verifier
#cgo linux LDFLAGS: ${SRCDIR}/lib/libsp1_verifier.a -ldl -lrt -lm -Wl,--allow-multiple-definition

#include "lib/sp1.h"
*/
import "C"
import (
	"unsafe"
)

func VerifySp1Proof(proofBuffer [MAX_PROOF_SIZE]byte, proofLen uint) bool {
	proofPtr := (*C.uchar)(unsafe.Pointer(&proofBuffer[0]))
	return (bool)(C.verify_sp1_proof_ffi(proofPtr, (C.uint)(proofLen)))
}
const (
	KB             = 1024
	MB             = KB * 1024
	MAX_PROOF_SIZE = MB + 512*KB
	MAX_ELF_SIZE   = 512 * KB
)

func VerifySp1ProofElf(proofBuffer [MAX_PROOF_SIZE]byte, elfBuffer [MAX_ELF_SIZE]byte, proofLen uint, elfLen uint) bool {
	proofPtr := (*C.uchar)(unsafe.Pointer(&proofBuffer[0]))
	elfPtr := (*C.uchar)(unsafe.Pointer(&elfBuffer[0]))
	return (bool)(C.verify_sp1_proof_with_elf_ffi(proofPtr, elfPtr, (C.uint)(proofLen), (C.uint)(elfLen)))
}

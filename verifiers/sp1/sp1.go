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

func VerifySp1ProofElf(proof []byte, elf []byte) bool {
	proofPtr := (*C.uchar)(unsafe.Pointer(&proof[0]))
	elfPtr := (*C.uchar)(unsafe.Pointer(&elf[0]))

	return (bool)(C.verify_sp1_proof_with_elf_ffi(proofPtr, elfPtr, (C.uint)(len(proof)), (C.uint)(len(elf))))
}

use sp1_core::SP1Verifier;

#[no_mangle]
pub extern "C" fn verify_sp1_proof_with_elf_ffi(
    proof_ptr: *const u8,
    elf_ptr: *const u8,
    proof_len: usize,
    elf_len: usize,
) -> bool {
    let proof_buffer = unsafe { std::slice::from_raw_parts(proof_ptr, proof_len) };
    let elf_buffer = unsafe { std::slice::from_raw_parts(elf_ptr, elf_len) };

    verify_sp1_proof_with_elf(proof_buffer, elf_buffer)
}

fn verify_sp1_proof_with_elf(proof_buffer: &[u8], elf_buffer: &[u8]) -> bool {
    let Ok(proof) = bincode::deserialize(proof_buffer) else {
        return false;
    };

    SP1Verifier::verify(elf_buffer, &proof).is_ok()
}

#[cfg(test)]
mod tests {
    use super::*;

    const PROOF: &[u8] = include_bytes!("../../../../prover_examples/sp1/example/fibonacci.proof");
    const ELF: &[u8] = include_bytes!("../elf/riscv32im-succinct-zkvm-elf");

    #[test]
    fn can_verify_good_proof() {
        let result = verify_sp1_proof_with_elf(PROOF, ELF);
        assert!(result)
    }

    #[test]
    #[should_panic]
    fn wont_verify_bad_proof() {
        let mut bad_proof = ELF.to_vec();
        bad_proof.push(3);

        let result = verify_sp1_proof_with_elf(&bad_proof, ELF);
        assert!(result)
    }

    #[test]
    fn can_verify_with_ffi() {
        let proof_ptr = PROOF.as_ptr();
        let elf_ptr = ELF.as_ptr();
        let proof_size = PROOF.len();
        let elf_size = ELF.len();

        let result = verify_sp1_proof_with_elf_ffi(proof_ptr, elf_ptr, proof_size, elf_size);
        assert!(result)
    }
}

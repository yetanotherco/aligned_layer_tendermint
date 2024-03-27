use sp1_core::SP1Verifier;

const ELF: &[u8] = include_bytes!("../elf/riscv32im-succinct-zkvm-elf");
pub const MAX_PROOF_SIZE: usize = 1024 * 1024;
pub const MAX_ELF_BUFFER_SIZE: usize = 1024 * 1024;

#[no_mangle]
pub extern "C" fn verify_sp1_proof_with_elf_ffi(
    proof_bytes: &[u8; MAX_PROOF_SIZE],
    elf_bytes: &[u8; MAX_ELF_BUFFER_SIZE],
    proof_len: usize,
    elf_len: usize,
) -> bool {
    let real_elf = &elf_bytes[0..elf_len];

    if let Ok(proof) = bincode::deserialize(&proof_bytes[..proof_len]) {
        return SP1Verifier::verify(real_elf, &proof).is_ok();
    }

    false
}

#[no_mangle]
pub extern "C" fn verify_sp1_proof_ffi(
    proof_bytes: &[u8; MAX_PROOF_SIZE],
    proof_len: usize,
) -> bool {
    if let Ok(proof) = bincode::deserialize(&proof_bytes[..proof_len]) {
        return SP1Verifier::verify(ELF, &proof).is_ok();
    }

    false
}

#[cfg(test)]
mod tests {
    use super::*;

    const PROOF: &[u8; 1040380] =
        include_bytes!("../../../../prover_examples/sp1/example/fibonacci.proof");

    #[test]
    fn verify_sp1_proof_works() {
        let mut proof_buffer = [0u8; super::MAX_PROOF_SIZE];
        let proof_size = PROOF.len();
        proof_buffer[..proof_size].clone_from_slice(PROOF);
        let result = verify_sp1_proof_ffi(&proof_buffer, proof_size);
        assert!(result)
    }

    #[test]
    fn verify_sp1_aborts_with_bad_proof() {
        let mut proof_buffer = [42u8; super::MAX_PROOF_SIZE];
        let proof_size = PROOF.len();
        proof_buffer[..proof_size].clone_from_slice(PROOF);
        let result = verify_sp1_proof_ffi(&proof_buffer, proof_size - 1);
        assert!(!result)
    }
}

# !/bin/sh
#
# This script send two dummy transactions from an <account> with a test keyring. It should be run from the repository root.

if [ $# -ne 1 ]; then
  echo "Usage: $0 <account>"
  echo "accepts 1 arg(s), received $#"
  exit 1
else
  ACCOUNT=$1
fi

CHAIN_ID=alignedlayer

alignedlayerd tx verification verify \
  $(cat ./prover_examples/gnark_plonk/example/proof.base64.example) \
  $(cat ./prover_examples/gnark_plonk/example/public_inputs.base64.example) \
  $(cat ./prover_examples/gnark_plonk/example/verifying_key.base64.example) \
  --keyring-backend test \
  --from $ACCOUNT \
  --chain-id $CHAIN_ID \
  --fees 20stake \
  --yes

sleep 6

alignedlayerd tx verification verify \
  $(cat ./prover_examples/gnark_plonk/example/bad_proof.base64.example) \
  $(cat ./prover_examples/gnark_plonk/example/public_inputs.base64.example) \
  $(cat ./prover_examples/gnark_plonk/example/verifying_key.base64.example) \
  --keyring-backend test \
  --from $ACCOUNT \
  --chain-id $CHAIN_ID \
  --fees 20stake \
  --yes

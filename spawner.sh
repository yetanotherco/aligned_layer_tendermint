# !/bin/sh
#
# This script send two dummy transactions from an <account> with a test keyring. It should be run from the repository root.

if [ $# -ne 1 ]; then
	echo "Usage: $0 <account>"
  echo "accepts 1 arg(s), received $#"
	exit 1
else
	ACCOUNT=$1
  echo $ACCOUNT
fi

CHAIN_ID=alignedlayer
FEES=50000stake

alignedlayerd tx verification verify \
  $(cat ./prover_examples/gnark_plonk/example/proof.base64.example) \
  $(cat ./prover_examples/gnark_plonk/example/public_inputs.base64.example) \
  $(cat ./prover_examples/gnark_plonk/example/verifying_key.base64.example) \
  --keyring-backend test \
  --from $ACCOUNT \
  --chain-id $CHAIN_ID \
  --fees $FEES \
  --yes

sleep 1

alignedlayerd tx verification verify \
  $(cat ./prover_examples/gnark_plonk/example/proof_2.base64.example) \
  $(cat ./prover_examples/gnark_plonk/example/public_inputs_2.base64.example) \
  $(cat ./prover_examples/gnark_plonk/example/verifying_key_2.base64.example) \
  --keyring-backend test \
  --from $ACCOUNT \
  --chain-id $CHAIN_ID \
  --fees $FEES \
  --yes

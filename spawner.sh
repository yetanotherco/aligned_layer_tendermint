#!/usr/bin/bash
#
# This script sends dummy transactions from an <account> with a test keyring. It should be run from the repository root.
#
# To create the account, run: alignedlayerd keys add <account> --keyring-backend test

if [ $# -ne 1 ]; then
  echo "Usage: $0 <account>"
  echo "accepts 1 arg, received $#"
  exit 1
else
  ACCOUNT=$1
fi

CHAIN_ID=alignedlayer

# New elements can be added to the array to send more transactions
PROOFS=(
  "$(cat ./prover_examples/gnark_plonk/example/proof.base64.example)
  $(cat ./prover_examples/gnark_plonk/example/public_inputs.base64.example)
  $(cat ./prover_examples/gnark_plonk/example/verifying_key.base64.example)" 

  "$(cat ./prover_examples/gnark_plonk/example/bad_proof.base64.example)
  $(cat ./prover_examples/gnark_plonk/example/public_inputs.base64.example)
  $(cat ./prover_examples/gnark_plonk/example/verifying_key.base64.example)" 
)

ADDRESS=$(
  alignedlayerd keys show $ACCOUNT \
    --keyring-backend test --output json \
    | jq .address | tr -d \"
)
ACCOUNT_INFO=$(alignedlayerd q auth account $ADDRESS --output json)
NUMBER=$(echo $ACCOUNT_INFO | jq .account.value.account_number | tr -d \")
sequence=$(echo $ACCOUNT_INFO | jq .account.value.sequence | tr -d \")

for ((i = 0; i < ${#PROOFS[@]}; i++))
do
  proof=${PROOFS[$i]}
  echo $proof | xargs alignedlayerd tx verification verify \
    --keyring-backend test --from $ACCOUNT \
    --chain-id $CHAIN_ID \
    --fees 20stake \
    --offline \
    --sequence $sequence \
    --account-number $NUMBER \
    --yes
  let sequence++
done

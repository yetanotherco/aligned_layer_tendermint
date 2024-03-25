#!/bin/bash

set -e

if [ $# -ne 3 ]; then
  echo "Usage: $0 <account-name> <proof-file>"
  echo "accepts 2 arg, received $#"
  exit 1
else
  PROVER=$1
  ACCOUNT=$2
  PROOF_FILE=$3
fi


if [ $1 = "cairo" ]; then 
  PROVER="verify-cairo"
elif [ $1 =  "sp1" ]; then 
  PROVER="verify-sp-1"
else 
  echo "Usage: $0 <prover> <proof-file>"
  echo "Provers accepted: cairo & sp1"
  exit 1
fi

CHAIN_ID=alignedlayer

: ${NODE:="tcp://localhost:26657"}
: ${FEES:=20stake}
: ${GAS:=5000000}

NEW_PROOF_FILE=$(mktemp)
base64 -i $PROOF_FILE | tr -d '\n' > $NEW_PROOF_FILE

TRANSACTION=$(mktemp)
alignedlayerd tx verification ${PROVER} "PLACEHOLDER" \
  --from $ACCOUNT --chain-id $CHAIN_ID --generate-only \
  --gas $GAS --fees $FEES \
  | jq '.body.messages[0].proof=$proof' --rawfile proof $NEW_PROOF_FILE \
  > $TRANSACTION

SIGNED=$(mktemp)
alignedlayerd tx sign $TRANSACTION \
  --from $ACCOUNT --node $NODE \
  > $SIGNED

alignedlayerd tx broadcast $SIGNED --node $NODE

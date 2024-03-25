#!/bin/bash

set -e

if [ $# -ne 3 ]; then
  echo "Usage: $0 <verifier> <account-name> <proof-file>"
  echo "accepts 3 arg, received $#"
  exit 1
else
  VERIFIER=$1
  ACCOUNT=$2
  PROOF_FILE=$3
fi

CHAIN_ID=alignedlayer

: ${NODE:="tcp://localhost:26657"}
: ${FEES:=20stake}
: ${GAS:=5000000}

NEW_PROOF_FILE=$(mktemp)
base64 -i $PROOF_FILE | tr -d '\n' > $NEW_PROOF_FILE

TRANSACTION=$(mktemp)
alignedlayerd tx verify $VERIFIER "PLACEHOLDER" \
  --from $ACCOUNT --chain-id $CHAIN_ID --generate-only \
  --gas $GAS --fees $FEES \
  | jq '.body.messages[0].proof=$proof' --rawfile proof $NEW_PROOF_FILE \
  > $TRANSACTION

SIGNED=$(mktemp)
alignedlayerd tx sign $TRANSACTION \
  --from $ACCOUNT --node $NODE \
  > $SIGNED

alignedlayerd tx broadcast $SIGNED --node $NODE

#!/bin/bash

set -e

if [ $# -ne 2 ]; then
  echo "Usage: $0 <account-name> <proof-file>"
  echo "accepts 2 arg, received $#"
  exit 1
else
  ACCOUNT=$1
  PROOF_FILE=$2
fi

CHAIN_ID=alignedlayer

: ${NODE:="tcp://localhost:26657"}
: ${FEES:=20stake}
: ${GAS:=5000000}

NEW_PROOF_FILE=$(mktemp)
base64 -i $PROOF_FILE | tr -d '\n' > $NEW_PROOF_FILE

TRANSACTION=$(mktemp)
alignedlayerd tx verify cairo-platinum "PLACEHOLDER" \
  --from $ACCOUNT --chain-id $CHAIN_ID --generate-only \
  --gas $GAS --fees $FEES \
  | jq '.body.messages[0].proof=$proof' --rawfile proof $NEW_PROOF_FILE \
  > $TRANSACTION

SIGNED=$(mktemp)
alignedlayerd tx sign $TRANSACTION \
  --from $ACCOUNT --node $NODE \
  > $SIGNED

alignedlayerd tx broadcast $SIGNED --node $NODE

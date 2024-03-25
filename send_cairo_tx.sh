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

NEW_PROOF_FILE=$(mktemp)
base64 -i $PROOF_FILE | tr -d '\n' > $NEW_PROOF_FILE

TRANSACTION=$(mktemp)
alignedlayerd tx verification verify-cairo "PLACEHOLDER" \
  --from $FROM --chain-id $CHAIN_ID --generate-only --fees $FEES \
  | jq '.body.messages[0].proof=$proof' --rawfile proof $NEW_PROOF_FILE \
  > $TRANSACTION

SIGNED=$(mktemp)
alignedlayerd tx sign $TRANSACTION \
  --from $FROM --node $NODE \
  > $SIGNED

alignedlayerd tx broadcast $SIGNED --node $NODE

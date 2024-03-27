#!/bin/bash

set -e

if [ $# -ne 3 ]; then
  echo "Usage: $0 <account-name> <proof-file> <elf-file>"
  echo "accepts 3 arg, received $#"
  exit 1
else
  ACCOUNT=$1
  PROOF_FILE=$2
  ELF_FILE=$3
fi

CHAIN_ID=alignedlayer

: ${NODE:="tcp://localhost:26657"}
: ${FEES:=2000stake}
: ${GAS:=20000000}

NEW_PROOF_FILE=$(mktemp)
base64 -i $PROOF_FILE | tr -d '\n' > $NEW_PROOF_FILE

NEW_ELF_FILE=$(mktemp)
base64 -i $ELF_FILE | tr -d '\n' > $NEW_ELF_FILE

TRANSACTION=$(mktemp)
alignedlayerd tx verify sp-1 "PLACEHOLDER" "PLACEHOLDER" \
  --from $ACCOUNT --chain-id $CHAIN_ID --generate-only \
  --gas $GAS --fees $FEES \
  | jq '.body.messages[0].proof=$proof' --rawfile proof $NEW_PROOF_FILE \
  | jq '.body.messages[0].elf=$elf' --rawfile elf $NEW_ELF_FILE \
  > $TRANSACTION

SIGNED=$(mktemp)
alignedlayerd tx sign $TRANSACTION \
  --from $ACCOUNT --node $NODE \
  > $SIGNED

alignedlayerd tx broadcast $SIGNED --node $NODE

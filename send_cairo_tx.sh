#!/bin/bash

set -e

if [ $# -ne 1 ]; then
  echo "Usage: $0 <proof-file>"
  echo "accepts 1 arg, received $#"
  exit 1
else
  PROOF_FILE=$1
fi

CHAIN_ID=alignedlayer
METHOD=cosmos.tx.v1beta1.Service/BroadcastTx

: ${FROM:="alice"}
: ${NODE:="tcp://localhost:26657"}
: ${NODE_RPC:="localhost:9090"}
: ${GAS:=5000000}


TRIMMED_PROOF_FILE=$(mktemp)
cat $PROOF_FILE | tr -d '\n' > $TRIMMED_PROOF_FILE

TRANSACTION=$(mktemp)
alignedlayerd tx verification verify-cairo "PLACEHOLDER" \
  --from $FROM --chain-id $CHAIN_ID --generate-only --gas $GAS \
  | jq '.body.messages[0].proof=$proof' --rawfile proof $TRIMMED_PROOF_FILE \
  > $TRANSACTION

SIGNED=$(mktemp)
alignedlayerd tx sign $TRANSACTION \
  --from $FROM --node $NODE \
  > $SIGNED

alignedlayerd tx broadcast $SIGNED --node $NODE

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

TRANSACTION=$(mktemp)
TRIMMED_PROOF_FILE=$(mktemp)

cat $PROOF_FILE | tr -d '\n' > $TRIMMED_PROOF_FILE

alignedlayerd tx verification verify-cairo "PLACEHOLDER" \
  --from $FROM --chain-id $CHAIN_ID --generate-only --gas 5000000\
  | jq '.body.messages.[0].proof=$proof' --rawfile proof $TRIMMED_PROOF_FILE \
  > $TRANSACTION

SIGNED=$(mktemp)
alignedlayerd tx sign $TRANSACTION --chain-id $CHAIN_ID --from $FROM --node $NODE \
  --overwrite > $SIGNED

ENCODED=$(mktemp)
alignedlayerd tx encode $SIGNED > $ENCODED

echo '{"tx_bytes": "", "mode": "BROADCAST_MODE_SYNC"}' \
  | jq '.tx_bytes=$bytes' --rawfile bytes $ENCODED \
  | grpcurl -plaintext -d "@" $NODE_RPC $METHOD \
  | jq .txResponse

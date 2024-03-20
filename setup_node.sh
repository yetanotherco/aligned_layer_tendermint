#!/bin/bash
set -e

if [ $# -lt 1 ]; then
	echo "Usage: $0 <moniker>"
	exit 1
else
    MONIKER=$1
fi

NODE_HOME=$HOME/.alignedlayer
CHAIN_BINARY=alignedlayerd
CHAIN_ID=alignedlayer

PEER_ADDRESSES=("91.107.239.79" "116.203.81.174" "88.99.174.203" "128.140.3.188")
: ${MINIMUM_GAS_PRICES="0.0001stake"}

ignite chain build

$CHAIN_BINARY comet unsafe-reset-all
$CHAIN_BINARY init $MONIKER \
    --chain-id $CHAIN_ID --overwrite

for ADDR in "${PEER_ADDRESSES[@]}"; do
    GENESIS=$(curl -f "$ADDR:26657/genesis" | jq '.result.genesis')
    if [ -n "$GENESIS" ]; then
        echo "$GENESIS" > $NODE_HOME/config/genesis.json;
        break;
    fi
done

PEERS=()

for ADDR in "${PEER_ADDRESSES[@]}"; do
    PEER_ID=$(curl -s "$ADDR:26657/status" | jq -r '.result.node_info.id')
    if [ -n "$PEER_ID" ]; then
        PEERS+=("$PEER_ID@$ADDR:26656")
    fi
done

PEER_LIST=$(IFS=,; echo "${PEERS[*]}")

$CHAIN_BINARY config set config p2p.persistent_peers "$PEER_LIST" --skip-validate
$CHAIN_BINARY config set app minimum-gas-prices "$MINIMUM_GAS_PRICES" --skip-validate

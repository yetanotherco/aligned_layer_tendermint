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

: ${PEER_ADDR:="91.107.239.79"}
: ${MINIMUM_GAS_PRICES="0.25stake"}

ignite chain build

$CHAIN_BINARY comet unsafe-reset-all
$CHAIN_BINARY init $MONIKER \
    --chain-id $CHAIN_ID --overwrite

curl $PEER_ADDR:26657/genesis \
    | jq '.result.genesis' \
    > $NODE_HOME/config/genesis.json

PEER_IR=$(
    curl -s $PEER_ADDR:26657/status \
    | jq -r '.result.node_info.id'
)

$CHAIN_BINARY config set config p2p.seeds "$PEER_IR@$PEER_ADDR:26656" \
    --skip-validate
$CHAIN_BINARY config set config p2p.persistent_peers "$PEER_IR@$PEER_ADDR:26656" \
    --skip-validate
$CHAIN_BINARY config set app minimum-gas-prices $MINIMUM_GAS_PRICES \
    --skip-validate

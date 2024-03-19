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

: ${PEER_ADDR1:="91.107.239.79"}
: ${PEER_ADDR2:="116.203.81.174"}
: ${PEER_ADDR3:="88.99.174.203"}
: ${PEER_ADDR4:="128.140.3.188"}
: ${MINIMUM_GAS_PRICES="0.25stake"}

ignite chain build

$CHAIN_BINARY comet unsafe-reset-all
$CHAIN_BINARY init $MONIKER \
    --chain-id $CHAIN_ID --overwrite

curl $PEER_ADDR1:26657/genesis \
    | jq '.result.genesis' \
    > $NODE_HOME/config/genesis.json

PEER_IR1=$(
    curl -s $PEER_ADDR1:26657/status \
    | jq -r '.result.node_info.id'
)

PEER_IR2=$(
    curl -s $PEER_ADDR2:26657/status \
    | jq -r '.result.node_info.id'
)

PEER_IR3=$(
    curl -s $PEER_ADDR3:26657/status \
    | jq -r '.result.node_info.id'
)

PEER_IR4=$(
    curl -s $PEER_ADDR4:26657/status \
    | jq -r '.result.node_info.id'
)

$CHAIN_BINARY config set config p2p.persistent_peers "$PEER_IR1@$PEER_ADDR1:26656,$PEER_IR2@$PEER_ADDR2:26656,$PEER_IR3@$PEER_ADDR3:26656,$PEER_IR4@$PEER_ADDR4:26656" \
    --skip-validate
$CHAIN_BINARY config set app minimum-gas-prices $MINIMUM_GAS_PRICES \
    --skip-validate

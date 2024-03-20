#!/bin/bash
set -e

if [ $# -lt 1 ]; then
	echo "Usage: $0 <account> <staking_amount>"
	exit 1
else
	VALIDATOR=$1
	STAKING_AMOUNT=$2
fi

NODE_HOME=$HOME/.alignedlayer
CHAIN_BINARY=alignedlayerd
CHAIN_ID=alignedlayer

: ${FEES:="50stake"}
: ${PEER_ADDR:="91.107.239.79"}

VALIDATOR_KEY=$($CHAIN_BINARY tendermint show-validator)
MONIKER=$($CHAIN_BINARY config get config moniker)

cat << EOF > $NODE_HOME/config/validator.json
{
	"pubkey": $VALIDATOR_KEY,
	"amount": "$STAKING_AMOUNT",
	"moniker": $MONIKER,
	"commission-rate": "0.1",
	"commission-max-rate": "0.2",
	"commission-max-change-rate": "0.01",
	"min-self-delegation": "1"
}
EOF

$CHAIN_BINARY tx staking create-validator $NODE_HOME/config/validator.json \
	--from $VALIDATOR --chain-id $CHAIN_ID \
	--node tcp://$PEER_ADDR:26657 --fees $FEES

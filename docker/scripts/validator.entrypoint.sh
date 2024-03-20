#!/bin/bash
set -e

if [ $# -lt 2 ]; then
	echo "Usage: $0 <account> <staking_amount>"
	exit 1
else
	VALIDATOR=$1
	STAKING_AMOUNT=$2
fi

apt update && apt install -y jq

NODE_HOME=/root/.alignedlayer
CHAIN_BINARY=alignedlayerd
CHAIN_ID=alignedlayer
PEER_ADDRESSES=("91.107.239.79" "116.203.81.174" "88.99.174.203" "128.140.3.188")

: ${MINIMUM_GAS_PRICES="0.0001stake"}
: ${PASSWORD:="password"}
: ${FEES:="50stake"}
: ${PEER_ADDR:="91.107.239.79"}

(echo $PASSWORD; echo $PASSWORD) | alignedlayerd keys add $VALIDATOR
ADDRESS=$(echo $PASSWORD | alignedlayerd keys show $VALIDATOR --address)

until $(curl --output /dev/null --silent --head --fail https://faucet.alignedlayer.com/send/alignedlayer/$ADDRESS); do
    printf '.'
    sleep 1
done


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
	"min-self-delegation": "1",
    "website": "https://droak.sh",
    "security": "me+aligned@droak.sh"
}
EOF

(echo $PASSWORD) | $CHAIN_BINARY tx staking create-validator \
  $NODE_HOME/config/validator.json \
  --chain-id $CHAIN_ID \
  --fees $FEES \
  --from $VALIDATOR \
  --node tcp://$PEER_ADDR:26657 \
  --yes

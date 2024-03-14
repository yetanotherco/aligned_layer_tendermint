#!/bin/bash

: "${PASSWORD:=password}"
token="stake"
initial_balance=1000000000
initial_stake=60000000


if [ $# -lt 1 ]; then
    echo "Usage: $0 <node1> [<node2> ...]"
    exit 1
fi

echo "Creating directories for nodes..."
rm -rf prod-sim
for node in "$@"; do
    mkdir -p prod-sim/$node
done

node_ids=()

for node in "$@"; do
    echo "Initializing $node..."
    docker run -v $(pwd)/prod-sim/$node:/root/.alignedlayer -it alignedlayerd_i init alignedlayer_$node --chain-id alignedlayer > /dev/null
    
    docker run --rm -it -v $(pwd)/prod-sim/$node:/root/.alignedlayer --entrypoint sed alignedlayerd_i -i 's/"stake"/"'$token'"/g' /root/.alignedlayer/config/genesis.json 
    docker run -v $(pwd)/prod-sim/$node:/root/.alignedlayer -it alignedlayerd_i config set app minimum-gas-prices "0.1$token"
    docker run -v $(pwd)/prod-sim/$node:/root/.alignedlayer -it alignedlayerd_i config set app pruning "nothing" 


    node_id=$(docker run --rm -i -v $(pwd)/prod-sim/$node:/root/.alignedlayer alignedlayerd_i tendermint show-node-id)
    node_ids+=($node_id)

    echo "Node ID for $node: $node_id"
done


for (( i=1; i <= "$#"; i++ )); do
    echo "Creating key for ${!i} user..."
    printf "$PASSWORD\n$PASSWORD\n" | docker run --rm -i -v $(pwd)/prod-sim/${!i}:/root/.alignedlayer alignedlayerd_i keys --keyring-backend file --keyring-dir /root/.alignedlayer/keys add val_${!i} > /dev/null 2> ./prod-sim/${!i}/mnemonic.txt

    val_address=$(echo $PASSWORD | docker run --rm -i -v $(pwd)/prod-sim/${!i}:/root/.alignedlayer alignedlayerd_i keys --keyring-backend file --keyring-dir /root/.alignedlayer/keys show val_${!i} --address)
    echo "val_${!i} address: $val_address"
    echo "val_${!i} mnemonic: $(cat ./prod-sim/${!i}/mnemonic.txt)"

    echo "Giving val_${!i} some tokens..."
    docker run --rm -it -v $(pwd)/prod-sim/${!i}:/root/.alignedlayer alignedlayerd_i genesis add-genesis-account $val_address $initial_balance$token

    if [ $((i+1)) -le "$#" ]; then
        j=$((i+1))
        cp prod-sim/${!i}/config/genesis.json prod-sim/${!j}/config/genesis.json
    else
        cp prod-sim/${!i}/config/genesis.json prod-sim/$1/config/genesis.json
    fi      
done



for (( i=1; i <= "$#"; i++ )); do
    echo "Giving val_${!i} some stake..."
    echo $PASSWORD | docker run --rm -i -v $(pwd)/prod-sim/${!i}:/root/.alignedlayer alignedlayerd_i genesis gentx val_${!i} $initial_stake$token --keyring-backend file --keyring-dir /root/.alignedlayer/keys --account-number 0 --sequence 0 --chain-id alignedlayer --gas 1000000 --gas-prices 0.1$token

    if [ $i -gt 1 ]; then
        cp prod-sim/${!i}/config/gentx/* prod-sim/$1/config/gentx/
    fi

    if [ $((i+1)) -le "$#" ]; then
        j=$((i+1))
        cp prod-sim/${!i}/config/genesis.json prod-sim/${!j}/config/genesis.json
    else
        cp prod-sim/${!i}/config/genesis.json prod-sim/$1/config/genesis.json
    fi   
done

echo "Collecting genesis transactions..."
docker run --rm -it -v $(pwd)/prod-sim/$1:/root/.alignedlayer alignedlayerd_i genesis collect-gentxs > /dev/null

if ! docker run --rm -it -v $(pwd)/prod-sim/$1:/root/.alignedlayer alignedlayerd_i genesis validate-genesis; then
    echo "Invalid genesis"
    exit 1
fi

echo "Copying genesis file to other nodes..."
for node in "${@:2}"; do
    cp prod-sim/$1/config/genesis.json prod-sim/$node/config/genesis.json
done

echo "Setting node addresses in config..."
for (( i=1; i <= "$#"; i++ )); do
    other_addresses=()
    for (( j=1; j <= "$#"; j++ )); do
        if [ $j -ne $i ]; then
            other_addresses+=("${node_ids[$j - 1]}@${!j}:26656")
        fi
    done
    other_addresses=$(IFS=,; echo "${other_addresses[*]}")
    docker run -v $(pwd)/prod-sim/${!i}:/root/.alignedlayer -it alignedlayerd_i config set config p2p.seeds "$other_addresses" --skip-validate
done

docker run -v $(pwd)/prod-sim/$1:/root/.alignedlayer -it alignedlayerd_i config set config rpc.laddr "tcp://0.0.0.0:26657" --skip-validate

echo "Setting up faucet files..."
mkdir -p prod-sim/faucet/.faucet
mkdir -p prod-sim/faucet/config
cp faucet/config/config.js prod-sim/faucet/config/config.js
sed -n '6p' ./prod-sim/$1/mnemonic.txt | tr -d '\n' > temp.txt && mv temp.txt ./prod-sim/faucet/.faucet/mnemonic.txt
sed -i '' 's|\(rpc_endpoint: \).*"|\1"http://'$1':26657"|' prod-sim/faucet/config/config.js

echo "Setting up docker compose..."
rm -f ./prod-sim/docker-compose.yml
printf "version: '3.7'\nnetworks:\n  net-public:\nservices:\n" > ./prod-sim/docker-compose.yml
for node in "$@"; do
    printf "  alignedlayerd-$node:\n    command: start\n    image: alignedlayerd_i\n    container_name: $node\n    volumes:\n      - ./$node:/root/.alignedlayer\n    networks:\n      - net-public\n" >> ./prod-sim/docker-compose.yml
    if [ $node == "$1" ]; then
        printf "    ports:\n      - 0.0.0.0:26657:26657\n" >> ./prod-sim/docker-compose.yml
    fi
    printf "\n" >> ./prod-sim/docker-compose.yml
done
printf "  alignedlayerd-faucet:\n    command: faucet.js\n    image: alignedlayerd_faucet\n    container_name: faucet\n    volumes:\n      - ./faucet/config:/faucet/config\n      - ./faucet/.faucet:/faucet/.faucet\n    networks:\n      - net-public\n    ports:\n      - 8088:8088\n" >> ./prod-sim/docker-compose.yml

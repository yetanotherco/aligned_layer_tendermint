#!/bin/bash

if [ "$1" = "prod" ]; then
    nodes=("node0" "node1" "node2" "node3")
    nodes_ips=("10.0.0.2" "10.0.0.3" "10.0.0.4" "10.0.0.6")
    servers=("admin@blockchain-1" "admin@blockchain-2" "admin@blockchain-3" "admin@blockchain-4")

    read -p "Are you sure you want to deploy in production? (y/n): " answer
    if [ "$answer" != "y" ]; then
        exit 0
    fi
elif [ "$1" = "test" ]; then
    nodes=("node0" "node1" "node2")
    nodes_ips=("10.0.0.2" "10.0.0.3" "10.0.0.4")
    servers=("admin@testing-blockchain-1" "admin@testing-blockchain-2" "admin@testing-blockchain-3")
else
    echo "Usage: $0 [prod|test]"
    exit 1
fi

rm -rf server-setup

echo "Building binary..."

if [[ "$OSTYPE" == "linux-gnu"* ]]; then
    make build-cairo-ffi-linux
    make build-sp1-ffi-linux
elif [[ "$OSTYPE" == "darwin"* ]]; then
    make build-sp1-ffi-from_macos-to_linux
    make build-cairo-ffi-from_macos-to_linux
else
    echo "Unsupported OS"
    exit 0
fi

ignite chain build --release -t linux:amd64
## TODO change ignite with makefile
cd release
tar -xzf alignedlayer_linux_amd64.tar.gz
for server in "${servers[@]}"; do
    scp alignedlayerd $server:/usr/local/bin
done
cd ..

mkdir -p server-setup
cd server-setup

echo "Calling setup script..."
bash ../multi_node_setup.sh "${nodes[@]}"

echo "Setting node addresses in config..."
for i in "${!nodes[@]}"; do 
    echo $(pwd)
    seeds=$(docker run -v "$(pwd)/prod-sim/${nodes[$i]}:/root/.alignedlayer" -it alignedlayerd_i config get config p2p.persistent_peers)
    for j in "${!nodes[@]}"; do  
        seeds=${seeds//${nodes[$j]}/${nodes_ips[$j]}}
    done
    
    docker run -v "$(pwd)/prod-sim/${nodes[$i]}:/root/.alignedlayer" -it alignedlayerd_i config set config p2p.persistent_peers $seeds --skip-validate  
done

echo "Configuring transaction sizes"

MAX_BODY_BYTES=20971520
MAX_TX_BYTES=20971520
MAX_TXS_BYTES=25165824

for i in "${!nodes[@]}"; do
    docker run -v "$(pwd)/prod-sim/${nodes[$i]}:/root/.alignedlayer" -it alignedlayerd_i config set config rpc.max_body_bytes $MAX_BODY_BYTES --skip-validate
    docker run -v "$(pwd)/prod-sim/${nodes[$i]}:/root/.alignedlayer" -it alignedlayerd_i config set config rpc.max_tx_bytes $MAX_TX_BYTES --skip-validate  
    docker run -v "$(pwd)/prod-sim/${nodes[$i]}:/root/.alignedlayer" -it alignedlayerd_i config set config rpc.max_txs_bytes $MAX_TXS_BYTES --skip-validate 
done
 
echo "Sending directories to servers..."
for i in "${!servers[@]}"; do  
    ssh ${servers[$i]} "rm -rf /home/admin/.alignedlayer"
    scp -r "prod-sim/${nodes[$i]}" "${servers[$i]}:/home/admin/.alignedlayer"
done


ssh ${servers[0]} "rm -rf /home/admin/faucet/.faucet"
scp -p -r "prod-sim/faucet/.faucet" "${servers[0]}:/home/admin/faucet/.faucet"

cd ..

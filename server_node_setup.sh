#!/bin/bash

if [ $# -lt 2 ]; then
    echo "Usage: $0 [prod|test] binary_release_tag"
    exit 1
fi

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
    echo "Usage: $0 [prod|test] binary_release_tag"
    exit 1
fi

rm -rf server-setup

echo "Downloading binaries into servers..."
for server in "${servers[@]}"; do
    ssh $server "rm -rf /home/admin/alignedlayerd"
    ssh $server "curl -L --output alignedlayerd https://github.com/yetanotherco/aligned_layer_tendermint/releases/download/$2/alignedlayerd && chmod +x alignedlayerd"
done

mkdir -p server-setup
cd server-setup

export FAUCET_DIR="../faucet"
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

echo "Sending directories to servers..."
for i in "${!servers[@]}"; do  
    ssh ${servers[$i]} "rm -rf /home/admin/.alignedlayer"
    scp -r "prod-sim/${nodes[$i]}" "${servers[$i]}:/home/admin/.alignedlayer"
done


ssh ${servers[0]} "rm -rf /home/admin/aligned_layer_tendermint/faucet/.faucet"
scp -p -r "prod-sim/faucet/.faucet" "${servers[0]}:/home/admin/aligned_layer_tendermint/faucet/.faucet"

cd ..

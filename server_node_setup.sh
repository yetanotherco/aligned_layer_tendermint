#!/bin/bash

nodes=("node0" "node1" "node2")
nodes_ips=("10.0.0.2" "10.0.0.3" "10.0.0.4")
servers=("admin@blockchain-1" "admin@blockchain-2" "admin@blockchain-3")

rm -rf server-setup

echo "Building binary..."
ignite chain build --release -t linux:amd64

cd release
tar -xzf alignedlayer_linux_amd64.tar.gz
for server in "${servers[@]}"; do
    scp alignedlayerd $server:/home/admin
done
cd ..

mkdir -p server-setup
cd server-setup

echo "Calling setup script..."
bash ../multi_node_setup.sh ${nodes[@]}

echo "Setting node addresses in config..."
for (( i=0; i<3; i++ )); do
    echo $(pwd)
    seeds=$(docker run -v $(pwd)/prod-sim/${nodes[$i]}:/root/.alignedlayer -it alignedlayerd_i config get config p2p.seeds)
        for ((j=0; j<3; j++)); do
        seeds=${seeds//${nodes[$j]}/${nodes_ips[$j]}}
    done
    
    docker run -v $(pwd)/prod-sim/${nodes[$i]}:/root/.alignedlayer -it alignedlayerd_i config set config p2p.seeds $seeds --skip-validate    
done

echo "Sending directories to servers..."
for ((i=0; i<3; i++)); do
    ssh ${servers[i]} "rm -rf /home/admin/.alignedlayer"
    scp -r prod-sim/${nodes[i]} ${servers[i]}:/home/admin/.alignedlayer
done

ssh ${servers[0]} "rm -rf /home/admin/faucet/.faucet"
scp -p -r prod-sim/faucet/.faucet ${servers[0]}:/home/admin/faucet/.faucet

cd ..

#!/bin/bash

# Clone the GitHub repository
git clone -b block_monitor https://github.com/yetanotherco/aligned_layer_tendermint.git

# Move the desired folder to the current directory
mv aligned_layer_tendermint/monitor .

# Clean up (optional)
rm -rf aligned_layer_tendermint

cd monitor && make setup

sudo cp monitor/block_monitor.service /etc/systemd/system/monitor.service
sudo systemctl start monitor
sudo systemctl enable monitor

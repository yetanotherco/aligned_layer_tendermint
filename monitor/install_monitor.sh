#!/bin/bash

# Clone the GitHub repository
git clone -b block_monitor https://github.com/yetanotherco/aligned_layer_tendermint.git

# Move the desired folder to the current directory
mv aligned_layer_tendermint/monitor .

sudo echo "SLACK_URL=$SLACK_URL" > monitor/.env

# Clean up (optional)
rm -rf aligned_layer_tendermint

# Create python venv
cd monitor && make setup

# Setup systemd
sudo cp ~/monitor/block_monitor.service /etc/systemd/system/monitor.service
sudo systemctl daemon-reload
sudo systemctl start monitor
sudo systemctl enable monitor

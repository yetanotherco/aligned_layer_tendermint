#!/bin/bash

# Clone the GitHub repository
git clone https://github.com/yetanotherco/aligned_layer_tendermint.git

# Move the desired folder to the current directory
mv aligned_layer_tendermint/monitor .

# Clean up (optional)
rm -rf aligned_layer_tendermint

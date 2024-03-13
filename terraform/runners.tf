variable "seed_ip" {
  default = "91.107.239.79"
}

variable "staking_amount" {
  default = 50000
}

# Create a server
resource "hcloud_server" "alignedlayer-runner" {
  count = 1

  name        = "alignedlayer-${count.index}"
  image       = "debian-12"
  server_type = "cx11"
  public_net {
    ipv4_enabled = true
    ipv6_enabled = true
  }

  user_data = <<EOF
    package_update: true
    package_upgrade: true
    packages:
      - git
      - curl
      - jq
      - go
    runcmd:
      - curl https://get.ignite.com/cli! | bash
      - git clone https://github.com/yetanotherco/aligned_layer_tendermint.git
      - cd aligned_layer_tendermint
      - ignite chain build --output /usr/local/bin
      - alignedlayerd init "alignedlayer-${count.index}" --chain-id alignedlayer
      - curl -s ${var.seed_ip}:26657/genesis | jq '.result.genesis' > ~/.alignedlayer/config/genesis.json
      - curl -s ${var.seed_ip}:26657/status | jq '.result.node_info.id' > .seed_id
      - alignedlayerd config set config seeds "$(cat .seed_id)@${seed_ip}:26656" --skip-validate
      - alignedlayerd config set config persistent_peers "$(cat .seed_id)@${seed_ip}:26656" --skip-validate
      - alignedlayerd config set app minimum-gas-prices "0.0025stake"
      - alignedlayerd keys add user-${count.index}
      - # Here we need to get stake tokens
      - cat > validator.json <<EOL
        {
        	"pubkey": $(alignedlayerd tendermint show-validator),
        	"amount": "${var.staking_amount}stake",
        	"moniker": "alignedlayer-${count.index}",
        	"commission-rate": "0.1",
        	"commission-max-rate": "0.2",
        	"commission-max-change-rate": "0.01",
        	"min-self-delegation": "1"
        }
      - alignedlayerd tx staking create-validator validator.json --from alignedlayerd-${count.index} --node tcp://${seed_ip}:26656
      - alignedlayerd start
  EOF
}

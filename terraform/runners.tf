variable "seed_ip" {
  deafult = "91.107.239.79"
}

# Create a server
resource "hcloud_server" "alignedlayer-runner" {
  name        = "alignedlayer-${count.index}"
  image       = "debian-12"
  server_type = "cx11"
  public_net {
    ipv4_enabled = true
    ipv6_enabled = true
  }
  user_data = <<EOF
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
      - alignedlayerd config set seeds "$(cat .seed_id)@${seed_ip}:26656"
  EOF
}

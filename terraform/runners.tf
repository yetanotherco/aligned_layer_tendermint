variable "staking_amount" {
  default = 1000000000
}

variable "staking_token" {
  default = "stake"
}

variable "genesis_initial_balance" {
  default = 2000000000
}

variable "chain_id" {
  default = "alignedlayer-1"
}

variable "instances" {
  default = 1
}

variable "password" {
  default = "password"
}

resource "hcloud_network" "private_net" {
  name     = "alignedlayer-net"
  ip_range = "10.0.0.0/16"
}

resource "hcloud_network_subnet" "private_subnet" {
  type         = "cloud"
  network_id   = hcloud_network.private_net.id
  network_zone = "eu-central"
  ip_range     = "10.0.1.0/24"
}

resource "hcloud_server" "alignedlayer-genesis-runner" {
  name = "alignedlayer-genesis"
  image = "debian-12"
  server_type = "cx21"

  ssh_keys = ["manubilbao", "tomyrd"]

  public_net {
    ipv4_enabled = true
    ipv6_enabled = true
  }

  network {
    network_id = hcloud_network.private_net.id
    ip = "10.0.1.2"
  }

  depends_on = [
    hcloud_network_subnet.private_subnet
  ]

  user_data = <<EOF
    #cloud-config
    package_update: true
    package_upgrade: true
    packages:
      - curl
      - git
      - npm
    write_files:
      - path: /etc/systemd/system/faucet.service
        content: |
          [Unit]
          Description=Faucet
          After=network.target
          StartLimitIntervalSec=0

          [Service]
          Type=simple
          Restart=always
          RestartSec=1
          User=root
          WorkingDirectory=/root/aligned_layer_tendermint/faucet
          ExecStart=node faucet.js

          [Install]
          WantedBy=multi-user.target
      - path: /etc/systemd/system/alignedlayer.service
        content: |
          [Unit]
          Description=Aligned Layer
          After=network.target
          StartLimitIntervalSec=0

          [Service]
          Type=simple
          Restart=always
          RestartSec=1
          User=root
          ExecStart=alignedlayerd start

          [Install]
          WantedBy=multi-user.target
    runcmd:
      - curl -L -o /root/alignedlayer.tar.gz https://github.com/yetanotherco/aligned_layer_tendermint/releases/download/v0.1/alignedlayer_linux_amd64.tar.gz
      - tar -C /usr/local/bin -xzf /root/alignedlayer.tar.gz
      - git clone https://github.com/yetanotherco/aligned_layer_tendermint /root/aligned_layer_tendermint
      - cd /root/aligned_layer_tendermint/faucet
      - npm install
      - mkdir .faucet
      - export HOME=/root
      - alignedlayerd init victor-node --chain-id ${var.chain_id}
      - sed -i 's/"stake"/"${var.staking_token}"/g' /root/.alignedlayer/config/genesis.json
      - alignedlayerd config set app minimum-gas-prices 0.1${var.staking_token}
      - alignedlayerd config set app pruning "nothing"
      - printf "${var.password}\n${var.password}\n" | alignedlayerd keys add victor 2>&1 >/dev/null | tail -n1 > /root/aligned_layer_tendermint/faucet/.faucet/mnemonic.txt
      - export ADDRESS=$(printf "${var.password}\n" | alignedlayerd keys show victor --address)
      - alignedlayerd genesis add-genesis-account $ADDRESS ${var.genesis_initial_balance}${var.staking_token}
      - printf "${var.password}\n" | alignedlayerd genesis gentx victor ${var.staking_amount}${var.staking_token} --account-number 0 --sequence 0 --chain-id ${var.chain_id} --gas 1000000 --gas-prices 0.1${var.staking_token}
      - alignedlayerd genesis collect-gentxs
      - systemctl enable faucet
      - systemctl start faucet
      - systemctl enable alignedlayer
      - systemctl start alignedlayer
  EOF
}

# Create a server
resource "hcloud_server" "alignedlayer-runner" {
  count = var.instances - 1 # -1 because genesis runner is already a validator

  name        = "alignedlayer-${count.index}"
  image       = "debian-12"
  server_type = "cx21"
  public_net {
    ipv4_enabled = true
    ipv6_enabled = true
  }

  network {
    network_id = hcloud_network.private_net.id
    ip = "10.0.1.${count.index+3}"
  }

  ssh_keys = ["manubilbao", "tomyrd"]

  depends_on = [
    hcloud_server.alignedlayer-genesis-runner,
    hcloud_network_subnet.private_subnet
  ]

  user_data = <<EOF
    #cloud-config
    package_update: true
    package_upgrade: true
    packages:
      - curl
      - jq
    runcmd:
      - curl -L -o /root/alignedlayer.tar.gz https://github.com/yetanotherco/aligned_layer_tendermint/releases/download/v0.1/alignedlayer_linux_amd64.tar.gz
      - tar -C /usr/local/bin -xzf /root/alignedlayer.tar.gz
      - export HOME=/root
      - alignedlayerd init "node${count.index}" --chain-id alignedlayer
      - while [ ! "$(curl -s 10.0.1.2:26657/health)" ]; do sleep 1; done  # Wait until genesis node is ready
      - curl -s '10.0.1.2:26657/genesis' | jq '.result.genesis' > ~/.alignedlayer/config/genesis.json
      - curl -s '10.0.1.2:26657/status' | jq '.result.node_info.id' > .seed_id
      - alignedlayerd config set config p2p.seeds "$(cat .seed_id)@10.0.1.2:26656" --skip-validate
      - alignedlayerd config set config p2p.persistent_peers "$(cat .seed_id)@10.0.1.2:26656" --skip-validate
      - alignedlayerd config set app minimum-gas-prices "0.0025${var.staking_token}"
      - printf "${var.password}\n${var.password}\n" | alignedlayerd keys add node${count.index}
      - # Here we need to get stake tokens
      - cat > validator.json <<EOL
        {
        	"pubkey": $(alignedlayerd tendermint show-validator),
        	"amount": "${var.staking_amount}${var.staking_token}",
        	"moniker": "${random_string.random.result}",
        	"commission-rate": "0.1",
        	"commission-max-rate": "0.2",
        	"commission-max-change-rate": "0.01",
        	"min-self-delegation": "1"
        }
      - alignedlayerd tx staking create-validator validator.json --from ${random_string.random.result} --node tcp://${var.seed_ip}:26656
      - alignedlayerd start
  EOF
}

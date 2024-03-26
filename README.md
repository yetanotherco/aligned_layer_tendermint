# Aligned Layer Blockchain

An application-specific blockchain built using [Cosmos SDK](https://docs.cosmos.network/) and created with [Ignite CLI](https://ignite.com/). The blockchain offers a variety of zkSNARK implementations to verify proofs sent over transactions, and stores their results.

Cosmos SDK provides a framework to build an application layer on top of a consensus layer interacting via ABCI (Application BlockChain Interface). By default, [CometBFT](https://cometbft.com/) (a fork of Tendermint) is used in the consensus and network layer.

Ignite CLI is used to generate boilerplate code for a Cosmos SDK application, making it easier to deploy a blockchain to production.

## Table of Contents

- [Aligned Layer Blockchain](#aligned-layer-blockchain)
  - [Table of Contents](#table-of-contents)
  - [Requirements](#requirements)
  - [Example Local Blockchain](#example-local-blockchain)
  - [Verifiers](#verifiers)
    - [Gnark Plonk](#gnark-plonk)
    - [Cairo Platinum](#cairo-platinum)
    - [Kimchi](#kimchi)
    - [Sp1](#sp1)
  - [Trying our testnet](#trying-our-testnet)
  - [Joining Our Testnet](#joining-our-testnet)
    - [Requirements](#requirements-1)
      - [Hardware](#hardware)
      - [Software](#software)
    - [With Docker](#with-docker)
    - [Node Setup](#node-setup)
      - [The fast way](#the-fast-way)
      - [Manual step by step](#manual-step-by-step)
    - [Creating an Account](#creating-an-account)
    - [Registering as a Validator](#registering-as-a-validator)
      - [The fast way](#the-fast-way-1)
      - [Manual step by step](#manual-step-by-step-1)
  - [Testnet public IPs](#testnet-public-ips)
  - [How It Works](#how-it-works)
    - [Project Anatomy](#project-anatomy)
    - [Transaction Lifecycle](#transaction-lifecycle)
    - [Interacting with a Node](#interacting-with-a-node)
      - [gRPC](#grpc)
      - [REST](#rest)
      - [CometBFT RPC](#cometbft-rpc)
  - [Tutorials](#tutorials)
    - [Setting up a local network with multiple nodes](#setting-up-a-local-network-with-multiple-nodes)
    - [Setup the Faucet Locally](#setup-the-faucet-locally)
    - [Claiming Staking Rewards](#claiming-staking-rewards)
      - [Querying Outstanding Rewards](#querying-outstanding-rewards)
      - [Querying Validator Distribution Info](#querying-validator-distribution-info)
      - [Withdraw All Rewards](#withdraw-all-rewards)
    - [Bank](#bank)
      - [Querying Account Balances](#querying-account-balances)
    - [Slashing](#slashing)
      - [Querying Slashing Params](#querying-slashing-params)
      - [Querying Signing info](#querying-signing-info)
      - [Querying Slashes](#querying-slashes)
      - [Sending Unjail Transaction](#sending-unjail-transaction)
    - [Staking](#staking)
  - [Acknowledgements](#acknowledgements)

## Requirements

- [Go v1.22](https://go.dev/dl/)
- [Ignite v28.2](https://docs.ignite.com/welcome/install)
- [Rust v1.76](https://www.rust-lang.org/tools/install)

## Example Local Blockchain

To run a single node blockchain, run:

```sh
make run-macos # or 
make run-linux
```

This command installs dependencies, builds, initializes, and starts your blockchain in development.

You can try to send an example proof used in the repo with the following command:

```sh
alignedlayerd tx verify gnark-plonk --from alice --chain-id alignedlayer \
    $(cat ./prover_examples/gnark_plonk/example/proof.base64.example) \
    $(cat ./prover_examples/gnark_plonk/example/public_inputs.base64.example) \
    $(cat ./prover_examples/gnark_plonk/example/verifying_key.base64.example)
```

This will output the transaction result (usually containing default values as it doesn't wait for the blockchain to execute it), and the transaction hash.

```txt
...
txhash: F105EAD99F96289914EF16CB164CE43A330AEDB93CAE2A1CFA5FAE013B5CC515
```

To get the transaction result, run:

```sh
alignedlayerd query tx <txhash> | grep proof_verifies -B 10
```

## Verifiers

Information on the parameters received by the CLI when sending transactions can be found by running:

```sh
alignedlayerd tx verify --help
```
Currently, verify supports four proof systems: gnark-plonk, cairo-platinum, sp1 and kimchi. 

Upon verification, the transaction produces an event called `verification_finished` which contains a boolean attriute `proof_verifies` indicating the result.

We also provide the script send_verify_tx.sh to send verification transactions. You may use it according to the following syntax:

```sh
bash send_verify_tx.sh <verifier> <account> <proof_file>
```

### Gnark Plonk

If you want to generate a Gnark Plonk proof by yourself, you must edit the circuit definition and witness in `./prover_examples/gnark_plonk/gnark_plonk.go` and run the following command:

```sh
go run ./prover_examples/gnark_plonk/gnark_plonk.go
```

This will compile the circuit and create a proof in the root folder that is ready to be sent with:

```sh
alignedlayerd tx verify gnark-plonk --from alice --chain-id alignedlayer \
    $(cat proof.base64) \
    $(cat public_inputs.base64) \
    $(cat verifying_key.base64)
```

### Cairo Platinum

To send a Cairo Platinum verification transaction, we can use the following script, which generates the proof manually by reading the file in order to bypass the shell limit (the size of Cairo proofs tends to be large). 

```sh
bash send_verify_tx.sh cairo-platinum alice ./prover_examples/cairo_platinum/example/fibonacci_10.proof
```

If we need, we can set GAS and FEES as env vars before running the script.

>[!TIP]
> The script already converts the `.proof` to `.proof.base64`, but `base64` can be used as follows to encode the proofs:
> ```sh
> base64 -i ./prover_examples/cairo_platinum/example/fibonacci_10.proof -o ./prover_examples/cairo_platinum/example/fibonacci_10.base64
> ```

To create your own proofs, visit [CairoVM](https://github.com/lambdaclass/cairo-vm).

### Kimchi
To send a Kimchi verification transaction, run the following command: 

```sh
bash send_verify_tx.sh kimchi alice ./prover_examples/kimchi/example/kimchi_ec_add.proof
```

### Sp1

To send a Sp1 verification transaction, run the following command: 

```sh
bash send_verify_tx.sh sp-1 alice ./prover_examples/sp1/example/fibonacci.proof
```

## Trying our testnet

Compile with:

```sh
make build-macos # or
make build-linux
```

Create some keys:

```sh
alignedlayerd keys add <your_key_name> --node tcp://91.107.239.79:26657
```

After adding the keys you will get an address, use it in the [faucet](https://faucet.alignedlayer.com/) to get more gas for paying fees.

If you forgot your address, you can get it again with:

```sh
alignedlayerd keys list
```

To send a gnark-plonk proof to the blockchain, use:

```sh
alignedlayerd tx verify gnark-plonk --from <your_key_name> \
	--chain-id alignedlayer \
	--node tcp://rpc-node.alignedlayer.com:26657 \
	--fees 50stake \
	$(cat ./prover_examples/gnark_plonk/example/proof.base64.example) \
	$(cat ./prover_examples/gnark_plonk/example/public_inputs.base64.example) \
	$(cat ./prover_examples/gnark_plonk/example/verifying_key.base64.example)
```

## Joining Our Testnet

### Requirements

#### Hardware

- CPU: 4 cores
- Memory: 16GB
- Disk: 160GB

#### Software

- [jq](https://jqlang.github.io/jq/download/)
- [sponge](https://linux.die.net/man/1/sponge)

### With Docker

If you want to run a node on Docker, you first need to build the image by running:

```sh
docker build . -t alignedlayerd_i
```

Then go into the `docker` directory and run the following command to setup the node:

```sh
bash docker.sh setup <your_node_name>
```

After that, you can decide if you want to be a validator or not. To start the node, run:

```sh
bash docker.sh run[-validator] <your_node_name>
```

Once you do a `run-validator`, that node name will be a validator, even if you later run it with `run`.

### Node Setup

To join our network as a full-node, you need a list of public nodes to first connect to. This must be set on a PEER_ADDR env variable:

```sh
export PEER_ADDR=91.107.239.79,116.203.81.174,88.99.174.203,128.140.3.188
```

A list of our testnet public IP addresses can be found [below](#publicips).

#### The fast way

The fastest way to setup a new node is with our script. It receives the new node's moniker as argument:

```sh
bash setup_node.sh <your-node-name>
```

Then we can start the node with:

```sh
alignedlayerd start
```

#### Manual step by step
<details>
 <summary>Steps</summary>
If you want to do a more detailed step by step setup, follow this instructions:

First, build the app:
```sh
make build_<macos or linux>
```

To make sure the installation was successful, run the following command:
```sh
alignedlayerd version
```

To initialize the node, run
```sh
alignedlayerd init <your-node-name> --chain-id alignedlayer
```
If you have already run this command, you can use the `-o` flag to overwrite previously generated files.

You now need to download the blockchain genesis file and replace the one which was automatically generated for you. Running this command gets the genesis from the first address in `$PEER_ADDR`:
```sh
curl -s $(echo $PEER_ADDR | cut -d, -f1):26657/genesis | jq '.result.genesis' > ~/.alignedlayer/config/genesis.json
```

You now need to build a initial node list. This is the list of nodes you will first connect to, preferablly you should use add all of our public nodes. The list should have this structure:
```
<node1_ID>@<node1_IP>:26656,<node2_ID>@<node2_IP>:26656,...
```

You can get the initial node list by running:
```sh
export INIT_NODES=""; for ip in $(echo $PEER_ADDR | sed 's/,/ /g'); do export INIT_NODES="$INIT_NODES$(curl -s $ip:26657/status | jq -r '.result.node_info.id')@$ip:26656,"; done; export INIT_NODES=${INIT_NODES%?}
```

To check if the list was created correctly you can print the list:
```sh
echo $INIT_NODES
```

To configure persistent peers and gas prices, run the following commands:
```sh
alignedlayerd config set config p2p.persistent_peers "$INIT_NODES" --skip-validate
alignedlayerd config set app minimum-gas-prices 0.0001stake --skip-validate
```

The two most important ports are 26656 and 26657.

The former is used to establish P2P communication with other nodes. This port should be open to world, in order to allow others to communicate with you. Check that the `$HOME/.alignedlayer/config/config.toml` file contains the right address in the p2p section:

```toml
laddr = "tcp://0.0.0.0:26656"
```

The second port is used for the RPC server. If you want to allow remote conections to your node to make queries and transactions, open this port. Note that by default the config sets the address (`rpc.laddr`) to `tcp://127.0.0.1:26657`, you might change the IP to.

Finally, start your node:
```sh
alignedlayerd start
```

You should keep this shell session attached to this process.

The node will start to sync up with the blockchain. To check if your node is already synced:
```sh
curl -s localhost:26657/status |  jq '.result.sync_info.catching_up'
```

It should return `false`. If not, try again some minutes later.
</details>

### Creating an Account

The following command shows all the possible operations regarding keys:

```sh
alignedlayerd keys --help
```

Set a new key:

```sh
alignedlayerd keys add <account-name>
```

This commands will return the following information:
```
address: alignedxxxxxxxxxxxx
name: your-account-name
pubkey: '{"@type":"xxxxxx","key":"xxxxxx"}'
type: local
```

You'll be encouraged to save a mnemomic in case you need to recover your account.

> [!TIP]
> If you don't remember the address, you can do the following:
> `alignedlayerd keys show <address>` or `alignedlayerd keys list`

To check the balance of an address using the binary:

```sh
alignedlayerd query bank balances <account-address-or-name>
```

To ask for tokens, connect to our [faucet](https://faucet.alignedlayer.com) with your browser. You'll be asked to specify your account address `alignedxxxxxxxxxxxx`, which you obtained in the previuos step.

### Registering as a Validator

#### The fast way

The fastest way to setup a new node is with our script. It receives the amount to stake as an argument:

```sh
bash setup_validator.sh <account-name-or-address> 1050000stake
```

This will configure your node and send a transaction for creating a validator.

#### Manual step by step

<details>
  <summary>Steps</summary>
If you want to do a more detailed step by step registering, follow this instructions:

First, obtain your validator pubkey:

```sh
alignedlayerd tendermint show-validator
```

Now create the validator.json file:
```json
{
	"pubkey": {"@type": "...", "key": "..."}, // <-- Replace this with your pubkey
	"amount": "XXXXXstake", // <-- Replace the XXXXX with the amount you want to stake
	"moniker": "your-node-name", // <-- Replace this with your validator name
	"commission-rate": "0.1",
	"commission-max-rate": "0.2",
	"commission-max-change-rate": "0.01",
	"min-self-delegation": "1"
}
```

Now, run:
```sh
alignedlayerd tx staking create-validator validator.json --from <account-name-or-address> --node tcp://$PEER_ADDR:26657 --fees 50stake --chain-id alignedlayer
```

Check whether your validator was accepted with:
```sh
alignedlayerd query tendermint-validator-set | grep $(alignedlayerd tendermint show-address)
```

It should return something like:

```
- address: alignedvalcons1yead8vgxnmtvmtfrfpleuntslx2jk85drx3ug3
```
</details>

## Testnet public IPs

Our public nodes have the following IPs. Please be aware that they are in development stage, so expect inconsistency.

```
91.107.239.79
116.203.81.174
88.99.174.203
128.140.3.188
```

## How It Works

### Project Anatomy

The core of the state machine `App` is defined in [app.go](https://github.com/lambdaclass/aligned_layer_tendermint/blob/main/app/app.go). The application inherits from Cosmos' `BaseApp`, which routes messages to the appropriate module for handling. A transaction contains any number of messages.

Cosmos SDK provides an Application Module interface to facilitate the composition of modules to form a functional unified application. Custom modules are defined in the [x](https://github.com/lambdaclass/aligned_layer_tendermint/blob/main/x/) directory.

A module defines a message service for handling messages. These services are defined in a [protobuf file](https://github.com/lambdaclass/aligned_layer_tendermint/blob/main/proto/alignedlayer/verify/tx.proto). The methods are then implemented in a [message server](https://github.com/lambdaclass/aligned_layer_tendermint/blob/main/x/verify/keeper/msg_server.go), which is registered in the main application.

Each message's type is identified by its fully-qualified name. For example, the _verify_ message has the type `/alignedlayer.verify.MsgVerify`.

A module usually defines a [keeper](https://github.com/lambdaclass/aligned_layer_tendermint/blob/main/x/verify/keeper/keeper.go) which encapsulates the sub-state of each module, tipically through a key-value store. A reference to the keeper is stored in the message server to be accesed by the handlers.

<p align="center">
  <img src="imgs/Diagram_Cosmos.svg">
</p>

The boilerplate for creating custom modules and messages can be generated using Ignite CLI. To generate a new module, run:

```sh
ignite scaffold module <module-name>
```

To generate a message handler for the module, run:

```sh
ignite scaffold message --module <module-name> <message-name> \
    <parameters...> \
    --response <response-fields...>
```

See the [Ignite CLI reference](https://docs.ignite.com/references/cli) to learn
about other scaffolding commands.

### Transaction Lifecycle

A transaction can be created and sent with protobuf with ignite CLI. A JSON representation of the transaction can be obtained with the `--generate-only` flag. It contains transaction metadata and a set of messages. A **message** contains the fully-qualified type to route it correctly, and its parameters.

```json
{
    "body": {
        "messages": [
            {
                "@type": "/alignedlayer.verify.MsgName",
                "creator": "aligned1524vzjchy064rr98d2de7u6uvl4qr3egfq67xn",
                "parameter1": "argument1"
                "parameter2": "argument2"
                ...
            }
        ],
        "memo": "",
        "timeout_height": "0",
        "extension_options": [],
        "non_critical_extension_options": []
    },
    "auth_info": {
        "signer_infos": [],
        "fee": {
            "amount": [],
            "gas_limit": "200000",
            "payer": "",
            "granter": ""
        },
        "tip": null
    },
    "signatures": []
}
```

After Comet BFT receives the transaction, its relayed to the application through the ABCI methods `checkTx` and `deliverTx`.

- `checkTx`: The default `BaseApp` implementation does the following.
    - Checks that a handler exists for every message based on its type.
    - A `ValidateBasic` method (optionally implemented for each message type) is executed for every message, allowing stateless validation. This step is deprecated and should be avoided.
    - The `AnteHandler`'s are executed, by default verifying transaction authentication and gas fees.
- `deliverTx`: In addition to the `checkTx` steps previously mentioned, the following is executed to.
    - The corresponding handler is called for every message.
    - The `PostHandler`'s are executed.

The response is then encoded in the transaction result, and added to the blockchain.

### Interacting with a Node

The full-node exposes three different types of endpoints for interacting with it.

#### gRPC

The node exposes a gRPC server on port 9090.

To get a list with all services, run:

```sh
grpcurl -plaintext localhost:9090 list
```

The requests can be made programatically with any programming language containing the protobuf definitions.

#### REST

The node exposes REST endpoints via gRPC-gateway on port 1317. An OpenAPI specification can be found [here](https://docs.cosmos.network/api)

To get the status of the server, run:

```sh
curl "http://localhost:1317/cosmos/base/node/v1beta1/status" 
```

#### CometBFT RPC

The CometBFT layer exposes a RPC server on port 26657. An OpenAPI specification can be found in [here](https://docs.cometbft.com/v0.38/rpc/).

When sending the transaction, it must be sent serialized with protobuf and encoded in base64, like the following example:

```json
{
    "jsonrpc": "2.0",
    "id": 2,
    "method": "broadcast_tx_sync",
    "params": {
        "tx": "CloKWAoeL2xhbWJjaGFpbi5sYW1iY2hhaW4uTXNnVmVyaWZ5EjYKLWNvc21vczE1MjR2empjaHkwNjRycjk4ZDJkZTd1NnV2bDRxcjNlZ2ZxNjd4bhIFcHJvb2YSWApQCkYKHy9jb3Ntb3MuY3J5cHRvLnNlY3AyNTZrMS5QdWJLZXkSIwohAn0JsZxYl0K5OPEcDNS6nTDsERXapNMidfDtTtrsjtGwEgQKAggBGA0SBBDAmgwaQIzdKrUQB9oMGpFTbPJgLMbcGDvteJ+KIShE7FlUxcipS9i8FslYSqPoZ0RUg9LAGl4/PMD8s/ooEpzO4N7XqLs="
    }
}
```

This is the format used by the CLI.

## Tutorials

### Setting up a local network with multiple nodes

Sets up a network of docker containers each with a validator node and a faucet account.

Build docker images:
```sh
docker build . -t alignedlayerd_i
docker build . -t alignedlayerd_faucet -f node.Dockerfile
```

After building the image we need to set up the files for each cosmos validator node.
The steps are:
- Creating and initializing each node working directory with cosmos files.
- Add users for each node with sufficient funds.
- Create and distribute inital genesis file.
- Set up addresses between nodes.
- Set up faucet files.
- Build docker compose file.

Run script (replacing node names eg. `bash multi_node_setup.sh node0 node1 node2`).

```sh
bash multi_node_setup.sh <node1_name> [<node2_name> ...]
```

The script retrives the password from the **PASSWORD** env_var. 
'password' is set as the default.

Start nodes:
```sh
docker-compose --project-name alignedlayer -f ./prod-sim/docker-compose.yml up --detach
```
This command creates a docker container for each node. Only the first node (`<node1_name>`) has the 26657 port open to receive RPC requests.

It also creates an image that runs the faucet frontend in `localhost:8088`.

You can verify that it works by running (replacing `<node1_name>` by the name of the first node chosen in the bash script):
```sh
docker run --rm -it --network alignedlayer_net-public alignedlayerd_i status --node "tcp://<node1_name>:26657"
```

### Setup the Faucet Locally

The dir `/faucet` has the files needed to setup the client.

Requirements:

- npm
- node

Instructions:

Include the mnemonic at `faucet/.faucet/mnemonic.txt` to reconstruct the address responsible for generating transactions, ensuring that the address belongs to a validator.

Change the parameters defined by the `config.js` file as needed, such as:
- The node's endpoint with: `rpc_endpoint`
- How much it is given per request: `tx.amount`

```
cd faucet
npm install
node faucet.js
```

Then the express server is started at `localhost:8088`
Note: The Tendermint Node(Blockchain) has to be running.

Now the web view can used to request tokens or curl can be used as follows:
```sh
curl http://localhost:8088/send/alignedlayer/:address
```
### Claiming Staking Rewards

Validators and delegators can use the following commands to claim their rewards:

#### Querying Outstanding Rewards
The **validator-outstanding-rewards** command allows users to query all outstanding (un-withdrawn) rewards for a validator and all their delegations.

```sh
alignedlayerd query distribution validator-outstanding-rewards [validator] [flags]
```

Example:
```sh
alignedlayerd query distribution validator-outstanding-rewards alignedvaloper1...
```
Example Output:
```sh
rewards:
- amount: "1000000.000000000000000000"
  denom: stake
```

#### Querying Validator Distribution Info
The **validator-distribution-info** command allows users to query validator commission and self-delegation rewards for validator.

Example:
```sh
alignedlayerd query distribution validator-distribution-info alignedvaloper1...
```
Example output:
```sh
commission:
- amount: "100000.000000000000000000"
  denom: stake
operator_address: alignedvaloper1...
self_bond_rewards:
- amount: "100000.000000000000000000"
  denom: stake
```

#### Withdraw All Rewards
The **withdraw-rewards** command allows users to withdraw all rewards from a given delegation address, and optionally withdraw validator commission if the delegation address given is a validator operator and the user proves the **--commission** flag.
```sh
alignedlayerd tx distribution withdraw-rewards [validator-addr] [flags]
```

Example:
```sh
alignedlayerd tx distribution withdraw-rewards alignedvaloper1... --from aligned1... --commission
```

See the Cosmos' [documentation](https://docs.cosmos.network/main/build/modules/distribution) to learn
about other distribution commands.

### Bank 
#### Querying Account Balances
You can use the **balances** command to query account balances by address.
```sh
alignedlayerd query bank balances [address] [flags]
```
Example:
```sh
alignedlayerd query bank balances aligned1..
```

### Slashing
You can use the slashing CLI commands to query slashing state
```
alignedlayerd query slashing --help
```
#### Querying Slashing Params
To query genesis parameters for the slashing module:
```
alignedlayerd query slashing params [flags]
```
#### Querying Signing info
- To query signing infos of all validators:
    ```sh
    alignedlayerd query slashing signing-infos [flags]
    ```
    Example output:
    ```
    info:
    - address: alignedvalcons15gc...
    index_offset: "147"
    jailed_until: "1970-01-01T00:00:00Z"
    - address: alignedvalcons14xa...
    index_offset: "147"
    jailed_until: "1970-01-01T00:00:00Z"
    - address: alignedvalcons14nz...
    index_offset: "147"
    jailed_until: "1970-01-01T00:00:00Z"
    - address: alignedvalcons1a34...
    index_offset: "147"
    jailed_until: "1970-01-01T00:00:00Z"
    pagination:
    total: "4"
    ```
- To query signing-info of the validator using consensus public key.

    Example:
    ```sh
    alignedlayerd query slashing signing-info alignedvalcons15gc...
    ```

    Example output:
    ```
    val_signing_info:
        address: alignedvalcons15gc...
        index_offset: "255"
        jailed_until: "1970-01-01T00:00:00Z"
        missed_blocks_counter: "16"
    ```

#### Querying Slashes
To query all slashes for a given block range:
```
alignedlayerd query distribution slashes [validator-addr] [start-height] [end-height] [flags]
```

#### Sending Unjail Transaction
To send a transaction to unjail yourself, after the JailPeriod, and thus rejoin the validator set:
```
alignedlayerd tx slashing unjail --from <account_name> --chain-id alignedlayer --fees 20stake
```

### Staking 
You may stake additional tokens after registering your validator with the following command: 
```
alignedlayerd tx staking delegate <valoperaddr> <amount> --from <account_name> --chain-id alignedlayer --fees 20stake
```

You can obtain your validator `valoperaddr` by doing:

```
alignedlayerd keys show <account_name> --bech val --address
```

## Acknowledgements 
We are most grateful to [Cosmos SDK](https://github.com/cosmos/cosmos-sdk), [Ignite CLI](https://github.com/ignite/cli), [CometBFT](https://github.com/cometbft/cometbft) and Ping.pub for their [faucet](https://github.com/ping-pub/faucet) and [explorer](https://github.com/ping-pub/explorer).

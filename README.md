# Aligned Layer Blokchain

An application-specific blockchain built using [Cosmos SDK](https://docs.cosmos.network/) and created with [Ignite CLI](https://ignite.com/). The blockchain offers a variety of zkSNARK implementations to verify proofs sent over transactions, and stores their results.

Cosmos SDK provides a framework to build an application layer on top of a consensus layer interacting via ABCI (Application BlockChain Interface). By default, [CometBFT](https://cometbft.com/) (a fork of Tendermint) is used in the consensus and network layer.

Ignite CLI is used to generate boilerplate code for a Cosmos SDK application, making it easier to deploy a blockchain to production.

## Requirements

- Go (v1.22)
- Ignite (v28.2)

## Example Application Usage with Local Blockchain 

To run a single node blockchain, run:

```sh
ignite chain serve
```

This command installs dependencies, builds, initializes, and starts your blockchain in development.

To send a verify message (transaction), use the following command:

```sh
alignedlayerd tx verification verify --from alice --chain-id alignedlayer <proof> <public_inputs> <verifying_key>
```

You can try with an example proof used in the repo with the following command:

```sh
alignedlayerd tx verification verify --from alice --chain-id alignedlayer \
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
alignedlayerd query tx <txhash>
```
If you want to generate a gnark proof by yourself, you must edit the circuit definition and soltion in `./prover_examples/gnark_plonk/gnark_plonk.go` and run the following command:

```sh
go run ./prover_examples/gnark_plonk/gnark_plonk.go
```

This will compile the circuit and create a proof in the root folder that is ready to be sent with:

```sh
alignedlayerd tx verification verify --from alice --chain-id alignedlayer \
    $(cat proof.base64) \
    $(cat public_inputs.base64) \
    $(cat verifying_key.base64)
```

## How to join as validator

### Requirements

#### Hardware

- CPU: 4 cores
- Memory: 16GB
- Disk: 160GB

#### Software

- jq

### Steps
To set up a validator node, you can either run the provided script setup_validator.sh, or manually run the step by step instructions (see below). The script receives three command line parameters: the name for the validator and the stake amount. For example:
```sh
bash setup_validator.sh myValidator 6000000 stake
```

In order to join the blockchain, you need a known public node to first connect to. As an example, we will name it `blockchain-1`.

1. Get the code and build the app:
```sh
git clone https://github.com/yetanotherco/aligned_layer_tendermint.git
cd aligned_layer_tendermint
ignite chain build --output OUTPUT_BIN 
```

To make sure the installation was successful, run the following command:
```sh
alignedlayerd version
```

2. To create the node, run
```sh
alignedlayerd init <your-node-name> --chain-id alignedlayer
```
If you have already run this command, you can use the -o flag to overwrite previously generated files. 

3. You now need to download the blockchain genesis file and replace the one which was automatically generated for you:
```sh
curl -s blockchain-1:26657/genesis | jq '.result.genesis' > ~/.alignedlayer/config/genesis.json
```

4. Obtain your NODEID by running:
```sh
curl -s blockchain-1:26657/status | jq -r '.result.node_info.id'
```
To configure persistent peers, seeds and gas prices, run the following commands:
```sh
alignedlayerd config set config p2p.seeds "NODEID@blockchain-1:26656" --skip-validate
alignedlayerd config set config p2p.persistent_peers "NODEID@blockchain-1:26656" --skip-validate
alignedlayerd config set app minimum-gas-prices 0.25stake --skip-validate
``` 

Alternatively, you can update the configuration manually:

Add to $HOME/.alignedlayer/config/config.toml the following address to the [p2p] seeds and persistent_peers:
```txt
seeds = "NODEID@blockchain-1:26656"
persistent_peers = "NODEID@blockchain-1:26656"
```

Choose and specify in $HOME/.alignedlayer/config/app.toml the minimum gas price the validator is willing to accept for processing a transaction:
```txt
minimum-gas-prices = "0.25stake"
```

5. The two most important ports are 26656 and 26657.

The former is used to establish p2p communication with other nodes. This port should be open to world, in order to allow others to communicate with you. Check that the $HOME/.alignedlayer/config/config.toml file contains the right address in the p2p section:

```txt
laddr = "tcp://0.0.0.0:26656"
```

The second port is used for the RPC server. If you want to allow remote conections to your node to make queries and transactions, open this port. Note that by default the config sets the address (`rpc.laddr`) to `tcp://127.0.0.1:26657`, you should change the IP to.

6. Start your node:
```sh
alignedlayerd start
```

7. Check if your node is already synced:
```sh
curl -s 127.0.0.1:26657/status |  jq '.result.sync_info.catching_up'
```
It should return false. 

8. Make an account:
```sh
alignedlayerd keys add <your-validator>
```
This commands will return the following information:
```txt
address: cosmosxxxxxxxxxxxx
 name: your-validator
 pubkey: '{"@type":"xxxxxx","key":"xxxxxx"}'
 type: local
```
You'll be encouraged to save a mnemomic in case you need to recover your account. 

9. Ask for tokens. To do so, connect to http://91.107.239.79:8088/. You'll be asked to specify your account address cosmosxxxxxxxxxxxx, which you obtained in step 8. 

10. To create the validator, you need to create a validator.json file. First, obtain your validator pubkey:

```sh
alignedlayerd tendermint show-validator
```

Now create the validator.json file:
```json
{
	"pubkey": {"@type": "...", "key": "..."},
	"amount": "xxxxxxstake",
	"moniker": "your-validator",
	"commission-rate": "0.1",
	"commission-max-rate": "0.2",
	"commission-max-change-rate": "0.01",
	"min-self-delegation": "1"
}
```

Now, run:
```sh
alignedlayerd tx staking create-validator validator.json --from <your-validator-address> --node tcp://blockchain-1:26656 --fees 20000stake
```

Your validator address is the one you obtained in step 8.

11. Check whether your validator was accepted:
```sh
alignedlayerd query tendermint-validator-set
```

Our public nodes have the following IPs. Please be aware that they are in development stage, so expect inconsistency.

```
91.107.239.79
116.203.81.174
88.99.174.203
```

## How It Works

### Project Anatomy

The core of the state machine `App` is defined in [app.go](https://github.com/lambdaclass/aligned_layer_tendermint/blob/main/app/app.go). The application inherits from Cosmos' `BaseApp`, which routes messages to the appropriate module for handling. A transaction contains any number of messages.

Cosmos SDK provides an Application Module interface to facilitate the composition of modules to form a functional unified application. Custom modules are defined in the [x](https://github.com/lambdaclass/aligned_layer_tendermint/blob/main/x/) directory.

A module defines a message service for handling messages. These services are defined in a [protobuf file](https://github.com/lambdaclass/aligned_layer_tendermint/blob/main/proto/alignedlayer/verification/tx.proto). The methods are then implemented in a [message server](https://github.com/lambdaclass/aligned_layer_tendermint/blob/main/x/verification/keeper/msg_server.go), which is registered in the main application.

Each message's type is identified by its fully-qualified name. For example, the _verify_ message has the type `/alignedlayer.verification.MsgVerify`.

A module usually defines a [keeper](https://github.com/lambdaclass/aligned_layer_tendermint/blob/main/x/verification/keeper/keeper.go) which encapsulates the sub-state of each module, tipically through a key-value store. A reference to the keeper is stored in the message server to be accesed by the handlers.

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
                "@type": "/alignedlayer.verification.MsgName",
                "creator": "cosmos1524vzjchy064rr98d2de7u6uvl4qr3egfq67xn",
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

## Setting up multiple local nodes using docker

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

## Tutorials

### How to Create a new Address

The following command shows all the possible operations regarding keys:

```sh
alignedlayerd keys --help
```

Set a new key:

```sh
alignedlayerd keys add <id_string>
```

> [!TIP]
> If you don't remember the address, you can do the following:
> `alignedlayerd keys show <address>` or `alignedlayerd keys list`

Use the faucet in order to have some balance.

To check the balance of an address using the binary: 

```sh
alignedlayerd query bank balances <address or id_string>
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
alignedlayerd query distribution validator-outstanding-rewards cosmosvaloper1...
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
alignedlayerd query distribution validator-distribution-info cosmosvaloper1...
```
Example output:
```sh
commission:
- amount: "100000.000000000000000000"
  denom: stake
operator_address: cosmosvaloper1...
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
alignedlayerd tx distribution withdraw-rewards cosmosvaloper1... --from cosmos1... --commission
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
alignedlayerd query bank balances cosmos1..
```

# Acknowledgements
We are most grateful to [Cosmos SDK](https://github.com/cosmos/cosmos-sdk), [Ignite CLI](https://github.com/ignite/cli), [CometBFT](https://github.com/cometbft/cometbft) and [Ping.pub](https://github.com/ping-pub/faucet).



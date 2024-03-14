# Aligned Layer Blokchain

An application-specific blockchain built using [Cosmos SDK](https://docs.cosmos.network/) and created with [Ignite CLI](https://ignite.com/). The blockchain offers a variety of zkSNARK implementations to verify proofs sent over transactions, and stores their results.

Cosmos SDK provides a framework to build an application layer on top of a consensus layer interacting via ABCI (Application BlockChain Interface). By default, [CometBFT](https://cometbft.com/) (a fork of Tendermint) is used in the consensus and network layer.

Ignite CLI is used to generate boilerplate code for a Cosmos SDK application, making it easier to deploy a blockchain to production.

## Requirements

- Go
- Ignite

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

Sets up a network of docker containers each with a validator node.

Build docker image:
```sh
docker build . -t alignedlayerd_i
```

After building the image we need to set up the files for each cosmos validator node.
The steps are:
- Creating and initializing each node working directory with cosmos files.
- Add users for each node with sufficient funds.
- Create and distribute inital genesis file.
- Set up addresses between nodes.
- Build docker compose file.

Run script (replacing node names eg. `bash multi_node_setup.sh node0 node1 node2`)
```sh
bash multi_node_setup.sh <node1_name> [<node2_name> ...]
```

Start nodes:
```sh
docker-compose --project-name alignedlayer -f ./prod-sim/docker-compose.yml up --detach
```
This command creates a docker container for each node. Only the first node (`<node1_name>`) has the 26657 port open to receive RPC requests.

You can verify that it works by running (replacing `<node1_name>` by the name of the first node chosen in the bash script):
```sh
docker run --rm -it --network alignedlayer_net-public alignedlayerd_i status --node "tcp://<node1_name>:26657"
```

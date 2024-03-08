# Bootcamp Verifying Lambchain (WIP)

An application-specific blockchain built using [Cosmos SDK](https://docs.cosmos.network/) and created with [Ignite CLI](https://ignite.com/). The blockchain offers a variety of zkSNARK implementations to verify proofs sent over transactions, and stores their results.

Cosmos SDK provides a framework to build an application layer on top of a consensus layer interacting via ABCI (Application BlockChain Interface). By default, [CometBFT](https://cometbft.com/) (a fork of Tendermint) is used in the consensus and network layer.

Ignite CLI is used to generate boilerplate code for a Cosmos SDK application, making it easier to deploy a blockchain to production.

## Requirements

- Go
- Ignite

## Example Application Usage 

To run a single node blockchain, run:

```sh
ignite chain serve
```

This command installs dependencies, builds, initializes, and starts your blockchain in development.

To send a verify message (transaction), run:

```sh
lambchaind tx lambchain verify --from alice --chain-id lambchain <proof>
```

This will output the transaction result (usually containing default values as it doesn't wait for the blockchain to execute it), and the transaction hash.

```txt
...
txhash: F105EAD99F96289914EF16CB164CE43A330AEDB93CAE2A1CFA5FAE013B5CC515
```

To get the transaction result, run:

```sh
lambchaind query tx <txhash>
```

## How It Works

### Project Anatomy

The core of the state machine `App` is defined in [app.go](https://github.com/lambdaclass/lambchain/blob/main/app/app.go). The application inherits from Cosmos' `BaseApp`, which routes messages to the appropriate module for handling. A transaction contains any number of messages.

Cosmos SDK provides an Application Module interface to facilitate the composition of modules to form a functional unified application. Custom modules are defined in the [x](https://github.com/lambdaclass/lambchain/blob/main/x/) directory.

A module defines a message service for handling messages. These services are defined in a [protobuf file](https://github.com/lambdaclass/lambchain/blob/main/proto/lambchain/lambchain/tx.proto). The methods are then implemented in a [message server](https://github.com/lambdaclass/lambchain/blob/main/x/lambchain/keeper/msg_server.go), which is registered in the main application.

Each message's type is identified by its fully-qualified name. For example, the _verify_ message has the type `/lambchain.lambchain.MsgVerify`.

A module usually defines a [keeper](https://github.com/lambdaclass/lambchain/blob/main/x/lambchain/keeper/keeper.go) which encapsulates the sub-state of each module, tipically through a key-value store. A reference to the keeper is stored in the message server to be accesed by the handlers.

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

A transaction can be created and sent with protobuf with ignite CLI, using the following command:

```sh
lambchaind tx lambchain verify --from alice --chain-id lambchain "base64-encoded proof"
```

A JSON representation of the transaction can be obtained with the `--generate-only` flag. It contains transaction metadata and a set of messages. A **message** contains the fully-qualified type to route it correctly, and its parameters.

```json
{
    "body": {
        "messages": [
            {
                "@type": "/lambchain.lambchain.MsgVerify",
                "creator": "cosmos1524vzjchy064rr98d2de7u6uvl4qr3egfq67xn",
                "proof": "base64-encoded proof"
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


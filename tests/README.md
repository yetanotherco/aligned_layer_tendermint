## Testing

This Makefile makes it easier to test node behaviour under different transaction messages.

### Available targets

- `generate PROOF=<proof>`: Generate a `Verify` transaction with a parameter `<proof>` and save it in `transaction.json`.
- `sign`: Sign `transaction.json` with _Alice_ key and save the signed transaction in `signed.json`.
- `encode`: Serialize `signed.json` using Protobuf and save it in `encoded.txt` as base64.
- `send-without-encode`: Broadcast the `encode.txt` transaction to the blockchain.
- `send`: A shortcut for `encode` and `send-without-encode`.
- `result HASH=<tx_hash>`: Get the result code of the transaction with hash `<tx_hash>` (`0` means success).
- `clean`: Remove all generated files.

This targets help testing transasctions by allowing us to modify messages at -_almost_- every step on the process.

#### Example

```sh
make generate PROOF="Hello, world!"
make sign
make send
# ...
# txhash: 6E4...
# ...
make result HASH=6E4...
# 0
```

```sh
make generate PROOF="An invalid sentence"
make sign
make send
# ...
# txhash: D46...
# ...
make result HASH=D46...
# 1101
```
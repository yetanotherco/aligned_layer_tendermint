# Gnark Plonk

The Gnark Plonk verifier in the blockchain needs the following base64 encoded elements:

- Proof
- Public Inputs
- Verifying Key

The serialization is performed by the Gnark Plonk library.

## Sending a Proof to Local Blockchain

Change the circuit definition and solution inside `gnark_plonk.go`:

Generate the proof and necessary elements for the verification by running:

```sh
make generate-proof
```

This will generate the necessary files in the current directory. These files
will be used by the makefile in the next step.

Send the proof to the local blockchain. 

```sh
make send-proof
```

This will output the transaction hash.

To clean the generated files, run:

```sh
make clean
```

The last three steps can be executed in a single command with:

```sh
make prove
```

To query the result, run:

```sh
HASH=63a... make query-tx
```

We should see an event called `verifiaction_finished` containing a `proof_verifies` attribute.

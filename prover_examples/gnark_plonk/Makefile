.PHONY: prove generate-proof send-proof query-tx clean

CHAIN ?= alignedlayer
MODULE ?= verification
MESSAGE ?= verify
PROVER ?= gnark_plonk.go

PROOF_FILE ?= proof.base64
PUBLIC_INPUTS_FILE ?= public_inputs.base64
VERIFYING_KEY_FILE ?= verifying_key.base64

prove: | generate-proof send-proof clean

generate-proof:
	go run gnark_plonk.go

send-proof:
	$(CHAIN)d tx $(MODULE) $(MESSAGE) \
		$$(cat $(PROOF_FILE)) \
		$$(cat $(PUBLIC_INPUTS_FILE)) \
		$$(cat $(VERIFYING_KEY_FILE)) \
		--from alice --chain-id $(CHAIN) \

query-tx:
	$(CHAIN)d query tx $(HASH) \

clean:
	rm *.base64

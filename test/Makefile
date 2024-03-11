PROOF ?= $(PROOF)
PW ?= $(PW)
generate:
	lambchaind tx lambchain verify $(PROOF) $(PW) \
		--from alice \
		--chain-id lambchain \
		--generate-only > transaction.json

sign: transaction.json
	lambchaind tx sign transaction.json \
		--chain-id lambchain \
		--keyring-backend test \
		--from alice \
		--overwrite \
		--output-document signed.json

encode: signed.json
	lambchaind tx encode signed.json > encoded.txt

send-without-encode:
	grpcurl -plaintext \
	    -d "{\"tx_bytes\":\"`cat encoded.txt`\",\"mode\":\"BROADCAST_MODE_SYNC\"}" \
	    localhost:9090 \
	    cosmos.tx.v1beta1.Service/BroadcastTx

send: | encode send-without-encode

# Output result code for HASH transaction
result:
	lambchaind query tx $(HASH) \
		--output json | jq .code

clean:
	rm -f transaction.json signed.json encoded.txt

plonk-test:
	lambchaind tx lambchain verify 5+qp3poS52nTTOlBoKrAQ20QT9C+qTIj1N2sGKQ7wOHLPLP1UmnQuGv56qpYqKLeqoFq8nGzlTtgaIUCRSm5vIqGrmF0tjZTBcLgsRtCd6KStBtsFeCLP7SnWuKI/l6irE9w65Wu8kZg0W83yjUMyx9hA3N2LaT1eXPa+lhW0PLXBqNpfouTKPAykWnbrT5dmdnFx3ilwNS8NpVGqWLX9ozePVoZ15L1eMbGfp8N2vol93cOnYAcgDT4uku4Glgw5yu0QUfvSVcpEGhurodpRd/I5Qh+jGYv1dmQODfEyFOEmHxX4WQKbxvHTusXCj64mNzVAxYhqNxbtuUz1RjwLAAAAAcMJbP+2806jQu2QkcKdCqJc3u/l0LhivVdcQ/f0hRqfR8fAiOIign63aNXfm21KaJR3u80jctkFBMRi9IQlA7VLoxR+SSbDFBhO4F6c8i6MgAQU8qJjMGIGL4Xl2aAwyADSiq7qQuOXHwJKeQEZyH3wxLaGmtBVRZZSHj5Q7gPCgyYkaeU0s/y0oY/eguyoQOa7AiwTwf8oSrSj31mHPsYCFW7L6Zrs3xNUpni+VWGzFUj5F4hsOUdvanW/xQOYW4PwvvYNw2MHOCe7oxreftVkWJkj6CNBW8xeAe/3WMJf+qf1y5PQsjt6Ch5L0eXqCu9yMj23gPx6rnoTrieon3LEpWd6D5RUCL6SkHT+8Hkc6hSd+DonrdW0d91UruByNkAAAAA AAAAAgAAAAAAAAACAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAIAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAABA== --from alice --chain-id lambchain

plonk-bad-test:
	lambchaind tx lambchain verify 5+qppoS52nTTOlBoKrAQ20QT9C+qTIj1N2sGKQ7wOHLPLP1UmnQuGv56qpYqKLeqoFq8nGzlTtgaIUCRSm5vIqGrmF0tjZTBcLgsRtCd6KStBtsFeCLP7SnWuKI/l6irE9w65Wu8kZg0W83yjUMyx9hA3N2LaT1eXPa+lhW0PLXBqNpfouTKPAykWnbrT5dmdnFx3ilwNS8NpVGqWLX9ozePVoZ15L1eMbGfp8N2vol93cOnYAcgDT4uku4Glgw5yu0QUfvSVcpEGhurodpRd/I5Qh+jGYv1dmQODfEyFOEmHxX4WQKbxvHTusXCj64mNzVAxYhqNxbtuUz1RjwLAAAAAcMJbP+2806jQu2QkcKdCqJc3u/l0LhivVdcQ/f0hRqfR8fAiOIign63aNXfm21KaJR3u80jctkFBMRi9IQlA7VLoxR+SSbDFBhO4F6c8i6MgAQU8qJjMGIGL4Xl2aAwyADSiq7qQuOXHwJKeQEZyH3wxLaGmtBVRZZSHj5Q7gPCgyYkaeU0s/y0oY/eguyoQOa7AiwTwf8oSrSj31mHPsYCFW7L6Zrs3xNUpni+VWGzFUj5F4hsOUdvanW/xQOYW4PwvvYNw2MHOCe7oxreftVkWJkj6CNBW8xeAe/3WMJf+qf1y5PQsjt6Ch5L0eXqCu9yMj23gPx6rnoTrieon3LEpWd6D5RUCL6SkHT+8Hkc6hSd+DonrdW0d91UruByNkAAAAA AAAAAgAAAAAAAAACAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAIAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAABA== --from alice --chain-id lambchain

__KIMCHI_FFI__: ## 
build-kimchi-macos:
		@cd operators/kimchi/lib \
				&& cargo build --release \
				&& cp target/release/libkimchi_verifier_ffi.dylib ./libkimchi_verifier.dylib \
				&& cp target/release/libkimchi_verifier_ffi.a ./libkimchi_verifier.a

build-kimchi-linux:
		@cd operators/kimchi/lib \
				&& cargo build --release \
				&& cp target/release/libkimchi_verifier_ffi.so ./libkimchi_verifier.so \
				&& cp ./lib/target/release/libkimchi_verifier_ffi.a ./libkimchi_verifier.a

test-kimchi-ffi: 
	go test -v ./operators/kimchi

proof-to-base64:
	@cd ./to_base64 && go run proof_to_base64.go 

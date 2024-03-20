__KIMCHI_FFI__: ## 
build-kimchi-macos:
	@cd ./lib && cargo build --release
	@cp ./lib/target/release/libkimchi_verifier_ffi.dylib ./lib/libkimchi_verifier.dylib

build-kimchi-linux:
	@cd ./lib && cargo build --release
	@cp ./lib/target/release/libkimchi_verifier_ffi.so operator/kimchi/lib/libkimchi_verifier.so

test-kimchi-ffi: 
	go test ./... -v

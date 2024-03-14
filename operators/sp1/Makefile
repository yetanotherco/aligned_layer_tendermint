__SP1_FFI__: ## 
build-sp1-macos:
	@cd lib && cargo build --release
	@cp lib/target/release/libsp1_verifier_wrapper.dylib lib/libsp1_verifier.dylib

build-sp1-linux:
	@cd lib && cargo build --release
	@cp lib/target/release/libsp1_verifier_wrapper.so lib/libsp1_verifier.so

test-ffi-sp1: 
	go test -v
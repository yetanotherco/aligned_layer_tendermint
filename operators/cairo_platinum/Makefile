
__LAMBDAWORKS_FFI__: ## 
build-lambdaworks-macos:
	@cd lib && cargo build --release
	@cp lib/target/release/libcairo_platinum_ffi.a lib/libcairo_platinum.a

build-lambdaworks-linux:
	@cd lib && cargo build --release
	@cp lib/target/release/libcairo_platinum_ffi.a lib/libcairo_platinum.a

test-ffi-lambdaworks: 
	go test -v
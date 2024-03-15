
__LAMBDAWORKS_FFI__:
build-cairo-ffi:
	@cd operators/cairo_platinum/lib && cargo build --release && cp target/release/libcairo_platinum_ffi.dylib ./libcairo_platinum.dylib

test-ffi-cairo: 
	go test -v ./operators/cairo_platinum 


__SP1_FFI__: ## 
build-sp1-ffi:
	@cd operators/sp1/lib && cargo build --release && cp target/release/libsp1_verifier_wrapper.dylib ./libsp1_verifier.dylib

test-ffi-sp1: 
	go test -v ./operators/sp1 


__COSMOS_BLOCKCHAIN__:
run_chain: build-sp1-ffi build-cairo-ffi
	ignite chain serve


test_cairo:

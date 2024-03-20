__LAMBDAWORKS_FFI_LINUX__:
build-cairo-ffi-macos:
	@cd operators/cairo_platinum/lib \
		&& cargo build --release \
		&& cp target/release/libcairo_platinum_ffi.dylib ./libcairo_platinum.dylib \
		&& cp target/release/libcairo_platinum_ffi.a ./libcairo_platinum.a 

build-cairo-ffi-linux:
	@cd operators/cairo_platinum/lib \
		&& cargo build --release \
		&& cp target/release/libcairo_platinum_ffi.so ./libcairo_platinum.so \
		&& cp target/release/libcairo_platinum_ffi.a ./libcairo_platinum.a 

test-ffi-cairo: 
	go test -v ./operators/cairo_platinum 

__SP1_FFI__: 
build-sp1-ffi-macos:
	@cd operators/sp1/lib \
		&& cargo build --release \
		&& cp target/release/libsp1_verifier_wrapper.dylib ./libsp1_verifier.dylib \
		&& cp target/release/libsp1_verifier_wrapper.a ./libsp1_verifier.a

build-sp1-ffi-linux:
	@cd operators/sp1/lib \
		&& cargo build --release \
		&& cp target/release/libsp1_verifier_wrapper.so ./libsp1_verifier.so \
		&& cp target/release/libsp1_verifier_wrapper.a ./libsp1_verifier.a

test-ffi-sp1:
	go test -v ./operators/sp1 

__COSMOS_BLOCKCHAIN__:
run-macos: build-sp1-ffi-macos build-cairo-ffi-macos
	ignite chain serve

build-macos: build-sp1-ffi-macos build-cairo-ffi-macos
	ignite chain build

run-linux: build-sp1-ffi-linux build-cairo-ffi-linux
	ignite chain serve

build-linux: build-sp1-ffi-linux build-cairo-ffi-linux
	ignite chain build

__LOCAL_TEST__:
ltest-cairo-true: 
	alignedlayerd tx verification verifycairo \
		--from alice \
		--gas 4000000 \
		--chain-id alignedlayer \
		$$(cat prover_examples/cairo_platinum/fibonacci_10.base64.example)

ltest-cairo-false:
	alignedlayerd tx verification verifycairo \
		--from alice \
		--chain-id alignedlayer \
		SHOULDFAIL

clean-ffi:
	rm -rf operators/sp1/lib/target/release/libsp1_verifier*
	rm -rf operators/cairo_platinum/lib/libcairo_platinum*
	rm -rf operators/sp1/lib/target/release/libsp1_verifier*
	rm -rf operators/cairo_platinum/lib/target/release/libcairo_platinum*

clean:
	rm -rf ~/.alignedlayer

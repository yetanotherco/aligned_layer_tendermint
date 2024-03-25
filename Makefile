__LAMBDAWORKS_FFI__:
build-cairo-ffi-macos:
	@cd verifiers/cairo_platinum/lib \
		&& cargo build --release \
		&& cp target/release/libcairo_platinum_ffi.dylib ./libcairo_platinum.dylib \
		&& cp target/release/libcairo_platinum_ffi.a ./libcairo_platinum.a 

build-cairo-ffi-linux:
	@cd verifiers/cairo_platinum/lib \
		&& cargo build --release \
		&& cp target/release/libcairo_platinum_ffi.so ./libcairo_platinum.so \
		&& cp target/release/libcairo_platinum_ffi.a ./libcairo_platinum.a 

test-ffi-cairo: 
	go test -v ./verifiers/cairo_platinum 

__SP1_FFI__:
build-sp1-ffi-macos:
	@cd verifiers/sp1/lib \
		&& cargo build --release \
		&& cp target/release/libsp1_verifier_wrapper.dylib ./libsp1_verifier.dylib \
		&& cp target/release/libsp1_verifier_wrapper.a ./libsp1_verifier.a

build-sp1-ffi-linux:
	@cd verifiers/sp1/lib \
		&& cargo build --release \
		&& cp target/release/libsp1_verifier_wrapper.so ./libsp1_verifier.so \
		&& cp target/release/libsp1_verifier_wrapper.a ./libsp1_verifier.a

test-ffi-sp1:
	go test -v ./verifiers/sp1

__COSMOS_BLOCKCHAIN__:
build-macos: build-sp1-ffi-macos build-cairo-ffi-macos
	ignite chain build

run-macos: build-macos
	ignite chain serve

build-linux: build-sp1-ffi-linux build-cairo-ffi-linux
	ignite chain build

run-linux: build-linux
	ignite chain serve

__LOCAL_TEST__:
clean-ffi:
	rm -rf verifiers/sp1/lib/target/release/libsp1_verifier*
	rm -rf verifiers/cairo_platinum/lib/libcairo_platinum*
	rm -rf verifiers/sp1/lib/target/release/libsp1_verifier*
	rm -rf verifiers/cairo_platinum/lib/target/release/libcairo_platinum*

clean:
	rm -rf ~/.alignedlayer

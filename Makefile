__LAMBDAWORKS_FFI_LINUX__:
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

__COSMOS_BLOCKCHAIN__:
build-macos: build-cairo-ffi-macos
	ignite chain build

run-macos: build-macos
	ignite chain serve

build-linux: build-cairo-ffi-linux
	ignite chain build

run-linux: build-linux
	ignite chain serve

clean-ffi:
	rm -rf verifiers/cairo_platinum/lib/libcairo_platinum*
	rm -rf verifiers/cairo_platinum/lib/target/release/libcairo_platinum*

clean:
	rm -rf ~/.alignedlayer
	rm ${HOME}/go/bin/alignedlayerd

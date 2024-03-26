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

__KIMCHI_FFI__: ## 
build-kimchi-macos:
		@cd verifiers/kimchi/lib \
		&& cargo build --release \
		&& cp target/release/libkimchi_verifier_ffi.dylib ./libkimchi_verifier.dylib \
		&& cp target/release/libkimchi_verifier_ffi.a ./libkimchi_verifier.a

build-kimchi-linux:
		@cd verifiers/kimchi/lib \
		&& cargo build --release \
		&& cp target/release/libkimchi_verifier_ffi.so ./libkimchi_verifier.so \
		&& cp target/release/libkimchi_verifier_ffi.a ./libkimchi_verifier.a

test-kimchi-ffi: 
	go test -v ./verifiers/kimchi

__COSMOS_BLOCKCHAIN__:
build-macos: build-cairo-ffi-macos build-sp1-ffi-macos build-kimchi-macos
	ignite chain build

run-macos: build-macos
	ignite chain serve

build-linux: build-cairo-ffi-linux build-sp1-ffi-linux build-kimchi-linux
	ignite chain build

run-linux: build-linux
	ignite chain serve

clean-ffi:
	rm -rf verifiers/cairo_platinum/lib/libcairo_platinum*
	rm -rf verifiers/cairo_platinum/lib/target/release/libcairo_platinum*
	rm -rf verifiers/kimchi/lib/libkimchi_verifier*
	rm -rf verifiers/kimchi/lib/target/release/libcairo_verifier*

clean:
	rm -rf ~/.alignedlayer
	rm ${HOME}/go/bin/alignedlayerd

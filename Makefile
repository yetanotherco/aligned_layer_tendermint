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
	base64 -i ./operators/kimchi/example/kimchi_ec_add.proof.example -o ./operators/kimchi/example/kimchi_ec_add.proof.example.base64

__COSMOS_BLOCKCHAIN__:
build-macos: build-cairo-ffi-macos build-kimchi-macos
	ignite chain build

run-macos: build-macos
	ignite chain serve

build-linux: build-cairo-ffi-linux build-kimchi-macos
	ignite chain build

run-linux: build-linux build-kimchi-linux
	ignite chain serve

clean-ffi:
	rm -rf operators/cairo_platinum/lib/libcairo_platinum*
	rm -rf operators/cairo_platinum/lib/target/release/libcairo_platinum*

clean:
	rm -rf ~/.alignedlayer

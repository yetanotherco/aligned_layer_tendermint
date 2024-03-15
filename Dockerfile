FROM ignitehq/cli as builder

USER root
WORKDIR /root

# Get Ubuntu packages
RUN apt-get install -y \
    build-essential \
    curl
# Get Rust
RUN curl https://sh.rustup.rs -sSf | sh -s -- -y
# Add .cargo/bin to PATH
ENV PATH="/root/.cargo/bin:${PATH}"


COPY . .

RUN make build-sp1-ffi
RUN make build-cairo-ffi

#RUN cp operators/cairo_platinum/lib/libcairo_platinum.dylib /usr/bin/ld
#RUN cp operators/sp1/lib/libsp1_verifier.dylib /usr/bin/ld

#RUN ignite chain build

ENTRYPOINT [ "alignedlayerd" ]

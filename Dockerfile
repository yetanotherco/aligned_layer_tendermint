FROM ignitehq/cli AS builder

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

COPY . ./alignedlayer

WORKDIR /root/alignedlayer

RUN make clean-ffi
RUN make build-linux

FROM debian:stable-slim

RUN apt-get update && apt-get install -y curl jq

COPY --from=builder /go/bin/alignedlayerd /bin/alignedlayerd

ENTRYPOINT [ "alignedlayerd" ]

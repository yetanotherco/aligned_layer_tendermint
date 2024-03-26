FROM rust AS lib-compiler

USER root
WORKDIR /root

COPY ./verifiers ./verifiers
COPY ./Makefile .

RUN make clean-ffi
RUN make build-libs-linux

FROM ignitehq/cli AS cosmos-builder

USER root
WORKDIR /root/alignedlayer

# Get Ubuntu packages
RUN apt-get install -y build-essential

COPY . .
COPY --from=lib-compiler /root/verifiers ./verifiers

RUN make build

FROM debian:stable-slim

RUN apt-get update && apt-get install -y curl jq

COPY --from=cosmos-builder /go/bin/alignedlayerd /bin/alignedlayerd

ENTRYPOINT [ "alignedlayerd" ]

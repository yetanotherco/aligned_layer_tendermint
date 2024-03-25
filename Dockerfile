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

RUN make clean_ffi
RUN make build_linux

ENTRYPOINT [ "alignedlayerd" ]

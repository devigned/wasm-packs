FROM rust

# Install packages that we want to make available at both build and run time
RUN apt update && \
  apt install -y build-essential curl && \
  rm -rf /var/lib/apt/lists/*

# Create user and group
ARG cnb_uid=1001
ARG cnb_gid=1001
RUN groupadd cnb --gid ${cnb_gid} && \
  useradd --uid ${cnb_uid} --gid ${cnb_gid} -m -s /bin/bash cnb

# Set user and group
USER ${cnb_uid}:${cnb_gid}
ENV HOME=/home/cnb

# Get Rust
ENV CARGO_HOME=/home/cnb/.cargo
ENV RUSTUP_HOME=/home/cnb/.rustup
RUN curl https://sh.rustup.rs -sSf | bash -s -- -y -target=wasm32-wasip2
RUN echo 'source $HOME/.cargo/env' >> $HOME/.bashrc

RUN cargo install \
  --profile='fastest-runtime' \
  --config='profile.fastest-runtime.strip="symbols"' \
  --config='profile.fastest-runtime.panic="abort"' \
  wasmtime-cli

RUN cargo install --locked wasm-tools

USER root
RUN cp $CARGO_HOME/bin/wasmtime /wasmtime

# Set required CNB target information
ARG distro_name
LABEL io.buildpacks.base.distro.name=${distro_name}
ARG distro_version
LABEL io.buildpacks.base.distro.version=${distro_version}

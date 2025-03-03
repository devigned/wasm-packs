ARG BASE_IMAGE=heroku/heroku:24-build
FROM ${BASE_IMAGE} AS base

# Install packages that we want to make available at build time
USER root
RUN apt update && \
  apt install -y git wget jq && \
  rm -rf /var/lib/apt/lists/*

USER heroku
# Get Rust
ENV CARGO_HOME=/home/heroku/.cargo
ENV RUSTUP_HOME=/home/heroku/.rustup
RUN curl https://sh.rustup.rs -sSf | bash -s -- -y -target=wasm32-wasip2
RUN echo 'source $HOME/.cargo/env' >> $HOME/.bashrc
ENV PATH="$PATH:/home/heroku/.cargo/bin"

RUN cargo install \
  --profile='fastest-runtime' \
  --config='profile.fastest-runtime.strip="symbols"' \
  --config='profile.fastest-runtime.panic="abort"' \
  wasmtime-cli

RUN cargo install --locked wasm-tools

USER root
RUN cp $CARGO_HOME/bin/wasmtime /wasmtime

USER heroku
LABEL io.buildpacks.base.distro.name="ubuntu"
LABEL io.buildpacks.base.distro.version="24.04"
LABEL io.buildpacks.base.homepage="https://github.com/devigned/wasm-packs"
LABEL io.buildpacks.base.maintainer="devigned"
ENV CNB_USER_ID=1000
ENV CNB_GROUP_ID=1000

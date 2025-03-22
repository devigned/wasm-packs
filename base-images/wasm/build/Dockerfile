ARG BASE_IMAGE=heroku/heroku:24-build
FROM --platform=${BUILDPLATFORM} ${BASE_IMAGE} AS base

ARG TARGETOS
ARG TARGETARCH

# Install packages that we want to make available at build time
USER root
RUN apt update && \
  apt install -y git wget jq musl-tools && \
  rm -rf /var/lib/apt/lists/*

ENV GO_RELEASE=1.24.1
RUN wget https://dl.google.com/go/go${GO_RELEASE}.${TARGETOS}-${TARGETARCH}.tar.gz && \
  tar xfv go${GO_RELEASE}.${TARGETOS}-${TARGETARCH}.tar.gz -C /usr/local && \
  find /usr/local/go -mindepth 1 -maxdepth 1 ! -name 'src' ! -name 'VERSION' ! -name 'bin' ! -name 'pkg' ! -name 'go.env' -exec rm -rf {} +

ENV TINYGO_RELEASE=0.37.0
RUN wget https://github.com/tinygo-org/tinygo/releases/download/v${TINYGO_RELEASE}/tinygo${TINYGO_RELEASE}.${TARGETOS}-${TARGETARCH}.tar.gz && \
  tar xfv tinygo${TINYGO_RELEASE}.${TARGETOS}-${TARGETARCH}.tar.gz -C /usr/local

ENV PATH=${PATH}:/usr/local/tinygo/bin
ENV PATH=${PATH}:/usr/local/go/bin

RUN wget https://github.com/mikefarah/yq/releases/latest/download/yq_linux_${TARGETARCH} -O /usr/local/bin/yq && chmod +x /usr/local/bin/yq

USER heroku
# Get Rust
ENV CARGO_HOME=/home/heroku/.cargo
ENV RUSTUP_HOME=/home/heroku/.rustup
RUN curl https://sh.rustup.rs -sSf | bash -s -- -y --target=wasm32-wasip2,$(uname -m)-unknown-linux-musl
RUN echo 'source $HOME/.cargo/env' >> $HOME/.bashrc
ENV PATH="$PATH:/home/heroku/.cargo/bin"

LABEL io.buildpacks.base.distro.name="ubuntu"
LABEL io.buildpacks.base.distro.version="24.04"
LABEL io.buildpacks.base.homepage="https://github.com/devigned/wasm-packs"
LABEL io.buildpacks.base.maintainer="devigned"
ENV CNB_USER_ID=1000
ENV CNB_GROUP_ID=1000

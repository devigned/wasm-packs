ARG BASE_IMAGE=wasm/demo-base-build:wasm
FROM ${BASE_IMAGE} AS base

FROM gcr.io/distroless/static:nonroot AS final
COPY --from=base --chown=nonroot:nonroot --chmod=0755 /wasmtime /wasmtime
LABEL io.buildpacks.stack.id="heroku-24"

ARG base_image
FROM ${base_image} AS base

FROM gcr.io/distroless/static:nonroot AS final

COPY --from=base --chown=nonroot:nonroot --chmod=0755 /wasmtime /wasmtime

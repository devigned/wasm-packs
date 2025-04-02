# Example of Cloud Native Buildpacks for Wasm

This repository contains a set of build packs and example Wasm applications to demonstrate how to build and run Wasm applications using Cloud Native Buildpacks.

## The Parts
### Base Images
The `base-images` directory contains the Dockerfiles for the base images used by the buildpacks. The `build.sh` script builds the base images used for build
and run in the buildpacks. They are based on the Heroku stack images.

### Builders
The `builders` directory contains the builder for the Wasm buildpacks.

### Buildpacks
The `buildpacks` directory contains the buildpacks used in the builder.

#### Meta-buildpacks
The `meta-buildpacks` directory contains the meta-buildpacks used in the builder.

### Apps
The `apps` directory contains the example Wasm applications. The `js` directory contains a simple Wasm application written in Rust and compiled to Wasm. The `compose` directory contains a multi-component Wasm microservice application using both JS and Golang to demonstrate the use of multiple languages in a single application.

## Building and Running the Example
```bash
./base-images/build.sh wasm
pack builder create wasm/demo-builder:wasm --config ../builders/wasm/builder.toml
pack build wasm-compose --builder wasm/demo-builder:wasm --path apps/compose/
docker run --rm -it -p 8080:8080 -e OPENAI_API_KEY wasm-compose
```

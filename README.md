# wasm-packs

```bash
./base-images/build.sh wasm
pack builder create wasm/demo-builder:wasm --config ../builders/wasm/builder.toml
pack build test-wasm-js --builder wasm/demo-builder:wasm --path apps/js/
docker run --rm -it -p 8080:8080 test-wasm-js
```

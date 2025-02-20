# wasm-packs

```bash
pack builder create wasm/demo-builder:wasm --config ../builders/wasm/builder.toml
pack build test-wasm-js --builder wasm/demo-builder:wasm --path apps/js/
```

//go:generate wkg wit fetch
//go:generate wit-bindgen-go generate -o ./internal ./wit

package main

import (
	"github.com/devigned/wasm-packs/compose/internal/example/domain/adder"
)

func init() {
	adder.Exports.Add = func(x int32, y int32) int32 {
		// This is where you would implement the logic for the add function.
		// For example, you could return the sum of x and y.
		return x + y
	}
}

// main is required for the `wasi` target, even if it isn't used.
func main() {
}

// Code generated by wit-bindgen-go. DO NOT EDIT.

// Package preopens represents the imported interface "wasi:filesystem/preopens@0.2.3".
package preopens

import (
	"github.com/devigned/wasm-packs/compose/internal/wasi/filesystem/types"
	"go.bytecodealliance.org/cm"
)

// Descriptor represents the imported type alias "wasi:filesystem/preopens@0.2.3#descriptor".
//
// See [types.Descriptor] for more information.
type Descriptor = types.Descriptor

// GetDirectories represents the imported function "get-directories".
//
//	get-directories: func() -> list<tuple<descriptor, string>>
//
//go:nosplit
func GetDirectories() (result cm.List[cm.Tuple[Descriptor, string]]) {
	wasmimport_GetDirectories(&result)
	return
}

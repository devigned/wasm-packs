// Code generated by wit-bindgen-go. DO NOT EDIT.

// Package stdout represents the imported interface "wasi:cli/stdout@0.2.0".
package stdout

import (
	"github.com/devigned/wasm-packs/compose/internal/wasi/io/v0.2.0/streams"
	"go.bytecodealliance.org/cm"
)

// OutputStream represents the imported type alias "wasi:cli/stdout@0.2.0#output-stream".
//
// See [streams.OutputStream] for more information.
type OutputStream = streams.OutputStream

// GetStdout represents the imported function "get-stdout".
//
//	get-stdout: func() -> output-stream
//
//go:nosplit
func GetStdout() (result OutputStream) {
	result0 := wasmimport_GetStdout()
	result = cm.Reinterpret[OutputStream]((uint32)(result0))
	return
}

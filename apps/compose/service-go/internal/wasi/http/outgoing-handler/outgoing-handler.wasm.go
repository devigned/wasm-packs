// Code generated by wit-bindgen-go. DO NOT EDIT.

package outgoinghandler

import (
	"go.bytecodealliance.org/cm"
)

// This file contains wasmimport and wasmexport declarations for "wasi:http@0.2.3".

//go:wasmimport wasi:http/outgoing-handler@0.2.3 handle
//go:noescape
func wasmimport_Handle(request0 uint32, options0 uint32, options1 uint32, result *cm.Result[ErrorCodeShape, FutureIncomingResponse, ErrorCode])

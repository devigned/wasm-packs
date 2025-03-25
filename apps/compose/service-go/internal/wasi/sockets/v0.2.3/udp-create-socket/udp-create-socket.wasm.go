// Code generated by wit-bindgen-go. DO NOT EDIT.

package udpcreatesocket

import (
	"go.bytecodealliance.org/cm"
)

// This file contains wasmimport and wasmexport declarations for "wasi:sockets@0.2.3".

//go:wasmimport wasi:sockets/udp-create-socket@0.2.3 create-udp-socket
//go:noescape
func wasmimport_CreateUDPSocket(addressFamily0 uint32, result *cm.Result[UDPSocket, UDPSocket, ErrorCode])

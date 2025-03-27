// Code generated by wit-bindgen-go. DO NOT EDIT.

// Package terminalinput represents the imported interface "wasi:cli/terminal-input@0.2.3".
//
// Terminal input.
//
// In the future, this may include functions for disabling echoing,
// disabling input buffering so that keyboard events are sent through
// immediately, querying supported features, and so on.
package terminalinput

import (
	"go.bytecodealliance.org/cm"
)

// TerminalInput represents the imported resource "wasi:cli/terminal-input@0.2.3#terminal-input".
//
// The input side of a terminal.
//
//	resource terminal-input
type TerminalInput cm.Resource

// ResourceDrop represents the imported resource-drop for resource "terminal-input".
//
// Drops a resource handle.
//
//go:nosplit
func (self TerminalInput) ResourceDrop() {
	self0 := cm.Reinterpret[uint32](self)
	wasmimport_TerminalInputResourceDrop((uint32)(self0))
	return
}

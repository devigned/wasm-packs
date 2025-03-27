// Code generated by wit-bindgen-go. DO NOT EDIT.

// Package terminalstderr represents the imported interface "wasi:cli/terminal-stderr@0.2.3".
//
// An interface providing an optional `terminal-output` for stderr as a
// link-time authority.
package terminalstderr

import (
	terminaloutput "github.com/devigned/wasm-packs/compose/internal/wasi/cli/terminal-output"
	"go.bytecodealliance.org/cm"
)

// TerminalOutput represents the imported type alias "wasi:cli/terminal-stderr@0.2.3#terminal-output".
//
// See [terminaloutput.TerminalOutput] for more information.
type TerminalOutput = terminaloutput.TerminalOutput

// GetTerminalStderr represents the imported function "get-terminal-stderr".
//
// If stderr is connected to a terminal, return a `terminal-output` handle
// allowing further interaction with it.
//
//	get-terminal-stderr: func() -> option<terminal-output>
//
//go:nosplit
func GetTerminalStderr() (result cm.Option[TerminalOutput]) {
	wasmimport_GetTerminalStderr(&result)
	return
}

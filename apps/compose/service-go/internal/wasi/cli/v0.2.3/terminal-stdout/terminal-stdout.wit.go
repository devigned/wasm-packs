// Code generated by wit-bindgen-go. DO NOT EDIT.

// Package terminalstdout represents the imported interface "wasi:cli/terminal-stdout@0.2.3".
//
// An interface providing an optional `terminal-output` for stdout as a
// link-time authority.
package terminalstdout

import (
	terminaloutput "github.com/devigned/wasm-packs/compose/internal/wasi/cli/v0.2.3/terminal-output"
	"go.bytecodealliance.org/cm"
)

// TerminalOutput represents the imported type alias "wasi:cli/terminal-stdout@0.2.3#terminal-output".
//
// See [terminaloutput.TerminalOutput] for more information.
type TerminalOutput = terminaloutput.TerminalOutput

// GetTerminalStdout represents the imported function "get-terminal-stdout".
//
// If stdout is connected to a terminal, return a `terminal-output` handle
// allowing further interaction with it.
//
//	get-terminal-stdout: func() -> option<terminal-output>
//
//go:nosplit
func GetTerminalStdout() (result cm.Option[TerminalOutput]) {
	wasmimport_GetTerminalStdout(&result)
	return
}

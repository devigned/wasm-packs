// Code generated by wit-bindgen-go. DO NOT EDIT.

// Package wallclock represents the imported interface "wasi:clocks/wall-clock@0.2.3".
package wallclock

import (
	"go.bytecodealliance.org/cm"
)

// DateTime represents the record "wasi:clocks/wall-clock@0.2.3#datetime".
//
//	record datetime {
//		seconds: u64,
//		nanoseconds: u32,
//	}
type DateTime struct {
	_           cm.HostLayout `json:"-"`
	Seconds     uint64        `json:"seconds"`
	Nanoseconds uint32        `json:"nanoseconds"`
}

// Now represents the imported function "now".
//
//	now: func() -> datetime
//
//go:nosplit
func Now() (result DateTime) {
	wasmimport_Now(&result)
	return
}

// Resolution represents the imported function "resolution".
//
//	resolution: func() -> datetime
//
//go:nosplit
func Resolution() (result DateTime) {
	wasmimport_Resolution(&result)
	return
}

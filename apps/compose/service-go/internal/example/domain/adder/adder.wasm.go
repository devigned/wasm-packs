// Code generated by wit-bindgen-go. DO NOT EDIT.

package adder

// This file contains wasmimport and wasmexport declarations for "example:domain@0.3.0".

//go:wasmexport example:domain/adder@0.3.0#add
//export example:domain/adder@0.3.0#add
func wasmexport_Add(x0 uint32, y0 uint32) (result0 uint32) {
	x := (int32)((uint32)(x0))
	y := (int32)((uint32)(y0))
	result := Exports.Add(x, y)
	result0 = (uint32)(result)
	return
}

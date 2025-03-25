package gen

import "strings"

// Ident represents a package-level Go declaration.
type Ident struct {
	Package *Package
	Name    string
}

// ParseSelector parses string s into a package path and short name.
// It does not validate the input or resulting values. Examples:
// "io" -> "io", "io"
// "encoding/json" -> "encoding/json", "json"
// "encoding/json#Decoder" -> "encoding/json", "Decoder"
// "wasi/clocks/wall#DateTime" -> "wasi/clocks/wall", "DateTime"
func ParseSelector(s string) (path, name string) {
	path, name, _ = strings.Cut(s, "#")
	if name == "" {
		i := strings.LastIndex(path, "/")
		if i >= 0 && i < len(path)-1 {
			name = path[i+1:] // encoding/json -> json
		} else {
			name = path // encoding -> encoding
		}
	}
	return path, name
}

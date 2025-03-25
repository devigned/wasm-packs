//go:build !wasip1

package relpath

import "os"

// Getwd returns best-effort path to the current working directory, even on WebAssembly.
// The path may be absolute, or it may be relative, prefixed with one or more ".." segments.
func Getwd() (string, error) {
	return os.Getwd()
}

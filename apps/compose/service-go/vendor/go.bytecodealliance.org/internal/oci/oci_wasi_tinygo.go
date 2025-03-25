//go:build wasip1 || wasip2 || tinygo

package oci

import (
	"context"
	"errors"
)

func IsOCIPath(path string) bool {
	return false
}

func PullWIT(ctx context.Context, path string) ([]byte, error) {
	return nil, errors.New("OCI not supported on WASI or TinyGo")
}

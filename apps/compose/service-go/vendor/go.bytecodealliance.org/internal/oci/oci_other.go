//go:build !wasip1 && !wasip2 && !tinygo

package oci

import (
	"bytes"
	"context"
	"fmt"
	"io"
	"os"

	"github.com/regclient/regclient"
	"github.com/regclient/regclient/types/manifest"
	"github.com/regclient/regclient/types/ref"
)

// IsOCIPath checks if a given path is an OCI path
func IsOCIPath(path string) bool {
	if _, err := os.Stat(path); err == nil {
		return false
	}

	_, err := ref.New(path)
	return err == nil
}

// PullWIT fetches an OCI artifact from a given OCI path
// It invokes "regclient" APIs to pull the artifact and then
// processes it with `wasm-tools`.
// The output is returned as raw bytes.
func PullWIT(ctx context.Context, path string) ([]byte, error) {
	r, err := ref.New(path)
	if err != nil {
		return nil, fmt.Errorf("failed to parse ref: %v", err)
	}

	rc := regclient.New()
	defer rc.Close(ctx, r)

	m, err := rc.ManifestGet(ctx, r)
	if err != nil {
		return nil, fmt.Errorf("failed to get manifest: %v", err)
	}

	mi, ok := m.(manifest.Imager)
	if !ok {
		return nil, fmt.Errorf("manifest does not support image methods")
	}

	layers, err := mi.GetLayers()
	if err != nil {
		return nil, fmt.Errorf("failed to get layers: %v", err)
	}
	if len(layers) == 0 {
		return nil, fmt.Errorf("no layers found in the artifact")
	}

	layer := layers[0] // assuming the WIT file is the first layer

	if err = layer.Digest.Validate(); err != nil {
		return nil, fmt.Errorf("layer contains invalid digest: %s: %v", string(layer.Digest), err)
	}

	rdr, err := rc.BlobGet(ctx, r, layer)
	if err != nil {
		return nil, fmt.Errorf("failed to fetch blob: %v", err)
	}
	defer rdr.Close()

	// Read the blob content into a buffer
	var buf bytes.Buffer
	_, err = io.Copy(&buf, rdr)
	if err != nil {
		return nil, fmt.Errorf("failed to read blob content: %v", err)
	}
	return buf.Bytes(), nil
}

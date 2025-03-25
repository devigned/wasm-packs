package witcli

import (
	"errors"
	"fmt"
	"io/fs"
	"os"
)

// FindOrCreateDir attempts to load file information for the provided path and
// the provided path does not already exist, it attempts to create the directory
// and then return the file information for it.
func FindOrCreateDir(path string) (fs.FileInfo, error) {
	info, err := os.Stat(path)
	if errors.Is(err, fs.ErrNotExist) {
		err = os.MkdirAll(path, os.FileMode(0o755))
		if err != nil {
			return nil, err
		}
		info, err = os.Stat(path)
	}
	if err != nil {
		return nil, err
	}
	if !info.IsDir() {
		return nil, fmt.Errorf("%s is not a directory", path)
	}
	return info, nil
}

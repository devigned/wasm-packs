//go:build wasip1

package relpath

import (
	"io/fs"
	"os"
	"path/filepath"
)

// Getwd returns best-effort path to the current working directory, even on WebAssembly.
// The path may be absolute, or it may be relative, prefixed with one or more ".." segments.
func Getwd() (string, error) {
	var path string
	var target string
	parent := "."
outer:
	for {
		target = parent
		// println("statting target:", target)
		dirInfo, err := os.Stat(target)
		if err != nil {
			return path, err
		}
		parent = filepath.Join("..", target)

		dir, err := os.Open(parent)
		if err != nil {
			return filepath.Clean(filepath.Join(target, path)), nil
		}
		names, err := dir.Readdirnames(-1)
		dir.Close()
		if err != nil {
			return filepath.Clean(filepath.Join(target, path)), nil
		}
		// println("NAMES: ", strings.Join(names, " "))
		for _, name := range names {
			info, err := os.Stat(filepath.Join(parent, name))
			if err != nil {
				continue
			}
			if info.IsDir() && os.SameFile(dirInfo, info) {
				path = filepath.Join(name, path)
				continue outer
			}
		}
		break
	}
	return filepath.Clean(filepath.Join("/", path)), nil
}

func sameFile(a, b fs.FileInfo) bool {
	// fmt.Printf("A: %s %d %v %v %t\n", a.Name(), a.Size(), a.Mode(), a.ModTime(), a.IsDir())
	// fmt.Printf("B: %s %d %v %v %t\n", b.Name(), b.Size(), b.Mode(), b.ModTime(), b.IsDir())
	return a.Size() == b.Size() &&
		a.Mode() == b.Mode() &&
		a.ModTime().Equal(b.ModTime()) &&
		a.IsDir() == b.IsDir()
}

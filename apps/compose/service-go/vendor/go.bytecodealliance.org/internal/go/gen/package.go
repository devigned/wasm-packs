package gen

// Package represents a Go package, containing zero or more files
// of generated code, along with zero or more declarations.
type Package struct {
	// Path is the Go package path, e.g. "encoding/json"
	Path string

	// Name is the short Go package name, e.g. "json"
	Name string

	// Files is the list of Go and non-Go files in this package.
	Files map[string]*File

	// Declared tracks declared package-scoped identifiers,
	// including constants, variables, and functions.
	Scope
}

// NewPackage returns a newly instantiated Package for path.
// The local name may optionally be specified with a "#name" suffix.
func NewPackage(path string) *Package {
	p := &Package{
		Files: make(map[string]*File),
		Scope: NewScope(nil),
	}
	p.Path, p.Name = ParseSelector(path)
	return p
}

// File finds or adds a new file named name to pkg.
func (pkg *Package) File(name string) *File {
	file := pkg.Files[name]
	if file != nil {
		return file
	}
	file = NewFile(pkg, name)
	pkg.Files[name] = file
	return file
}

// HasPackageDocs returns true if pkg contains at least 1 [File]
// with a non-empty PackageDocs field.
func (pkg *Package) HasPackageDocs() bool {
	for _, file := range pkg.Files {
		if file.IsGo() && file.PackageDocs != "" {
			return true
		}
	}
	return false
}

// HasContent returns true if pkg contains at least 1 [File]
// with non-empty content.
func (pkg *Package) HasContent() bool {
	for _, file := range pkg.Files {
		if file.HasContent() {
			return true
		}
	}
	return false
}

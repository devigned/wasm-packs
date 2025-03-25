package gen

import "go/token"

// UniqueName tests name against filters and modifies name until all filters return false.
// Use IsReserved to filter out Go keywords and predeclared identifiers.
//
// Exported names that start with a capital letter will be appended with an underscore.
// Non-exported names are prefixed with an underscore.
func UniqueName(name string, filters ...func(string) bool) string {
	filter := func(name string) bool {
		for _, f := range filters {
			if f(name) {
				return true
			}
		}
		return false
	}
	for filter(name) {
		name += "_"
	}
	return name
}

// Scope represents a Go name scope, like a package, file, interface, struct, or function blocks.
type Scope interface {
	// DeclareName declares name within this scope, modifying it as necessary to avoid
	// colliding with any preexisting names. It returns the unique generated name.
	// Subsequent calls to GetName will return the unique name.
	// Subsequent calls to HasName with the returned name will return true.
	// Subsequent calls to DeclareName will return a different name.
	DeclareName(name string) string

	// GetName returns the first declared unique name for name, if declared.
	GetName(name string) string

	// HasName returns true if this scope or any of its parent scopes contains name.
	HasName(name string) bool
}

type scope struct {
	parent Scope
	exists map[string]bool
	first  map[string]string
}

// NewScope returns an initialized [Scope] that's ready to use.
// If parent is nil, [Reserved] will be used.
func NewScope(parent Scope) Scope {
	if parent == nil {
		parent = Reserved()
	}
	return &scope{
		parent: parent,
		exists: make(map[string]bool),
		first:  make(map[string]string),
	}
}

func (s *scope) DeclareName(name string) string {
	unique := UniqueName(name, s.HasName)
	s.exists[unique] = true
	if _, ok := s.first[name]; !ok {
		s.first[name] = unique
	}
	return unique
}

func (s *scope) GetName(name string) string {
	first := s.parent.GetName(name)
	if first != "" {
		return first
	}
	return s.first[name]
}

func (s *scope) HasName(name string) bool {
	return s.exists[name] || s.parent.HasName(name)
}

type reservedScope struct{}

// Reserved returns a preset [Scope] with the default [Go keywords] and [predeclared identifiers].
// Calls to its UniqueName method will panic, as the scope is immutable.
//
// [Go keywords]: https://go.dev/ref/spec#Keywords
// [predeclared identifiers]: https://go.dev/ref/spec#Predeclared_identifiers
func Reserved() Scope {
	return reservedScope{}
}

func (reservedScope) DeclareName(name string) string {
	panic("cannot add a name to reserved scope")
}

func (reservedScope) GetName(name string) string {
	if IsReserved(name) {
		return name
	}
	return ""
}

func (reservedScope) HasName(name string) bool {
	return IsReserved(name)
}

// IsReserved returns true for any name that is a Go keyword or predeclared identifier.
func IsReserved(name string) bool {
	return token.IsKeyword(name) || reserved[name]
}

var reserved = mapWords(
	// Types
	"any",
	"bool",
	"byte",
	"comparable",
	"complex64",
	"complex128",
	"error",
	"float32",
	"float64",
	"int",
	"int8",
	"int16",
	"int32",
	"int64",
	"rune",
	"string",
	"uint",
	"uint8",
	"uint16",
	"uint32",
	"uint64",
	"uintptr",

	// Constants
	"true",
	"false",
	"iota",

	// Zero value
	"nil",

	// Functions
	"append",
	"cap",
	"clear",
	"close",
	"complex",
	"copy",
	"delete",
	"imag",
	"len",
	"make",
	"max",
	"min",
	"new",
	"panic",
	"print",
	"println",
	"real",
	"recover",
)

// Initialisms is a set of common initialisms.
var Initialisms = mapWords(
	"acl",
	"abi",
	"api",
	"ascii",
	"cabi",
	"cpu",
	"css",
	"cwd",
	"dns",
	"eof",
	"fifo",
	"guid",
	"html",
	"http",
	"https",
	"id",
	"imap",
	"io",
	"ip",
	"js",
	"json",
	"lhs",
	"mime",
	"ok",
	"posix",
	"qps",
	"ram",
	"rhs",
	"rpc",
	"sla",
	"smtp",
	"sql",
	"ssh",
	"tcp",
	"tls",
	"ttl",
	"tty",
	"udp",
	"ui",
	"uid",
	"uuid",
	"uri",
	"url",
	"utf8",
	"vm",
	"xml",
	"xmpp",
	"xsrf",
	"xss",
)

func mapWords(words ...string) map[string]bool {
	m := make(map[string]bool, len(words))
	for _, word := range words {
		m[word] = true
	}
	return m
}

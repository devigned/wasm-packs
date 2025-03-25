package wit

import (
	"go.bytecodealliance.org/wit/iterate"
	"go.bytecodealliance.org/wit/ordered"
)

// An Interface represents a [collection of types and functions], which are imported into
// or exported from a [WebAssembly component].
// It implements the [Node], and [TypeOwner] interfaces.
//
// [collection of types and functions]: https://github.com/WebAssembly/component-model/blob/main/design/mvp/WIT.md#wit-interfaces.
// [WebAssembly component]: https://github.com/WebAssembly/component-model/blob/main/design/mvp/WIT.md#wit-worlds
type Interface struct {
	_typeOwner

	Name      *string
	TypeDefs  ordered.Map[string, *TypeDef]
	Functions ordered.Map[string, *Function]
	Package   *Package  // the Package this Interface belongs to
	Stability Stability // WIT @since or @unstable (nil if unknown)
	Docs      Docs
}

// WITPackage returns the [Package] this [Interface] belongs to.
func (i *Interface) WITPackage() *Package {
	return i.Package
}

// Match returns true if [Interface] i matches pattern, which can be one of:
// "name", "namespace:package/name" (qualified), or "namespace:package/name@1.0.0" (versioned).
func (i *Interface) Match(pattern string) bool {
	if i.Name == nil {
		return false
	}
	if pattern == *i.Name {
		return true
	}
	id := i.Package.Name
	id.Extension = *i.Name
	if pattern == id.String() {
		return true
	}
	id.Version = nil
	return pattern == id.String()
}

// AllFunctions returns a [sequence] that yields each [Function] in an [Interface].
// The sequence stops if yield returns false.
//
// [sequence]: https://github.com/golang/go/issues/61897
func (i *Interface) AllFunctions() iterate.Seq[*Function] {
	return func(yield func(*Function) bool) {
		i.Functions.All()(func(_ string, f *Function) bool {
			return yield(f)
		})
	}
}

func (i *Interface) dependsOn(dep Node) bool {
	if dep == i || dep == i.Package {
		return true
	}
	// _, depIsInterface := dep.(*Interface)
	var done bool
	i.TypeDefs.All()(func(_ string, t *TypeDef) bool {
		done = DependsOn(t, dep)
		// A type alias transitively pulls in the dependencies of its owner
		if root := t.Root(); !done && root != t && root.Owner != nil && root.Owner != i {
			done = DependsOn(root.Owner, dep)
		}
		return !done
	})
	if done {
		return true
	}
	i.Functions.All()(func(_ string, f *Function) bool {
		done = DependsOn(f, dep)
		return !done
	})
	return done
}

// An InterfaceRef represents a reference to an [Interface] with a [Stability] attribute.
// It implements the [Node] and [WorldItem] interfaces.
type InterfaceRef struct {
	_worldItem

	Interface *Interface
	Stability Stability
}

func (ref *InterfaceRef) dependsOn(dep Node) bool {
	return dep == ref || DependsOn(ref.Interface, dep)
}

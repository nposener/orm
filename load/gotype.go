package load

import (
	"path/filepath"
	"strings"
)

type GoType struct {
	Type string
	// ImportPath is a path to add to the import section for this type
	ImportPath string
}

func newGoType(fullName string) GoType {
	return GoType{
		Type:       typeName(fullName),
		ImportPath: importPath(fullName),
	}
}

func (g GoType) String() string {
	if g.ImportPath != "" {
		return pointer(g.Type) + g.ImportPath + "." + g.NonPointer()
	}
	return g.Type
}

// ExtTypeName is the full type of the imported type, as used in a go code
// outside the defining package. For example: "example.Person"
func (g GoType) ExtTypeName() string {
	if g.Package() != "" {
		return pointer(g.Type) + g.Package() + "." + g.NonPointer()
	}
	return g.Type
}

// ExtNonPointer is the full type of the imported type in it's non-pointer form,
// as used in a go code outside the defining package.
// For example: "example.Person"
func (g GoType) ExtNonPointer() string {
	if g.Package() != "" {
		return g.Package() + "." + g.NonPointer()
	}
	return g.NonPointer()
}

// Package is the package name of the type
// for example, type in "github.com/posener/orm/example" has the package
// name: "example"
func (g GoType) Package() string {
	_, pkg := filepath.Split(g.ImportPath)
	return pkg
}

// IsPointer returns true if field is a pointer
func (g *GoType) IsPointer() bool {
	return len(g.Type) > 0 && g.Type[0] == '*'
}

// NonPointer returns the non-pointer type of a filed.
// ex, if the type is `*int`, this function will return `int`
func (g *GoType) NonPointer() string {
	if g.IsPointer() {
		return g.Type[1:]
	}
	return g.Type
}

// import path returns the import statement of a type
// If a full type name is 'github.com/posener/orm/example.Person', this
// function will return 'github.com/posener/orm/example' which is what you
// would write in the `import` statement.
func importPath(fullName string) string {
	fullName = strings.TrimLeft(fullName, "*")
	i := strings.LastIndex(fullName, ".")
	if i == -1 {
		return ""
	}
	return fullName[:i]
}

// typeName returns the type string from a full type name.
// If a full type name is 'github.com/posener/orm/example.Person', this
// function will return 'Person' which is how you would use this
// struct in a go file
func typeName(fullName string) string {
	i := strings.LastIndex(fullName, ".")
	if i == -1 {
		return fullName
	}
	return pointer(fullName) + fullName[i+1:]
}

func pointer(typeName string) string {
	if len(typeName) > 0 && typeName[0] == '*' {
		return "*"
	}
	return ""
}
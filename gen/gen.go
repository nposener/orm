package gen

import (
	"fmt"
	"go/types"
	"log"
	"os"
	"path/filepath"
	"strings"
	"text/template"
)

const suffix = "orm"

type TplProps struct {
	// Table is the table name of the given struct
	Table string
	// The name	of the new created package
	PackageName string
	// Type describes the type of the given struct to generate code for
	Type Type
}

// Type describes the type of the given struct to generate code for
type Type struct {
	// ImportPath is the import path of the package of the given struct.
	// for example: "github.com/posener/orm/example"
	ImportPath string
	// Name is the full type of the imported type, as used in a go code
	// outside the defining package. For example: "example.Person"
	Name string
	// Fields is the list of exported fields
	Fields []Field
}

type Field struct {
	Name       string
	ColumnName string
	Type       string
}

var tpls = template.Must(template.ParseGlob("gen/*.go.tpl"))

func Gen(pkg *types.Package, st *types.Struct, name string) error {
	// get the package dir on disk
	dir, err := packagePath(pkg.Path())
	if err != nil {
		return err
	}

	// the new created package name is the name of the struct with "orm" suffix
	packageName := strings.ToLower(name + suffix)

	// the files will be generated in a sub package
	dir = filepath.Join(dir, packageName)
	log.Printf("Generating code to directory: %s", dir)
	err = os.MkdirAll(dir, 0775)
	if err != nil {
		return fmt.Errorf("creating directory %s: %s", dir, err)
	}

	props := TplProps{
		Table: strings.ToLower(name),
		Type: Type{
			ImportPath: pkg.Path(),
			Name:       fmt.Sprintf("%s.%s", pkg.Name(), name),
			Fields:     collectFields(st),
		},
		PackageName: packageName,
	}

	for _, tpl := range tpls.Templates() {
		err := writeTpl(tpl, props, dir)
		if err != nil {
			return err
		}
	}

	log.Println(props)
	return nil
}

func packagePath(pkg string) (string, error) {
	for _, gopath := range filepath.SplitList(os.Getenv("GOPATH")) {
		pkgPath := filepath.Join(gopath, "src", pkg)
		f, err := os.Stat(pkgPath)
		if err == nil && f.IsDir() {
			return pkgPath, nil
		}
	}
	return "", fmt.Errorf("package path was not found")
}

func writeTpl(tpl *template.Template, props TplProps, dir string) error {
	// remove the ".tpl" suffix
	fileName := tpl.Name()[:len(tpl.Name())-4]
	filePath := filepath.Join(dir, fileName)
	log.Printf("Writing file: %s", filePath)
	f, err := os.Create(filePath)
	if err != nil {
		return fmt.Errorf("creating file %s: %s", filePath, err)
	}
	defer f.Close()
	return tpl.Execute(f, props)
}

func collectFields(s *types.Struct) []Field {
	var f []Field
	for i := 0; i < s.NumFields(); i++ {
		v := s.Field(i)
		if !v.Exported() {
			continue
		}
		f = append(f, Field{
			Name:       v.Name(),
			ColumnName: strings.ToLower(v.Name()),
			Type:       v.Type().String(),
		})
	}
	return f
}

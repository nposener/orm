package gen

import (
	"fmt"
	"go/types"
	"log"
	"os"
	"os/exec"
	"path/filepath"
	"strings"
	"text/template"
)

const suffix = "orm"

const header = `// Autogenerated by github.com/posener/orm
`

type TplProps struct {
	// Table is the table name of the given struct
	Table string
	// The name	of the new created package
	PackageName string
	// Type describes the type of the given struct to generate code for
	Type Type
}

var templates = template.Must(
	template.New("").
		Funcs(template.FuncMap{
			"plus1": func(x int) int { return x + 1 },
		}).
		ParseGlob("gen/*.go.tpl"))

// Gen generates all the ORM files for a given struct in a given package.
// pkg is the package of the struct
// st is the type descriptor of the struct
// name is the name of the type of the struct
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
		Table:       strings.ToLower(name),
		Type:        NewType(name, pkg, st),
		PackageName: packageName,
	}
	log.Printf("Template configuration: %+v", props)

	for _, tpl := range templates.Templates() {
		err := writeTemplate(tpl, props, dir)
		if err != nil {
			return err
		}
	}
	fmtOut, err := exec.Command("gofmt", "-s", "-w", dir).CombinedOutput()
	if err != nil {
		log.Printf("Failed formatting package: %s", err)
	}
	log.Printf("Format package output: %s", string(fmtOut))
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

func writeTemplate(tpl *template.Template, props TplProps, dir string) error {
	// remove the ".tpl" suffix
	fileName := tpl.Name()[:len(tpl.Name())-4]
	filePath := filepath.Join(dir, fileName)
	log.Printf("Writing file: %s", filePath)
	f, err := os.Create(filePath)
	if err != nil {
		return fmt.Errorf("creating file %s: %s", filePath, err)
	}
	defer f.Close()
	f.Write([]byte(header))
	return tpl.Execute(f, props)
}

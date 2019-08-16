package main

import (
	"fmt"
	"go/ast"
	"go/token"
	"os"
	"os/exec"
	"sort"
	"strings"
	"text/template"

	"golang.org/x/tools/go/packages"
)

// 1 OMIT
func main() {
	typeName := os.Args[1]
	var names []string
	// 2 OMIT
	cfg := &packages.Config{
		Mode: packages.NeedTypes | packages.NeedTypesInfo |
			packages.NeedSyntax | packages.NeedName,
	}
	pkgs, err := packages.Load(cfg)
	if err != nil {
		panic(err)
	}
	if len(pkgs) != 1 {
		panic(fmt.Errorf("got unexpected number of packages %v", len(pkgs)))
	}
	pkg := pkgs[0]
	// 3 OMIT
	targetType := pkg.Types.Scope().Lookup(typeName)
	if targetType == nil {
		panic(fmt.Errorf("failed to find type declaration for %v", typeName))
	}
	// 4 OMIT
	for _, file := range pkg.Syntax {
		for _, decl := range file.Decls {
			gd, ok := decl.(*ast.GenDecl) // type, const, var
			if !ok {
				continue
			}
			if gd.Tok != token.CONST {
				continue
			}
			for _, spec := range gd.Specs {
				spec := spec.(*ast.ValueSpec)
				for _, name := range spec.Names {
					if pkg.TypesInfo.Defs[name].Type() == targetType.Type() {
						names = append(names, name.Name)
					}
				}
			}
		}
	}
	sort.Strings(names)
	// 5 OMIT
	const outTmpl = `package {{.PkgName}}

import "fmt"

func (v {{.TypeName}}) String() string {
	switch v {
	{{range .ConstVals -}}
	case {{.}}:
		return "{{.}}"
	{{end -}}
	default:
		panic(fmt.Errorf("unknown {{.TypeName}} value %d", v))
	}
}
`
	// 6 OMIT
	var tmplVals = struct {
		PkgName   string
		TypeName  string
		ConstVals []string
	}{
		PkgName:   pkg.Name,
		TypeName:  typeName,
		ConstVals: names,
	}
	// 7 OMIT
	outFileName := fmt.Sprintf("gen_%v_simplestringer.go", typeName)
	outFile, err := os.Create(outFileName)
	if err != nil {
		panic(err)
	}

	tmpl := template.Must(template.New("out").Parse(outTmpl))
	if err := tmpl.Execute(outFile, tmplVals); err != nil {
		panic(err)
	}
	if err := outFile.Close(); err != nil {
		panic(err)
	}

	cmd := exec.Command("gofmt", "-w", outFileName)
	if errOut, err := cmd.CombinedOutput(); err != nil {
		panic(fmt.Errorf("failed to run %v: %v\n%s", strings.Join(cmd.Args, " "), err, errOut))
	}
}

// 8 OMIT

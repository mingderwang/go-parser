package main

import (
	"fmt"
	"github.com/davecgh/go-spew/spew"
	"go/ast"
	"go/parser"
	"go/token"
)

type Slack struct {
	Name string
}

var src = `
//go:generate ginger $GOFILE
package main
// @gigner
type Slack struct {
	Name string 
}
func main() {
	println("hello ginger")
}
`

func main() {
	fset := token.NewFileSet()
	f, err := parser.ParseFile(fset, "", src, parser.ParseComments)
	if err != nil {
		panic(err)
	}
	ast.Print(fset, f)
	fmt.Println("-----------------")
	for index, decl := range f.Decls {
		fmt.Println(index)
		spew.Dump(decl)
		genDecl, ok := decl.(*ast.GenDecl)
		if !ok {
			fmt.Println("!ok")
		} else {
			if genDecl.Doc == nil {
				fmt.Println("genDecl.Doc == nil")
			} else {
				spew.Dump(genDecl.Doc)

			}
		}
	}
}

package main

import (
	"fmt"
	//	"github.com/davecgh/go-spew/spew"
	"go/ast"
	"go/parser"
	"go/token"
	"strings"
)

var src = `
//go:generate ginger $GOFILE
package main
// @gigner
type SlackMessage struct {
	Name string 
}
type dontParse struct {
	Name string 
}
// @gigner
type SlackUser struct {
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
	//ast.Print(fset, f)
	for index, decl := range f.Decls {
		fmt.Println(index)
		//spew.Dump(decl)
		if genDecl, ok := decl.(*ast.GenDecl); ok {
			if genDecl.Doc == nil {
				fmt.Println("genDecl.Doc == nil")
			} else {
				//spew.Dump(genDecl.Doc)
				for _, comment := range genDecl.Doc.List {
					if strings.Contains(comment.Text, "@gigner") {
						for _, spec := range genDecl.Specs {
							if typeSpec, ok := spec.(*ast.TypeSpec); ok {
								if typeSpec.Name != nil {
									fmt.Println(typeSpec.Name.Name)
								}
							}
						}
					}
				}
			}
		}
	}
}

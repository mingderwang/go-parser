package main

import (
	"fmt"
	"go/parser"
	"go/token"
)

type Slack struct {
	Name string
}

var src = `
package main
type Slack struct {
	Name string
}
func main() {
	println("Hello World")
}
`

func main() {
	fset := token.NewFileSet()
	f, err := parser.ParseFile(fset, "", src, 0)
	if err != nil {
		panic(err)
	}
	//ast.Print(fset, f)
	for _, decl := range f.Decls {
		fmt.Println(decl)
	}
}

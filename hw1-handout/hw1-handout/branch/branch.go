package branch

import (
	"go/ast"
	"go/parser"
	"go/token"
	//"fmt"
)
//var pow = []int{1, 2, 4, 8, 16, 32, 64, 128}
/* 
var x int

var gocode =
`
    package example
    import ("fmt")
    func foo(s string) {
        	fmt.Printf("foo: %s\n", s)
    		switch x.(type){
			case int:
				fmt.Println("hey 1")
		}
    }
    func bar(s string) {
        fmt.Printf("bar: %s\n", s)
    }
    func main() {
        foo("blah")
    }
    `
*/
// bacnchCount should count the number of branching statements in the function.
func branchCount(fn *ast.FuncDecl) uint {
	//TODO write this function
	
	var count uint
	count = 0
    ast.Inspect(fn, func (node ast.Node) bool {
        switch node.(type) {
        case *ast.IfStmt:
            count += 1
    	case *ast.ForStmt:
	    count+=1
    	case *ast.RangeStmt:
	    count+=1
        case *ast.SwitchStmt:
	    count+=1	
        case *ast.TypeSwitchStmt:
    	    count+=1
	}

        // If we return true, we keep recursing under this AST node.
        // If we return false, we won't visit anything under this AST node.
        return true
    })
    //fmt.Printf("function calls: %d\n",count)
/*
    
  
	 fset := token.NewFileSet()
    f, err := parser.ParseFile(fset, "gocode.go", gocode, 0)
    if err != nil {
      panic(err)
    }
    ast.Print(fset, f)
*/
    return count
}

// ComputeBranchFactors returns a map from the name of the function in the given
// Go code to the number of branching statements it contains.
func ComputeBranchFactors(src string) map[string]uint {
	fset := token.NewFileSet()
	f, err := parser.ParseFile(fset, "src.go", src, 0)
	if err != nil {
		panic(err)
	}

	m := make(map[string]uint)
	for _, decl := range f.Decls {
		switch fn := decl.(type) {
		case *ast.FuncDecl:
			m[fn.Name.Name] = branchCount(fn)
		}
	}

	return m
}

package rewrite

import (
	"bytes"
	"go/ast"
	"go/format"
	"go/parser"
	"go/token"
	//"github.com/nf/expr"
	//"github.com/nf/simplify"
	"hw1/simplify"
	"hw1/expr"
//	"fmt"
	"strconv" // This package may be helpful...
)

// rewriteCalls should modify the passed AST
func rewriteCalls(node ast.Node) {
	//TODO Write the rewriteCalls function
//	panic("TODO: implement this!")

//testing expr and simplify
/*
	x:=expr.Binary{X:expr.Literal(2),Y:expr.Literal(3),Op:'+'}
	var env expr.Env
	y:=simplify.Simplify(x,env)
	fmt.Println(x)
	fmt.Println(y)
*/
	ast.Inspect(node,func(node ast.Node) bool{
		switch callExpr := node.(type){
		case *ast.CallExpr:
			switch selectorExpr:=callExpr.Fun.(type){
			case *ast.SelectorExpr:
				switch exprIdent:=selectorExpr.X.(type){
				case *ast.Ident:
					if exprIdent.Name == "expr"{
						 parseAndEvalIdent:=selectorExpr.Sel
						 if parseAndEvalIdent.Name == "ParseAndEval"{
							 //parseAndEvalIdent.Name="changed"
							 if len(callExpr.Args)==2{
							 	switch basicLiteral:=callExpr.Args[0].(type){
								case *ast.BasicLit:
									 stringLiteral:=basicLiteral.Kind
									 if(stringLiteral==token.STRING){
									 	//fmt.Println("we are sucessful")
										//basicLiteral.Value="\"3\""
										//next step is to parse the string into an expr.Expr 
										//var input expr.Expr
										//var env expr.Env

										inputString:=basicLiteral.Value[1:len(basicLiteral.Value)-1]
										parsedInput,e:=expr.Parse(inputString)
										if(e==nil){
											var env expr.Env
											simplifiedExpr:=simplify.Simplify(parsedInput,env)
											stringOutput:=expr.Format(simplifiedExpr)
											quotedStringOutput:=strconv.Quote(stringOutput)
											basicLiteral.Value=quotedStringOutput
										}
									}
								}
							 }
						 }
					}
				}
			}
		}
	return true
	})

}

func SimplifyParseAndEval(src string) string {
	fset := token.NewFileSet()
	f, err := parser.ParseFile(fset, "src.go", src, 0)
	if err != nil {
		panic(err)
	}

	rewriteCalls(f)

//	ast.Print(fset,f)
	var buf bytes.Buffer
	format.Node(&buf, fset, f)
//	fmt.Println(buf.String())
	return buf.String()
}

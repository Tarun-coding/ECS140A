package depth

import (
	//"github.com/nf/expr"
	"hw1/expr"
	
)


// Depth should return the maximum number of AST nodes between the root of the
// tree and any leaf (literal or variable) in the tree.
func Depth(e expr.Expr) uint {
	// TODO: implement this function
	//fmt.Printf("type of e: %T\n",e)
	/*
	if l, ok:=e.(expr.Binary);ok{
		fmt.Println("binary expression left ",l.X)
		if z,ok := l.X.(expr.Binary);ok{
			fmt.Println("further binary expression Right ",z.Y)
		}
	}

	return 1234
	*/
	pass:=0
	if e,ok:=e.(expr.Binary);ok{
		e=e
		pass=1
	}
	
	if e,ok:=e.(expr.Literal);ok{
		e=e
		pass=1
	}

	if e,ok:=e.(expr.Unary);ok{
		e=e
		pass=1
	}

	if e,ok:=e.(expr.Var);ok{
		e=e
		pass=1
	}
	if pass == 0{
		panic("oh shit")
	}
	var currentDepth uint
	currentDepth = 1
	var nodesAtDepth []expr.Expr=[]expr.Expr{e}
	//fmt.Printf("type of e: %T\n",nodesAtDepth[0])
	//fmt.Println(len(nodesAtDepth))
	for true{
		var nextDepthNodes []expr.Expr
		for arrayPosition:=0;arrayPosition<len(nodesAtDepth);arrayPosition++{
			if b, ok:=nodesAtDepth[arrayPosition].(expr.Binary);ok{
				nextDepthNodes=append(nextDepthNodes,b.X)
				nextDepthNodes=append(nextDepthNodes,b.Y)
			}
		}
		if len(nextDepthNodes)>0{
			nodesAtDepth=nextDepthNodes
			currentDepth++
		}else{
			//unary check
			for positionIndex:=0;positionIndex<len(nodesAtDepth);positionIndex++{
				if  u,ok:=nodesAtDepth[positionIndex].(expr.Unary);ok{
					e=u
					currentDepth++
				}
			}
			return currentDepth
		}
	}
	return 1234
}

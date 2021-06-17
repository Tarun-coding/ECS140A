package simplify

import (
	"hw1/expr"
	//"github.com/nf/expr"
	//"fmt"
	
)
var result expr.Expr
type BinaryNode struct {
	left *BinaryNode
	right *BinaryNode
	parent *BinaryNode
	data expr.Expr
	operand rune
	unary bool
	numberofTimesVisited int
	firstBinaryNode bool
}
/*
type BinaryTree struct{
	root *BinaryNode
}
*/
/*
func(t *BinaryTree) insert(data expr.Binary) *BinaryTree{
	if t.root == nil{
		t.root = &BinaryNode{left:data.X,right:data.Y,operand:data.Op}
	}
	return t
}
*/
func (n *BinaryNode)recurseUpwards(env expr.Env) {
	//modify the vars to literals
	//fmt.Println("we are in recurseUpward")
	
	//simplification
	if v,ok := n.data.(expr.Var);ok{
		
		if val,okTwo := env[v];okTwo{//v.Eval(env); okTwo{
			//fmt.Println("we are in if")
			n.data=expr.Literal(val)
		}
	}else if u,ok := n.data.(expr.Unary);ok{
		u=u	
		n.data=expr.Unary{Op:n.operand, X:n.left.data}
		//fmt.Println(n.data)
		if l, okTwo := n.left.data.(expr.Literal);okTwo{
		//	if operand == '-'{
		//		n.data=expr.Literal(float64(-1)*float64(n.left.data))
		//	}
			l=l
			n.data=expr.Literal(n.data.Eval(env))
		}
	}else if b,ok := n.data.(expr.Binary);ok{
		//fmt.Println("we should be here 3 times")
		b=b
		n.data=expr.Binary{Op:n.operand,X:n.left.data,Y:n.right.data}
		//fmt.Println(n.data)
		if n.left.data == expr.Literal(0){
			//fmt.Println("does this happen")
			if n.operand == '*'{
				n.data=expr.Literal(0)
			}else if n.operand == '+'{
				n.data = n.right.data
			}
		}else if n.right.data ==expr.Literal(0){
			
		//	fmt.Println("does this happen")
			if n.operand=='*'{
				n.data=expr.Literal(0)
			}else if n.operand=='+'{
				n.data = n.left.data
			}
		}else if n.left.data ==expr.Literal(1) && n.operand=='*'{
			
			//fmt.Println("does this happen")
			if n.operand =='*'{
				n.data=n.right.data
			}
		}else if n.right.data == expr.Literal(1) && n.operand=='*'{
			
			//fmt.Println("does this happen")
			if n.operand == '*'{
				n.data=n.left.data
			}
		}else{
		//	fmt.Println(n.data)
			if lOne,okTwo := n.left.data.(expr.Literal);okTwo{
				//fmt.Println("left is literal")
				lOne=lOne
				if lTwo,okThree := n.right.data.(expr.Literal);okThree{
					
					lTwo=lTwo
					//fmt.Println(n.data)
					n.data=expr.Literal(n.data.Eval(env))
					
				     
				}
		}//else{
//			fmt.Println(n.left.data)
//			fmt.Printf("left is: %T\n",n.left.data)
	//	}
	}

	}
	//fmt.Println("we are at end of if blocks")
	if n.firstBinaryNode{
		//fmt.Println("we are in first binary node")
		//return n.data
		//fmt.Println("first binary node if call")
		
	//	fmt.Println(n.data)
	//	return n.data
		result=n.data
	}else{
		
		n.parent.numberofTimesVisited++
	
		if n.parent.unary{
		//fmt.Println("unary if call")
		//n.parent.data=expr.Unary(Op:n.parent.operand,X:n.parent.left.data)
		n.parent.recurseUpwards(env)
		}else if n.parent.numberofTimesVisited == 2{
		//fmt.Println("number of times visited")
		//	fmt.Println(n.data)
			n.parent.recurseUpwards(env)
		
		}
	}
	
}
func (n *BinaryNode) insert(env expr.Env){
	if c,ok := n.data.(expr.Binary);ok{
	  n.numberofTimesVisited = 0
	  n.unary=false
	  n.operand= c.Op
	  n.left = &BinaryNode{left:nil,right:nil,parent:n,data:c.X,firstBinaryNode:false}
	  n.right = &BinaryNode{left:nil,right:nil,parent:n,data:c.Y,firstBinaryNode:false}
	  n.left.insert(env)
	  n.right.insert(env)
  }else if u,ok := n.data.(expr.Unary);ok{
	//	fmt.Println(n.data)
//		fmt.Println(n.parent.data)
		n.unary=true
		n.numberofTimesVisited =0
		n.operand = u.Op
		n.left = &BinaryNode{left:nil,right:nil,parent:n,data:u.X,firstBinaryNode:false}
		n.left.insert(env)
	}else{
	//	fmt.Println(n.data)
		n.recurseUpwards(env)
	//	fmt.Println(n.data)
	}
}
//next step is to recurse back up
// Simplify should return the simplified expresion
func Simplify(e expr.Expr, env expr.Env) expr.Expr {
	//TODO implement the simplify
//	panic("TODO: implement this!")
	//fmt.Printf("Type: %T",e)
	
	//var x string = ast.Ident(e)
	//fmt.Println(e)
	//x  := expr.Binary{X:e,Y:e,Op:'+'}	
	//if b, ok:=e.(expr.Binary);ok{
	//	fmt.Println(b.X)
	//		}
	
	
//	f:=e.Eval(env)
	//fmt.Println(f)
	//fmt.Printf("Type: %T",f)
	
//	var finalValue expr.Expr
//	finalValue=expr.Literal(f)
	//var x expr.Literal
	//x=-2
	
//	x  := expr.Binary{X:expr.Literal(5),Y:expr.Literal(3),Op:'+'}	
//	x:=expr.Var("X")
	//x:=expr.Unary{Op:'-',X:expr.Literal(2)}	
	
	//x:=expr.Binary{Op:'-',X:expr.Literal(9),Y:expr.Literal(1)}
	//y:=expr.Binary{Op:'-',X:expr.Var("X"),Y:expr.Var("Y")}
	//z:=expr.Binary{Op:'+',X:x,Y:y}
	
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
	var root BinaryNode
	root.data = e
	root.firstBinaryNode=true
	root.insert(env)
	//return finalValue
	//return e
	return result
}

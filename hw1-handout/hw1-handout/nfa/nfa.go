package nfa
//import "fmt"
// A state in the NFA is labeled by a single integer.
type state uint

// TransitionFunction tells us, given a current state and some symbol, which
// other states the NFA can move to.
//
// Deterministic automata have only one possible destination state,
// but we're working with non-deterministic automata.
type TransitionFunction func(st state, act rune) []state

func Reachable(
	// `transitions` tells us what our NFA looks like
	transitions TransitionFunction,
	// `start` and `final` tell us where to start, and where we want to end up
	start, final state,
	// `input` is a (possible empty) list of symbols to apply.
	input []rune,
) bool {
	// TODO Write the Reachable function,
	// return true if the nfa accepts the input and can reach the final state with that input,
	// return false otherwise
	var arr []state
	if len(input)==0{
		arr = append(arr,start)
	}
	for inputCount := 0;inputCount < len(input);inputCount++{
		if len(arr)==0{
			arr=transitions(start,input[inputCount])
			//fmt.Println(len(arr))
		//	fmt.Println(arr)
		//	fmt.Println(len(arr))
		}else{
			var arrTwo []state
			for arrCount := 0;arrCount< len(arr);arrCount++{
				newTransitions:=transitions(arr[arrCount],input[inputCount])
				for transitionCount:=0;transitionCount<len(newTransitions);transitionCount++{
					arrTwo = append(arrTwo,newTransitions[transitionCount])
				}
			}
			arr = arrTwo
		}
	}
	//this is to search for the end state in our final state array
	for finalArrayCount := 0;finalArrayCount < len(arr) ; finalArrayCount++{
		if arr[finalArrayCount]==final{
			return true
		}
	}
	return false
	//fmt.Println("hello")
	//panic("TODO: implement this!")
}

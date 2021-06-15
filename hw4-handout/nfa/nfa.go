package nfa
import "sync"
//import "fmt"
// A state in the NFA is labeled by a single integer.
type state uint

type TransitionFunction func(st state, act rune) []state
var mu sync.Mutex
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
	c:=make(chan []state)

	//concurrent function
	go allFinalStatesAccessable(start,input,arr,c,transitions)
	finalStateArray:=<-c

	for finalArrayCount := 0;finalArrayCount < len(finalStateArray) ; finalArrayCount++{
		if finalStateArray[finalArrayCount]==final{
			return true
		}
	}
	return false
	//for i:= range c{
	//	finalStateArray:=<-c	
	//}

}
func allFinalStatesAccessable(start state,input []rune,arr []state,c chan []state,transitions TransitionFunction){

	if len(input)==0{
		arr = append(arr,start)

	}else{
		for inputCount := 0;inputCount < len(input);inputCount++{
			if len(arr)==0{
				arr=transitions(start,input[inputCount])
			}else{
				var arrTwo []state
				var wg sync.WaitGroup
				for arrCount := 0;arrCount< len(arr);arrCount++{
					wg.Add(1)
					go func(arrCount int){
						newTransitions:=transitions(arr[arrCount],input[inputCount])
						for transitionCount:=0;transitionCount<len(newTransitions);transitionCount++{
							mu.Lock()
							arrTwo = append(arrTwo,newTransitions[transitionCount])
							mu.Unlock()
						}
						wg.Done()
					}(arrCount)

				}
				wg.Wait()
				arr = arrTwo
			}
		}
	}
	c<-arr
}

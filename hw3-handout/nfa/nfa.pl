transitionIteration([],CurrentInput,[]).


transitionIteration([H|TransitionStates],CurrentInput,TransitionResult):-
	transitionIteration(TransitionStates,CurrentInput,NextResult),transition(H,CurrentInput,NextStates),append(NextStates,NextResult,TransitionResult).	

inputIteration(TransitionStates,FinalState,[]):-
	%%write(TransitionStates),
	member(FinalState,TransitionStates).

inputIteration(TransitionStates,FinalState,[H|Input]):-
	transitionIteration(TransitionStates,H,TransitionResult),inputIteration(TransitionResult,FinalState,Input).		

reachable(StartState, FinalState, Input) :-
	inputIteration([StartState],FinalState,Input).
	
	%%Iterate through every input. For every input iterated provide a new array of transitions from the former array(iterating through this aswell perhaps).Finally check for the finalState in the final array 	

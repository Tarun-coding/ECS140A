; You may define helper functions here

(defun transitionArray(transition state input)
  
     
     (funcall transition state (car input))	

)
(defun newStateArray(transition stateArray input)
	;;iterate through each element in state array

    	(cond
	  ((null stateArray) NIL)
	  ((atom (car stateArray)) (append (transitionArray transition (car stateArray) input)(newStateArray transition (cdr stateArray)input))) 
	)   
)
(defun inputRecursion(transition stateArray input)
	;;iterate through each input	  
	;;(print stateArray)	
    	(cond
	  ((null input) stateArray)
	  ((atom (car input)) (inputRecursion transition (newStateArray transition stateArray input) (cdr input))) 
	)   

)
(defun existence (final reachableStates)
	(cond
	  ((null reachableStates) NIL)
	  ((equal (list(car reachableStates)) (list final))t)
	  (t (existence final (cdr reachableStates)))
	
	)
)
(defun reachable (transition start final input)
	(setq firstArrayOfStates (transitionArray transition start input))
	;;(inputRecursion transition firstArrayOfStates (cdr input))

	(cond
	   ((null input) (if (equal (list start)(list final)) t NIL)) 
	  ;;(if (null input) (car xs) (smallest(cdr xs )))  
	   (t(existence final (inputRecursion transition firstArrayOfStates (cdr input))))
	;;fuction to check if the final state exists
	)
)



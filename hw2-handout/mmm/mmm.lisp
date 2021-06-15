; You may define helper functions here
(defun elementsInList(xs)
  
	;;code to get number of elements in the list
    	(cond
	  ((null xs) 0)
	  ((atom (car xs)) (+ 1 (elementsInList(cdr xs)))) 
	)   
)
(defun addAllElements(xs)
	(cond
	  ((null xs) 0)
	  ((atom (car xs)) (+ (car xs) (addAllElements(cdr xs))))
	)
	  
)
(defun largest(xs)
	(cond
	  ((null xs) -10000)
	  ((atom (car xs)) (if (> (car xs) (largest(cdr xs))) (car xs) (largest(cdr xs ))))  
	)
	
)

(defun smallest(xs)
	(cond
	  ((null xs) 100000)
	  ((atom (car xs)) (if (< (car xs) (smallest(cdr xs))) (car xs) (smallest(cdr xs ))))  
	)
	
)
(defun min-mean-max (xs)
    	;;this is for calculating the mean
	(setq a (/ (addAllElements xs)(elementsInList xs)))
	;;finding biggest number
	(setq b (largest xs))
	(setq c (smallest xs))
	 (list c a b)
)

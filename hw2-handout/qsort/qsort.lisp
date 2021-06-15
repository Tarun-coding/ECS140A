
;;You may define helper functions here

(defun allNumbersGreater (n xs)
		  
	(cond
	  ((null xs) NIL)
	  ((atom (car xs)) (if (>= (car xs) n) (append (list (car xs)) (allNumbersGreater n (cdr xs))) (allNumbersGreater n (cdr xs ))))  
	)
)
(defun allNumbersLess (n xs)
	
	(cond
	  ((null xs) NIL)
	  ((atom (car xs)) (if (< (car xs) n) (append (list (car xs)) (allNumbersLess n (cdr xs))) (allNumbersLess n (cdr xs ))))  
	)
)
(defun pivot (n xs)
  ;; TODO: Incomplete function.
  ;; The next line should not be in your solution.
  
;;  (list 'incomplete)

;;append(NIL '(a b))
;;(append '(x) NIL)
;;(list allNumbersGreater n xs)

;;(print (allNumbersGreater n xs))

;;(print (allNumbersLess n xs))
(list (allNumbersLess n xs) (allNumbersGreater n xs))
)
(defun algo(xs)
	
	
	(cond
	  ((null xs) NIL)
	  ((atom (car xs)) (append (algo (allNumbersLess (car xs)(cdr xs))) (list (car xs)) (algo(allNumbersGreater (car xs) (cdr xs)))))  
	)
 )
(defun quicksort (xs)
  ;;(list 'incomplete)
	;;car xs is the pivot
	(algo xs)	
)
;;(print((pivot (3 '(a b)))))
;;(print (pivot 3 '(a b)))
;;(print (allNumbersGreater 3 '(2 3)))

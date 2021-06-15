; You may define helper functions here
(defun addElementRecursor(rowOne rowTwo)
	;;this has to recurse through each element and add them
	(cond
	  ((null rowOne) NIL)
	  (t(append (list (+ (car rowOne) (car rowTwo))) (addElementRecursor (cdr rowOne)(cdr rowTwo))))
	)
)
(defun addRowRecursor(mat1 mat2)
	
  
    	(cond
	  ((null mat1)NIL )
	  (t (cons (addElementRecursor (car mat1)(car mat2)) (addRowRecursor(cdr mat1)(cdr mat2)))) 
	)   
)
(defun recurseElementRecursor(mat)
	(cond 
 	  ((null (car(car mat)))NIL)
	  (t (append (list(car(car mat))) (recurseElementRecursor(cdr mat))))
	)
)
(defun createNewMatrix(mat)
	(cond
	  ((null(car mat))NIL)
	  (t(cons (cdr (car mat)) (createNewMatrix (cdr mat))))
	  )
)
(defun transposing(mat)
	(cond
	  ((null(car mat))NIL)
	  (t (cons (recurseElementRecursor mat) (transposing (createNewMatrix mat))))
)
	)
(defun eachRowMultiply(row1 row2)
	  (cond
	    ((null (car row1))0)
	    (t (+(* (car row1)(car row2))(eachRowMultiply (cdr row1)(cdr row2))))
	  )
)
(defun eachColMultiply(row1 mat2)
  	(cond
	  ((null (car mat2))NIL)
	  (t (append (list (eachRowMultiply row1 (car mat2))) (eachColMultiply row1 (cdr mat2))))
	  )
 )
(defun multiplyIteratorRow(mat1 mat2)
	(cond
	  ((null(car mat1))NIL)
	  (t (cons (eachColMultiply (car mat1) mat2) (multiplyIteratorRow(cdr mat1)mat2)))
	)
)
(defun matrix-add (mat1 mat2)
  ;; TODO: Incomplete function.
  ;; The next line should not be in your solution.
;;  (list 'incomplete)
	(addRowRecursor mat1 mat2)
)

(defun matrix-transpose (mat)
  ;; TODO: Incomplete function.
  ;; The next line should not be in your solution.
 ;; (list 'incomplete)
;;iterate through each row in the matrix and take its first element
	;;(recurseElementRecursor mat)
	(transposing mat)	
)

(defun matrix-multiply (mat1 mat2)
  ;; TODO: Incomplete function.
  ;; The next line should not be in your solution.
  ;;(list 'incomplete)
  ;;find the transpose of mat2: the number of rows of the transpose=number of cols of result
  ;;multiply each row from mat1 with each row in mat2. append to same list until iterated through all rows of mat 2. and call cons at the end
  (multiplyIteratorRow mat1 (transposing mat2))
)
;;(print(matrix-add '((1 2)(2 3)) '((3 4)(5 6))))
;;(print(recurseElementRecursor(createNewMatrix '((1 2 3)(4 5 6)))))

;;(print(transposing '((3 1 5)(6 9 7))))
;;(print(eachColMultiply '(1 2) '((3 4)(5 6))))

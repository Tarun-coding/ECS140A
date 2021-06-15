; You may define helper functions here

(defun questionMarks (pattern assertion)
	(cond
	  ((null pattern) NIL)
	  ((null assertion) '(NIL))
	  
	  ((atom (car assertion)) (if (equal (list(car pattern)) '(?)) (append (list (car assertion)) (questionMarks (cdr pattern)(cdr assertion)))
									    
	  		 		(append (list (car pattern)) (questionMarks (cdr pattern)(cdr assertion)))
				    )
	  )
	)
  
  )
(defun withinExclamation(pattern assertion)
	
  	(setq myAssertion assertion)
	(setq myPattern pattern)
	;;(print  pattern)
	(cond
	  ;;((null pattern) NIL)
	  ;;((null assertion) '(NIL))
	  ((null assertion) NIL)
	  ((equal (list(car pattern)) '(?))(append(list(car assertion))(withinExclamation(cdr pattern)(cdr assertion))) )

	  
	  ((equal (list(car pattern)) '(!))(append(list(car assertion))(withinExclamation(cdr pattern)(cdr assertion))) )
	  (t (if (equal (car pattern) (car assertion)) NIL
									    
	  		 		(append (list (car assertion)) (withinExclamation pattern (cdr assertion)))
				    )
	  )
	)

)
(defun exclamationMarks(pattern assertion)
 	;; (print myAssertion)
	(cond
	  ((null pattern) nil)
	  ((null assertion) '(NIL))
	  
	  ((equal (list(car pattern)) '(!)) (append (list(car assertion))(withinExclamation (cdr pattern) (cdr assertion)) (exclamationMarks  myPattern myAssertion)))
	  
	  ((equal (list(car pattern)) '(?)) (append (list (car assertion)) (exclamationMarks (cdr pattern) (cdr assertion))))
	  (t (append (list (car pattern)) (exclamationMarks (cdr pattern)(cdr assertion))))
	)

  
)
(defun match (pattern assertion)
;;	(print (exclamationMarks pattern assertion))
	(equal (exclamationMarks pattern assertion) assertion)
	
)

;;(print(exclamationMarks '(! B ? B A B A ! ? !) '(B B A B A B A B A B A)))

;;(print(exclamationMarks '(! ?) '(B B A B A B A B A B A)))
;;(print(equal (exclamationMarks '(!) '(a a NIL)) '(a a NIL)))

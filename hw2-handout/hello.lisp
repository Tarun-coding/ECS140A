;function that counts the number of lists
;(setq a '(A (G G) B (F F)))

(defun numberOfLists(x)
	(cond
	  ((null x) 0)
	  ((listp (car x))(+ 1 (numberOfLists(cdr x))))
	  (+ 0 (numberOfLists(cdr x)))
	  )

)
;(print (numberOfLists a))
;(print (listp(zerop 1)))
;(print (list 5))
;(print
;	(let* ((a 1) (b 4))
;	  (let ((a 3) (b (+ a b)))
;	  (+ a b)))

 ; )
;(setq a (cons 3 ''(+ 5 6)))
;(print(eval(cdr a)))
;(print (cons 1 (cons 2 nil)))
;(print (apply (lambda (m n) (+ m n)) '(3 2))
 ;      )

;(print(maplist #'(lambda (x y) (mapcar  #'- x y)) '(1 3 8) '(0 1 5)))
 
 ; (let ((a 10))
  ;  (defun foo(b c)
   ;   (setq a(- a (+ b c)))))
;(print (list(foo 2 4) (foo 4 6)))  

;(defun my-sequence (seed)
 ; (function     
  ;   (lambda () (setq seed (+ seed 2)))))

;(setq seq1 (my-sequence 1))
;(setq seq2 (my-sequence -2))
;(print(list (funcall seq1) (funcall seq1) (funcall seq2) (funcall seq1)))  
 
;(print  (car(cdr ( cdr '(((1 3) 5 7) 9 (11 13))))) 
;	)
(print (cons 1 (cons 2 (cons 4 nil))))
  

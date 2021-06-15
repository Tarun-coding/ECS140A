/* All novels published during the year 1953 or 1996*/
year_1953_1996_novels(Book) :-
    %% remove fail and add body/other cases for this predicate
    %%fail.
    novel(Book,1953).
    %%novel(Book,1996).	
year_1953_1996_novels(Book) :-
    %% remove fail and add body/other cases for this predicate
    %%fail.
    novel(Book,1996).
/* List of all novels published during the period 1800 to 1900*/
period_1800_1900_novels(Book) :-
    %% remove fail and add body/other cases for this predicate
    %%fail.
    novel(Book,X),X>=1800,X=<1900.

/* Characters who are fans of LOTR */
%%member(E,[E|_]).
%%member(E,[_|R]):- member(E,R).
memberTwo([A|_],B):- member(A,B).
memberTwo([_|A],B):- memberTwo(A,B).
lotr_fans(Fan) :-
    %% remove fail and add body/other cases for this predicate
    %%fail.
    fan(Fan,X),member(the_lord_of_the_rings,X).

/* Authors of the novels owned by Chandler */
author_names(Author) :-
	%%%% re
    %%move fail and add body/other cases for this predicate
    %%so iterate through each element of authors books and iterate through chandler's liked books
    %%X=favbooklist, Y=author's booklist
    	fan(chandler,X),author(Author,Y),memberTwo(X,Y).
    
    %%fail.

/* Characters who are fans of Brandon Sanderson's novels */
fans_names(Fan) :-
    %% remove fail and add body/other cases for this predicate
    %%fail.
    	author(brandon_sanderson,X),fan(Fan,Y),memberTwo(X,Y).

/* Novels common between either of Phoebe, Ross, and Monica */
mutual_novels(Book) :-
    %% remove fail and add body/other cases for this predicate
    %%fail.
    	fan(phoebe,X),fan(ross,Y),member(Book,X),member(Book,Y).

mutual_novels(Book):-

    	fan(phoebe,X),fan(monica,Y),member(Book,X),member(Book,Y).

mutual_novels(Book):-

	fan(ross,X),fan(monica,Y),member(Book,X),member(Book,Y).

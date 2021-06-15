union(A,[],C):-append(A,[],C).
union([],B,C):-append([],B,C).
union(A,[H|B],C):- member(H,A),union(A,B,C).
union(A,[H|B],C):-append(A,[H],RT),union(RT,B,C).
%%memberTwo(A,B,C):-append(B,A,C).
%%memberSet([],B).
%%memberSet([A|C],B):- member(A,B),memberSet(C,B).
intersection(A,[],C):-append([],[],C).
intersection([],B,C):-append([],[],C).
intersection(A,[H|B],C):-member(H,A),append([H],RT,C),intersection(A,B,RT).
intersection(A,[H|B],C):-intersection(A,B,C).

equalLeft([],B).
%%equalLeft(A,[]).
equalLeft([H|A],B):-member(H,B),equalLeft(A,B).

%%equalRight([],B).
equalRight(B,[]).
equalRight(A,[H|B]):-member(H,A),equalRight(A,B).
isUnion(Set1,Set2,Union) :-
    %% remove fail and add body/other cases for this predicate
    %%fail.
    %%append(Set1,Set2,Union).
    %%check each element of set 1 and set 2 are members of union
    %%memberSet(Set1,Union),memberSet(Set2,Union).
    %%fail.
    %%memberSet(Set1,Union).
    %%append(Set1,Set2,Union).
   union(Set1,Set2,Union).
    %%true.
	

isIntersection(Set1,Set2,Intersection) :-
    %% remove fail and add body/other cases for this predicate
    %%fail.
    intersection(Set1,Set2,Intersection).

isEqual(Set1,Set2) :-
    %% remove fail and add body/other cases for this predicate
    %%every element from set1 is a member of set2, every element from set 2 is a member of set1
    equalLeft(Set1,Set2),equalRight(Set1,Set2).

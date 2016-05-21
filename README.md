# go-fundamentals  
![Golang Gopher](https://golang.org/doc/gopher/bumper640x360.png)  

NOTE: for single file use and demo purposes, I have multiple main files because this is not a conventional package. I am aware that a go package can only have one main Func.

Learnings from going through the go-fundamentals course by Nigel Poulton

##Variables  
-  Uninitailized variables get a Zero value
-  Package level variables are global
-  Package-level variabes must be declared with var
-  Declare function variables with :=
-  Go supports type inference
-  Go passes arguments by value
- '&' References a pointer
- '*' De-References a pointer
- Go Passes arguments by VALUE 


##Functions  
- Functions make your code concise,re-usable, easy to test & debug
- all statements within a function is enclosed within {}
-  Params and returns are defined within the func signature
-  The return keyword exits a function 
- Go functions can return multiple values
- Variadic functions - functions that can have an unknown number of paramters

## Conditionsals
- 'If' conditions evaulate a Boolean expression for true or false
- Curly brance must start on the same line of the if
- No Parentheses around the if condition.
-  Simple initlaization Statements in if blocks. This allows block scoped variables


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
-  `*` De-References a pointer
- Go Passes arguments by VALUE


##Functions  
- Functions make your code concise,re-usable, easy to test & debug
- all statements within a function is enclosed within {}
-  Params and returns are defined within the func signature
-  The return keyword exits a function
- Go functions can return multiple values
- Variadic functions - functions that can have an unknown nu
mber of paramters


## Conditionsals  
- 'If' conditions evaulate a Boolean expression for true or false
- Curly brance must start on the same line of the if,else if, and else statements
- No Parentheses around the if condition.
-  Simple initlaization Statements in if blocks. This allows block scoped variables
- `Switch` statements have an optional simple statement and an expression to be evaulated.
-  Switches have cases to be evaluated and a default statement at the end
- A switches case statement has an implicit break statements
- Switches have a fallthrough keyword to run case blocks underneath it. this is case by case basis meaning that a fallthrough keyword only falls through to the next if block
- Switch statements can more than one value for a case
- In Go, it is expected to have an error as the last return from functions and emethods
- In errrors, nil is used to indicate success
- It is expected to check returned err variables

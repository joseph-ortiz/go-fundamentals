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


## Conditionals  
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
- In errors, nil is used to indicate success
- It is best pratice to check returned err variables

##Loops
- has one keyword for loops. a "for" loop
- for with a blank expression = infinite loop
- boolean expression
- for...range
- Pre and Post statements
- Nested
- Continue and Break
- Go has a break to a label feature

##Arrays and Slices
- Arrays are a numbered list of a single type and has a fixed length
- Slices are also a numbered list of a single type but can be resized.
- Slices are built on top of an Array.
- Slices are references that pass portions of an array.
- Data is never slices in an array!
- Slices are left side inclusive. The right side is excluded.
- an array is sliced with [0:4], you will get the value at index 0,1,2,3 and 3.
- when appending to a slice, once we hit capacity, the underlying array is doubled. Thus we are doubling the capacity.
- Refrencing a slice by it's variable name, it will show the entire slice
- for range loops iterate slices
- for range returns two values, an index and data
- Can append() slices to slices with ellippses
- Use slices instead of an array for flexibility

##Maps
- Unordered list
- Similar to slices and pointers
- Maps are dynamic
- Maps are references
- Needs keytype and val type when creating a map.
- Passed to functions by reference
- Changes made by functions visible caller
- Unsafe for concurrency
- Cheapt to pass around.
- Specify size for large maps.
- make(map<keyType><valueType>, size) 
- Can improve peformance

##Structs
- Go's way of defining Custom data types.
- Go has no objects,classes, nor inheritance.

##Concurrency
- Go's concurrency model
- Communication Sequential Processes (CSP)
- Uses channels for safe communication
- What is Concurrency? Creating multiple "processes" that execute INDEPENDENTLY.
>"... concurrency is the _composition_ of  
>independently executing processes, while  
>parallelism is the simultaneous _execution_ of  
>(possibly related) computations. Concurrency is  
>about _dealing_ with lots of things at once.  
>Parallelism is about _doing_ lots of things at once."  
>-- Rob Pike
- Sending multiple emails while waiting for a response.
- Go's doesn't touch threads directly
- goroutines handle threads for us.
- goroutines vs os Threads
- + multiple goroutines on a single thread
- + light weight. (kb in size).
- + Go manages goroutines.(simpler for programmers)
- + Less switching
- + Faster start-up times
- + Safe communication 

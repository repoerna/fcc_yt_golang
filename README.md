# fcc_yt_golang
Learn Go Programming - Golang Tutorial for Beginners

# Go (Golang)
- Strong and sttically typed
    - type of variable cant changing over time
- Excellent community
- Key Features
    - Simplicity
    - Fast compile times
    - Garbage collected
    - Built-in concurrency
    - Compile to stand alone binaries
        - all dependencies is there
    

# Resources
- golang.org
    - doc: for refrence
    - effective go: best use of go
    - pkg: list all of go package
    - help: mliling list, community, gopher slack

- golangbridge.org
community that aadvocate go developer

- play.golang.org
golang sandbox


# Go Program Stack
- Package
- Import 
- Main function
    entry point of thr program


# Setting Local Development 
- installation instructor
    https://golang.org/doc/install


# Setup Editor
VSCode
    - Extension


# go build
    - take the actual package path


# go install


# Into the Programming
## Variable
- variable declaration
- Redeclaration and shadowing
- Visibility
- Naming Convention
- Type Convertion


## Primitive Types
- Boolean Type
- Numeric Type
    - integers
    - Floating Point
    - Complex numbers
- Text Type

## Constant
- Naming Convention
- Type constant
- Untype constant
- Enumerate constant
- Enumeration expressions

## Array and Slices
- Array
    - Creation
    - Built in function
    - Working with arrays

- Slice
    - Creation
    - Built in function
    - Working eith function

## Map and Struct
- Map
    - What they are?
    - Creating
    - Manipulation

- Struct
    - What they are?
    - Creating
    - Naming conventions
    - Embedding
    - Tags

## Control Flow
- If Statement
    - Operator
    - If - else and if -else if statement

- Switch statement
    - Simple cases
    - Cases with multiple test
    - Falling through
    - Type Switches

## Looping
- For statement
    - simple loops
    - exiting early
    - looping trough colections 

## Control Flow - Defer, Panic, Recover
- Defer
    Invoke a function but delay the execution
- Panic
    make go program in state that cant continue to run,
    and how go runtime can triger that, and triger it on our own
- Recover
    when panic occur, recover can be used to saved program to not bailed out completely

## Pointer 
- Crerating pointer
- Derefrencing pointers
- the new function
- working with nil
- type with internal pointer
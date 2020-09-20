package main

// import (
// 	"fmt"
// )

// func main() {
// 	fmt.Println(`
// 	***********************************
// 	-------------FUNCTION--------------
// 	***********************************
// 	`)

// 	fmt.Println("========== Basic Syntax ==========")
// 	// * same with this entry point
// 	// * <func keyword> <func name> () {}
// 	// * func main () {}
// 	// * name can use PascalCase or camelCase
// 	// * opening curly braces must be inline with func keyword

// 	sayMessage("test")
// 	sayMessage2("test2", 2)

// 	fmt.Println("========== call function in loop ==========")
// 	// * further implementation of function
// 	for i := 0; i < 5; i++ {
// 		sayMessage2("in loop", i)
// 	}

// 	fmt.Println("========== function with more than 1 parameter ==========")
// 	sayGreeting("Hello", "Purna")
// 	sayGreeting2("Hello", "Purna", 1, 2)

// 	fmt.Println("========== create function to passing pointer ==========")
// 	name := "Stacey"
// 	greeting := "Hello"
// 	sayGreeting3(greeting, name)
// 	sayGreeting4(&greeting, &name) // * pass the address of variable
// 	fmt.Println(name)              // check if value is changed

// 	fmt.Println("========== create function using periodic parameter ==========")
// 	sum("The sum is", 1, 2, 3, 4, 5)

// 	fmt.Println("========== return value ==========")
// 	s := sum2(1, 2, 3, 4, 5)
// 	fmt.Println("The sum is", s)

// 	fmt.Println("========== return value with pointer ==========")
// 	ss := sum3(1, 2, 3, 4, 5)
// 	fmt.Println("The sum is", *ss)

// 	fmt.Println("========== using name return ==========")
// 	sss := sum4(1, 2, 3, 4, 5)
// 	fmt.Println("The sum is", sss)

// 	fmt.Println("========== multiple output function ==========")
// 	d, err := divide(5.0, 1.0)
// 	if err != nil {
// 		fmt.Println(err)
// 		return
// 	}
// 	fmt.Println(d)

// 	fmt.Println("========== Anonymous Function ==========")
// 	func() {
// 		fmt.Println("this is from anonymous function")
// 	}() // * end parenthesis is need to call the function / it will invoke the function

// 	for i := 0; i < 5; i++ {
// 		func() {
// 			fmt.Println(i)
// 		}()
// 	}

// 	for i := 0; i < 5; i++ {
// 		func(i int) {
// 			fmt.Println(i)
// 		}(i)
// 	}

// 	fmt.Println("========== function as type ==========")

// 	f := func() {
// 		fmt.Println("variable f")
// 	}

// 	f()

// 	var ff func() = func() {
// 		fmt.Println("type signature function")
// 	}

// 	ff()

// 	// var divide2 func(float64 float64) (float64, error)
// 	divide2 := func(a, b float64) (float64, error) {
// 		if b == 0.0 {
// 			return 0.0, fmt.Errorf("Cannot divide by zero")
// 		}
// 		return a / b, nil
// 	}

// 	dd, err := divide2(5.0, 2.0)
// 	if err != nil {
// 		fmt.Println(err)
// 		return
// 	}
// 	fmt.Println(dd)

// 	fmt.Println("========== method ==========")
// 	g := greeter{
// 		greeting: "Hallo",
// 		name:     "Ola",
// 	}
// 	g.greet()
// 	g.greet2()
// 	fmt.Println(g.name)

// }

// // TODO: basic with parameter
// // * basic func with paramater
// // * parameter --> (<variable> <type>)
// func sayMessage(msg string) {
// 	fmt.Println(msg)
// }

// // TODO: with 2 parameter
// // * basic func with 2 paramater
// // * parameter --> (<variable> <type>)
// func sayMessage2(msg string, idx int) {
// 	fmt.Println(msg)
// 	fmt.Println("The value of index is", idx)
// }

// func sayGreeting(greeting string, name string) {
// 	fmt.Println(greeting, name)
// }

// // TODO: with 2 parameter with syntatic sugar
// // * in sayGreeting function, it have 2 parameter with same type
// // * we can use go syntatic sugar
// func sayGreeting2(greeting, name string, int1, int2 int) { // * it means every
// 	// * variable in comma delimiterwill have same type
// 	fmt.Println(greeting, name, int1, int2)
// }

// // TODO: using pointer
// func sayGreeting3(greeting, name string) {
// 	fmt.Println(greeting, name)
// 	// ? what if we change the variable value inside the function
// 	name = "Ted"
// 	fmt.Println(name)
// }

// // ? what if we use pointer
// func sayGreeting4(greeting, name *string) {
// 	fmt.Println(*greeting, *name) // * dereference the addres
// 	*name = "Ted"
// 	fmt.Println(*name)
// }

// // * using this, will change variable not only in scope of function
// // * but on the calling scope as well
// // ? why we need this
// // * ussualy function do need add on parameter that passed in
// // * passing pointer more efficient / for performance benefit

// // TODO: usong periodic parameter '...'
// func sum(msg string, values ...int) {
// 	// * we can use only one periodic parameter, and place it in the end
// 	fmt.Println(values)
// 	result := 0
// 	for _, v := range values {
// 		result += v
// 	}
// 	fmt.Println(msg, result)
// }

// // TODO: return value
// // * to do this we must add type of return value
// // * func <name func> (<variable name> <variable type) <return type> {}
// func sum2(values ...int) int {
// 	// * we can use only one periodic parameter, and place it in the end
// 	fmt.Println(values)
// 	result := 0
// 	for _, v := range values {
// 		result += v
// 	}
// 	return result
// }

// // TODO: return value using pointer
// // * to do this we must add type of return value
// // * func <name func> (<variable name> <variable type) <return type> {}
// func sum3(values ...int) *int {
// 	// * we can use only one periodic parameter, and place it in the end
// 	fmt.Println(values)
// 	result := 0
// 	for _, v := range values {
// 		result += v
// 	}
// 	return &result
// }

// // TODO: using name return variable and using pointer
// // * to do this we must add type of return value
// // * func <name func> (<variable name> <variable type) <return name return type> {}
// func sum4(values ...int) (result int) {
// 	// * we can use only one periodic parameter, and place it in the end
// 	fmt.Println(values)
// 	for _, v := range values {
// 		result += v
// 	}
// 	return
// }

// // * this method not much used

// // TODO: Multiple output
// func divide(a, b float64) (float64, error) {
// 	if b == 0.0 {
// 		return 0.0, fmt.Errorf("Cannot divide by zero")
// 	}
// 	return a / b, nil
// }

// // TODO: Method
// type greeter struct {
// 	greeting string
// 	name     string
// }

// func (g greeter) greet() {
// 	fmt.Println(g.greeting, g.name)
// }

// // * it really same with function
// // * method is function that executing in a noun/object contex
// // * noun is any type, int, struct, or anything
// // * func (<value receiver) <method name>()
// // * value receiver is the reiceive object
// func (g *greeter) greet2() {
// 	fmt.Println(g.greeting, g.name)
// 	g.name = "aaaaaa"
// }

// * Functions
// *	Basic Sytnax
// *		func foo() {
// *			...
// *		}

// *	Parameters
// *		Comma deimited list of variable and type
// *			func foo(bar string, baz int)
// *		Paramters of same type list type once
// *			func foo (bar, baz int)
// *		When pointer are passed in, the function can change the value in the caller
// *			this is always true for data of slice and map
// *		Use variadic parameters to send list of same type in
// *			Must be last parameter
// *			Received as a slice
// *			func foo(bar string, baz ..int)

// *	Return Value
// *		Single return value just list type
// *			func foo() int
// *		Multiple return value list type surrounded by parenthese
// *			func foo() (int, error)
// *			The (result type, error) paradigm is a very common idiom in go
// *		Can use named return values
// *			Initializez return value
// *			Return using 'return' keyword on its own
// *		Can return the address of local variable
// *			Automatically promoted from local memory (stack) to shared memory (heap)

// *	Anonymous Function
// *		Functions dont have name if they are:
// *			Immediately invoked
// *				func() {
// *					...
// *				}()
// *		Assigned to a variable or passed as an argument to a function
// *			a := func() {
// *				...
// *			}
// *			a()

// *	Function as type
// *		Can assign function to variable or use as arguments and return value in functions
// *		Type Signature s like function signature. with no parameter names
// *			var f func(string, string, int) (int, error)

// *	Methods
// *		Function that execute in contex of a type
// *		Format
// *			func (g greeter) greer() {
// *				...
// *			}
// *		Receiver can be value or pointer
// *			Value receiver gets copy of type
// *			Pointer receiver gets pointer to type

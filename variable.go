package main

// // * naming convention
// /*
// 	# Variable Scope
// 		lowercase - uppercase (scope)
// 		lowercase first letter = locally in this package
// 		uppercase first letter = exported to outside the package (globally)

// 		block scope
// 			variable that declared inside block (ex main block)
// 			can be use only in thatt block

// 		no private scope
// 			instead use block scope

// 	# Variable Name
// 		- keep it simple and short
// 		ex: variable in for loop on will eventually used in that
// 		so no need to give it a complex name
// 		if we need to declare variable that will alway be used for
// 		a long time we can use a longer name to rember it

// 		- Pascal or camelCase

// 	# Variable Acronym
// 		ex var theURL, HTTPRequest
// 		always use UPPERCASE for the acronym for good readabillty
// */

// // * declaring variable in pkg level
// // ! we cant use colon-equal syntax
// // ! it need use full syntax
// var x int = 60

// //  * Declare Variable with block of code
// var (
// 	name      string = "Purna"
// 	age       int    = 26
// 	pekerjaan string = "coder"
// )

// func main() {
// 	// * Declare a Variable
// 	// * Type 1
// 	// ? When to use this?
// 	/*
// 		there are gonna be time that we want to declare variable but not
// 		but not ready to initialize it yet
// 		ex: in loop that we need to create local variable
// 	*/
// 	var i int
// 	// * How to read: variable 'i' with type 'int'
// 	i = 43
// 	i = 27 // 'i' value will change
// 	fmt.Println("type 1, i:", i)

// 	// * Type 2
// 	// ? When to use this?
// 	/*
// 		- if we want assaingn variable with specific type
// 	*/
// 	var j float32 = 40
// 	fmt.Println("type 2, j:", j)
// 	fmt.Printf("%v, %T", j, j)
// 	fmt.Println("")

// 	// * Type 3 - colon equal syntax
// 	// ? When to use this?
// 	/*
// 		When to use this?
// 		in this type we can't set  specific type
// 		go will determine it self
// 	*/
// 	k := 99.44
// 	fmt.Println("type 3, k:", k)
// 	fmt.Printf("%v, %T", k, k)
// 	fmt.Println("")

// 	// * PKG Level Result
// 	fmt.Println("pkg level, x:", x)

// 	// * Redeclaration
// 	// ! we can use this type to redeclare variable
// 	// i := 4 <--- example
// 	//  instead just use the equal
// 	i = 10
// 	fmt.Println("i:", i) // this call shadowing
// 	// * shadowing: variable with innermost scope actually takes precedence

// 	// * type convertion
// 	//  destinationType(variable)
// 	var a int
// 	a = 10
// 	fmt.Println("---- 1st type ----")
// 	fmt.Printf("%v, %T", a, a)
// 	fmt.Println("")

// 	var b float32
// 	b = float32(a) // conversion
// 	fmt.Println("---- 2nd type ----")
// 	fmt.Printf("%v, %T", b, b)

// 	// ! Remember if you want convert type
// 	// ! float32 to int will cause loosing information

// 	// * if we want convert int to string wee need to use pkg "strconv"
// 	var c int
// 	c = 42
// 	fmt.Println("---- 1st type ----")
// 	fmt.Printf("%v, %T", c, c)
// 	fmt.Println("")

// 	var d string
// 	d = string(a) // conversion
// 	fmt.Println("---- 2nd type ----")
// 	fmt.Printf("%v, %T", d, d)
// 	fmt.Println("")

// 	var e string
// 	e = strconv.Itoa(a) // conversion
// 	fmt.Println("---- 3rd type ----")
// 	fmt.Printf("%v, %T", e, e)
// }

// // ! variable in go always have to be used

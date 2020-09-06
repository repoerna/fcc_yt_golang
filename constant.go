package main

// import (
// 	"fmt"
// )

// const myConst3 int = 20

// // enumerated
// const a = iota // <--- untyped const
// // * iota is counter that we can use for create enumerated constant
// // * we can use it using const block, if we have 2 const block for iota, value will start with 0 in the secon block
// const (
// 	b = iota
// 	c = iota
// 	d = iota

// 	// * it also can write like below
// 	/*
// 		b = iota
// 		c
// 		d
// 	*/
// )

// // example use of iota / enumerated constant
// const (
// 	errorSpecialist = iota
// 	// _ = iota // --> write only variable, it wont asign memory (blank identifier)
// 	catSpecialist
// 	dogSpecialist
// 	snakeSpecialist
// )

// const (
// 	_  = iota             // ignore first value by assigning to blank identifier
// 	KB = 1 << (10 * iota) // exported KB
// 	MB
// 	GB
// 	TB
// 	PB
// 	EB
// 	ZB
// 	YB
// ) // size

// const (
// 	isAdmin = 1 << iota
// 	isHeadquarters
// 	canSeeFinancials

// 	canSeeAfrica
// 	canSeeAsia
// 	canSeeEurope
// 	canSeeNorthAmerica
// 	canSeeSouthAmerica
// )

// func main() {
// 	fmt.Println(`
// 	*************************
// 	------  CONTSTANT  ------
// 	*************************
// 	`)

// 	fmt.Println("========== NAMING CONVENTION ==========")
// 	// * start with 'const'
// 	const myConst int = 42
// 	// beware for lowercase and uppercase

// 	fmt.Println("========== TYPED CONSTANT ==========")
// 	const myConst0 int = 42
// 	// myConst0 = 20 // ! wont work, const is immutable
// 	fmt.Printf("%v, %T\n", myConst, myConst)
// 	// * const must be assignable at the compile time
// 	// const myConst2 float32 = math.Sin(1.5) // ! wont work, math is not constant, cant assigned when compiled
// 	// constant can have all primitive

// 	// * constant can be shadowing
// 	// ! shadowing have different package level
// 	const myConst3 int16 = 10
// 	fmt.Printf("%v, %T\n", myConst3, myConst3)

// 	fmt.Println("========== UNTYPED CONSTANT ==========")
// 	// * same with variable we can use it for math operation, but if its typed constant and have different type, we cant do it

// 	// ! different type constant
// 	// * the operation will result error
// 	// const myConst4 int16 = 42
// 	// const myConst5 int = 27
// 	// fmt.Printf("%v, %T\n", myConst4+myConst5, myConst4+myConst5)

// 	// * using typed and untyped constant
// 	const myConst4 = 42 // <--- untyped constant
// 	const myConst5 int = 27
// 	fmt.Printf("%v, %T\n", myConst4+myConst5, myConst4+myConst5)
// 	// it can be done, because compiler assign the type automatially
// 	//  program see it as 42 + myConst5
// 	//  its called implicit conversion

// 	fmt.Println("========== ENUMERATED CONSTANT ==========")
// 	// * ussually applied in package level -- see variable 'a'
// 	fmt.Printf("%v, %T\n", a, a)

// 	fmt.Printf("%v\n", b)
// 	fmt.Printf("%v\n", c)
// 	fmt.Printf("%v\n", d)

// 	var specialistType int = dogSpecialist
// 	fmt.Printf("%v\n", specialistType == dogSpecialist)

// 	specialistType = catSpecialist
// 	fmt.Printf("%v\n", specialistType == catSpecialist)

// 	fmt.Println("========== ENUMERATION EXPRESION ==========")
// 	fmt.Printf("%v\n", KB)
// 	fmt.Printf("%v\n", MB)
// 	fmt.Printf("%v\n", GB)
// 	fmt.Printf("%v\n", TB)
// 	fmt.Printf("%v\n", PB)
// 	fmt.Printf("%v\n", EB)
// 	// fmt.Printf("%v\n", ZB)
// 	// fmt.Printf("%v\n", YB)

// 	fileSize := 4000000000.
// 	fmt.Printf("%.2fGB\n", fileSize/GB)

// 	fmt.Printf("%v\n", isAdmin)
// 	fmt.Printf("%v\n", isHeadquarters)
// 	fmt.Printf("%v\n", canSeeFinancials)
// 	fmt.Printf("%v\n", canSeeAfrica)
// 	fmt.Printf("%v\n", canSeeAsia)
// 	fmt.Printf("%v\n", canSeeEurope)
// 	fmt.Printf("%v\n", canSeeNorthAmerica)
// 	fmt.Printf("%v\n", canSeeSouthAmerica)

// 	var roles byte = isAdmin | canSeeFinancials | canSeeEurope
// 	fmt.Printf("%b\n", roles)
// 	fmt.Printf("is admin? %v\n", isAdmin&roles == isAdmin)
// 	fmt.Printf("is HQ? %v\n", isHeadquarters&roles == isHeadquarters)
// 	const test byte = 1 << 0
// 	fmt.Printf("%b\n", test)

// }

// // * immutable, but can be shadowed

// // * replaceable by the compiler at compile time
// // *	value must bu calculated at compiled time

// // * name like variable
// // * 	PascalCase for exported constant
// // *	camelCase for internal constant

// // * Typed Constant work like immutable variable
// // *	can interoperate only with same type

// // * Untype constant work like literal
// // *	can interoperate with similar type

// // * Enumerated constant
// // *	Special symbol iota allows related constant to be created easily
// // *	Iota start at 0 in eachh block and increments by one
// // *	Watch out of constant values that match zero values for variable

// // * Enumerated Expression
// // *	Operation that can be determined at compiled time are
// // *		Arithmetic
// // *		Bitwise operator
// // *		Bitshifting

package main

// import (
// 	"fmt"
// )

// func main() {
// 	fmt.Println(`
// 		***********************************
// 		--------------POINTER--------------
// 		***********************************
// 	`)

// 	fmt.Println("========== DEFER ==========")

// 	a := 42
// 	b := a // * variable 'b' will copy value/ data of 'a'
// 	fmt.Println(a, b)
// 	a = 27
// 	fmt.Println(a, b) // * result 27 42
// 	// * b will not change because it just copy of the first 'a'
// 	// * not point in same memory

// 	var x int = 42
// 	var y *int = &x // * y is pointer (using asterix in type '*int'), it will holding memory location of 's'
// 	// * & <-- adress operator (ampersand)
// 	fmt.Println(x, y)
// 	x = 27
// 	// * derefrencing operator, to see data in memory location
// 	fmt.Println(x, *y) // * y* <-- dereferencing
// 	*y = 10
// 	fmt.Println(x, *y) // * y* <-- dereferencing

// 	// *pointer arithmatics -- using unsafe pkg
// 	i := [3]int{1, 2, 3}
// 	j := &i[0] //* pointing to first element of 'i'
// 	k := &i[1] //* pointing to second element of 'i'
// 	fmt.Println("%v %p %p\n", i, j, k)

// 	var ms *myStruct
// 	ms = &myStruct{foo: 42}
// 	fmt.Println(ms) // * result '&{42}'
// 	// * it will tell that the pointer is pointing to adress that have value {42}

// 	// * built in new function
// 	var ms2 *myStruct
// 	fmt.Println(ms2) // * special value <nil>
// 	// * if pointer not yet initialized, go will automaticly initialized it,
// 	// * and give default value,'nil'
// 	ms2 = new(myStruct)
// 	fmt.Println(ms2) // * result '&{0}'

// 	var ms3 *myStruct
// 	ms3 = new(myStruct)
// 	(*ms3).foo = 42
// 	fmt.Println((*ms3).foo)

// 	// * using syntaxtic sugar
// 	var ms4 *myStruct
// 	ms4 = new(myStruct)
// 	ms4.foo = 42
// 	fmt.Println(ms4.foo)
// }

// type myStruct struct {
// 	foo int
// }

// * Pointer

// *	Creatin pointer
// *		Pointer type use an asterix (*) as a prefix to type pointes to
// * 			*int - a pointer to an integer
// *		Use the adress of operator (&) to get address of variable

// *	Derefrencing pointer
// *		Dereference a pointer by preceding with an asterks (*)
// *		Complex types (e.g struct) are automatically dereferenced

// *	Create pointer to object
// *		can ude the address of operator (&) if value type already exist
// *			ms := myStruct{foo:42}
// *			p := &ms
// *		use address of opertor before initializer
// *			&myStruct{foo:42}
// *		use 'new' keyword
// *			cant initialize field at same time

// * 	Type with internal pointers
// *		All assignment operations in Go are copy operations
// *		Slice and maps contain internal pointer, so copies point to same underlying data

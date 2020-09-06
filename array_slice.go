package main

// import (
// 	"fmt"
// )

// func main() {
// 	fmt.Println(`
// 		*************************
// 		--  ARRAY AND SCLICES  --
// 		*************************
// 	`)

// 	fmt.Println("========== ARRAY ==========")
// 	// * initialization variable_name := [array_size]array_tipe
// 	grades := [3]int{97, 85, 93}
// 	// grades := [...]int{97, 85, 93} // * '...' --> literal syntax, use it when we didnt want to define array size
// 	fmt.Printf("Grades: %v\n", grades)

// 	var students [3]string
// 	fmt.Printf("students: %v\n", students) // will result empty array
// 	students[0] = "Lisa"
// 	fmt.Printf("students: %v\n", students)
// 	students[1] = "Ahmad"
// 	students[2] = "Arnold"
// 	fmt.Printf("students: %v\n", students)
// 	fmt.Printf("students #1: %v\n", students[1])
// 	fmt.Printf("number of students: %v\n", len(students))

// 	var identityMatrix [3][3]int = [3][3]int{[3]int{1, 0, 0}, [3]int{0, 1, 0}, [3]int{0, 0, 1}}
// 	fmt.Println(identityMatrix)

// 	//  or we write like this
// 	var identityMatrix2 [3][3]int
// 	identityMatrix2[0] = [3]int{1, 0, 0}
// 	identityMatrix2[1] = [3]int{0, 1, 0}
// 	identityMatrix2[2] = [3]int{0, 0, 1}
// 	fmt.Println(identityMatrix2)

// 	// * in go, array is considered as value
// 	// * when copying an array, it getting literal copy but different data
// 	// * so becareful
// 	a := [...]int{1, 2, 3}
// 	b := a
// 	b[1] = 9
// 	fmt.Println(a)
// 	fmt.Println(b)

// 	// * we can change value a, by pointing it
// 	c := &a
// 	c[1] = 9
// 	fmt.Println(a)
// 	fmt.Println(c)

// 	fmt.Println("========== SLICES ==========")
// 	d := []int{1, 2, 3}
// 	e := d // in slice, it will modify the same data, not copying it like an array
// 	e[1] = 9
// 	fmt.Println(d)
// 	fmt.Println(e)
// 	fmt.Printf("Length: %v\n", len(d))
// 	fmt.Printf("Capacity: %v\n", cap(d))
// 	// a slice is pojecting underlying array / slice is part of an array
// 	// capacity show max length of an array

// 	x := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
// 	x1 := x[:]   // slice of all elements
// 	x2 := x[3:]  // slice from 4th elemet to end
// 	x3 := x[:6]  // slice first 6 elemet
// 	x4 := x[3:6] // slice the 4th element to 6th element
// 	// *[x:y] x inclide/inclusive, y exclude/exclusive

// 	fmt.Println(x1)
// 	fmt.Println(x2)
// 	fmt.Println(x3)
// 	fmt.Println(x4)
// 	fmt.Println(cap(x3))

// 	// * we can create slice use 'make' built in function
// 	// * there are 2 way to use it
// 	// * 3 input
// 	// * 2 input/ argument
// 	n := make([]int, 3) // (slice_type, slice_length)
// 	fmt.Println(n)
// 	fmt.Printf("Length: %v\n", len(n))
// 	fmt.Printf("Capacity: %v\n", cap(n))

// 	m := make([]int, 3, 100) // (slice_type, slice_length, slice_capacity)
// 	fmt.Println(m)
// 	fmt.Printf("Length: %v\n", len(m))
// 	fmt.Printf("Capacity: %v\n", cap(m))
// 	// ? why use this
// 	// * its because unike array, slices dont have to have fix size of their entire life
// 	// * wee can add or remeve element from them

// 	o := []int{}
// 	fmt.Println(o)
// 	fmt.Printf("Length: %v\n", len(o))
// 	fmt.Printf("Capacity: %v\n", cap(o))
// 	o = append(o, 1)
// 	fmt.Println(o)
// 	fmt.Printf("Length: %v\n", len(o))
// 	fmt.Printf("Capacity: %v\n", cap(o)) // will result 2
// 	o = append(o, 2, 3, 4, 5, 6, 7)
// 	// * we cant use array in append, because the argument must mactch with the variable
// 	// * if we use append(o, []int{2,3,4,5,6}) , it will throw err
// 	// * but we can use spread operator append(o, []int{2,3,4,5,6}...)
// 	p := []int{8, 9, 10}
// 	o = append(o, p...)
// 	fmt.Println(o)
// 	fmt.Printf("Length: %v\n", len(o))
// 	fmt.Printf("Capacity: %v\n", cap(o))

// 	fmt.Println("========== SLICES - STACK OPERATION==========")
// 	// * push and pop element
// 	// * we can push using append
// 	// ? but how to pop element
// 	// * we can use shift operator

// 	// * to remove first elelmet we can
// 	i := []int{1, 2, 3, 4, 5}
// 	i1 := append(i[1:])
// 	fmt.Println(i1)

// 	// * to remove last elelmet we can
// 	j := []int{1, 2, 3, 4, 5}
// 	j1 := append(j[:len(j)-1])
// 	fmt.Println(j1)

// 	// * to remove in beetween elelment we can concat the slice
// 	k := []int{1, 2, 3, 4, 5}
// 	fmt.Println(k)
// 	k1 := append(k[:2], k[3:]...)
// 	fmt.Println(k1)
// 	fmt.Println(k) // * result : [1, 2, 3, 4, 5, 5]
// 	// * be careful when using slice
// 	// * we can deal with this by using for loop to copy the slice

// }

// // * Arrays
// // *	collection if item with same type
// // *	Fixed size
// // *	Declaration style
// // *		a := [3]int{1, 2, 3}
// // *		a := [...]int{1 ,2, 3}
// // *		var a [3]int
// // *	Access via zero-based index
// // *		a := [3]int {1, 2, 3} --> a[1] == 3
// // *	len function return sizeof an array
// // * 	Copies refer to different underlying data

// // * Slice
// // *	Backed by array
// // *	Creation tyles
// // *		Slice existing array or slice
// // *		Literal style
// // *		Via make function
// // *			a := make([]int, 10) // create slice with capacity and length == 10
// // *			a := make([]int 10, 100)  // create slice with length == 10 and capacity == 100
// // *	len function return length of slice
// // *	cap function return length of underlying array
// // *	append function to add element to slice
// // *		may cause expensive copy operation if underlying array is too small
// // *	Copies refer to same underlying array

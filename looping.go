package main

// import (
// 	"fmt"
// )

// func main() {
// 	fmt.Println(`
// 		*************************
// 		-----  LOOPING  ----
// 		*************************
// 	`)

// 	fmt.Println("========== FOR STATEMENT ==========")
// 	// * BASIC FOR LOOP, 3 statement
// 	// * initializer statement;  statement give boolean result; incremental statement
// 	for i := 0; i < 5; i++ {
// 		fmt.Println(i)
// 	}
// 	// * using 2 ariable
// 	for x, y := 0, 0; x < 5; x, y = x+1, y+1 {
// 		fmt.Println(x, y)
// 	}

// 	for i := 0; i < 5; i++ {
// 		fmt.Println(i)
// 		if i%2 == 0 {
// 			i /= 2
// 		} else {
// 			i = 2*i + 1
// 		}
// 	}

// 	// * for loop withh 2 statement or 1 statement
// 	//  * without initializer
// 	//  * must initate variable in func scope variable
// 	k := 0
// 	for ; k < 5; k++ { // ! semi colon still must be used
// 		fmt.Println(k)
// 	}

// 	// * while like in go
// 	// ! in go didnt have while loop
// 	// * without incremental statement and initializer statement
// 	// * it will become while loop
// 	for k < 5 {
// 		fmt.Println(k)
// 		k++ // * incremental statement mustbe initializzed in side the loop
// 		// ! if not it will be infinite loop
// 	}

// 	// * deal with infinite loop with break statement
// 	x := 0
// 	for {
// 		fmt.Println(x)
// 		x++
// 		if x == 5 {
// 			break
// 		}
// 	}

// 	// ! break will make program to exit the loop
// 	// * if we still want to do the loop and keep the process, we can use 'continue' statemnet
// 	for i := 0; i < 10; i++ {
// 		if i%2 == 0 {
// 			continue
// 		}
// 		fmt.Println(i)
// 	}
// 	//  * continue will exit the iterate and start back over

// 	// * nested loop
// 	for i := 1; i <= 3; i++ {
// 		for j := 1; j <= 3; j++ {
// 			fmt.Println(i * j)
// 		}
// 	}

// 	// Labell

// Loop:
// 	for i := 1; i <= 3; i++ {
// 		for j := 1; j <= 3; j++ {
// 			fmt.Println(i * j)
// 			if i*j >= 3 {
// 				break Loop
// 			}
// 		}
// 	}

// 	fmt.Println("========== Collection Loop ==========")

// 	// * creat loop with collection
// 	s := []int{1, 2, 3}
// 	fmt.Println(s)
// 	// * using range
// 	for k, v := range s {
// 		fmt.Println(k, v)
// 	}

// 	for k := range s {
// 		fmt.Println(k)
// 	}

// 	for _, v := range s {
// 		fmt.Println(v)
// 	}
// 	// * range can use with all collections and string type

// }

// // * For statement
// // *	simple loop
// // *		for initializer; test; incrementer{}
// // *		for test {}
// // *		for {}
// // *	exiting early
// // *		break
// // *		continue
// // *		Label
// // *	looping over collection
// // *		array, slices, maps, string, channels
// // *		for k, v := range collection {}

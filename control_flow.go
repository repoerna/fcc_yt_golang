package main

// import (
// 	"fmt"
// 	"math"
// )

// func main() {
// 	fmt.Println(`
// 		*************************
// 		-----  CONTROL FLOW  ----
// 		*************************
// 	`)

// 	fmt.Println("========== IF STATEMENT ==========")
// 	// * simple if statement
// 	if true {
// 		fmt.Println("The test is ttrue")
// 	}

// 	// * other way to use if statement
// 	statePopulation := map[string]int{
// 		"California":   39250017,
// 		"Texas":        27862596,
// 		"Florida":      20612439,
// 		"New York":     19745289,
// 		"Pennysylania": 12802503,
// 		"Illinois":     12801539,
// 		"Ohio":         11614373,
// 	}

// 	if pop, ok := statePopulation["Florida"]; ok { // if initializer; test boolean {}
// 		fmt.Println(pop)
// 	}
// 	// fmt.Println(pop) // ! this will result an error, because pop variable just can be used in if block

// 	number := 50
// 	guess := 50
// 	// * add more rule in here
// 	if guess < 1 || guess > 100 { // * or operator
// 		fmt.Println("The guess must be between 1 and 100!")
// 	}

// 	if guess >= 1 && guess <= 100 { // * and opoerator
// 		if guess < number {
// 			fmt.Println("Too low")
// 		}
// 		if guess > number {
// 			fmt.Println("Too High")
// 		}
// 		if guess == number {
// 			fmt.Println("You got it!")
// 		}
// 	}

// 	fmt.Println(number <= guess, number >= guess, number != guess)
// 	// * we can use < or > in int variable
// 	// * if we gonna use string for comparisson we can only use ==, !=

// 	fmt.Println(!true) // * not operator

// 	// * short circuit concept

// 	number1 := 50
// 	guess1 := 105 // -5
// 	// * add more rule in here
// 	if guess1 < 1 || returnTrue() || guess1 > 100 { // * or operator
// 		fmt.Println("The guess must be between 1 and 100!")
// 	}

// 	if guess1 >= 1 && guess1 <= 100 { // * and opoerator
// 		if guess < number {
// 			fmt.Println("Too low")
// 		}
// 		if guess1 > number1 {
// 			fmt.Println("Too High")
// 		}
// 		if guess1 == number1 {
// 			fmt.Println("You got it!")
// 		}
// 	}

// 	// * we an rewrite it like this
// 	if guess1 < 1 || returnTrue() || guess1 > 100 { // * or operator
// 		fmt.Println("The guess must be between 1 and 100!")
// 	} else { // * and opoerator
// 		if guess < number {
// 			fmt.Println("Too low")
// 		}
// 		if guess1 > number1 {
// 			fmt.Println("Too High")
// 		}
// 		if guess1 == number1 {
// 			fmt.Println("You got it!")
// 		}
// 	}

// 	// * or like this
// 	if guess1 < 1 {
// 		fmt.Println("The guess must be greater than 1!")
// 	} else if guess1 > 100 { // * or operator
// 		fmt.Println("The guess must be lest than 100!")
// 	} else { // * and opoerator
// 		if guess < number {
// 			fmt.Println("Too low")
// 		}
// 		if guess1 > number1 {
// 			fmt.Println("Too High")
// 		}
// 		if guess1 == number1 {
// 			fmt.Println("You got it!")
// 		}
// 	}

// 	// ! careful when using decimal number and equal operator
// 	// ! ex:
// 	myNum := 0.1
// 	myNum = 0.123                               // ! result different
// 	if myNum == math.Pow(math.Sqrt(myNum), 2) { // the logic must be true, because pow with 2 and then sqrt it
// 		fmt.Println("These are the same")
// 	} else {
// 		fmt.Println("These are different")
// 	}
// 	// * this hapeen because go treat myNum as floating point number

// 	// ? how to deal with that
// 	//  * generate error value that less than certain treshold
// 	if math.Abs(myNum/math.Pow(math.Sqrt(myNum), 2)-1) < 0.001 { // the logic must be true, because pow with 2 and then sqrt it
// 		fmt.Println("These are the same")
// 	} else {
// 		fmt.Println("These are different")
// 	}

// 	fmt.Println("========== SWITCH STATEMENT ==========")
// 	//  switch is special purpose if stratement
// 	switch 2 { // * value 2 called tag
// 	case 1: // * case value will compared with tag value
// 		fmt.Println("one")
// 	case 2:
// 		fmt.Println("two")
// 	default:
// 		fmt.Println("not one or two")
// 	}

// 	// ? falling two cases / range casess, we dont have it in golang

// 	// * so,
// 	switch 4 {
// 	case 1, 5, 10:
// 		fmt.Println("one, five, or ten")
// 	case 2, 4, 6:
// 		fmt.Println("two, four, or six")
// 	default:
// 		fmt.Println("another number")
// 	}
// 	// ! test cases must unique

// 	switch i := 2 + 3; i { // * we can use initializer as tag
// 	case 1, 5, 10:
// 		fmt.Println("one, five, or ten")
// 	case 2, 4, 6:
// 		fmt.Println("two, four, or six")
// 	default:
// 		fmt.Println("another number")
// 	}

// 	// * we can use another syntax called tagless syntax
// 	j := 10
// 	switch {
// 	case j <= 10:
// 		fmt.Println("less than or equal ten")
// 	case j <= 20:
// 		fmt.Println("less than or equal twenty")
// 	default:
// 		fmt.Println("greater than twenty")
// 	}
// 	// * in this switch, it wil only excecute first case
// 	// * it is because the rule is overlaping
// 	// * implied / implicit, switch in go use break

// 	// * to deal with it, we can use 'falltrough'
// 	switch {
// 	case j <= 10:
// 		fmt.Println("less than or equal ten")
// 		fallthrough
// 	case j <= 20:
// 		fmt.Println("less than or equal twenty")
// 	default:
// 		fmt.Println("greater than twenty")
// 	}
// 	// ! falltrough is logic less, so it will pass the logic

// 	// * type switch
// 	var n interface{} = 1
// 	switch n.(type) {
// 	case int:
// 		fmt.Println("n is an int")
// 	case float64:
// 		fmt.Println("n is a float64")
// 	case string:
// 		fmt.Println("n is string")
// 	default:
// 		fmt.Println("n is another type")
// 	}

// 	// * break
// 	switch n.(type) {
// 	case int:
// 		fmt.Println("n is an int")
// 		break
// 		fmt.Println("without break, this will be printed")
// 	case float64:
// 		fmt.Println("n is a float64")
// 	case string:
// 		fmt.Println("n is string")
// 	default:
// 		fmt.Println("n is another type")
// 	}
// }

// func returnTrue() bool {
// 	fmt.Println("returning true")
// 	return true
// }

// // * if statement
// // *	Initializer
// // *	Comparisson operator
// // *	Logical operator
// // *	Short Circuiting
// // *	if-else statement
// // *	if-else if statement
// // *	Equality and floats

// // * Switch
// // *	Switching on tag
// // *	Cases multiple tests
// // * 	initializer
// // *	Switches with no tag
// // *	falltrough
// // *	type switches
// // *	Breaking out early

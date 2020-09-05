package main

// func main() {
// 	// * BOOLEAN
// 	// it represent 2 state, true or false
// 	// ero value is false
// 	var n bool = true
// 	fmt.Printf("%v, %T\n", n, n)

// 	//  comparison
// 	m := 1 == 1
// 	o := 1 == 2
// 	fmt.Printf("%v, %T\n", m, m)
// 	fmt.Printf("%v, %T\n", o, o)

// 	var p bool
// 	fmt.Printf("%v, %T\n", p, p)
// 	// ! in go, even dont initialize or set value to a variable, it still will output a value
// 	// ! because a variable type have its default value

// 	// * NUMERIC

// 	//  signed integer
// 	//  have varying size but min 32 bit
// 	/*
// 		int8	-128 - 127
// 		int16	-32 768 - 32 767
// 		int32	-2 147 483 648 - 2 147 483 647
// 		int64	-9 223 372 036 854 775 808 - 9 223 372 036 854 775 807
// 	*/

// 	// int & unsigned integer
// 	/*
// 		uint8	0 - 255
// 		uint16	0 - 65 535
// 		uint32	0 - 4 294 967 295
// 	*/
// 	var q uint16 = 42
// 	fmt.Printf("%v, %T\n", q, q)

// 	// * BYTE or unsigned 8 bit unsigned integer

// 	// * BIT OPERATOR
// 	r := 10             // 1010
// 	s := 3              // 0011
// 	fmt.Println(r & s)  // 0010 - 2
// 	fmt.Println(r | s)  // 1011 - 11
// 	fmt.Println(r ^ s)  // 1001 - 9
// 	fmt.Println(r &^ s) // 0100 - 8

// 	// * BIT SHIFTTING
// 	t := 10             // 2^3
// 	fmt.Println(t << 2) // 2^3 * 2^2 = 2^5 = 32
// 	fmt.Println(t >> 2) // 2^3 / 2^2 = 2^1 = 2

// 	// * FLOATING POINT
// 	// standar IEEE-754
// 	/*
// 		float32	+-1.18E-38 - +-3.4E38
// 		float64	+-2.23E-308 - +-1.8E308
// 	*/
// 	//  zero value is 0
// 	//  literal style
// 	//  - decimal 3.14
// 	//  - exponential 12e13
// 	//  - mixed 3.14e34
// 	var u float32 = 3.14
// 	// u = 13.7e72 // this will causing error because overflow float32
// 	u = 2.1e14
// 	fmt.Printf("%v, %T\n", u, u)

// 	// * Complex NUmber
// 	//  zero number is (0+0i)
// 	//  64 and 128 bit version
// 	//  built in function
// 	//  - complex - mkae complex number from two floats
// 	//  - real - get real part as float
// 	//  - imag - get imaginary part as float

// 	var v complex64 = 1 + 2i
// 	v = complex(5, 12)
// 	fmt.Printf("%v, %T\n", v, v)
// 	fmt.Printf("%v, %T\n", real(v), real(v))
// 	fmt.Printf("%v, %T\n", imag(v), imag(v))

// 	// text type
// 	//  string and rune
// 	// string
// 	// - utf-8
// 	// - immutable
// 	// -  can be concatenates with plus (+) operator
// 	// - can be converted into {}byte

// 	//  rune
// 	// - utf-32
// 	// - alias for int32
// 	// - special methods normally requires to process e.g. strings.Reader#ReadRune

// 	// * string or utf-8, its immutable

// 	w := "this is a string "
// 	fmt.Printf("%v, %T\n", w, w)
// 	fmt.Printf("%v, %T\n", w[2], w[2])

// 	// string concat
// 	w2 := "new string"
// 	fmt.Println(w + w2)

// 	x := []byte(w)
// 	fmt.Printf("%v, %T\n", x, x)

// 	// * rune or utf32 or int32
// 	var y rune = 'a'
// 	fmt.Printf("%v, %T\n", y, y)

// }

// // ! Cant mix types in the same familiy (uint8 + uint32 = error)

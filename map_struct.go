package main

// import (
// 	"fmt"
// )

// // * struct example
// // ? how to create struct
// // * 'type struct_name struct { structure_name structure_type}
// type Doctor struct {
// 	Number     int
// 	ActorName  string `required,max:"100"` //* <--- its called 'TAG',  give rule
// 	Companions []string
// }

// // * usually we will use Uppercase in first letter so the struct can be accessed globaly

// func main() {
// 	fmt.Println(`
// 		*************************
// 		---  MAP AND STRUCT  ----
// 		*************************
// 	`)

// 	fmt.Println("========== MAP ==========")
// 	// * key value type
// 	// ? how to create
// 	// * map[key_type]value_type{}
// 	statePopulation := map[string]int{
// 		"California":   39250017,
// 		"Texas":        27862596,
// 		"Florida":      20612439,
// 		"New York":     19745289,
// 		"Pennysylania": 12802503,
// 		"Illinois":     12801539,
// 		"Ohio":         11614373,
// 	}

// 	fmt.Println(statePopulation)

// 	// * we can using array as key in the map
// 	m := map[[3]int]string{}
// 	fmt.Println(m) // result --> map[]
// 	// ! we cannot using slices as a key

// 	// * we also can create map using built in function 'make'
// 	statePopulation1 := make(map[string]int)
// 	statePopulation1 = statePopulation
// 	fmt.Println(statePopulation1)

// 	// * in map, we can manupulate it using key
// 	// * get value of particular key
// 	fmt.Println(statePopulation1["Ohio"])

// 	// * add new key value
// 	statePopulation1["Georgia"] = 10310371
// 	fmt.Println(statePopulation1["Georgia"])
// 	// // ! the order of map always changing

// 	// * delete key-value
// 	fmt.Println(statePopulation1)
// 	delete(statePopulation1, "Georgia")
// 	fmt.Println(statePopulation1)
// 	// * if the key not found in the map, result will be 0
// 	delete(statePopulation1, "Ohi")
// 	fmt.Println(statePopulation1["Ohi"])

// 	// * so we can use that with ', ok' variable to check wheter the data is in the map or not
// 	pop, ok := statePopulation1["Ohio"]
// 	fmt.Println(pop, ok) // result --> valeu_of_ohio, true

// 	pop1, ok := statePopulation1["Ohi"]
// 	fmt.Println(pop1, ok) // result --> 0, true
// 	// * best practice is to use blank identifier
// 	_, ok = statePopulation1["Ohio"]
// 	fmt.Println(ok)

// 	// * using 'len' built-in function to check the length
// 	fmt.Println(len(statePopulation))
// 	fmt.Println(len(statePopulation1))

// 	// * just like slice, when we have multiple assignment to a map
// 	// * underlying data is pass by refrence, it means manipulating
// 	// * one map can effct on other map that have same refrence
// 	statePopulation2 := make(map[string]int)
// 	statePopulation2 = statePopulation
// 	fmt.Println(statePopulation)
// 	fmt.Println(statePopulation2)
// 	delete(statePopulation2, "Ohio")
// 	fmt.Println(statePopulation)
// 	fmt.Println(statePopulation2)

// 	fmt.Println("========== STRUCT ==========")
// 	// * struct is different with array, slice, and map that only have one type

// 	// ? then, how to use it
// 	// * initiate a struct
// 	aDoctor := Doctor{
// 		Number:    3,
// 		ActorName: "Jon Pertwee",
// 		Companions: []string{
// 			"Liz Shaw",
// 			"Jo Grant",
// 			"Sarah",
// 			"Jane Smith",
// 		},
// 	}
// 	fmt.Println(aDoctor)

// 	// * we can access the struct by using'.'
// 	fmt.Println(aDoctor.Companions)
// 	fmt.Println(aDoctor.Companions[1])

// 	// * we can also use positional syntax
// 	aDoctor1 := Doctor{
// 		3,
// 		"Jon Pertwee",
// 		[]string{
// 			"Liz Shaw",
// 			"Jo Grant",
// 			"Sarah",
// 			"Jane Smith",
// 		},
// 	}
// 	fmt.Println(aDoctor1) // same result with aDoctor

// 	// ! it is not recommended to use positional syntax
// 	// * let say the struct change, there are addional item in the struct
// 	type Doctor2 struct {
// 		Number     int
// 		ActorName  string
// 		Companions []string
// 		Episode    []string
// 	}

// 	// aDoctor3 := Doctor2{
// 	// 	3,
// 	// 	"Jon Pertwee",
// 	// 	[]string{
// 	// 		"Liz Shaw",
// 	// 		"Jo Grant",
// 	// 		"Sarah",
// 	// 		"Jane Smith",
// 	// 	},
// 	// }

// 	// fmt.Println(aDoctor3) // ! it will be error
// 	// ! because not enough argument

// 	// * or if the order of struct change it can give you error or misinformation
// 	// * example:

// 	aDoctor4 := Doctor2{
// 		3,
// 		"Jon Pertwee",
// 		[]string{
// 			"test1",
// 			"test2",
// 		},
// 		[]string{
// 			"Liz Shaw",
// 			"Jo Grant",
// 			"Sarah",
// 			"Jane Smith",
// 		},
// 	}

// 	fmt.Println(aDoctor4.Companions) // ! no error, but the value is wrong
// 	//  TODO: so it recomended to use the normal way

// 	// * we can also can create anonymous struct
// 	anonStruct := struct{ name string }{name: "anonymous"}
// 	fmt.Println(anonStruct)
// 	// * there are few cases where we want use anonymous struct
// 	// * when we want to structured a data that didnt have a formal struct
// 	// * because it didnt have a name, so we can use the struct globally

// 	// * unlike map, struct is value based
// 	anonStruct2 := anonStruct
// 	anonStruct2.name = "non anon"
// 	fmt.Println(anonStruct)
// 	fmt.Println(anonStruct2)

// 	// * in go we dont support OOP principal
// 	// * but we have another

// 	// * Embedding

// 	type Animal struct {
// 		Name   string
// 		Origin string
// 	}

// 	type Bird struct {
// 		Animal // * <-- embedded struct in another struct, it looks like inheritance
// 		// * it call composition / embed
// 		SpeedKPH float32
// 		CanFly   bool
// 	}

// 	// ? So, how to use it
// 	b := Bird{}
// 	b.Name = "emu"
// 	b.Origin = "Australia"
// 	b.SpeedKPH = 48
// 	b.CanFly = true

// 	fmt.Println(b)
// 	fmt.Println(b.Name)

// 	// ! Bird is np inheriting from Animal
// 	// ! it still dependence type

// 	// * if we use literal syntax
// 	b1 := Bird{
// 		Animal: Animal{
// 			Name:   "aa",
// 			Origin: "bb",
// 		},
// 		SpeedKPH: 48,
// 		CanFly:   true,
// 	}
// 	fmt.Println(b1)
// }

// // * Maps
// // *	Collection of value type that are accessed via keys
// // *	Created via literals or via make fuction
// // *	Members accessed via [key] syntax
// // *		myMap["key"] = "value"
// // *	Check for presence with "value, ok" form of result
// // *	Multiple assignment refer to same underlying data

// // * Struct
// // *	Collection of disparate data types that describe a single concept
// // *	Keyed by name fields
// // *	Normally Created as types, but anonymous struct are allowed
// // *	Struct are value types
// // *	No inheritance, but can use composition via embedding
// // *	Tags can be added to struct fields to escribe field

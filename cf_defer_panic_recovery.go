package main

// import (
// 	"fmt"
// 	"log"
// )

// func main() {
// 	fmt.Println(`
// 		***********************************
// 		-----  DEFER, PANIC, RECOVERY  ----
// 		***********************************
// 	`)

// 	fmt.Println("========== DEFER ==========")

// 	// fmt.Println("start")
// 	// fmt.Println("middle")
// 	// fmt.Println("end")

// 	fmt.Println("=== using defer ===")

// 	// fmt.Println("start")
// 	// defer fmt.Println("middle")
// 	// fmt.Println("end")

// 	fmt.Println("=== LIFO ===")
// 	// defer fmt.Println("start")
// 	// defer fmt.Println("middle")
// 	// defer fmt.Println("end")

// 	// * usually defer is used to close resourced
// 	// * and some resources can be dependent with others

// 	// * resource request example
// 	// res, err := http.Get("http://www.google.com/robots.txt")
// 	// if err != nil {
// 	// 	log.Fatal(err)
// 	// }
// 	// robots, err := ioutil.ReadAll(res.Body)
// 	// res.Body.Close()
// 	// if err != nil {
// 	// 	log.Fatal(err)
// 	// }
// 	// fmt.Printf("%s\n", robots)

// 	// * use defer to handle with closing a resource
// 	// res, err := http.Get("http://www.google.com/robots.txt")
// 	// if err != nil {
// 	// 	log.Fatal(err)
// 	// }
// 	// defer res.Body.Close()
// 	// robots, err := ioutil.ReadAll(res.Body)
// 	// if err != nil {
// 	// 	log.Fatal(err)
// 	// }
// 	// fmt.Printf("%s\n", robots)

// 	// ! if we open resource within in loop, dont use defer!

// 	// a := "start"
// 	// defer fmt.Println(a)
// 	// a = "end"
// 	// * result of a will be "start"
// 	// * when call a defer, it will take argument at the time when defer is called
// 	// * not when it execute

// 	fmt.Println("========== PANIC ==========")
// 	// * go didn't have exception

// 	// * natural panic in go
// 	// a, b := 1, 0
// 	// ans := a / b
// 	// fmt.Println(ans)

// 	// * built in panic
// 	// fmt.Println("start")
// 	// panic("it is bad")
// 	// fmt.Println("end") // * this line will not excecuted

// 	// * some practical example
// 	// http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
// 	// 	w.Write([]byte("Hello Go!"))
// 	// })
// 	// err := http.ListenAndServe(":8080", nil)
// 	// if err != nil {
// 	// 	panic(err.Error())
// 	// }

// 	// * panic statemnet will be exceuted after defer executed
// 	fmt.Println("========== RECOVER ==========")
// 	// fmt.Println("start")
// 	// defer fmt.Println("this was defered")
// 	// panic("something bad")
// 	// fmt.Println("end")

// 	// fmt.Println("start")
// 	// defer func() {
// 	// 	if err := recover(); err != nil { // recover wiill return nill if no panic occur
// 	// 		log.Println("Error:", err)
// 	// 	}
// 	// }() // * parenthesis are making function excecuted, it wil invoke the function
// 	// // ! defer statement does not tak efunction statement, it take function call
// 	// panic("something bad")
// 	// fmt.Println("end")

// 	// * another deeper example
// 	fmt.Println("start")
// 	panicker()
// 	fmt.Println("end")

// }

// func panicker() {
// 	fmt.Println("about to panic")
// 	defer func() {
// 		if err := recover(); err != nil {
// 			log.Println("Error:", err)
// 			panic(err)
// 		}
// 	}()
// 	panic("something bad")
// 	fmt.Println("done panicking")
// }

// // * Defer
// // *	used to delay execution of a statemnt until function exits
// // *	useful to group "open" and "close" function together
// // *		becareful in loops
// // *	Run in LIFO order
// // * 	Argument evaluated at time defer is executeed, not at time of called function

// // * Panic
// // *	Occur when program cannot continue at all
// // *		Dont use when file cant be opened, unless it critical
// // *		Use for unrecoverable event - cannot obtain TCP port for web server
// // *	Function will stop executon
// // *		deferred function will still fire
// // *	If nothing handles panic, program will exit

// // * Recover
// // *	Used to recover from panics
// // *	Only useful in deferred function
// // *	Current function will not attempt to continue, but higher function in call stack will

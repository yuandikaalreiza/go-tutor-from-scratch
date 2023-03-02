package main

import "fmt"

func main() {
	// 1. BASIC
	//============================================================
	// 1.1. variable
	/*
		var a = "initial"
		fmt.Println(a)

		var b, c int = 1, 2
		fmt.Println(b, c)

		var d = true
		fmt.Println(d)

		var e int
		fmt.Println(e)

		f := "apple"
		fmt.Println(f)
	*/

	// 1.2. data type
	var a bool = true    // Boolean
	var b int = 5        // Integer
	var c float32 = 3.14 // Floating point number
	var d string = "Hi!" // String

	fmt.Println("Boolean: ", a)
	fmt.Println("Integer: ", b)
	fmt.Println("Float:   ", c)
	fmt.Println("String:  ", d)

	// 		bool
	var b1 bool = true // typed declaration with initial value
	var b2 = true      // untyped declaration with initial value
	var b3 bool        // typed declaration without initial value
	b4 := true         // untyped declaration with initial value

	fmt.Println(b1) // Returns true
	fmt.Println(b2) // Returns true
	fmt.Println(b3) // Returns false
	fmt.Println(b4) // Returns true

	// 		sign int
	var x int = 500
	var y int = -4500
	fmt.Printf("Type: %T, value: %v", x, x)
	fmt.Printf("Type: %T, value: %v", y, y)

	//		unsign int
	var x uint = 500
	var y uint = 4500
	fmt.Printf("Type: %T, value: %v", x, x)
	fmt.Printf("Type: %T, value: %v", y, y)

}

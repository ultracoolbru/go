package main

import "fmt"

func main() {
	// Variables
	var a int = 10
	var b float32 = 3.14
	const c string = "Hello World"

	var (
		d = 10
		e = 7
	)

	x, y := 10, 20

	fmt.Println(a, b, c, d, e, x, y)

	// Pointers
	var p1 *int = &a
	fmt.Println(p1, *p1)

	p2 := 5
	fmt.Println(p2, &p2)

	changePointerValue(&p2)
	fmt.Println(p2)

	// Printf
	printFExamples()
}

func changePointerValue(x *int) {
	*x = 10
}

func printFExamples() {
	const pi float64 = 3.14159265358979323846
	x := 5
	isbool := true

	fmt.Printf("Value of pi: %f \n", pi)
	fmt.Printf("Value of pi: %t \n", pi)
	fmt.Printf("Value of x: %d \n", x)
	fmt.Printf("Value of isbool: %t \n", isbool)
	fmt.Printf("Value of isbool: %T \n", isbool)
}

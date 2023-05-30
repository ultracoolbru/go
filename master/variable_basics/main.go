package main

import (
	"fmt"
)

func main() {
	// Variables
	variables()

	// Types & Zero Values
	typesAndZeroValues()

	// Naming Conventions
	namingConventions()

	// Go fmt Package
	// https://golang.org/pkg/fmt/
	// Printf
	printFExamples()

	// Pointers
	pointers()
}

func variables() {
	var a int = 10
	var b float32 = 3.14
	const c string = "Hello World"

	var (
		d = 10
		e = 7
	)

	x, y := 10, 20

	fmt.Println(a, b, c, d, e, x, y)

	// In Go you cannot declare a variable and not use it.
	var name = "Edward"
	_ = name

	car, cost := "Audi", 50000
	fmt.Println(car, cost)

	var (
		salary    float64 = 10000.0
		firstName string  = "Edward"
		gender    bool    = true
	)

	// salary = 10000.0
	// firstName = "Edward"
	// gender = true

	fmt.Println(salary, firstName, gender)

	// Variable Swapping
	var firstVar, SecondVar int = 5, 10
	firstVar, SecondVar = SecondVar, firstVar
	fmt.Println(firstVar, SecondVar)
}

func typesAndZeroValues() {
	var a int = 4
	var b float64 = 5.2

	fmt.Println("Types & Zero Vals", a, b)

	a = int(b)
	fmt.Println("Types & Zero Vals", a, b)
}

func namingConventions() {
	// It is better to use short versions of variable names.
	var mv int = 10        // mv = max value
	var max_value int = 10 // not good practice.
	_ = max_value

	var packetsReceived int // not good practice.
	_ = packetsReceived
	var pr int // pr = packets received

	const MAX_VALUE = 100 // not good practice.
	const maxVal = 100    // maxVal = max value

	fmt.Println(mv, pr, maxVal)

	// If you declare a variable starting with an uppercase letter, it will be exported.
	var Max = 100
	fmt.Println(Max)

	writeToDB := true  // ok
	writeToDb := false // not ok

	fmt.Println(writeToDB, writeToDb)
}

func printFExamples() {
	const pi float64 = 3.14159265358979323846
	x := 5
	isbool := true

	fmt.Printf("Value of pi: %f \n", pi)         // %f = float
	fmt.Printf("Value of pi: %T \n", pi)         // %t = type
	fmt.Printf("Value of x: %d \n", x)           // %d = decimal
	fmt.Printf("Value of isbool: %t \n", isbool) // %t = boolean
	fmt.Printf("Value of isbool: %T \n", isbool) // %T = type
	fmt.Printf("Value of pi: %.7f \n", pi)       // %.2f = float with 2 decimal places
	fmt.Printf("Value of pi: %9.2f \n", pi)      // %9.2f = float with 2 decimal places and 9 characters (spaces) wide
	fmt.Printf("Value of pi: %9.f \n", pi)       // %9.f = float with 9 characters (spaces) wide
}

func pointers() {
	var anInt int = 42
	var p1 *int = &anInt
	fmt.Println(p1, *p1)

	p2 := 5
	fmt.Println(p2, &p2)

	changePointerValue(&p2)
	fmt.Println(p2)
}

func changePointerValue(x *int) {
	*x = 10
}

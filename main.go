package main

import (
	"fmt"

	"learning-go/class"
)

func main() {

	// print
	fmt.Println("=== Print ===")
	class.Prints()
	fmt.Println()

	// strings
	fmt.Println("=== String ===")
	class.StringVariables()
	fmt.Println()

	// ints
	fmt.Println("=== Int ===")
	class.IntVariables()
	fmt.Println()

	// floats
	fmt.Println("=== Float ===")
	class.FloatVariables()
	fmt.Println()

	// bits & memory
	fmt.Println("=== Bits ===")
	class.BitsVariables()
	fmt.Println()

	// functions
	fmt.Println("=== Function ===")
	class.Sum(10, 50)
	fmt.Println()

	// arrays
	fmt.Println("=== Arrays ===")
	class.Arrays()
	fmt.Println()

	// errors
	fmt.Println("=== Errors ===")
	class.HandlerError()

	// standard library
	fmt.Println("=== Standard Library ===")
	class.StringStandards()
	class.SortStandards()
	fmt.Println()

	// loops
	fmt.Println("=== Loops ===")
	class.ForLoop()
	fmt.Println()

	// operators
	fmt.Println("=== Operators ===")
	class.Boolean()
	class.Conditionals()
	fmt.Println()

	// structs
	fmt.Println("=== Structs ===")
	class.Structs()
	fmt.Println()

}

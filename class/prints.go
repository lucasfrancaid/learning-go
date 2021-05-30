package class

import "fmt"

func Prints() {
	// print line
	fmt.Println("[print line] -> Hello Go!")

	// print format
	age := 23
	firstName := "Lucas"

	// variables
	fmt.Printf("[variables] -> My age is %v and my first name is %v \n", age, firstName)

	// explicit int and string
	fmt.Printf("[explicit int and string] -> My age is %d and my first name is %s \n", age, firstName)

	// explicit int and string with double quotes
	fmt.Printf("[explicit int and string with double quotes] -> My age is %d and my first name is %q \n", age, firstName)

	// format showing type
	fmt.Printf("[format showing type] -> My age is %T and my first name is %T \n", age, firstName)

	// format float and float formated
	fmt.Printf("[format float and float formated] -> PI is %f and I have U$ %6.2f in my wallet \n", 3.14159, 553.32)

	// save formated print
	var formatedPrint = fmt.Sprintf("My age is %v and my first name is %v", age, firstName)
	fmt.Println("[save formated print] -> My formated print is:", formatedPrint)
}

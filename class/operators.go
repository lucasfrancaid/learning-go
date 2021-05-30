package class

import "fmt"

func Boolean() {
	age := 45

	// less than or equal
	formated := fmt.Sprintf("[boolean] -> The age %v is less than or equal to %v? %v", age, 50, age <= 50)
	fmt.Println(formated)

	// more than or equal
	formated = fmt.Sprintf("[boolean] -> The age %v is more than or equal to %v? %v", age, 50, age >= 50)
	fmt.Println(formated)

	// equal
	formated = fmt.Sprintf("[boolean] -> The age %v is equal to %v? %v", age, 45, age == 45)
	fmt.Println(formated)

	// diffent
	formated = fmt.Sprintf("[boolean] -> The age %v is different to %v? %v", age, 50, age != 50)
	fmt.Println(formated)

}

func Conditionals() {
	age := 45

	// if else if
	if age < 30 {
		fmt.Println("[conditional if else if else] -> Age is less than to 30")
	} else if age < 40 {
		fmt.Println("[conditional if else if else] -> Age is less than to 40")
	} else {
		fmt.Println("[conditional if else if else] -> Age is 45 and not less than to 30 and 40")
	}
}

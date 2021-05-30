package class

import "fmt"

func ForLoop() {
	// for i < 5
	i := 0
	fmt.Print("[for loop] -> Value of iterate sample (for i < 5) is ")
	for i < 5 {
		fmt.Print(i, ", ")
		i++
	}
	fmt.Println()

	// for i := 0; i < 5; i++
	fmt.Print("[for loop] -> Value of iterate sample (for i := 0; i < 5; i++) is ")
	for i := 0; i < 5; i++ {
		fmt.Print(i, ", ")
	}
	fmt.Println()

	// for i := 0; i < len(fruits); i++
	fruits := []string{"strawberry", "orange", "banana", "apple", "pineapple", "grape"}
	fmt.Print("[for loop] -> Value of iterate sample (for i := 0; i < len(fruits); i++) is ")
	for i := 0; i < len(fruits); i++ {
		fmt.Print(fruits[i], ", ")
	}
	fmt.Println()

	// for index, value := range fruits
	fmt.Print("[for loop] -> Value of iterate sample (for index, value := range fruits) is ")
	for index, value := range fruits {
		fmt.Printf("(index: %v value: %v), ", index, value)
	}
	fmt.Println()

	// for index, value := range fruits
	fmt.Print("[for loop break] -> Value of iterate sample (for index, value := range fruits) is ")
	for index, value := range fruits {
		fmt.Printf("(index: %v value: %v), ", index, value)
		if value == "banana" {
			fmt.Print("Break for loop")
			break
		}
	}
	fmt.Println()

}

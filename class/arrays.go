package class

import "fmt"

func Arrays() {

	// verbose array adding by index limit length
	var myArray [2]string
	myArray[0] = "Hello"
	myArray[1] = "Go"
	fmt.Println("[verbose array adding by index with limit length] ->", myArray)

	// creating an array of strings
	fruits := []string{"strawberry", "orange", "banana", "apple"}
	fmt.Println("[creating an array of strings] ->", fruits, "length:", len(fruits))

	// creating an array of ints
	numbers := []int{1, 2, 3, 4}
	fmt.Println("[creating an array of ints] ->", numbers, "length:", len(numbers))

	// creating an array of strings with explict length
	var fruitsExplict [3]string = [3]string{"strawberry", "orange", "banana"}
	fmt.Println("[creating an array of strings with explict length] ->", fruitsExplict, "length:", len(fruitsExplict))

	// creating an array of ints with explict length
	var numbersExplict [2]int = [2]int{1, 2}
	fmt.Println("[creating an array of ints with explict length] ->", numbersExplict, "length:", len(numbersExplict))

	// changing string array value by index
	fruits[0] = "pineapple"
	fmt.Println("[changing string array value by index] ->", fruits, "length:", len(fruits))

	// changing int array value by index
	numbers[0] = 5
	fmt.Println("[changing int array value by index] ->", numbers, "length:", len(numbers))

	// append value in string array
	fruits = append(fruits, "strawberry")
	fmt.Println("[append value in string array] ->", fruits, "length:", len(fruits))

	// append value in int array
	numbers = append(numbers, 1)
	fmt.Println("[append value in int array] ->", numbers, "length:", len(numbers))

	// slice fruits [1:5]
	slicedFruits := fruits[1:5]
	fmt.Println("[slice fruits [1:5] (obs: len include the first index [1] but not the last index [5])] ->", slicedFruits, "length:", len(fruits))

	// slice numbers [2:3]
	slicedNumbers := numbers[2:3]
	fmt.Println("[slice number [2:3] (obs: len include the first index [2] but not the last index [3])] ->", slicedNumbers, "length:", len(numbers))

	// slice fruits [3:] from third to end
	slicedFruits = fruits[3:5]
	fmt.Println("[slice fruits [3:] from third to end] ->", slicedFruits, "length:", len(fruits))

	// slice numbers [3:] from third to end
	slicedNumbers = numbers[:2]
	fmt.Println("[slice numbers [:2] from zero to first (obs: after double point is index-1)] ->", slicedNumbers, "length:", len(numbers))

	customs := append(fruits, "tomato")
	fmt.Println("[custom list created from other list using append] ->", customs, "length:", len(customs))

}

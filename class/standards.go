package class

import (
	"fmt"
	"sort"
	"strings"
)

func StringStandards() {
	greetings := "Hello my friend!"

	// verify with string part contains in variable
	fmt.Println("[(strings) verify with string part contains in variable] ->", strings.Contains(greetings, "friend!"))

	// replace string matched with number of replaces
	fmt.Println("[(strings) replace string with number of replaces] ->", strings.Replace(greetings, "friend", "friends", 1))

	// replace all strings matched
	fmt.Println("[(strings) replace all strings matched] ->", strings.ReplaceAll(greetings, "friends", "friend"))

	// string to upper
	fmt.Println("[(strings) string to upper] ->", strings.ToUpper(greetings))

	// string to lower
	fmt.Println("[(strings) string to lower] ->", strings.ToLower(greetings))

	// string to title
	fmt.Println("[(strings) string to title] ->", strings.ToTitle(greetings))

	// verify if contains any character specified
	fmt.Println("[(strings) verify if contains any character specified] ->", strings.ContainsAny(greetings, "al"))

	// return the position of substring
	fmt.Println("[(strings) return the position of substring] ->", strings.Index(greetings, "fri"))

	// split string
	fmt.Println("[(strings) split string] ->", strings.Split(greetings, " "))
}

func SortStandards() {
	numbers := []int{4, 3, 8, 9, 10, 5, 6, 7, 2, 1}
	fmt.Println("[(sort) set int list] ->", numbers)

	// sort int list
	sort.Ints(numbers)
	fmt.Println("[(sort) sort int list] ->", numbers)

	// search int position in list
	intPosition := sort.SearchInts(numbers, 4)
	fmt.Printf("[(sort) search int position in list] -> The number 4 are in %v index position", intPosition)
	fmt.Println()

	fruits := []string{"strawberry", "orange", "banana", "apple", "pineapple", "grape"}
	fmt.Println("[(sort) set string list] ->", fruits)

	// sort string list
	sort.Strings(fruits)
	fmt.Println("[(sort) sort string list] ->", fruits)

	// search string position in list
	stringPosition := sort.SearchStrings(fruits, "banana")
	fmt.Printf("[(sort) search string position in list] -> The fruit banana are in %v index position", stringPosition)
	fmt.Println()
}

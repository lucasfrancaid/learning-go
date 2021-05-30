package class

import "fmt"

func Sum(n1 int, n2 int) int {
	result := n1 + n2
	fmt.Printf("[sum function] -> Result of %d+%d is %d \n", n1, n2, result)
	return result
}

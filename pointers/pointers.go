package pointers

import (
	"fmt"
)

var pointer *string
var initial string = "Alou"

func SetPointer() {
	pointer = &initial
	fmt.Printf("[set pointer pointer] -> pointer: %v, initial: %v", *pointer, initial)
	fmt.Println()
}

func GetPointer() *string {
	return pointer
}

func UpdatePointer(another *string) {
	*another, *pointer = "Pega ai", "Uola"
	fmt.Printf("[set pointer pointer] -> pointer: %v, initial: %v", *pointer, initial)
	fmt.Println()
}

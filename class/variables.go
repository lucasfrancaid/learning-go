package class

import "fmt"

func StringVariables() {
	var firstName string = "John"
	var secondName = "Smith"
	var thirdName string

	fmt.Println("[string variables] ->", firstName, secondName, thirdName)

	secondName = "Brown"
	thirdName = "Smith"

	fmt.Println("[changing values] ->", firstName, secondName, thirdName)

	lastName := "Owen"

	fmt.Println("[self typing variable] ->", firstName, secondName, thirdName, lastName)
}

func IntVariables() {
	var ageOne int = 20
	var ageTwo = 30
	ageThree := 40

	fmt.Println("[adding int variables] ->", ageOne, ageTwo, ageThree)
}

func FloatVariables() {
	var scoreOne float32 = 25.32
	var scoreTwo float64 = 343493493439434.3
	scoreThree := 10.42
	scoreFour := 1.4
	fmt.Println("[adding float variables] ->", scoreOne, scoreTwo, scoreThree, scoreFour)
}

func BitsVariables() {
	var bitOne int8 = 25
	var bitTwo int16 = 128
	var bitThree int32 = 256
	var bitFour uint8 = 25
	var bitFive uint16 = 128
	var bitSix int32 = 256
	fmt.Println("[adding bits variables] ->", bitOne, bitTwo, bitThree, bitFour, bitFive, bitSix)
}

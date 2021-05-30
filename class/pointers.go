package class

import (
	"fmt"
	"learning-go/pointers"
)

type Car struct {
	Name string
	Year int
}

func Pointers() {
	// pointer set and get pointer
	pointers.SetPointer()
	pointer := pointers.GetPointer()
	fmt.Println("[pointer set and get pointer] ->", *pointer)

	// pointer another
	another := "Hello"
	pointers.UpdatePointer(&another)
	fmt.Println("[pointer another] ->", another)

	// pointer updatedPointer
	updatedPointer := pointers.GetPointer()
	fmt.Println("[pointer updatedPointer] ->", *updatedPointer)

	// changing value of variable as param of function
	param := "Ola amigo"
	fmt.Println("[updating param as a string - before] ->", param)
	param = updateParam(param)
	fmt.Println("[updating param as a string - after] ->", param)

	// changing value of variable as pointer of function
	pointerParam := "Ola amigo"
	fmt.Println("[updating pointer as a pointer - before] ->", pointerParam)
	updatePointerParam(&pointerParam)
	fmt.Println("[updating pointer as a pointer - after] ->", pointerParam)

	// changing value of struct with scope
	scopedCar := Car{
		Name: "Fiat",
		Year: 2014,
	}
	fmt.Println("[scoped car update - before] ->", scopedCar)
	scopedCar.updateStructScopedCar()
	fmt.Println("[scoped car update - after] ->", scopedCar)

	// changing value of struct with pointer
	pointerCar := Car{
		Name: "Fiat",
		Year: 2014,
	}
	fmt.Println("[pointer car update - before] ->", pointerCar)
	pointerCar.updateStructPointerCar()
	fmt.Println("[pointer car update - after] ->", pointerCar)

}

func updateParam(param string) string {
	param = param + " Ola irmao"
	return param
}

func updatePointerParam(param *string) {
	*param = "Ola irmao"
}

func (car Car) updateStructScopedCar() {
	car.Name = "Ferrari"
	car.Year = 2020
	fmt.Println("[updated scoped car] -> Name:", car.Name, "Year:", car.Year)
}

func (car *Car) updateStructPointerCar() {
	car.Name = "Porsche"
	car.Year = 2021
	fmt.Println("[updated pointer car] -> Name:", car.Name, "Year:", car.Year)
}

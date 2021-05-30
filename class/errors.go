package class

import (
	"errors"
	"fmt"
	"log"
)

func Greetings(name string) (string, error) {
	if name == "" {
		return "", errors.New("name can't be empty")
	}
	message := fmt.Sprintf("Hello, %v.", name)
	return message, nil
}

func HandlerError() {
	message, err := Greetings("")
	if err != nil {
		// log.Fatal(err) is used generally to stop process, but, is not our goal in this moment
		log.Default().Println("[handling error when call Greetings function with empty param] ->", err)
	}
	fmt.Println(message)
}

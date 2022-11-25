package greetings

import (
	"errors"
	"fmt"
)

func Hello(name string) (string, error) {
	// check if no name was given
	if name == "" {
		return "", errors.New("empty name")
	}

	// return the greeting to entered name
	message := fmt.Sprintf("hi, %v. Welcome!", name)
	return message, nil
}
package greetings

import (
	"errors"
	"fmt"
	"math/rand"
	"time"
)

// Hello function greet a person
func Hello(name string) (string, error) {
	// check if no name was given
	if name == "" {
		return "", errors.New("empty name")
	}

	// return the greeting to entered name
	message := fmt.Sprintf(randomFormat(), name)
	return message, nil
}

// Hellos function greets people
func Hellos(names []string) (map[string]string, error) {
	messages := make(map[string]string)
	for _, name := range names {
		message, error := Hello(name)
		if error != nil {
			return nil, error
		}
		messages[name] = message
	}
	return messages, nil
}

// init sets initial values for variables used in the function.
func init() {
	rand.Seed(time.Now().UnixNano())
}

// randomFormat returns one of a set of greeting messages. The returned
// message is selected at random.
func randomFormat() string {
	// A slice of message formats.
	formats := []string{
		"Hi, %v. Welcome!",
		"Great to see you, %v!",
		"Hail, %v! Well met!",
	}

	// Return a randomly selected message format by specifying
	// a random index for the slice of formats.
	return formats[rand.Intn(len(formats))]
}
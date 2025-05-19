package greetings

import (
	"errors"
	"fmt"
	"math/rand"
)

func Hello(name string) (string, error) {
	if name == "" {
		return "", errors.New("empty value")
	}
	message := fmt.Sprintf(randomMessage(), name)
	return message, nil
}

func randomMessage() string {
	messages := []string{
		"!Hello! %v !Wellcome!",
		"!Nice to see you, %v!",
		"!Greetings, %v! !Nice to meet you!",
	}
	return messages[rand.Intn(len(messages))]
}

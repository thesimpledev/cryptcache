package main

import (
	"errors"
	"fmt"
	"os"
)

var nouns = map[string]func(string, []string) error{
	"project": projectHandler,
	"profile": profileHandler,
	"secret":  secretHandler,
}

func parse() error {
	err := commandLenValidator(len(os.Args))
	if err != nil {
		return err
	}

	noun := os.Args[2]
	verb := os.Args[1]
	args := os.Args[3:]

	cmd, exists := nouns[noun]

	if !exists {
		return fmt.Errorf("unknown noun %s", noun)
	}

	err = cmd(verb, args)

	return err
}

func commandLenValidator(length int) error {
	errorMessages := map[int]string{
		1: "expected a noun (e.g., project, profile, value)",
		2: "expected a verb (e.g., create, delete, update, merge)",
		3: "expected one or more arguments",
	}

	if msg, ok := errorMessages[length]; ok {
		return errors.New(msg)
	}

	return nil

}

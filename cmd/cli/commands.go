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
	"keypair": keypairHandler,
}

var helpers = map[string]func(){
	"project": helpProject,
	"profile": helpProfile,
	"secret":  helpSecret,
	"keypair": helpKeypair,
}

func parse() error {
	length := len(os.Args)

	if length < 2 {
		return errors.New("expected a verb (e.g., create, delete, update, merge)")
	}

	verb := os.Args[1]

	if isHelpFlag(verb) {
		helpHandler()
		return nil
	}

	if length < 3 {
		return errors.New("expected a noun (e.g., project, profile, value)")
	}

	noun := os.Args[2]

	if isHelpFlag(noun) {
		helpCmd, exists := helpers[verb]
		if !exists {
			return fmt.Errorf("unknown noun %s", verb)
		}

		helpCmd()
	}

	args := os.Args[3:]

	cmd, exists := nouns[noun]

	if !exists {
		return fmt.Errorf("unknown noun %s", noun)
	}

	return cmd(verb, args)

}

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

// Command parsing follows a "verb noun" pattern (e.g., "create project")
// for better UX, even though code is organized by noun.
// This disconnect is intentional - commands should be memorable for users,
// while code should be organized logically for developers.
//
// Help flags are checked on both verb and noun positions to allow:
// - "cryptcache create help" (general create help)
// - "cryptcache help project" or "cryptcache project help" (noun-specific help)
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

	// This is a bit messy so let me explain the logic.
	// This is so users can pass in `cryptcache project help`
	// since we pass in `verb noun` it first checks if the noun
	// is help then if it is it looks up the help function for the verb
	// There is probably a cleaner way to do this but it works for now
	// while I work out what I want the end cli command setup to look like.
	if isHelpFlag(noun) {
		helpCmd, exists := helpers[verb]
		if !exists {
			return fmt.Errorf("unknown noun %s", verb)
		}

		helpCmd()
		return nil
	}

	args := os.Args[3:]

	cmd, exists := nouns[noun]

	if !exists {
		return fmt.Errorf("unknown noun %s", noun)
	}

	return cmd(verb, args)

}

func isHelpFlag(value string) bool {
	var helpFlags = []string{"help", "--help", "-h"}
	for _, flag := range helpFlags {
		if value == flag {
			return true
		}
	}
	return false
}

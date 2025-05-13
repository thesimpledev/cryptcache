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

func isHelpFlag(value string) bool {
	var helpFlags = []string{"help", "--help", "-h"}
	for _, flag := range helpFlags {
		if value == flag {
			return true
		}
	}
	return false
}

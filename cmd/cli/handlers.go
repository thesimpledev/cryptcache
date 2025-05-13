package main

import (
	"errors"
	"flag"
	"fmt"
	"strings"
)

type project struct {
	name        string
	profile     string
	private_key string
	public_key  string
}

type profile struct {
	name        string
	private_key string
	public_key  string
}

type secret struct {
	key         string
	value       string
	encrypt     bool
	private_key string
	public_key  string
	profile     string
}

// Handlers are responsible for:
// 1. Parsing noun-specific flags (each noun has different requirements)
// 2. Validating input for the specific noun/verb combination
// 3. Routing to the appropriate action function
//
// This design keeps flag parsing close to the business logic that needs it,
// rather than trying to create a generic flag parser that would be complex
// and harder to maintain.

func projectHandler(verb string, args []string) error {
	if cryptcacheExists() {
		return fmt.Errorf("project already exists")
	}
	projectOptions := &project{}
	fs := flag.NewFlagSet("project-"+verb, flag.ExitOnError)

	fs.StringVar(&projectOptions.name, "n", "", "Project Name")
	fs.StringVar(&projectOptions.profile, "p", "", "Initial Profile Name")
	fs.StringVar(&projectOptions.private_key, "pk", "", "Private Key Path")
	fs.StringVar(&projectOptions.public_key, "pub", "", "Public Key Path")

	if verb == "-h" || verb == "--help" || verb == "help" {
		helpProfile()
	}

	if err := fs.Parse(args); err != nil {
		return fmt.Errorf("failed to parse flags: %w", err)
	}

	if projectOptions.name == "" {
		return errors.New("project name cannot be empty set with -n flag")
	}

	if projectOptions.profile == "" {
		return errors.New("profile name cannot be empty set with -p flag")
	}

	if projectOptions.private_key == "" {
		return errors.New("profile Ed25519 private key path cannot be empty set with -pk flag")
	}

	if projectOptions.public_key == "" {
		return errors.New("profile Ed25519 public key cannot be empty set with -pub flag")
	}

	switch verb {
	case "create":
		return createProject(projectOptions)
	default:
		return fmt.Errorf("unknown project verb %s", verb)
	}
}

func profileHandler(verb string, args []string) error {
	if !cryptcacheExists() {
		return fmt.Errorf("no Project Found")
	}

	profileOptions := &profile{}
	fs := flag.NewFlagSet("project-"+verb, flag.ExitOnError)

	fs.StringVar(&profileOptions.name, "n", "", "Profile Name")
	fs.StringVar(&profileOptions.private_key, "pk", "", "Ed25519 Private Encryption Key")
	fs.StringVar(&profileOptions.public_key, "pub", "", "Ed25519 Public Encryption Key")

	if err := fs.Parse(args); err != nil {
		return fmt.Errorf("failed to parse flags: %w", err)
	}

	// TODO: Add check value so only the same key can modify the profile outside of rotate

	switch verb {
	case "create":
		// TODO: Create New Profile
		return nil
	case "rotate":
		// TODO: Roatete Encryption Keys, Update all Signatures and Values in Profile
		return nil
	case "delete":
		// TODO: Delete profile requires private key
		return nil
	case "export":
		// TODO: Export to .env requires public key
		// Will expand to other exports in the future
		return nil
	case "diff":
		// TODO: Compare two profiles to see missing keys
		return nil
	case "verify":
		// TODO: Verify Profiles Cryptographic Signature
		return nil
	case "set":
		// TODO: Set Active Profile
		return nil
	default:
		return fmt.Errorf("unknown profile verb %s", verb)

	}
}

func secretHandler(verb string, args []string) error {
	if !cryptcacheExists() {
		return fmt.Errorf("no Project Found")
	}

	secretOptions := &secret{}
	fs := flag.NewFlagSet("secret-"+verb, flag.ExitOnError)

	fs.StringVar(&secretOptions.key, "k", "", "Secret Key Name")
	fs.StringVar(&secretOptions.value, "v", "", "Secret Value")
	fs.BoolVar(&secretOptions.encrypt, "e", true, "Encrypt Secret - default true")

	fs.StringVar(&secretOptions.profile, "n", "", "Profile Name if not set or to override set")
	fs.StringVar(&secretOptions.public_key, "pub", "", "Ed25519 Public Key if not set or to override set")
	fs.StringVar(&secretOptions.private_key, "pk", "", "Ed25519 Private key if not set or to override set")

	if err := fs.Parse(args); err != nil {
		return fmt.Errorf("failed to parse flags: %w", err)
	}

	// encryption is a protected namespace
	// Private and Public Key Optional if set in profile
	// Has to have an active profile set or profile name passed in

	switch verb {
	case "create":
		// TODO: Create new Secret (Encrypted or Unencrypted)
		return nil
	case "update":
		// TODO: Update Secrete
		return nil
	case "delete":
		// TODO: Delete Secret - Requires private key
		return nil
	case "view":
		// TODO: View Secret - Requires public key
		return nil
	case "verify":
		// TODO: Verify Secret
		return nil
	default:
		return fmt.Errorf("unknown secret verb %s", verb)

	}
}

func keypairHandler(verb string, args []string) error {
	switch verb {
	case "create":
		return generateEd25519KeyPair()
	default:
		return fmt.Errorf("unknown keypair verb %s", verb)
	}
}

//TODO: Not currently in use but more robust error handling to be implemented.
// May want to abstract it more so we can pass in groups of items to validate
// Would be nice to send in a structure that has all of these items
// Maybe create a struct to hold the value, flag, and error message in an array of structs

func (p *project) validate() error {
	var errs []string

	if p.name == "" {
		errs = append(errs, "project name is required (-n flag)")
	}
	if p.profile == "" {
		errs = append(errs, "profile name is required (-p flag)")
	}
	if p.private_key == "" {
		errs = append(errs, "private key path is required (-pk flag)")
	}
	if p.public_key == "" {
		errs = append(errs, "public key path is required (-pub flag)")
	}

	if len(errs) > 0 {
		return fmt.Errorf("validation failed: %s", strings.Join(errs, "; "))
	}
	return nil
}

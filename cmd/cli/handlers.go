package main

import (
	"errors"
	"flag"
	"fmt"
)

type project struct {
	name             string
	profile          string
	private_key_path string
	public_key_path  string
}

type profile struct {
	name             string
	private_key_path string
	public_key_path  string
	private_key      string
	public_key       string
}

type secret struct {
	key   string
	value string
}

func projectHandler(verb string, args []string) error {
	if cryptcacheExists() {
		return fmt.Errorf("project already exists")
	}
	projectOptions := &project{}
	fs := flag.NewFlagSet("project-"+verb, flag.ExitOnError)

	fs.StringVar(&projectOptions.name, "n", "", "Project Name")
	fs.StringVar(&projectOptions.profile, "p", "", "Initial Profile Name")
	fs.StringVar(&projectOptions.private_key_path, "pk", "", "Private Key Path")
	fs.StringVar(&projectOptions.public_key_path, "pub", "", "Public Key Path")

	if err := fs.Parse(args); err != nil {
		return fmt.Errorf("failed to parse flags: %w", err)
	}

	if projectOptions.name == "" {
		return errors.New("project name cannot be empty set with -n flag")
	}

	if projectOptions.profile == "" {
		return errors.New("profile name cannot be empty set with -p flag")
	}

	if projectOptions.private_key_path == "" {
		return errors.New("profile Ed25519 private key path cannot be empty set with -pk flag")
	}

	if projectOptions.public_key_path == "" {
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

	//add check value in so only the same key can modify this profile
	//When updating key reencrypt all values

	switch verb {
	case "create":
		return nil
	case "update":
		//update public and private keys
		return nil
	case "delete":
		return nil
	case "export":
		return nil
	case "diff":
		return nil
	case "load-pk":
		//load private key into memory
		return nil
	case "verify":
		return nil
	default:
		return fmt.Errorf("unknown profile verb %s", verb)

	}

}

func secretHandler(verb string, args []string) error {
	if !cryptcacheExists() {
		return fmt.Errorf("no Project Found")
	}

	//encrypted vs not encrypted
	//encryption is a protected namespace

	switch verb {
	case "create":
		return nil
	case "update":
		//update public and private keys
		return nil
	case "delete":
		return nil
	case "view":
		return nil
	case "diff":
		return nil
	case "load-pk":
		//load private key into memory
		return nil
	case "verify":
		return nil
	default:
		return fmt.Errorf("unknown profile verb %s", verb)

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

func helpHandler() {
	fmt.Println("Commands:")
	fmt.Println("  create project    Create a new CryptCache project with the specified profile and keys")
	fmt.Println("      -n            Project name (required)")
	fmt.Println("      -p            Profile name (required)")
	fmt.Println("      -pk           Path to Ed25519 private key file (required)")
	fmt.Println("      -pub          Path to Ed25519 public key file or URL (required)")
	fmt.Println()
	fmt.Println("  create keypair    Generates an Ed25519 keypair in the current directory")
	fmt.Println("                    Files: 'id_ed25519' (private), 'id_ed25519.pub' (public)")
	fmt.Println()
	fmt.Println("⚠️  Make sure to secure 'id_ed25519' and avoid committing it to version control.")
}

package main

import "fmt"

func helpHandler() {
	fmt.Println("Usage: cryptcache <verb> <noun> [flags]")
	fmt.Println()
	fmt.Println("Available Commands:")
	fmt.Println()

	helpProject()
	helpKeypair()
	helpProfile()
	helpSecret()

	fmt.Println("Global Options:")
	fmt.Println("  --help     Show help for the CLI")
	fmt.Println()
	fmt.Println("Use '-h' after any verb+noun for more details, e.g.:")
	fmt.Println("  cryptcache create project -h")
}

func helpProject() {
	fmt.Println("create project    Create a new CryptCache project")
	fmt.Println("    -n   Project name (required)")
	fmt.Println("    -p   Profile name (required)")
	fmt.Println("    -pk  Ed25519 private key path (required)")
	fmt.Println("    -pub Ed25519 public key path or URL (required)")
	fmt.Println()
}

func helpKeypair() {
	fmt.Println("create keypair    Generate an Ed25519 keypair in the current directory")
	fmt.Println("                  Files: id_ed25519 (private), id_ed25519.pub (public)")
	fmt.Println()
}

func helpProfile() {
	fmt.Println("profile commands")
	fmt.Println("  create          Create a new profile with associated keys")
	fmt.Println("  rotate          Replace profile keys and update associated values")
	fmt.Println("  delete          Delete profile (requires private key)")
	fmt.Println("  export          Export to .env format")
	fmt.Println("  diff            Show key differences between profiles")
	fmt.Println("  verify          Validate cryptographic integrity of profile")
	fmt.Println("  set             Set profile as active")
	fmt.Println()
}

func helpSecret() {
	fmt.Println("secret commands")
	fmt.Println("  create          Store a new secret")
	fmt.Println("  update          Change a secret’s value")
	fmt.Println("  delete          Remove a secret (requires private key)")
	fmt.Println("  view            Show a secret (requires public key)")
	fmt.Println("  verify          Check a secret’s signature")
	fmt.Println()
}

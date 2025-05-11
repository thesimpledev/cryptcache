package main

import (
	"crypto/ed25519"
	"fmt"
	"os"
	"path/filepath"
)

func createProject(project *project) error {
	data := map[string]any{
		"title":    project.name,
		"profiles": []string{project.profile},
		fmt.Sprintf("%s.encryption", project.profile): map[string]any{
			"public_key_path":  project.public_key,
			"private_key_path": project.private_key,
		},
	}

	sortedBytes, err := sortTOML(data)
	if err != nil {
		return err
	}

	cwd, err := os.Getwd()
	if err != nil {
		return err
	}

	filePath := filepath.Join(cwd, "cryptcache.toml")
	err = os.WriteFile(filePath, sortedBytes, 0644)
	if err != nil {
		return err
	}

	return nil
}

func generateEd25519KeyPair() error {
	publicKey, privateKey, err := ed25519.GenerateKey(nil)
	if err != nil {
		return fmt.Errorf("failed to generate ed25519 key pair: %w", err)
	}

	err = os.WriteFile("id_ed25519", privateKey, 0600)
	if err != nil {
		return fmt.Errorf("failed to write private key: %w", err)
	}

	err = os.WriteFile("id_ed25519.pub", publicKey, 0644)
	if err != nil {
		return fmt.Errorf("failed to write public key: %w", err)
	}

	fmt.Println("✅ Ed25519 key pair generated.")
	fmt.Println("🔒 Make sure to secure 'id_ed25519' (private key) and do not commit it to source control.")
	return nil
}

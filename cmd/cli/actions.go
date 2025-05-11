package main

import (
	"fmt"
	"os"
	"path/filepath"
)

func createProject(project *project) error {
	data := map[string]any{
		"title":    project.name,
		"profiles": []string{project.profile},
		fmt.Sprintf("%s.encryption", project.profile): map[string]any{
			"public_key_path":  project.public_key_path,
			"private_key_path": project.private_key_path,
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

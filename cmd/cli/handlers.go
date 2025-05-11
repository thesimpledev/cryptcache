package main

import (
	"errors"
	"flag"
	"fmt"
	"os"
	"path/filepath"
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
	fs.StringVar(&projectOptions.profile, "p", "default", "Initial Profile Name")
	fs.StringVar(&projectOptions.private_key_path, "pk", "", "Private Key Path")
	fs.StringVar(&projectOptions.public_key_path, "pub", "", "Public Key Path")

	if err := fs.Parse(args); err != nil {
		return fmt.Errorf("failed to parse flags: %w", err)
	}

	if projectOptions.name == "" {
		return errors.New("project name cannot be empty")
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
		return fmt.Errorf("No Project Found")
	}
	//create, delete, export, diff, privaet key path, public key path, private key, public key

	return nil
}

func secretHandler(verb string, args []string) error {
	if !cryptcacheExists() {
		return fmt.Errorf("No Project Found")
	}
	//create, delete, update, view
	//encrypted vs not encrypted
	//encryption is a protected namespace
	//public key can be on https path or local file

	return nil
}

func cryptcacheExists() bool {
	cwd, err := os.Getwd()
	if err != nil {
		return false
	}

	path := filepath.Join(cwd, "cryptcache.toml")
	_, err = os.Stat(path)
	return err == nil
}

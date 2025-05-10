package main

import (
	"errors"
	"flag"
	"fmt"
)

type project struct {
	name string
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
		return fmt.Errorf("unknown Project Verb %s", verb)
	}

}

func profileHandler(verb string, args []string) error {
	//create, delete, export, diff, privaet key path, public key path, private key, public key

	return nil
}

func secretHandler(verb string, args []string) error {
	//create, delete, update, view
	//encrypted vs not encrypted

	return nil
}

func cryptcacheExists() bool {
	return false
}

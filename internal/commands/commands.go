package commands

import (
	"flag"
	"fmt"
	"os"
)

type Flags struct {
	InitProject bool
	NewProfile  bool
	Set         bool
	Get         bool
	Export      bool
	Name        string
	File        string
	Key         string
	Value       string
	Private     string
	Public      string
}

func ParseCommands(commands []string) {
	if len(commands) < 2 {
		fmt.Println("Expected a subcommand. Usage: cryptcache <command> [options]")
		os.Exit(1)
	}
	subCommand := commands[1]

	fmt.Println("Subcommand: ", subCommand)

	switch subCommand {
	case "init":
		projectInitialization(commands)
	case "profile":
		profileCommands(commands)
	default:
		fmt.Printf("Unknown command: %s\n", subCommand)
		fmt.Println("Available commands: init, profile, set, get")
		os.Exit(1)
	}
}

func projectInitialization(commands []string) {
	fmt.Printf("Initializing project %s\n", "")
	initCmd := flag.NewFlagSet("init", flag.ExitOnError)
	projectName := initCmd.String("n", "", "Project name")
	profileName := initCmd.String("p", "", "Initial profile name")
	initCmd.Parse(commands[2:])

	if *projectName == "" {
		fmt.Println("Project name is required")
		os.Exit(1)
	}

	if *profileName == "" {
		fmt.Println("Profile name is required")
		os.Exit(1)
	}

	fmt.Printf("Project Name: %s\n", *projectName)
	fmt.Printf("Profile Name: %s\n", *profileName)
	//TODO: Somehow Handle the initialization of the project
}

func profileCommands(commands []string) {
	fmt.Printf("Creating new profile %s\n", "")
	profileCmd := flag.NewFlagSet("profile", flag.ExitOnError)
	profileAction := commands[2]

	switch profileAction {
	case "add":
		addProfile(profileCmd, commands[3:])
	case "remove":
		removeProfile(profileCmd, commands[3:])
	case "list":
		list(profileCmd, commands[3:])
	case "checkout":
		checkoutProfile()
	case "set":
		setKeyValuePair(profileCmd, commands[3:])
	case "get":
		getKeyValuePair(profileCmd, commands[3:])
	case "update":
		updateKeyValuePair(profileCmd, commands[3:])
	case "export":
		exportToFile(profileCmd, commands[3:])
	default:
		fmt.Printf("Unknown profile command: %s\n", profileAction)
		fmt.Println("Available commands: checkout, add, remove, list")
		os.Exit(1)
	}
}

func checkoutProfile() {
	//TODO: Check if profile exists
	//TODO: Check if user has access to profile
	//TODO: Checkout a profile by setting it to active
}

func addProfile(profileCmd *flag.FlagSet, args []string) {

	profileName := profileCmd.String("n", "", "Profile name")
	if profileName == nil {
		fmt.Println("Profile name is required")
		os.Exit(1)
	}

	profileCmd.Parse(args)
	//TODO: Check if profile already exists and if so throw error
	//TODO: Create profile if it doesn't exist
	fmt.Printf("Adding profile %s\n", "")
}

func removeProfile(profileCmd *flag.FlagSet, args []string) {
	profileName := profileCmd.String("n", "", "Profile name")
	if profileName == nil {
		fmt.Println("Profile name is required")
		os.Exit(1)
	}

	profileCmd.Parse(args)

	//TODO: Check user has correct key to profile to delete
	//TODO: Delete profile

}

func list(profileCmd *flag.FlagSet, args []string) {
	if profileCmd.NArg() != 1 {
		fmt.Println("Error: Key name must be provided")
		fmt.Println("Usage: cryptcache get [flags] <key>")
		profileCmd.PrintDefaults()
		os.Exit(1)
	}

	path := profileCmd.Arg(0) // First non-flag argument

	switch path {
	case "profiles":
		fmt.Println("Listing all profiles")
	case "keys":
		fmt.Println("Listing all keys")
	default:
		fmt.Println("Unknown path: ", path)
		os.Exit(1)
	}

}

func setKeyValuePair(setCmd *flag.FlagSet, args []string) {
	if setCmd.NArg() != 2 {
		fmt.Println("Error: Both key and value must be provided")
		fmt.Println("Usage: cryptcache set [flags] <key> <value>")
		setCmd.PrintDefaults()
		os.Exit(1)
	}

	key := setCmd.Arg(0)   // First non-flag argument
	value := setCmd.Arg(1) // Second non-flag argument
	//TODO: Check if key exists in profile and throw error if it does

	fmt.Printf("Setting '%s' to '%s' (will be encrypted)\n", key, value)
	//TODO: Encrypt value and set key value pair in profile
}

func getKeyValuePair(getCmd *flag.FlagSet, args []string) {
	// Check if we have exactly 1 non-flag argument (key)
	if getCmd.NArg() != 1 {
		fmt.Println("Error: Key name must be provided")
		fmt.Println("Usage: cryptcache get [flags] <key>")
		getCmd.PrintDefaults()
		os.Exit(1)
	}

	// Get the key from arguments
	key := getCmd.Arg(0) // First non-flag argument

	fmt.Printf("Getting value for key '%s' (will be decrypted)\n", key)
	// TODO: Implement get logic
}

func updateKeyValuePair(updateCmd *flag.FlagSet, args []string) {
	if updateCmd.NArg() != 2 {
		fmt.Println("Error: Both key and value must be provided")
		fmt.Println("Usage: cryptcache set [flags] <key> <value>")
		updateCmd.PrintDefaults()
		os.Exit(1)
	}

	key := updateCmd.Arg(0)   // First non-flag argument
	value := updateCmd.Arg(1) // Second non-flag argument

	//TODO: Check if key exists in profile and throw error if it doesn't

	fmt.Printf("Setting '%s' to '%s' (will be encrypted)\n", key, value)
	//TODO: Encrypt value and set key value pair in profile
}

func exportToFile(exportCmd *flag.FlagSet, args []string) {
	fileName := exportCmd.String("f", "", "File name to export to")
	if fileName == nil {
		fmt.Println("File name is required")
		os.Exit(1)
	}

	exportCmd.Parse(args)
	fmt.Printf("Exporting profile to file: %s\n", *fileName)
}

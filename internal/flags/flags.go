package flags

import "flag"

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

func ParseFlags() Flags {
	flags := Flags{}
	flag.BoolVar(&flags.InitProject, "init", false, "Initialize the project")
	flag.BoolVar(&flags.NewProfile, "new", false, "Create a new profile")
	flag.StringVar(&flags.Name, "n", "Definitely not a Mimic", "Set the project name")

	flag.BoolVar(&flags.Set, "set", false, "Set/Update a key value pair on current profile")
	flag.BoolVar(&flags.Get, "get", false, "Get a value on current profile")
	flag.BoolVar(&flags.Export, "export", false, "Export the current profile to env file")

	flag.StringVar(&flags.Key, "k", "", "Set the key")
	flag.StringVar(&flags.Value, "v", "", "Set the value")
	flag.StringVar(&flags.File, "f", "", "Set the file")

	flag.StringVar(&flags.Private, "private", "~/.ssh/id_rsa", "Set the private key")
	flag.StringVar(&flags.Public, "public", "~/.ssh/id_rsa.pub", "Set the public key")
	flag.Parse()
	return flags
}

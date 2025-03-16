package main

import (
	"os"

	"github.com/thesimpledev/cryptcache/internal/commands"
)

func main() {
	commands.ParseCommands(os.Args)

	// if flags.InitProject {
	// 	fmt.Printf("Initializing project %s\n", flags.Name)
	// 	//TODO: Make sure to throw and error is the project already exists
	// 	return
	// }

	// if flags.NewProfile {
	// 	fmt.Printf("creating new profile %s\n", flags.Name)
	// 	//TODO: Make sure to throw an error if the profile already exists
	// 	return
	// }

	// if flags.Key != "" {
	// 	fmt.Printf("Setting key %s\n", flags.Key)
	// 	//TODO: Add to all profiles even if value empty
	// }

}

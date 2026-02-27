package cli

import (
	"fmt"
	"os"
)

func printUsage() {
	fmt.Println("Usage:")
}

func DispatchCommands(args []string) {
	if len(args) == 1 {
		return
	}
	cmd := args[1]
	switch cmd {
	case "help":
		printUsage()
		os.Exit(0)
	case "create-user":
		os.Exit(createUser(args[2:]))
	default:
		fmt.Fprintf(os.Stderr, "Unrecognized command '%s'\n", cmd)
		printUsage()
		os.Exit(1)
	}
}

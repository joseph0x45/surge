package cli

import (
	"fmt"
	"os"

	"github.com/joseph0x45/goutils"
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
	case "service-file":
		goutils.GenerateServiceFile("Surge, simple sitting time tracker")
    os.Exit(0)
	default:
		fmt.Fprintf(os.Stderr, "Unrecognized command '%s'\n", cmd)
		printUsage()
		os.Exit(1)
	}
}

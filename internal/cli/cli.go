package cli

import (
	"fmt"
	"os"

	"github.com/joseph0x45/goutils"
	"github.com/joseph0x45/surge/internal/buildinfo"
)

func printUsage() {
	fmt.Printf(`Surge %s

Usage:
  surge <command> [flags]

Commands:
  version                        Print the current version
  help                           Print this help message
  create-user                    Create a new user
    -username  string            Username
    -password  string            Password
    -limit     int               Max number of seconds to sit
  service-file                   Generate a systemd service file

`, buildinfo.Version)
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
	case "version":
		fmt.Println(buildinfo.AppName, buildinfo.Version)
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

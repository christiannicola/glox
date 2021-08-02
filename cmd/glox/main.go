package main

import (
	"fmt"
	"os"
)

var (
	// Version will be set with ldflags and represents the version of the interpreter
	Version string
	// BuildDate will be set with ldflags and shows the time stamp on when the interpreter was built
	BuildDate string
)

func main() {
	arguments := os.Args[1:]

	if len(arguments) > 1 {
		showHelp()
		os.Exit(64)
	}

	if len(arguments) == 1 {
		switch arguments[0] {
		case "-v":
			fallthrough
		case "--version":
			fmt.Println(version())
		default:
			if err := runFile(arguments[0]); err != nil {
				exitWithError(err)
			}
		}

		os.Exit(0)
	}

	if err := runPrompt(); err != nil {
		exitWithError(err)
	}
}

func showHelp() {
	fmt.Println("Usage: glox [options] [script]")
	fmt.Println()
	fmt.Println("Options:")
	fmt.Println("  -v, --version\t\t\tprint glox version")
}

func version() string {
	return fmt.Sprintf("%s (built %s)", Version, BuildDate)
}

func exitWithError(err error) {
	_, _ = fmt.Fprint(os.Stderr, err.Error())
	os.Exit(74)
}

func runFile(path string) error {
	return nil
}

func runPrompt() error {
	fmt.Printf("Welcome to glox %s", version())
	return nil
}

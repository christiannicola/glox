package main

import (
	"bufio"
	"fmt"
	"github.com/christiannicola/glox/internal/lox"
	"github.com/christiannicola/glox/internal/lox/ast"
	"io/ioutil"
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
				printErrorAndExit(err, 74)
			}
		}

		os.Exit(0)
	}

	if err := runPrompt(); err != nil {
		printErrorAndExit(err, 74)
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

func printErrorAndExit(err error, code int) {
	lox.PrintError(err)
	os.Exit(code)
}

func runFile(path string) error {
	file, err := os.OpenFile(path, os.O_RDONLY, 744)
	if err != nil {
		return err
	}

	defer file.Close()

	source, err := ioutil.ReadAll(file)
	if err != nil {
		return err
	}

	return run(source)
}

func runPrompt() error {
	fmt.Printf("Welcome to glox %s\n\n", version())

	streamScanner := bufio.NewReader(os.Stdin)

	for {
		fmt.Print("> ")
		line, err := streamScanner.ReadBytes('\n')
		if err != nil {
			return err
		}

		if err = run(line); err != nil {
			lox.PrintError(err)
			fmt.Println()
		}
	}
}

func run(source []byte) error {
	sourceScanner := lox.NewScanner(string(source))

	tokens, err := sourceScanner.ScanTokens()
	if err != nil {
		return err
	}

	parser := lox.NewParser(tokens)
	expression, err := parser.Parse()

	if err != nil {
		return err
	}

	printer := ast.NewPrinter()

	message, err := printer.Print(expression)

	if err != nil {
		return err
	}

	fmt.Println(message)

	return nil
}

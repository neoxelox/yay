package main

import (
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
	"runtime"
	"strings"
	"time"

	"github.com/mkideal/cli"
	"github.com/neoxelox/yay/mod"
)

func read(filepath string) (string, error) {
	program, err := os.ReadFile(filepath)

	if err != nil {
		return "", fmt.Errorf("failed to read program %s: %w", filepath, err)
	}

	return string(program), nil
}

func write(filepath string, program string) error {
	file, err := os.Create(filepath)
	if err != nil {
		return fmt.Errorf("failed to write program %s: %w", filepath, err)
	}
	defer file.Close()

	_, err = file.WriteString(program)
	if err != nil {
		return fmt.Errorf("failed to write program %s: %w", filepath, err)
	}

	return nil
}

func exit(ctx *cli.Context, err error) {
	ctx.String("%s %s\n", ctx.Color().Red("Error:"), err)
	os.Exit(1)
}

var root = &cli.Command{
	Desc: "YAY is an stack-oriented language that transpiles to Go code",
	Fn: func(ctx *cli.Context) error {
		ctx.WriteUsage()
		os.Exit(1)
		return nil
	},
}

var help = cli.HelpCommand("display help information")

var version = &cli.Command{
	Name: "version",
	Desc: "shows YAY version",
	Fn: func(ctx *cli.Context) error {
		ctx.String("YAY version %s %s/%s\n", ctx.Color().Yellow(mod.Version), runtime.GOOS, runtime.GOARCH)
		return nil
	},
}

type runArgs struct {
	cli.Helper
	InputFile string `cli:"*input" usage:"path of the file to run"`
}

var run = &cli.Command{
	Name: "run",
	Desc: "transpiles and executes a YAY program",
	Argv: func() interface{} { return new(runArgs) },
	Fn: func(ctx *cli.Context) error {
		args := ctx.Argv().(*runArgs)
		inputFilename := filepath.Base(args.InputFile)

		// Read input program

		inputProgram, err := read(args.InputFile)
		if err != nil {
			exit(ctx, err)
		}

		// Parse input program

		start := time.Now()

		tokens, err := parse(inputProgram, args.InputFile)
		if err != nil {
			exit(ctx, err)
		}

		parseElapsed := time.Since(start)

		// Transpile input program

		start = time.Now()

		outputProgram, err := transpile(tokens, args.InputFile)
		if err != nil {
			exit(ctx, err)
		}

		transpileElapsed := time.Since(start)

		// Write output program

		outputFile := filepath.Join(os.TempDir(), (strings.TrimSuffix(inputFilename, filepath.Ext(inputFilename)) + ".go"))
		// outputFile := strings.TrimSuffix(args.InputFile, filepath.Ext(args.InputFile)) + ".go"

		err = write(outputFile, outputProgram)
		if err != nil {
			exit(ctx, err)
		}

		// Format output program

		err = exec.Command("go", "fmt", outputFile).Run()
		if err != nil {
			exit(ctx, fmt.Errorf("failed to format program %s: %w", outputFile, err))
		}

		// Execute output program

		outputExecutable := exec.Command("go", "run", outputFile)
		outputExecutable.Stdout = os.Stdout
		outputExecutable.Stderr = os.Stderr
		outputExecutable.Stdin = os.Stdin

		start = time.Now()

		err = outputExecutable.Run()
		if err != nil {
			exit(ctx, fmt.Errorf("failed to execute program %s: %w", outputFile, err))
		}

		executeElapsed := time.Since(start)

		// Remove output program

		err = os.Remove(outputFile)
		if err != nil {
			exit(ctx, fmt.Errorf("failed to remove program %s: %w", outputFile, err))
		}

		ctx.String("Parsing     %s took %s\n", ctx.Color().Yellow(inputFilename), ctx.Color().Green(parseElapsed.String()))
		ctx.String("Transpiling %s took %s\n", ctx.Color().Yellow(inputFilename), ctx.Color().Green(transpileElapsed.String()))
		ctx.String("Executing   %s took %s\n", ctx.Color().Yellow(inputFilename), ctx.Color().Green(executeElapsed.String()))
		ctx.String("Run         %s total %s\n", ctx.Color().Yellow(inputFilename), ctx.Color().Cyan((parseElapsed + transpileElapsed + executeElapsed).String()))

		return nil
	},
}
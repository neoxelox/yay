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

var rootCmd = &cli.Command{
	Desc: "YAY is an stack-oriented language that transpiles to Go code",
	Fn: func(ctx *cli.Context) error {
		ctx.WriteUsage()
		os.Exit(1)
		return nil
	},
}

var helpCmd = cli.HelpCommand("display help information")

var versionCmd = &cli.Command{
	Name: "version",
	Desc: "shows YAY version",
	Fn: func(ctx *cli.Context) error {
		ctx.String("YAY version %s %s/%s\n", ctx.Color().Yellow(mod.Version), runtime.GOOS, runtime.GOARCH)
		return nil
	},
}

type runArgs struct {
	cli.Helper
	InputFile string `cli:"*input" usage:"path of the yay program to run"`
}

var runCmd = &cli.Command{
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

		err = write(outputFile, outputProgram)
		if err != nil {
			exit(ctx, err)
		}

		// Execute output program

		outputExecutable := exec.Command("go", "run", outputFile)
		outputExecutable.Stdout = os.Stdout
		outputExecutable.Stderr = os.Stderr
		outputExecutable.Stdin = os.Stdin

		start = time.Now()

		err = outputExecutable.Run()
		if err != nil {
			exit(ctx, fmt.Errorf("failed to execute program %s: %w", inputFilename, err))
		}

		executeElapsed := time.Since(start)

		// Remove output program

		err = os.Remove(outputFile)
		if err != nil {
			exit(ctx, fmt.Errorf("failed to remove program %s: %w", inputFilename, err))
		}

		ctx.String("Parsing     %s took %s\n", ctx.Color().Yellow(inputFilename), ctx.Color().Green(parseElapsed.String()))
		ctx.String("Transpiling %s took %s\n", ctx.Color().Yellow(inputFilename), ctx.Color().Green(transpileElapsed.String()))
		ctx.String("Executing   %s took %s\n", ctx.Color().Yellow(inputFilename), ctx.Color().Green(executeElapsed.String()))
		ctx.String("Run         %s total %s\n", ctx.Color().Yellow(inputFilename), ctx.Color().Cyan((parseElapsed + transpileElapsed + executeElapsed).String()))

		return nil
	},
}

type buildArgs struct {
	cli.Helper
	InputFile  string `cli:"*input" usage:"path of the yay program to build"`
	OutputFile string `cli:"output" usage:"path of the builded program"`
}

var buildCmd = &cli.Command{
	Name: "build",
	Desc: "transpiles and compiles a YAY program",
	Argv: func() interface{} { return new(buildArgs) },
	Fn: func(ctx *cli.Context) error {
		args := ctx.Argv().(*buildArgs)
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

		err = write(outputFile, outputProgram)
		if err != nil {
			exit(ctx, err)
		}

		// Build output program

		outputExecutable := strings.TrimSuffix(args.InputFile, filepath.Ext(args.InputFile))
		if len(args.OutputFile) > 0 {
			outputExecutable = args.OutputFile
		}

		start = time.Now()

		err = exec.Command("go", "build", "-o", outputExecutable, outputFile).Run()
		if err != nil {
			exit(ctx, fmt.Errorf("failed to build program %s: %w", inputFilename, err))
		}

		buildElapsed := time.Since(start)

		// Remove output program

		err = os.Remove(outputFile)
		if err != nil {
			exit(ctx, fmt.Errorf("failed to remove program %s: %w", inputFilename, err))
		}

		ctx.String("Parsing     %s took %s\n", ctx.Color().Yellow(inputFilename), ctx.Color().Green(parseElapsed.String()))
		ctx.String("Transpiling %s took %s\n", ctx.Color().Yellow(inputFilename), ctx.Color().Green(transpileElapsed.String()))
		ctx.String("Building    %s took %s\n", ctx.Color().Yellow(inputFilename), ctx.Color().Green(buildElapsed.String()))
		ctx.String("Build       %s total %s\n", ctx.Color().Yellow(inputFilename), ctx.Color().Cyan((parseElapsed + transpileElapsed + buildElapsed).String()))

		return nil
	},
}

type transpileArgs struct {
	cli.Helper
	InputFile  string `cli:"*input" usage:"path of the yay program to transpile"`
	OutputFile string `cli:"output" usage:"path of the transpiled program"`
}

var transpileCmd = &cli.Command{
	Name: "transpile",
	Desc: "transpiles to Go a YAY program",
	Argv: func() interface{} { return new(transpileArgs) },
	Fn: func(ctx *cli.Context) error {
		args := ctx.Argv().(*transpileArgs)
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

		outputFile := strings.TrimSuffix(args.InputFile, filepath.Ext(args.InputFile)) + ".go"
		if len(args.OutputFile) > 0 {
			outputFile = args.OutputFile
		}

		err = write(outputFile, outputProgram)
		if err != nil {
			exit(ctx, err)
		}

		// Format output program

		err = exec.Command("go", "fmt", outputFile).Run()
		if err != nil {
			exit(ctx, fmt.Errorf("failed to format program %s: %w", inputFilename, err))
		}

		ctx.String("Parsing     %s took %s\n", ctx.Color().Yellow(inputFilename), ctx.Color().Green(parseElapsed.String()))
		ctx.String("Transpiling %s took %s\n", ctx.Color().Yellow(inputFilename), ctx.Color().Green(transpileElapsed.String()))
		ctx.String("Transpile   %s total %s\n", ctx.Color().Yellow(inputFilename), ctx.Color().Cyan((parseElapsed + transpileElapsed).String()))

		return nil
	},
}

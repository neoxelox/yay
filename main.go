package main

import (
	"bytes"
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
	"strconv"
	"strings"
	"text/template"
	"time"
	"unicode"

	"github.com/mkideal/cli"
	"github.com/neoxelox/yay/mod"
	"github.com/neoxelox/yay/std"
)

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

func main() {
	if err := cli.Root(root,
		cli.Tree(help),
		cli.Tree(run),
	).Run(os.Args[1:]); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}

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

var Identifiers = map[string]mod.Identifier{
	std.LiteralAdd:   &std.Add{},
	std.LiteralSub:   &std.Sub{},
	std.LiteralMul:   &std.Mul{},
	std.LiteralDiv:   &std.Div{},
	std.LiteralPrint: &std.Print{},
}

func isNumber(literal string) bool {
	if _, err := strconv.ParseInt(literal, 10, 64); err == nil {
		return true
	}

	if _, err := strconv.ParseFloat(literal, 64); err == nil {
		return true
	}

	return false
}

func isIdentifier(literal string) bool {
	if _, ok := Identifiers[literal]; ok {
		return true
	}

	return false
}

func isComment(literal string) bool {
	return literal == mod.LiteralComment
}

func searchRune(r rune, runes []rune, start int) int {
	for i := start; i < len(runes); i++ {
		if runes[i] == r {
			return i
		}
	}

	return -1
}

func parse(program string, filepath string) ([]mod.Token, error) {
	var tokens []mod.Token
	var literal []rune
	var row int
	var col int
	runes := []rune(program)
	runes = append(runes, '\n') // Always have at least one end of line

	for cur := 0; cur < len(runes); cur++ {
		if unicode.IsSpace(runes[cur]) || isComment(string(runes[cur])) {
			if cur != 0 && !unicode.IsSpace(runes[cur-1]) {
				aCol := col - len(literal)
				aLiteral := string(literal)

				switch {
				case isNumber(aLiteral):
					tokens = append(tokens, mod.Token{
						Type:    mod.TypeNumber,
						Literal: aLiteral,
						File:    filepath,
						Row:     row,
						Col:     aCol,
					})
				case isIdentifier(aLiteral):
					token, err := Identifiers[aLiteral].Parse(aLiteral, filepath, row, aCol)
					if err != nil {
						return nil, fmt.Errorf("cannot parse identifier '%s' at %s:%d:%d", aLiteral, filepath, row+1, aCol+1)
					}
					tokens = append(tokens, token)
				default:
					return nil, fmt.Errorf("unknown token '%s' at %s:%d:%d", aLiteral, filepath, row+1, aCol+1)
				}

				literal = []rune{}
			}
		} else {
			literal = append(literal, runes[cur])
		}

		col++

		if runes[cur] == '\n' {
			row++
			col = 0
		} else if isComment(string(runes[cur])) {
			row++
			col = 0
			cur = searchRune('\n', runes, cur+1)
		}
	}

	return tokens, nil
}

func transpile(program []mod.Token, filepath string) (string, error) {
	var transpilation bytes.Buffer
	template := template.Must(template.New("").Parse(mod.TranspileBase))

	var imports []string
	var definitions []string
	var statements []string
	importSet := make(map[string]struct{})

	for _, token := range program {
		switch token.Type {
		case mod.TypeNumber:
			statements = append(statements, fmt.Sprintf(`push(&stack, %s)`, token.Literal))
		case mod.TypeIdentifier:
			iImports, iDefinitions, iStatements, err := Identifiers[token.Literal].Transpile(token)
			if err != nil {
				return "", fmt.Errorf("cannot transpile identifier '%s' at %s:%d:%d", token.Literal, token.File, token.Row+1, token.Col+1)
			}
			for _, iImport := range iImports {
				importSet[iImport] = struct{}{}
			}
			if len(iDefinitions) > 0 {
				definitions = append(definitions, iDefinitions)
			}
			if len(iStatements) > 0 {
				statements = append(statements, iStatements)
			}
		default:
			return "", fmt.Errorf("unknown token type '%s'", token.Type)
		}
	}

	for imprt := range importSet {
		imports = append(imports, imprt)
	}

	err := template.Execute(&transpilation, mod.TranspileData{
		Version:     mod.Version,
		Imports:     imports,
		Definitions: definitions,
		Statements:  statements,
	})

	if err != nil {
		return "", fmt.Errorf("cannot transpile file '%s': %w", filepath, err)
	}

	return transpilation.String(), nil
}

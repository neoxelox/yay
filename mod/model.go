package mod

const Version = "v0.0.1"

const (
	LiteralString  = '"'
	LiteralComment = '#'
)

type Type string

const (
	TypeNumber     = Type("NUMBER")
	TypeString     = Type("STRING")
	TypeIdentifier = Type("IDENTIFIER")
	TypeComment    = Type("COMMENT")
)

type Token struct {
	Type    Type
	Literal string
	File    string
	Row     int
	Col     int
	Meta    interface{}
}

type Identifier interface {
	Parse(literal string, file string, row int, col int) (Token, error)
	Transpile(token Token) ([]string, string, string, error)
}

type TranspileData struct {
	Version     string
	Imports     []string
	Definitions []string
	Statements  []string
}

// TODO: Automatically remove unused code
const TranspileBase = `// Generated automatically by YAY {{.Version}}
package main

import (
	{{- range .Imports}}
	"{{. -}}"
	{{- end}}
)

var (
	stack []interface{}
	a interface{}
	b interface{}
	aInt int64
	bInt int64
	aFloat float64
	bFloat float64
	aString string
	bString string
)

func push(value interface{}) {
	stack = append(stack, value)
}

func pop() interface{} {
	value := stack[len(stack)-1]
	stack = stack[:len(stack)-1]
	return value
}

func peek() interface{} {
	return stack[len(stack)-1]
}

func btoi(value bool) int64 {
    if value {
        return 1
    }
    return 0
}

{{range .Definitions}}
{{. -}}
{{- end}}

func main() {
	{{- range .Statements}}
	{{. -}}
	{{- end -}}
}
`

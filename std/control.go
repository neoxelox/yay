package std

import "github.com/neoxelox/yay/mod"

var reference []mod.Token

type If struct{}

const LiteralIf = "if"

func (self *If) Parse(literal string, file string, row int, col int) (mod.Token, error) {
	token := mod.Token{
		Type:    mod.TypeIdentifier,
		Literal: literal,
		File:    file,
		Row:     row,
		Col:     col,
		Meta: map[string]string{
			"name":    "if",
			"package": "std",
		},
	}

	reference = append(reference, token)

	return token, nil
}

func (self *If) Transpile(token mod.Token) ([]string, string, string, error) {
	// TODO: Typeless If
	return nil, "", `
	aInt = pop(&stack).(int)
	if pop(&stack) {
		goto if_%d
	}
	`, nil
}

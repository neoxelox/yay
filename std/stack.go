package std

import "github.com/neoxelox/yay/mod"

type Drop struct{}

const LiteralDrop = "drop"
const IdentifierDrop = "drop"

func (self *Drop) Parse(literal string, file string, row int, col int) (mod.Token, error) {
	return mod.Token{
		Type:    mod.TypeIdentifier,
		Literal: literal,
		File:    file,
		Row:     row,
		Col:     col,
		Meta: map[string]string{
			"name":    IdentifierDrop,
			"package": "std",
		},
	}, nil
}

func (self *Drop) Transpile(token mod.Token) ([]string, string, string, error) {
	return nil, "", `
	pop(&stack)
	`, nil
}

// TODO: dup identifier

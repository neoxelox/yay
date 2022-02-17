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
	pop()
	`, nil
}

type Dup struct{}

const LiteralDup = "dup"
const IdentifierDup = "dup"

func (self *Dup) Parse(literal string, file string, row int, col int) (mod.Token, error) {
	return mod.Token{
		Type:    mod.TypeIdentifier,
		Literal: literal,
		File:    file,
		Row:     row,
		Col:     col,
		Meta: map[string]string{
			"name":    IdentifierDup,
			"package": "std",
		},
	}, nil
}

func (self *Dup) Transpile(token mod.Token) ([]string, string, string, error) {
	return nil, "", `
	a = pop()
	push(a)
	push(a)
	`, nil
}

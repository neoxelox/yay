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

type Swap struct{}

const LiteralSwap = "swap"
const IdentifierSwap = "swap"

func (self *Swap) Parse(literal string, file string, row int, col int) (mod.Token, error) {
	return mod.Token{
		Type:    mod.TypeIdentifier,
		Literal: literal,
		File:    file,
		Row:     row,
		Col:     col,
		Meta: map[string]string{
			"name":    IdentifierSwap,
			"package": "std",
		},
	}, nil
}

func (self *Swap) Transpile(token mod.Token) ([]string, string, string, error) {
	return nil, "", `
	b = pop()
	a = pop()
	push(b)
	push(a)
	`, nil
}

type Len struct{}

const LiteralLen = "len"
const IdentifierLen = "len"

func (self *Len) Parse(literal string, file string, row int, col int) (mod.Token, error) {
	return mod.Token{
		Type:    mod.TypeIdentifier,
		Literal: literal,
		File:    file,
		Row:     row,
		Col:     col,
		Meta: map[string]string{
			"name":    IdentifierLen,
			"package": "std",
		},
	}, nil
}

func (self *Len) Transpile(token mod.Token) ([]string, string, string, error) {
	return nil, "", `
	push(int64(len(stack)))
	`, nil
}

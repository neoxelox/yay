package std

import "github.com/neoxelox/yay/mod"

type Add struct{}

const LiteralAdd = "+"

func (self *Add) Parse(literal string, file string, row int, col int) (mod.Token, error) {
	return mod.Token{
		Type:    mod.TypeIdentifier,
		Literal: literal,
		File:    file,
		Row:     row,
		Col:     col,
		Meta: map[string]string{
			"name":    "add",
			"package": "std",
		},
	}, nil
}

func (self *Add) Transpile(token mod.Token) ([]string, string, string, error) {
	// TODO: Typeless Add
	return nil, "", `
	aInt = pop(&stack).(int64)
	bInt = pop(&stack).(int64)
	push(&stack, aInt + bInt)
	`, nil
}

type Sub struct{}

const LiteralSub = "-"

func (self *Sub) Parse(literal string, file string, row int, col int) (mod.Token, error) {
	return mod.Token{
		Type:    mod.TypeIdentifier,
		Literal: literal,
		File:    file,
		Row:     row,
		Col:     col,
		Meta: map[string]string{
			"name":    "sub",
			"package": "std",
		},
	}, nil
}

func (self *Sub) Transpile(token mod.Token) ([]string, string, string, error) {
	// TODO: Typeless Sub
	return nil, "", `
	aInt = pop(&stack).(int64)
	bInt = pop(&stack).(int64)
	push(&stack, bInt - aInt)
	`, nil
}

type Mul struct{}

const LiteralMul = "*"

func (self *Mul) Parse(literal string, file string, row int, col int) (mod.Token, error) {
	return mod.Token{
		Type:    mod.TypeIdentifier,
		Literal: literal,
		File:    file,
		Row:     row,
		Col:     col,
		Meta: map[string]string{
			"name":    "mul",
			"package": "std",
		},
	}, nil
}

func (self *Mul) Transpile(token mod.Token) ([]string, string, string, error) {
	// TODO: Typeless Mul
	return nil, "", `
	aInt = pop(&stack).(int64)
	bInt = pop(&stack).(int64)
	push(&stack, aInt * bInt)
	`, nil
}

type Div struct{}

const LiteralDiv = "/"

func (self *Div) Parse(literal string, file string, row int, col int) (mod.Token, error) {
	return mod.Token{
		Type:    mod.TypeIdentifier,
		Literal: literal,
		File:    file,
		Row:     row,
		Col:     col,
		Meta: map[string]string{
			"name":    "div",
			"package": "std",
		},
	}, nil
}

func (self *Div) Transpile(token mod.Token) ([]string, string, string, error) {
	// TODO: Typeless Div
	return nil, "", `
	aInt = pop(&stack).(int64)
	bInt = pop(&stack).(int64)
	push(&stack, bInt / aInt)
	`, nil
}

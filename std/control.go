package std

import (
	"fmt"

	"github.com/neoxelox/yay/mod"
)

var reference []mod.Token

type End struct{}

const LiteralEnd = "end"
const IdentifierEnd = "end"

func (self *End) Parse(literal string, file string, row int, col int) (mod.Token, error) {
	if len(reference) < 1 {
		return mod.Token{}, fmt.Errorf("'end' does not close a control block at %s:%d:%d", file, row+1, col+1)
	}

	block := reference[len(reference)-1]
	reference = reference[:len(reference)-1]

	token := mod.Token{
		Type:    mod.TypeIdentifier,
		Literal: literal,
		File:    file,
		Row:     row,
		Col:     col,
		Meta: map[string]string{
			"name":    IdentifierEnd,
			"package": "std",
			"block":   fmt.Sprintf("addr%d%d", block.Row, block.Col),
		},
	}

	return token, nil
}

func (self *End) EndParse() error {
	if len(reference) > 0 {
		block := reference[len(reference)-1]
		return fmt.Errorf("'%s' block is not closed by an 'end' at %s:%d:%d", block.Literal, block.File, block.Row+1, block.Col+1)
	}

	return nil
}

func (self *End) Transpile(token mod.Token) ([]string, string, string, error) {
	block, ok := token.Meta.(map[string]string)["block"]

	if !ok {
		return nil, "", "", fmt.Errorf("'end' does not close a control block at %s:%d:%d", token.File, token.Row+1, token.Col+1)
	}

	return nil, "", fmt.Sprintf(`
%s:
	`, block), nil
}

type If struct{}

const LiteralIf = "if"
const IdentifierIf = "if"

func (self *If) Parse(literal string, file string, row int, col int) (mod.Token, error) {
	token := mod.Token{
		Type:    mod.TypeIdentifier,
		Literal: literal,
		File:    file,
		Row:     row,
		Col:     col,
		Meta: map[string]string{
			"name":    IdentifierIf,
			"package": "std",
		},
	}

	reference = append(reference, token)

	return token, nil
}

func (self *If) Transpile(token mod.Token) ([]string, string, string, error) {
	return nil, "", fmt.Sprintf(`
	aInt = pop().(int64)
	if aInt == 0 {
		goto addr%d%d
	}
	`, token.Row, token.Col), nil
}

type Else struct{}

const LiteralElse = "else"
const IdentifierElse = "else"

func (self *Else) Parse(literal string, file string, row int, col int) (mod.Token, error) {
	if len(reference) < 1 {
		return mod.Token{}, fmt.Errorf("'else' does not extend an 'if' block at %s:%d:%d", file, row+1, col+1)
	}

	block := reference[len(reference)-1]
	reference = reference[:len(reference)-1]

	// TODO: Allow elif or {} if ... else {} if ... else ... end

	if block.Meta.(map[string]string)["name"] != IdentifierIf {
		return mod.Token{}, fmt.Errorf("'else' does not extend an 'if' block at %s:%d:%d", file, row+1, col+1)
	}

	token := mod.Token{
		Type:    mod.TypeIdentifier,
		Literal: literal,
		File:    file,
		Row:     row,
		Col:     col,
		Meta: map[string]string{
			"name":    IdentifierElse,
			"package": "std",
			"block":   fmt.Sprintf("addr%d%d", block.Row, block.Col),
		},
	}

	reference = append(reference, token)

	return token, nil
}

func (self *Else) Transpile(token mod.Token) ([]string, string, string, error) {
	block, ok := token.Meta.(map[string]string)["block"]

	if !ok {
		return nil, "", "", fmt.Errorf("'else' does not extend an 'if' block at %s:%d:%d", token.File, token.Row+1, token.Col+1)
	}

	return nil, "", fmt.Sprintf(`
	goto addr%d%d

%s:
	`, token.Row, token.Col, block), nil
}

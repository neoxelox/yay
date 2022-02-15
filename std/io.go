package std

import "github.com/neoxelox/yay/mod"

type Print struct{}

const LiteralPrint = "print"
const IdentifierPrint = "print"

func (self *Print) Parse(literal string, file string, row int, col int) (mod.Token, error) {
	return mod.Token{
		Type:    mod.TypeIdentifier,
		Literal: literal,
		File:    file,
		Row:     row,
		Col:     col,
		Meta: map[string]string{
			"name":    IdentifierPrint,
			"package": "std",
		},
	}, nil
}

func (self *Print) Transpile(token mod.Token) ([]string, string, string, error) {
	return []string{"fmt"}, "", `
	fmt.Print(peek(&stack))
	`, nil
}

type Println struct{}

const LiteralPrintln = "println"
const IdentifierPrintln = "println"

func (self *Println) Parse(literal string, file string, row int, col int) (mod.Token, error) {
	return mod.Token{
		Type:    mod.TypeIdentifier,
		Literal: literal,
		File:    file,
		Row:     row,
		Col:     col,
		Meta: map[string]string{
			"name":    IdentifierPrintln,
			"package": "std",
		},
	}, nil
}

func (self *Println) Transpile(token mod.Token) ([]string, string, string, error) {
	return []string{"fmt"}, "", `
	fmt.Println(peek(&stack))
	`, nil
}

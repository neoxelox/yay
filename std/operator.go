package std

import "github.com/neoxelox/yay/mod"

type Add struct{}

const LiteralAdd = "+"
const IdentifierAdd = "add"

func (self *Add) Parse(literal string, file string, row int, col int) (mod.Token, error) {
	return mod.Token{
		Type:    mod.TypeIdentifier,
		Literal: literal,
		File:    file,
		Row:     row,
		Col:     col,
		Meta: map[string]string{
			"name":    IdentifierAdd,
			"package": "std",
		},
	}, nil
}

func (self *Add) Transpile(token mod.Token) ([]string, string, string, error) {
	// TODO: Typeless Add
	return nil, "", `
	bInt = pop().(int64)
	aInt = pop().(int64)
	push(aInt + bInt)
	`, nil
}

type Sub struct{}

const LiteralSub = "-"
const IdentifierSub = "sub"

func (self *Sub) Parse(literal string, file string, row int, col int) (mod.Token, error) {
	return mod.Token{
		Type:    mod.TypeIdentifier,
		Literal: literal,
		File:    file,
		Row:     row,
		Col:     col,
		Meta: map[string]string{
			"name":    IdentifierSub,
			"package": "std",
		},
	}, nil
}

func (self *Sub) Transpile(token mod.Token) ([]string, string, string, error) {
	// TODO: Typeless Sub
	return nil, "", `
	bInt = pop().(int64)
	aInt = pop().(int64)
	push(aInt - bInt)
	`, nil
}

type Mul struct{}

const LiteralMul = "*"
const IdentifierMul = "mul"

func (self *Mul) Parse(literal string, file string, row int, col int) (mod.Token, error) {
	return mod.Token{
		Type:    mod.TypeIdentifier,
		Literal: literal,
		File:    file,
		Row:     row,
		Col:     col,
		Meta: map[string]string{
			"name":    IdentifierMul,
			"package": "std",
		},
	}, nil
}

func (self *Mul) Transpile(token mod.Token) ([]string, string, string, error) {
	// TODO: Typeless Mul
	return nil, "", `
	bInt = pop().(int64)
	aInt = pop().(int64)
	push(aInt * bInt)
	`, nil
}

type Div struct{}

const LiteralDiv = "/"
const IdentifierDiv = "div"

func (self *Div) Parse(literal string, file string, row int, col int) (mod.Token, error) {
	return mod.Token{
		Type:    mod.TypeIdentifier,
		Literal: literal,
		File:    file,
		Row:     row,
		Col:     col,
		Meta: map[string]string{
			"name":    IdentifierDiv,
			"package": "std",
		},
	}, nil
}

func (self *Div) Transpile(token mod.Token) ([]string, string, string, error) {
	// TODO: Typeless Div
	return nil, "", `
	bInt = pop().(int64)
	aInt = pop().(int64)
	push(aInt / bInt)
	`, nil
}

type Mod struct{}

const LiteralMod = "%"
const IdentifierMod = "mod"

func (self *Mod) Parse(literal string, file string, row int, col int) (mod.Token, error) {
	return mod.Token{
		Type:    mod.TypeIdentifier,
		Literal: literal,
		File:    file,
		Row:     row,
		Col:     col,
		Meta: map[string]string{
			"name":    IdentifierMod,
			"package": "std",
		},
	}, nil
}

func (self *Mod) Transpile(token mod.Token) ([]string, string, string, error) {
	// TODO: Typeless Mod
	return nil, "", `
	bInt = pop().(int64)
	aInt = pop().(int64)
	push(aInt % bInt)
	`, nil
}

// Maybe do not include this as an operator
type Exp struct{}

const LiteralExp = "**"
const IdentifierExp = "exp"

func (self *Exp) Parse(literal string, file string, row int, col int) (mod.Token, error) {
	return mod.Token{
		Type:    mod.TypeIdentifier,
		Literal: literal,
		File:    file,
		Row:     row,
		Col:     col,
		Meta: map[string]string{
			"name":    IdentifierExp,
			"package": "std",
		},
	}, nil
}

func (self *Exp) Transpile(token mod.Token) ([]string, string, string, error) {
	// TODO: Typeless Exp
	return []string{"math"}, "", `
	bInt = pop().(int64)
	aInt = pop().(int64)
	push(int64(math.Pow(float64(aInt), float64(bInt))))
	`, nil
}

type Eq struct{}

const LiteralEq = "="
const IdentifierEq = "eq"

func (self *Eq) Parse(literal string, file string, row int, col int) (mod.Token, error) {
	return mod.Token{
		Type:    mod.TypeIdentifier,
		Literal: literal,
		File:    file,
		Row:     row,
		Col:     col,
		Meta: map[string]string{
			"name":    IdentifierEq,
			"package": "std",
		},
	}, nil
}

func (self *Eq) Transpile(token mod.Token) ([]string, string, string, error) {
	// TODO: Typeless Eq
	return nil, "", `
	bInt = pop().(int64)
	aInt = pop().(int64)
	push(btoi(aInt == bInt))
	`, nil
}

type Le struct{}

const LiteralLe = "<"
const IdentifierLe = "le"

func (self *Le) Parse(literal string, file string, row int, col int) (mod.Token, error) {
	return mod.Token{
		Type:    mod.TypeIdentifier,
		Literal: literal,
		File:    file,
		Row:     row,
		Col:     col,
		Meta: map[string]string{
			"name":    IdentifierLe,
			"package": "std",
		},
	}, nil
}

func (self *Le) Transpile(token mod.Token) ([]string, string, string, error) {
	// TODO: Typeless Le
	return nil, "", `
	bInt = pop().(int64)
	aInt = pop().(int64)
	push(btoi(aInt < bInt))
	`, nil
}

type Ge struct{}

const LiteralGe = ">"
const IdentifierGe = "ge"

func (self *Ge) Parse(literal string, file string, row int, col int) (mod.Token, error) {
	return mod.Token{
		Type:    mod.TypeIdentifier,
		Literal: literal,
		File:    file,
		Row:     row,
		Col:     col,
		Meta: map[string]string{
			"name":    IdentifierGe,
			"package": "std",
		},
	}, nil
}

func (self *Ge) Transpile(token mod.Token) ([]string, string, string, error) {
	// TODO: Typeless Ge
	return nil, "", `
	bInt = pop().(int64)
	aInt = pop().(int64)
	push(btoi(aInt > bInt))
	`, nil
}

type Leq struct{}

const LiteralLeq = "<="
const IdentifierLeq = "leq"

func (self *Leq) Parse(literal string, file string, row int, col int) (mod.Token, error) {
	return mod.Token{
		Type:    mod.TypeIdentifier,
		Literal: literal,
		File:    file,
		Row:     row,
		Col:     col,
		Meta: map[string]string{
			"name":    IdentifierLeq,
			"package": "std",
		},
	}, nil
}

func (self *Leq) Transpile(token mod.Token) ([]string, string, string, error) {
	// TODO: Typeless Leq
	return nil, "", `
	bInt = pop().(int64)
	aInt = pop().(int64)
	push(btoi(aInt <= bInt))
	`, nil
}

type Geq struct{}

const LiteralGeq = ">="
const IdentifierGeq = "geq"

func (self *Geq) Parse(literal string, file string, row int, col int) (mod.Token, error) {
	return mod.Token{
		Type:    mod.TypeIdentifier,
		Literal: literal,
		File:    file,
		Row:     row,
		Col:     col,
		Meta: map[string]string{
			"name":    IdentifierGeq,
			"package": "std",
		},
	}, nil
}

func (self *Geq) Transpile(token mod.Token) ([]string, string, string, error) {
	// TODO: Typeless Geq
	return nil, "", `
	bInt = pop().(int64)
	aInt = pop().(int64)
	push(btoi(aInt >= bInt))
	`, nil
}

type Not struct{}

const LiteralNot = "!"
const IdentifierNot = "not"

func (self *Not) Parse(literal string, file string, row int, col int) (mod.Token, error) {
	return mod.Token{
		Type:    mod.TypeIdentifier,
		Literal: literal,
		File:    file,
		Row:     row,
		Col:     col,
		Meta: map[string]string{
			"name":    IdentifierNot,
			"package": "std",
		},
	}, nil
}

func (self *Not) Transpile(token mod.Token) ([]string, string, string, error) {
	// TODO: Typeless Not
	return nil, "", `
	aInt = pop().(int64)
	push(btoi(aInt == 0))
	`, nil
}

type And struct{}

const LiteralAnd = "&"
const IdentifierAnd = "and"

func (self *And) Parse(literal string, file string, row int, col int) (mod.Token, error) {
	return mod.Token{
		Type:    mod.TypeIdentifier,
		Literal: literal,
		File:    file,
		Row:     row,
		Col:     col,
		Meta: map[string]string{
			"name":    IdentifierAnd,
			"package": "std",
		},
	}, nil
}

func (self *And) Transpile(token mod.Token) ([]string, string, string, error) {
	// TODO: Typeless And
	return nil, "", `
	bInt = pop().(int64)
	aInt = pop().(int64)
	push(aInt & bInt)
	`, nil
}

type Or struct{}

const LiteralOr = "|"
const IdentifierOr = "or"

func (self *Or) Parse(literal string, file string, row int, col int) (mod.Token, error) {
	return mod.Token{
		Type:    mod.TypeIdentifier,
		Literal: literal,
		File:    file,
		Row:     row,
		Col:     col,
		Meta: map[string]string{
			"name":    IdentifierOr,
			"package": "std",
		},
	}, nil
}

func (self *Or) Transpile(token mod.Token) ([]string, string, string, error) {
	// TODO: Typeless Or
	return nil, "", `
	bInt = pop().(int64)
	aInt = pop().(int64)
	push(aInt | bInt)
	`, nil
}

type Xor struct{}

const LiteralXor = "^"
const IdentifierXor = "xor"

func (self *Xor) Parse(literal string, file string, row int, col int) (mod.Token, error) {
	return mod.Token{
		Type:    mod.TypeIdentifier,
		Literal: literal,
		File:    file,
		Row:     row,
		Col:     col,
		Meta: map[string]string{
			"name":    IdentifierXor,
			"package": "std",
		},
	}, nil
}

func (self *Xor) Transpile(token mod.Token) ([]string, string, string, error) {
	// TODO: Typeless Xor
	return nil, "", `
	bInt = pop().(int64)
	aInt = pop().(int64)
	push(aInt ^ bInt)
	`, nil
}

type Neg struct{}

const LiteralNeg = "~"
const IdentifierNeg = "neg"

func (self *Neg) Parse(literal string, file string, row int, col int) (mod.Token, error) {
	return mod.Token{
		Type:    mod.TypeIdentifier,
		Literal: literal,
		File:    file,
		Row:     row,
		Col:     col,
		Meta: map[string]string{
			"name":    IdentifierNeg,
			"package": "std",
		},
	}, nil
}

func (self *Neg) Transpile(token mod.Token) ([]string, string, string, error) {
	// TODO: Typeless Neg
	return nil, "", `
	aInt = pop().(int64)
	push(^aInt)
	`, nil
}

type Lsh struct{}

const LiteralLsh = "<<"
const IdentifierLsh = "lsh"

func (self *Lsh) Parse(literal string, file string, row int, col int) (mod.Token, error) {
	return mod.Token{
		Type:    mod.TypeIdentifier,
		Literal: literal,
		File:    file,
		Row:     row,
		Col:     col,
		Meta: map[string]string{
			"name":    IdentifierLsh,
			"package": "std",
		},
	}, nil
}

func (self *Lsh) Transpile(token mod.Token) ([]string, string, string, error) {
	// TODO: Typeless Lsh
	return nil, "", `
	bInt = pop().(int64)
	aInt = pop().(int64)
	push(aInt << bInt)
	`, nil
}

type Rsh struct{}

const LiteralRsh = ">>"
const IdentifierRsh = "rsh"

func (self *Rsh) Parse(literal string, file string, row int, col int) (mod.Token, error) {
	return mod.Token{
		Type:    mod.TypeIdentifier,
		Literal: literal,
		File:    file,
		Row:     row,
		Col:     col,
		Meta: map[string]string{
			"name":    IdentifierRsh,
			"package": "std",
		},
	}, nil
}

func (self *Rsh) Transpile(token mod.Token) ([]string, string, string, error) {
	// TODO: Typeless Rsh
	return nil, "", `
	bInt = pop().(int64)
	aInt = pop().(int64)
	push(aInt >> bInt)
	`, nil
}

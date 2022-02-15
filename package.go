package main

import (
	"github.com/neoxelox/yay/mod"
	"github.com/neoxelox/yay/std"
)

// TODO: Enhance all of this odd system to a more clean one!

var Identifiers = map[string]mod.Identifier{
	std.LiteralDrop: &std.Drop{},

	std.LiteralAdd: &std.Add{},
	std.LiteralSub: &std.Sub{},
	std.LiteralMul: &std.Mul{},
	std.LiteralDiv: &std.Div{},
	std.LiteralMod: &std.Mod{},
	std.LiteralExp: &std.Exp{},
	std.LiteralEq:  &std.Eq{},
	std.LiteralLe:  &std.Le{},
	std.LiteralGe:  &std.Ge{},
	std.LiteralLeq: &std.Leq{},
	std.LiteralGeq: &std.Geq{},
	std.LiteralNot: &std.Not{},
	std.LiteralAnd: &std.And{},
	std.LiteralOr:  &std.Or{},
	std.LiteralXor: &std.Xor{},
	std.LiteralNeg: &std.Neg{},

	std.LiteralIf:   &std.If{},
	std.LiteralElse: &std.Else{},
	std.LiteralEnd:  &std.End{},

	std.LiteralPrint:   &std.Print{},
	std.LiteralPrintln: &std.Println{},
}

// TODO: Pass and allow to modify program
func beginParse() error {
	return nil
}

// TODO: Pass and allow to modify program
func endParse() error {
	var err error

	e := &std.End{}
	err = e.EndParse()
	if err != nil {
		return err
	}

	return nil
}

// TODO: Pass and allow to modify program
func beginTranspile() error {
	return nil
}

// TODO: Pass and allow to modify program
func endTranspile() error {
	return nil
}

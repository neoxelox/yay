package main

import (
	"github.com/neoxelox/yay/mod"
	"github.com/neoxelox/yay/std"
)

var Identifiers = map[string]mod.Identifier{
	std.LiteralAdd:     &std.Add{},
	std.LiteralSub:     &std.Sub{},
	std.LiteralMul:     &std.Mul{},
	std.LiteralDiv:     &std.Div{},
	std.LiteralPrint:   &std.Print{},
	std.LiteralPrintln: &std.Println{},
}

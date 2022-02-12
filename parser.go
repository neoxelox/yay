package main

import (
	"fmt"
	"strconv"
	"unicode"

	"github.com/neoxelox/yay/mod"
)

func isNumber(literal string) bool {
	if _, err := strconv.ParseInt(literal, 10, 64); err == nil {
		return true
	}

	if _, err := strconv.ParseFloat(literal, 64); err == nil {
		return true
	}

	return false
}

func isIdentifier(literal string) bool {
	if _, ok := Identifiers[literal]; ok {
		return true
	}

	return false
}

func isComment(literal string) bool {
	return literal == mod.LiteralComment
}

func searchRune(r rune, runes []rune, start int) int {
	for i := start; i < len(runes); i++ {
		if runes[i] == r {
			return i
		}
	}

	return -1
}

func parse(program string, filepath string) ([]mod.Token, error) {
	var tokens []mod.Token
	var literal []rune
	var row int
	var col int
	runes := []rune(program)
	runes = append(runes, '\n') // Always have at least one end of line

	for cur := 0; cur < len(runes); cur++ {
		if unicode.IsSpace(runes[cur]) || isComment(string(runes[cur])) {
			if cur != 0 && !unicode.IsSpace(runes[cur-1]) {
				aCol := col - len(literal)
				aLiteral := string(literal)

				switch {
				case isNumber(aLiteral):
					tokens = append(tokens, mod.Token{
						Type:    mod.TypeNumber,
						Literal: aLiteral,
						File:    filepath,
						Row:     row,
						Col:     aCol,
					})
				case isIdentifier(aLiteral):
					token, err := Identifiers[aLiteral].Parse(aLiteral, filepath, row, aCol)
					if err != nil {
						return nil, fmt.Errorf("cannot parse identifier '%s' at %s:%d:%d", aLiteral, filepath, row+1, aCol+1)
					}
					tokens = append(tokens, token)
				default:
					return nil, fmt.Errorf("unknown token '%s' at %s:%d:%d", aLiteral, filepath, row+1, aCol+1)
				}

				literal = []rune{}
			}
		} else {
			literal = append(literal, runes[cur])
		}

		col++

		if runes[cur] == '\n' {
			row++
			col = 0
		} else if isComment(string(runes[cur])) {
			row++
			col = 0
			cur = searchRune('\n', runes, cur+1)
		}
	}

	return tokens, nil
}

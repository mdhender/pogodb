/*
 * pogodb - a sql database engine
 *
 * Copyright (c) 2021 Michael D Henderson
 *
 * Permission is hereby granted, free of charge, to any person obtaining a copy
 * of this software and associated documentation files (the "Software"), to deal
 * in the Software without restriction, including without limitation the rights
 * to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
 * copies of the Software, and to permit persons to whom the Software is
 * furnished to do so, subject to the following conditions:
 *
 * The above copyright notice and this permission notice shall be included in all
 * copies or substantial portions of the Software.
 *
 * THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
 * IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
 * FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
 * AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
 * LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
 * OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN THE
 * SOFTWARE.
 */

package parser

import (
	"unicode"
	"unicode/utf8"
)

type PARSER struct{}

type Statement struct {
	Select *SELECT
}

type COLUMNALIAS struct {
	Name string
}
type EXPRESSION struct {
	Identifier string
}
type SELECT struct {
	Distinct   bool
	SelectList *SELECTLIST
	From       *TABLEREFERENCELIST
	Where      *SEARCHCONDITION
}
type SEARCHCONDITION struct{}
type SELECTCOLUMN struct {
	Expression *EXPRESSION
	Alias      *COLUMNALIAS
}
type SELECTLIST struct {
	Columns []*SELECTCOLUMN
}
type TABLEREFERENCELIST struct{}

//func skipSpaces(b []byte) []byte {
//	for !scanner.Eof(b) {
//		if lexeme, rest := scanner.Spaces(b); lexeme != nil {
//			b = rest
//		}
//		if lexeme, rest := scanner.Comment(b); lexeme != nil {
//			b = rest
//			continue
//		}
//		break
//	}
//	return b
//}

func bdup(src []byte) []byte {
	dst := make([]byte, len(src))
	copy(dst, src)
	return dst
}

func keyword(b, u, l []byte) (lexeme, rest []byte) {
	if !(len(u) == len(l) && len(u) <= len(b)) {
		return nil, b
	}
	i := 0
	for ; i < len(u) && (b[i] == u[i] || b[i] == l[i]); i++ {
	}
	if i == len(u) {
		r, _ := utf8.DecodeRune(b[i:])
		switch r {
		case '(', ')', '.', ',', '-', '+', '/', '*', ';':
			return b[:i-1], b[i-1:]
		case utf8.RuneError:
			return b[:i-1], b[i-1:]
		default:
			if unicode.IsSpace(r) {
				return b[:i-1], b[i-1:]
			}
		}
	}
	return nil, b
}

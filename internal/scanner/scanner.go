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

package scanner

import (
	"bytes"
	"unicode"
	"unicode/utf8"
)

func All(b []byte) (lexeme, rest []byte) {
	return Keyword(b, []byte("aAlLlL"))
}

func As(b []byte) (lexeme, rest []byte) {
	return Keyword(b, []byte("aAsS"))
}

func Comment(b []byte) (lexeme, rest []byte) {
	if bytes.HasPrefix(b, []byte{'-', '-'}) {
		r, w := utf8.DecodeRune(b)
		for !(r == '\n' || r == utf8.RuneError) {
			lexeme, b = append(lexeme, b[:w]...), b[w:]
		}
		return lexeme, b
	} else if bytes.HasPrefix(b, []byte{'/', '*'}) {
		w := bytes.Index(b[2:], []byte{'*', '/'})
		if w == -1 {
			w = len(b)
		}
		return b[:w+2], b[w+2:]
	}
	return nil, b
}

func Distinct(b []byte) (lexeme, rest []byte) {
	return Keyword(b, []byte("dDiIsStTiInNcCtT"))
}

func Eof(b []byte) bool {
	return len(b) == 0
}

func From(b []byte) (lexeme, rest []byte) {
	return Keyword(b, []byte("fFrRoOmM"))
}

func Glyph(b []byte, ch byte) (lexeme, rest []byte) {
	if len(b) == 0 || b[0] != ch {
		return nil, b
	}
	return b[:1], b[1:]
}

func Identifier(b []byte) (lexeme, rest []byte) {
	saved := b
	r, w := utf8.DecodeRune(b)
	if !(unicode.IsLetter(r) || r == '_') {
		return nil, saved
	}
	lexeme, b = append(lexeme, b[:w]...), b[w:]
	for r, w = utf8.DecodeRune(b); unicode.IsLetter(r) || unicode.IsDigit(r) || r == '_';  {
		lexeme, b = append(lexeme, b[:w]...), b[w:]
	}
	return lexeme, b
}

func Keyword(b, kw []byte) (lexeme, rest []byte) {
	saved := b
	for len(b) != 0 && len(kw) >= 2 {
		if !(b[0] == kw[0] || b[0] == kw[1]) {
			return nil, saved
		}
		lexeme, b = append(lexeme, b[0]), b[1:]
		if kw = kw[2:]; len(kw) == 0 {
			if r, _ := utf8.DecodeRune(b); unicode.IsSpace(r) || r == utf8.RuneError {
				return lexeme, b
			} else {
				switch r {
				case '(', ')', '.', ',', '-', '+', '/', '*', ';':
					return lexeme, b
				}
			}
		}
	}
	return nil, b
}

func Select(b []byte) (lexeme, rest []byte) {
	return Keyword(b, []byte("sSeElLeEcCtT"))
}

func SkipSpaces(b []byte) []byte {
	for !Eof(b) {
		if lexeme, rest := Spaces(b); lexeme != nil {
			b = rest
		}
		if lexeme, rest := Comment(b); lexeme != nil {
			b = rest
			continue
		}
		break
	}
	return b
}

func Spaces(b []byte) (lexeme, rest []byte) {
	r, w := utf8.DecodeRune(b)
	for unicode.IsSpace(r) {
		lexeme, b = append(lexeme, b[:w]...), b[w:]
	}
	return lexeme, b
}

func Where(b []byte) (lexeme, rest []byte) {
	return Keyword(b, []byte("wWhHeErReE"))
}

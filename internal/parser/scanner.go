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
	"bytes"
	"unicode"
	"unicode/utf8"
)

type scanner struct {
	buffer []byte
	lexeme []byte
}

func scannerFromString(s string) scanner {
	return scanner{buffer: []byte(s)}
}

func (s scanner) blockComment() scanner {
	s.lexeme = nil
	if !bytes.HasPrefix(s.buffer, []byte{'/', '*'}) {
		return s
	}
	length := 2
	if w := bytes.Index(s.buffer[length:], []byte{'*', '/'}); w == -1 {
		length = len(s.buffer)
	} else {
		length += w + 2 // the extra two captures the trailing `*/`
	}
	s.lexeme, s.buffer = s.buffer[:length], s.buffer[length:]
	return s
}

func (s scanner) comment() scanner {
	if s = s.lineComment(); s.lexeme == nil {
		s = s.blockComment()
	}
	return s

}

func (s scanner) eof() bool {
	return len(s.buffer) == 0
}

func (s scanner) glyph(ch byte) bool {
	return len(s.buffer) != 0 && s.buffer[0] == ch
}

func (s scanner) identifier() scanner {
	s.lexeme = nil
	r, w := utf8.DecodeRune(s.buffer)
	if !(unicode.IsLetter(r) || r == '_') {
		return s
	}
	length := w
	for r, w = utf8.DecodeRune(s.buffer[length:]); unicode.IsLetter(r) || unicode.IsDigit(r) || r == '_'; r, w = utf8.DecodeRune(s.buffer[length:]) {
		length += w
	}
	s.lexeme, s.buffer = s.buffer[:length], s.buffer[length:]
	return s
}

func (s scanner) keyword(kw []byte) scanner {
	s.lexeme = nil
	if !bytes.HasPrefix(s.buffer, kw) {
		return s
	} else if len(kw) < len(s.buffer) {
		// certain glyphs will terminate a keyword.
		// anything else means we didn't find a complete word.
		if bytes.IndexByte([]byte{'(', ')', '.', ',', '-', '+', '/', '*', ';'}, s.buffer[len(kw)]) == -1 {
			if r, _ := utf8.DecodeRune(s.buffer[len(kw):]); !(unicode.IsSpace(r) || r == utf8.RuneError) {
				return scanner{buffer: s.buffer}
			}
		}
	}
	s.lexeme, s.buffer = s.buffer[:len(kw)], s.buffer[len(kw):]
	return s
}

func (s scanner) lineComment() scanner {
	s.lexeme = nil
	if !bytes.HasPrefix(s.buffer, []byte{'-', '-'}) {
		return s
	}
	w := bytes.IndexByte(s.buffer, '\n')
	if w == -1 {
		w = len(s.buffer)
	}
	s.lexeme, s.buffer = s.buffer[:w], s.buffer[w:]
	return s
}

func (s scanner) spaces() scanner {
	s.lexeme = nil
	r, w := utf8.DecodeRune(s.buffer)
	if !unicode.IsSpace(r) {
		return s
	}
	length := w
	for r, w = utf8.DecodeRune(s.buffer[length:]); unicode.IsSpace(r); r, w = utf8.DecodeRune(s.buffer[length:]) {
		length += w
	}
	s.lexeme, s.buffer = s.buffer[:length], s.buffer[length:]
	return s
}

func (s scanner) _all() scanner {
	return s.keyword([]byte{'a', 'l', 'l'})
}

func (s scanner) _as() scanner {
	return s.keyword([]byte{'a', 's'})
}

func (s scanner) _distinct() scanner {
	return s.keyword([]byte{'d', 'i', 's', 't', 'i', 'n', 'c', 't'})
}

func (s scanner) _from() scanner {
	return s.keyword([]byte{'f', 'r', 'o', 'm'})
}

func (s scanner) _select() scanner {
	return s.keyword([]byte{'s', 'e', 'l', 'e', 'c', 't'})
}

func (s scanner) _where() scanner {
	return s.keyword([]byte{'w', 'h', 'e', 'r', 'e'})
}

func (s scanner) _whitespace() scanner {
	for !s.eof() {
		if s = s.spaces().comment(); s.lexeme == nil {
			break
		}
	}
	return s
}

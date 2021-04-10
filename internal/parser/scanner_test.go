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

import "testing"

// Specification: Scanner API
func TestScanner(t *testing.T) {
	// Given an empty scanner
	// When we test for end-of-input
	// Then the result should be true
	var s scanner
	if !s.eof() {
		t.Errorf("scanner: eof: expected true: got false")
	}

	// Given an empty string to scan
	// When we test for end-of-input
	// Then the result should be true
	s = scannerFromString("")
	if !s.eof() {
		t.Errorf("scanner: eof: expected true: got false")
	}

	// Given a string containing nothing but spaces
	// When we test for end-of-input
	// Then the result should be false
	input := " \t \r\n \n"
	s = scannerFromString(input)
	if s.eof() {
		t.Errorf("scanner: eof: expected false: got true")
	}

	// When we scan for spaces
	// Then the result should contain a lexeme
	// And the lexeme should be the entire input
	// And end-of-input should be true
	s = s.spaces()
	if s.lexeme == nil {
		t.Errorf("scanner: spaces: expected lexeme: got nil")
	} else if expected := input; expected != string(s.lexeme) {
		t.Errorf("scanner: spaces: expected %q: got %q", expected, string(s.lexeme))
	}
	if !s.eof() {
		t.Errorf("scanner: spaces: expected true: got false")
	}

	// Given the following input
	//   /* block\ncomment */\n
	// When we scan for block comment
	// Then the result should contain a lexeme
	// And the lexeme should be the entire input up to (but not including) the final new-line
	// And end-of-input should be false
	input = "/* block\ncomment */\n"
	s = scannerFromString(input).blockComment()
	if s.lexeme == nil {
		t.Errorf("scanner: block comment: expected lexeme: got nil")
	} else if expected := "/* block\ncomment */"; expected != string(s.lexeme) {
		t.Errorf("scanner: block comment: expected %q: got %q", expected, string(s.lexeme))
	}
	if s.eof() {
		t.Errorf("scanner: block comment: expected false: got true")
	}

	// Given the following input
	//   -- line comment\n
	// When we scan for line comment
	// Then the result should contain a lexeme
	// And the lexeme should be the entire input up to (but not including) the new-line
	// And end-of-input should be false
	input = "-- line comment\n"
	s = scannerFromString(input).lineComment()
	if s.lexeme == nil {
		t.Errorf("scanner: line comment: expected lexeme: got nil")
	} else if expected := "-- line comment"; expected != string(s.lexeme) {
		t.Errorf("scanner: line comment: expected %q: got %q", expected, string(s.lexeme))
	}
	if s.eof() {
		t.Errorf("scanner: line comment: expected false: got true")
	}

	// Given the following input
	//   /* block\ncomment */\n
	// When we scan for comment
	// Then the result should contain a lexeme
	// And the lexeme should be the entire input up to (but not including) the final new-line
	// And end-of-input should be false
	input = "/* block\ncomment */\n"
	s = scannerFromString(input).comment()
	if s.lexeme == nil {
		t.Errorf("scanner: comment: expected lexeme: got nil")
	} else if expected := "/* block\ncomment */"; expected != string(s.lexeme) {
		t.Errorf("scanner: comment: expected %q: got %q", expected, string(s.lexeme))
	}
	if s.eof() {
		t.Errorf("scanner: comment: expected false: got true")
	}

	// Given the following input
	//   one as _thr_3 fou(r
	// When we scan for an identifier
	// Then the result should contain a lexeme
	// And the lexeme should be the string `one`
	// When we scan for spaces
	// Then the result should contain a lexeme
	// When we scan for the keyword `as`
	// Then the result should contain a lexeme
	// And the lexeme should be the string `as`
	// When we skip whitespace and scan for an identifier
	// Then the result should contain a lexeme
	// And the lexeme should be the string `_thr_3`
	// When we skip whitespace and scan for an identifier
	// Then the result should contain a lexeme
	// And the lexeme should be the string `fou`
	// And the remainder of the buffer should be the string `(r`
	input = "one as _thr_3/* ignored */fou(r"
	s = scannerFromString(input).identifier()
	if s.lexeme == nil {
		t.Errorf("scanner: identifier: expected lexeme: got nil")
	} else if expected := "one"; expected != string(s.lexeme) {
		t.Errorf("scanner: identifier: expected %q: got %q", expected, string(s.lexeme))
	}
	if s.eof() {
		t.Errorf("scanner: identifier: expected false: got true")
	} else {
		s = s.spaces()
		if s.lexeme == nil {
			t.Errorf("scanner: identifier: expected lexeme: got nil")
		} else {
			s = s._as()
			if s.lexeme == nil {
				t.Errorf("scanner: identifier: expected lexeme: got nil")
			} else if expected := "as"; expected != string(s.lexeme) {
				t.Errorf("scanner: identifier: expected %q: got %q", expected, string(s.lexeme))
			} else {
				s = s._whitespace().identifier()
				if s.lexeme == nil {
					t.Errorf("scanner: identifier: expected lexeme: got nil")
				} else if expected := "_thr_3"; expected != string(s.lexeme) {
					t.Errorf("scanner: identifier: expected %q: got %q", expected, string(s.lexeme))
				} else {
					s = s._whitespace().identifier()
					if s.lexeme == nil {
						t.Errorf("scanner: identifier: expected lexeme: got nil")
					} else if expected := "fou"; expected != string(s.lexeme) {
						t.Errorf("scanner: identifier: expected %q: got %q", expected, string(s.lexeme))
					} else if expected := "(r"; expected != string(s.buffer) {
						t.Errorf("scanner: identifier: expected %q: got %q", expected, string(s.buffer))
					}
				}
			}
		}
	}

	// Given the following input
	//   select * from dummy;
	// When we skip whitespace and scan for the keyword `select`
	// Then the result should contain a lexeme
	// And the lexeme should be the string `select`
	s = scannerFromString("select * from dummy;")._whitespace()._select()
	if s.lexeme == nil {
		t.Errorf("scanner: select: expected lexeme: got nil")
	} else if expected := "select"; expected != string(s.lexeme) {
		t.Errorf("scanner: select: expected %q: got %q", expected, string(s.lexeme))
	}
}

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

//func acceptStatement(b []byte) (*Statement, []byte, error) {
//	b = skipSpaces(b)
//	if s, rest, err := acceptSelectStatement(b); err != nil {
//		return nil, rest, err
//	} else if s != nil {
//		return &Statement{Select: s}, rest, nil
//	}
//
//	return nil, b, nil
//}
//
//func acceptSelectStatement(b []byte) (*SELECT, []byte, error) {
//	saved, s := b, &SELECT{}
//
//	lexeme, rest := scanner.Select(b)
//	if lexeme == nil {
//		return nil, rest, nil
//	}
//	b = skipSpaces(rest)
//	if lexeme, rest = scanner.All(b); lexeme == nil {
//		if lexeme, rest = scanner.Distinct(b); lexeme != nil {
//			s.Distinct = true
//		}
//	}
//	b = skipSpaces(rest)
//	var err error
//	if s.SelectList, rest, err = expectSelectList(b); err != nil {
//		return nil, saved, err
//	}
//	b = skipSpaces(rest)
//	if lexeme, rest = scanner.From(b); lexeme == nil {
//		return nil, saved, fmt.Errorf("expected FROM")
//	}
//	b = skipSpaces(rest)
//	s.From, rest, err = expectTableReferenceList(b)
//	if err != nil {
//		return nil, saved, err
//	}
//	b = skipSpaces(rest)
//	if lexeme, rest = scanner.Where(b); lexeme == nil {
//		return s, rest, nil
//	}
//	b = skipSpaces(rest)
//	s.Where, rest, err = expectSearchCondition(b)
//	if err != nil {
//		return nil, saved, err
//	}
//	return s, rest, nil
//}
//
//
//func acceptSelectSublist(b []byte) (*EXPRESSION, []byte, error) {
//	b = skipSpaces(b)
//	saved := b
//	lexeme, rest := scanner.Glyph(b, ',')
//	if lexeme == nil {
//		return nil, saved, nil
//	}
//	b = rest
//	e, rest, err:= expectExpression(b)
//	if err != nil {
//		return nil, b, err
//	}
//	return e, rest, nil
//}
//
////func acceptSelectList(b []byte) (*SELECTLIST, []byte, error) {
////	b = skipSpaces(b)
////	e, rest, err:= expectExpression(b)
////	if err != nil {
////		return nil, b, err
////	}
////	s := SELECTLIST{}
////	return s, b, nil
////}
//
//
//
//func acceptSelectColumn(b []byte) (*SELECTCOLUMN, []byte, error) {
//	b = skipSpaces(b)
//	e, rest, err:= expectExpression(b)
//	if err != nil {
//		return nil, b, err
//	}
//	b = rest
//	c := SELECTCOLUMN{Expression: e}
//	lexeme, rest := scanner.As(b)
//	if lexeme == nil {
//		return &c, b, nil
//	}
//	b = rest
//	c.Alias, rest, err = expectColumnAlias(b)
//	if err != nil {
//		return nil, b, err
//	}
//	return &c, rest, nil
//}
//

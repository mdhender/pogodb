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
	"fmt"
	"github.com/mdhender/pogodb/internal/scanner"
	"strings"
)

func expectSelectList(b []byte) (*SELECTLIST, []byte, error) {
	saved, selectList := b, &SELECTLIST{}
	if lexeme, rest := scanner.Glyph(b, '*'); lexeme != nil {
		selectColumn := &SELECTCOLUMN{
			Expression: &EXPRESSION{Identifier: "*"},
			Alias:      nil,
		}
		selectList.Columns = append(selectList.Columns, selectColumn)
		return selectList, rest, nil
	}
	// expr ( COMMA expr )*
	for len(b) != 0 {
		expr, rest, err := expectExpression(b)
		if err != nil {
			return nil, saved, err
		}
		b = skipSpaces(rest)
		selectColumn := &SELECTCOLUMN{
			Expression: expr,
			Alias:      nil,
		}
		selectList.Columns = append(selectList.Columns, selectColumn)
		if lexeme, rest := scanner.Glyph(b, ','); lexeme != nil {
			b = skipSpaces(rest)
			continue
		}
		break
	}
	return selectList, b, nil
}

func expectExpression(b []byte) (*EXPRESSION, []byte, error) {
	//e, rest, err:= expectExpression(b)
	//if err != nil {
	//	return nil, b, err
	//}
	//b = rest
	//c := SELECTCOLUMN{Expression: e}
	//lexeme, rest := scanner.As(b)
	//if lexeme == nil {
	//	return &c, b, nil
	//}
	//b = rest
	//c.Alias, rest, err = expectColumnAlias(b)
	//if err != nil {
	//	return nil, b, err
	//}
	//return &c, rest, nil
	return nil, b, fmt.Errorf("expected EXPRESSION")
}

func expectTableReferenceList(b []byte) (*TABLEREFERENCELIST, []byte, error) {
	return nil, b, fmt.Errorf("expected TABLE_REFERENCE_LIST")
}

func expectSearchCondition(b []byte) (*SEARCHCONDITION, []byte, error) {
	return nil, b, fmt.Errorf("expected SEARCH_CONDITION")
}

func expectColumnAlias(b []byte) (*COLUMNALIAS, []byte, error) {
	b = skipSpaces(b)
	lexeme, rest := scanner.Identifier(b)
	if lexeme == nil {
		return nil, b, fmt.Errorf("expected COLUMN ALIAS")
	}
	return &COLUMNALIAS{Name: strings.ToUpper(string(lexeme))}, rest, nil
}

//func expectFrom(b []byte) (*FROM, []byte, error) {
//	b = skipSpaces(b)
//	lexeme, rest := scanner.From(b)
//	if lexeme == nil {
//		return nil, rest, fmt.Errorf("expected FROM")
//	}
//	return &FROM{}, rest, nil
//}


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

//import "github.com/mdhender/pogodb/internal/scanner"
//
//type NODE struct {
//	Type NODETYPE
//	DataType DATATYPE
//}
//type NODETYPE int
//const (
//	ntError NODETYPE = iota
//	ntDataType
//)
//
//func (p *PARSER) Statement() {
//	p.Or(p.SelectStatement, p.CreateTableStatement, p.DeleteStatementSearched, p.DropTableStatement, p.InsertStatement, p.UpdateStatementSearched)
//}
//func (p *PARSER) SelectStatement() {
//	p.Keyword("SELECT")
//	p.Or(p.Keyword("ALL"), p.Keyword("DISTINCT"))
//	p.SelectList()
//	p.Keyword("FROM")
//	p.TableReferenceList()
//	p.Keyword("WHERE")
//	p.SearchCondition()
//	p.OrderByClause()
//}
//func (p *PARSER) CreateTableStatement(b []byte) {
//	p.Keyword("CREATE")
//	p.Keyword("TABLE")
//	p.BaseTableName()
//	p.Literal("(")
//	for {
//		p.ColumnIdentifier()
//		dt, rest, err := p.DataType(b)
//		if err != nil {
//		} else if dt == nil {
//
//		}
//		b = skipSpaces(rest)
//		if !p.Literal(",") {
//			break
//		}
//	}
//	p.Literal(")")
//}
//func (p *PARSER) DeleteStatementSearched() {
//	p.Keyword("DELETE")
//	p.Keyword("FROM")
//	p.TableName()
//	p.Keyword("WHERE")
//	p.SearchCondition()
//}
//func (p *PARSER) DropTableStatement() {
//	p.Keyword("DROP")
//	p.Keyword("TABLE")
//	p.BaseTableName()
//}
//func (p *PARSER) InsertStatement() {
//	p.Keyword("INSERT")
//	p.Keyword("INTO")
//	p.TableName()
//	if p.Literal("(") {
//		for {
//			p.ColumnIdentifier()
//			if !p.Literal(",") {
//				break
//			}
//		}
//		p.Literal(")")
//	}
//	p.Keyword("VALUES")
//	p.Literal("(")
//	for {
//		p.InsertValue()
//		if !p.Literal(",") {
//			break
//		}
//	}
//	p.Literal(")")
//}
//func (p *PARSER) UpdateStatementSearched() {
//	p.Keyword("UPDATE")
//	p.TableName()
//	p.Keyword("SET")
//	for {
//		p.ColumnIdentifier()
//		p.Or(p.Expression, p.Keyword("NULL"))
//		if !p.Literal(",") {
//			break
//		}
//	}
//	p.Keyword("WHERE")
//	p.SearchCondition()
//}
//func (p *PARSER) SelectList() {
//	if p.Literal("*") {
//		return
//	}
//	for {
//		p.SelectSublist()
//		if !p.Literal(",") {
//			break
//		}
//	}
//}
//func (p *PARSER) TableReferenceList() {
//	for {
//		p.TableReference()
//		if !p.Literal(",") {
//			break
//		}
//	}
//}
//func (p *PARSER) SearchCondition() {
//	p.Expression()
//}
//func (p *PARSER) SelectSublist() {
//	p.Expression()
//}
//func (p *PARSER) Expression() {
//	p.Term()
//	p.Expression(); p.Literal("+"); p.Literal("-"); p.Term()
//}
//func (p *PARSER) Term() {
//	p.Factor()
//	p.Term(); p.Literal("*"); p.Literal("/"); p.Factor()
//}
//func (p *PARSER) Factor() {
//	p.Literal("+"); p.Literal("-")
//	p.Primary()
//}
//func (p *PARSER) TableReference() {
//	p.TableName()
//}
//func (p *PARSER) TableName() {
//	p.TableIdentifier()
//}
//func (p *PARSER) TableIdentifier() {
//	p.UserDefinedName()
//}
//func (p *PARSER) UserDefinedName() {}
//func (p *PARSER) BaseTableName() {
//	p.BaseTableIdentifier()
//}
//func (p *PARSER) ColumnIdentifier() {
//	p.UserDefinedName()
//}
//
//type DATATYPE int
//const (
//	dtError  DATATYPE = iota
//	dtFloat64
//	dtInt64
//	dtString
//	dtTimestamp
//)
//func (p *PARSER) DataType(b []byte) (*NODE, []byte, error) {
//	if lexeme, rest := keyword(b, []byte("FLOAT64"), []byte("float64")); lexeme != nil {
//		return &NODE{Type: ntDataType, DataType: dtFloat64}, rest, nil
//	}
//	if lexeme, rest := keyword(b, []byte("INT64"), []byte("int64")); lexeme != nil {
//		return &NODE{Type: ntDataType, DataType: dtInt64}, rest, nil
//	}
//	if lexeme, rest := keyword(b, []byte("STRING"), []byte("string")); lexeme != nil {
//		return &NODE{Type: ntDataType, DataType: dtString}, rest, nil
//	}
//	if lexeme, rest := keyword(b, []byte("TIMESTAMP"), []byte("timestamp")); lexeme != nil {
//		return &NODE{Type: ntDataType, DataType: dtTimestamp}, rest, nil
//	}
//	return nil, nil, nil
//}
//func (p *PARSER) InsertValue() {
//	p.DynamicParameter()
//	p.Literal()
//	p.Literal("NULL")
//	p.Literal("USER")
//}
//func (p *PARSER) BaseTableIdentifier() {
//	p.UserDefinedName()
//}
//func (p *PARSER) BooleanFactor() {
//	p.Literal("NOT")
//	p.BooleanPrimary()
//}
//func (p *PARSER) BooleanPrimary() {
//	p.ComparisonPredicate()
//	p.Literal("("); p.SearchCondition(); p.Literal(")")
//}
//func (p *PARSER) ComparisonPredicate() {
//	p.Expression()
//	p.ComparisonOperator()
//	p.Expression()
//}
//func (p *PARSER) BooleanTerm() {
//	p.BooleanFactor()
//	p.Literal("AND")
//	p.BooleanTerm()
//}
//func (p *PARSER) CharacterStringLiteral() {}
//func (p *PARSER) ComparisonOperator() {}
//func (p *PARSER) Digit() {}
//func (p *PARSER) DynamicParameter() {
//	p.Literal("?")
//}
//func (p *PARSER) Primary() {}
//func (p *PARSER) Literal(s string) bool {return false}
//func (p *PARSER) OrderByClause() {
//	p.Literal("ORDER")
//	p.Literal("BY")
//	for {
//		p.SortSpecification()
//		if !p.Literal(",") {
//			break
//		}
//	}
//}
//func (p *PARSER) SortSpecification() {
//	p.Or(p.UnsignedInteger, p.ColumnName)
//	p.Or(p.Literal("ASC"), p.Literal("DESC"))
//}
//func (p *PARSER) UnsignedInteger() {}
//func (p *PARSER) ColumnName() {}
//
//func (p *PARSER) Keyword(s string) {}
//func (p *PARSER) Or(n ...func()) {}

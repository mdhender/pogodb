# Supported Grammar

Using [a minimal grammar from Microsoft](https://docs.microsoft.com/en-us/sql/odbc/reference/appendixes/sql-minimum-grammar?view=sql-server-ver15) as a base.

statement ::=
      selectStatement
    | createTableStatement
    | deleteStatementSearched
    | dropTableStatement
    | insertStatement
    | updateStatementSearched .

selectStatement ::=
    SELECT (ALL | DISTINCT)?
    selectList
    FROM tableReferenceList
    (WHERE searchCondition)?
    orderByClause? .

    // selectList cannot contain parameters

selectList ::=
      STAR
    | selectSublist ( COMMA selectSublist )* .

selectSublist ::= expression .

expression ::=
      term
    | expression {+|-} term .

term ::=
      factor
    | term {*|/} factor .

tableReferenceList ::= tableReference ( COMMA tableReference )* .

tableReference ::= tableName .

tableName ::= tableIdentifier .

tableIdentifier ::= userDefinedName .






    // As a dataType in a createTableStatement, applications must use a
    // data type from the TYPE_NAME column of the result set returned by
    // SQLGetTypeInfo.

createTableStatement ::=
    CREATE TABLE baseTableName
    OPENPAREN
    columnIdentifier dataType ( COMMA columnIdentifier dataType )*
    CLOSEPAREN .

deleteStatementSearched ::=
    DELETE FROM tableName ( WHERE searchCondition ) .

dropTableStatement ::=
    DROP TABLE baseTableName .

insertStatement ::=
    INSERT INTO tableName
    (OPENPAREN columnIdentifier (COMMA columnIdentifier)* CLOSEPAREN)?
    VALUES OPENPAREN insertValue (COMMA insertValue)* CLOSEPAREN .

updateStatementSearched ::=
    UPDATE tableName
    SET columnIdentifier = (expression | NULL)
    (COMMA columnIdentifier = (expression | NULL))*
    (WHERE searchCondition)? .

baseTableIdentifier ::= userDefinedName .

baseTableName ::= baseTableIdentifier .

booleanFactor ::= NOT? booleanPrimary .

booleanPrimary ::= comparisonPredicate | PARENOPEN searchCondition PARENCLOSE .

booleanTerm ::= booleanFactor (AND booleanTerm)? .

    // character is any character in the character set of the driver/data source.
    // To include a single literal quote character ('') in a characterStringLiteral,
    // use two literal quote characters [''''].)

characterStringLiteral ::= ''{character}...'' .

columnIdentifier ::= userDefinedName .

columnName ::= (tableName DOT)? columnIdentifier .

comparisonOperator ::= < | > | <= | >= | = | <> .

comparisonPredicate ::= expression comparisonOperator expression .

    // characterStringType is any data type for which the ""DATA_TYPE""
    // column in the result set returned by SQLGetTypeInfo is either SQL_CHAR or SQL_VARCHAR.

dataType ::= characterStringType .

digit ::= 0 | 1 | 2 | 3 | 4 | 5 | 6 | 7 | 8 | 9 .

dynamicParameter ::= QUESTIONMARK .

factor ::= (PLUS|MINUS) primary .

insertValue ::=
      dynamicParameter
    | literal
    | NULL
    | USER .

letter ::= lowerCaseLetter | upperCaseLetter .

literal ::= characterStringLiteral .

lowerCaseLetter ::= a | b | c | d | e | f | g | h | i | j | k | l | m | n | o | p | q | r | s | t | u | v | w | x | y | z .

orderByClause ::= ORDER BY sortSpecification (COMMA sortSpecification)* .

primary ::= columnName
    | dynamicParameter
    | literal
    | ( expression ) .

searchCondition ::= booleanTerm (OR searchCondition)? .

sortSpecification ::= (unsignedInteger | columnName) (ASC | DESC)? .

unsignedInteger ::= DIGIT+ .

upperCaseLetter ::= A | B | C | D | E | F | G | H | I | J | K | L | M | N | O | P | Q | R | S | T | U | V | W | X | Y | Z .

userDefinedName ::= letter (digit | letter | UNDERSCORE)* .

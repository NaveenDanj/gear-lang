package lib

type Lexeme struct {
	LexType string
	Value   string
}

type LexemeDriver struct {
	LexList []Lexeme
}

type Token struct {
	Type  string
	Value string
}

type TokenDriver struct {
	TokenList   []Token
	KeyWordList [30]string
	Operators   map[string]int
	Numbers     map[byte]int
	Validator   map[byte]int
}

type Program struct {
	Statements []Statement
}

type Statement struct {
	StatementType string
	Value         interface{}
}

type PrintStatement struct {
	Expression interface{}
}

type LetStatement struct {
	VariableName string
	DataType     string
	Expression   interface{}
}

type Expression struct {
	Value    interface{}
	Left     *Expression
	Right    *Expression
	Operator string
}

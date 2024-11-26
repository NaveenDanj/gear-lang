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
	Other interface{}
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

type ImportStatement struct {
	ImportPath string
}

type IfStatement struct {
	Condition *Expression
	ThenBlock Statement
	ElseBlock Statement
}

type Expression struct {
	Left     *Expression
	Operator string
	Right    *Expression
	Value    interface{}
}

type StatementBlock struct {
	Type       string
	Statements []Statement
}

type WhileStatement struct {
	Condition *Expression
	Body      Statement
}

type FunctionDeclarationStatement struct {
	FunctionName string
	Parameters   []FormalParameter
	Body         Statement
	ReturnType   string
	IsExported   bool
}

type FormalParameter struct {
	Name     string
	DataType string
}

type ReturnStatement struct {
	Expression *Expression
}

type StructDeclarationStatement struct {
	Name       string
	Fields     Statement
	IsExported bool
}

type StructField struct {
	Name     string
	DataType string
	Body     Statement
}

type VaribleAssignmentStatement struct {
	VariableName interface{}
	Expression   *Expression
}

type ObjectPropertyAccessExpression struct {
	ObjectName   string
	PropertyName *ObjectPropertyAccessExpression
	Value        string
}

type ArrayExpressionElement struct {
	Elements interface{}
}

type ExpressionGenericType struct {
	Expression interface{}
}

type FunctionCallExpression struct {
	FunctionName string
	Arguments    []*Expression
}

type ArrayIndexAccessExpression struct {
	ArrayName       string
	IndexExpression []*Expression
}

// -------------------------------- new types ---------------------------------------


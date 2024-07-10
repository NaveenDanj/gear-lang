package nodes

type LetStatement struct {
	TokenType    string
	VariableName string
	Expression   *Expression
}

type Expression struct {
	NodeType string
	Value    interface{}
	Left     *Expression
	Right    *Expression
	Operator string
}

package nodes

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

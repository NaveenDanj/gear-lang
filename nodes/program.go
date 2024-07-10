package nodes

type Program struct {
	Statements []Statement
}

type Statement struct {
	StatementType string
	Value         interface{}
}

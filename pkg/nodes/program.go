package nodes

type Program struct {
	Statements []Statement
}

type Statement struct {
	NodeType string
	Value    interface{}
}

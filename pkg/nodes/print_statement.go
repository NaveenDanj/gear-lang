package nodes

import "gear-lang/pkg/lib"

func HandlePrintStatement(tokenList []lib.Token, index int) (int, lib.Statement) {
	expressionStr := ""
	counter := index + 1

	for i := index + 1; i < len(tokenList) && tokenList[i].Type != "SEMICOLON"; i++ {
		expressionStr += tokenList[i].Value
		counter += 1
	}

	newPrintStatement := lib.PrintStatement{
		Expression: expressionStr,
	}

	newStatement := lib.Statement{
		StatementType: "PRINT",
		Value:         newPrintStatement,
	}

	return counter, newStatement
}

package nodes

import (
	"gear-lang/pkg/lib"
	"gear-lang/pkg/util"
)

func HandlePrintStatement(tokenList []lib.Token, index int) (int, lib.Statement) {
	expressionStr := ""
	counter := index + 1

	for i := index + 1; i < len(tokenList) && tokenList[i].Type != "SEMICOLON"; i++ {
		expressionStr += tokenList[i].Value
		counter += 1
	}

	expr, err := util.ParseExpressionTokens(tokenList[index+1 : counter])

	if err != nil {
		panic("Error : Expression parsing error: " + err.Error())
	}

	newPrintStatement := lib.PrintStatement{
		Expression: expr,
	}

	// expr.PrintExpression("", true)

	newStatement := lib.Statement{
		StatementType: "PRINT",
		Value:         newPrintStatement,
	}

	return counter, newStatement
}

package nodes

import (
	"gear-lang/pkg/lib"
	"gear-lang/pkg/util"
)

func HandleWhileStatementCondition(tokenList []lib.Token, index int) (int, lib.WhileStatement) {
	counter := index + 1

	for i := index + 1; i < len(tokenList)-1 && tokenList[i].Type != "LEFT_BRACE"; i++ {
		counter += 1
	}

	l := tokenList[index+1 : counter]

	expr, err := util.ParseBooleanExpressionTokens(l)

	if err != nil {
		panic("Error: Error parsing boolean expression " + err.Error())
	}

	whilestmt := lib.WhileStatement{
		Condition: expr,
	}

	return counter, whilestmt
}

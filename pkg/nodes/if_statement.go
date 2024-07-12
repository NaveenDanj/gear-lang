package nodes

import (
	"fmt"
	"gear-lang/pkg/lib"
	"gear-lang/pkg/util"
)

func HandleIfStatement(tokenList []lib.Token, index int) (int, lib.Statement) {

	counter := index + 1

	for i := index + 1; i < len(tokenList)-1 && tokenList[i].Type != "LEFT_BRACE"; i++ {
		counter += 1
	}

	l := tokenList[index+1 : counter]

	fmt.Println("Data - ", l)

	expr, err := util.ParseBooleanExpressionTokens(l)

	fmt.Printf("%#v\n", expr.Right)

	if err != nil {
		panic("Error: Error parsing boolean expression " + err.Error())
	}

	return counter, lib.Statement{
		StatementType: "IF",
		Value:         lib.IfStatement{},
	}

}

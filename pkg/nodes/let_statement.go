package nodes

import (
	"gear-lang/pkg/lib"
	"gear-lang/pkg/util"
)

func HandleVariableDeclarationStatement(tokenList []lib.Token, index int) (int, lib.Statement) {

	dataType := tokenList[index+1].Value
	varName := tokenList[index+2].Value
	expressionStr := ""
	counter := index + 4

	for i := index + 4; i < len(tokenList) && tokenList[i].Type != "SEMICOLON"; i++ {
		expressionStr += tokenList[i].Value
		counter += 1
	}

	// TODO: have to handle expression strings
	expr, err := util.ParseExpressionTokens(tokenList[index+4 : counter])

	// expr.PrintExpression("", true)

	if err != nil {
		panic("Error : Expression parsing error: " + err.Error())
	}

	newLetStatement := lib.LetStatement{
		VariableName: varName,
		DataType:     dataType,
		Expression:   expr,
	}

	newStatement := lib.Statement{
		StatementType: "VARIABLE_DECLARATION",
		Value:         newLetStatement,
	}

	return counter, newStatement

}

func HandleVariableAssignmentStatement(tokenList []lib.Token, index int) (int, lib.Statement) {
	varName := tokenList[index-1].Value
	expressionStr := ""
	counter := index + 1

	for i := index + 1; i < len(tokenList) && tokenList[i].Type != "SEMICOLON"; i++ {
		expressionStr += tokenList[i].Value
		counter += 1
	}

	// TODO: have to handle expression strings
	expr, err := util.ParseExpressionTokens(tokenList[index+1 : counter])

	if err != nil {
		panic("Error: Error in parsing variable assingment expression")
	}

	st := lib.VaribleAssignmentStatement{
		VariableName: varName,
		Expression:   expr,
	}

	newStatement := lib.Statement{
		StatementType: "VARIABLE_ASSIGNMENT",
		Value:         st,
	}

	return counter, newStatement

}

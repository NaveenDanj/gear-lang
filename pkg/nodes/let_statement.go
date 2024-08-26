package nodes

import (
	"fmt"
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
	if tokenList[index+4].Type == "LEFT_BRACKET" {
		fmt.Println("--------------Found list ------> ")
		expr, _ := util.ParseArrayExpressionWrapper(tokenList, index+4)

		newLetStatement := lib.LetStatement{
			VariableName: varName,
			DataType:     dataType,
			Expression:   expr,
		}

		newStatement := lib.Statement{
			StatementType: "VARIABLE_DECLARATION",
			Value:         newLetStatement,
		}

		// fmt.Print("Expr is read: ")
		// val, _ := newLetStatement.Expression.(lib.ArrayExpressionStatement)
		// fmt.Printf("%#v\n", val.Body)

		return counter, newStatement

	}

	expr, err := util.ParseExpressionTokens(tokenList[index+4 : counter])
	fmt.Printf("%#v\n", expr)

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

	expr, err := util.ParseExpressionTokens(tokenList[index+1 : counter])

	if err != nil {
		panic("Error: Error in parsing variable assingment expression")
	}

	if util.IsPropertyExpressions(varName) {
		prop_exper := util.HandleParsePropertyExpressions(varName, 0, "")
		varName := prop_exper
		fmt.Println("Property assignment")
		fmt.Printf("%#v\n", prop_exper)

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

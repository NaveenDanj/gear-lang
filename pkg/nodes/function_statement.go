package nodes

import (
	"gear-lang/pkg/lib"
	"gear-lang/pkg/util"
)

func HandleFunctionDeclarationStatement(tokenList []lib.Token, index int, isDefault bool) (int, lib.FunctionDeclarationStatement) {
	returnType := tokenList[index+1].Value
	functionName := tokenList[index+2].Value
	stepper := 4

	if !isDefault {
		returnType = tokenList[index+3].Value
		functionName = tokenList[index].Value
		stepper = 3
	}

	paramList := make([]lib.FormalParameter, 0)
	i := index + stepper

	for i = index + stepper; i < len(tokenList) && tokenList[i].Type != "LEFT_BRACE"; i++ {

		if tokenList[i].Type == "COMMA" || (tokenList[i].Type == "RIGHT_PARANTHESES" && tokenList[i-1].Type == "IDENTIFIER") {
			paramName := tokenList[i-1].Value
			paramDataType := tokenList[i-2].Value

			param := lib.FormalParameter{
				Name:     paramName,
				DataType: paramDataType,
			}
			paramList = append(paramList, param)
		}
	}

	newFunction := lib.FunctionDeclarationStatement{
		ReturnType:   returnType,
		FunctionName: functionName,
		Parameters:   paramList,
		IsExported:   false,
	}

	return i, newFunction

}

func HandleReturnStatement(tokenList []lib.Token, index int) (int, lib.Statement) {
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

	newReturStatement := lib.ReturnStatement{
		Expression: expr,
	}

	newStatement := lib.Statement{
		StatementType: "RETURN",
		Value:         newReturStatement,
	}

	return counter, newStatement
}

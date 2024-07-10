package nodes

import "gear-lang/pkg/lib"

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

	newLetStatement := lib.LetStatement{
		VariableName: varName,
		DataType:     dataType,
		Expression:   nil,
	}

	newStatement := lib.Statement{
		StatementType: "VARIABLE_DECLARATION",
		Value:         newLetStatement,
	}

	return counter, newStatement

}

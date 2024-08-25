package util

import (
	"gear-lang/pkg/lib"
)

func HandlePreProcessFunctionCallExpression(tokens []lib.Token, index int, closeParan int) (lib.FunctionCallExpression, int) {
	funcName := tokens[index-2].Value
	// lastSepIndex := index
	argumentList := make([]*lib.Expression, 0)
	funcTokenList := make([]lib.Token, 0)

	for index <= closeParan {

		token := tokens[index]

		if token.Type == "LEFT_PARANTHESES" && tokens[index-1].Type == "IDENTIFIER" {
			// get the matching closing parantheses of the function call
			newCloseParan := GetFunctionCallerMatchingParan(tokens, index)
			// function call token = recursivly call the function call tokenizer parser
			newToken, newIndex := HandlePreProcessFunctionCallExpression(tokens, index+1, newCloseParan)

			// remove the token range from [index : matching_closing_parantheses+1] and add function call token
			funcTokenList = append(funcTokenList, lib.Token{
				Type:  "FunctionCallExpressionToken",
				Value: "",
				Other: newToken,
			})

			// argumentList = append(argumentList, newToken)

			index = newIndex
			continue

		} else if token.Type == "RIGHT_PARANTHESES" && index == closeParan {
			// create the function call token and return it with the updated index

			expr, err := ParseExpressionTokens(funcTokenList)

			if err != nil {
				panic("Error while parsing argument expression!")
			}

			argumentList = append(argumentList, expr)

			// funcTokenList = make([]lib.Token, 0)
			funcCallExpressionToken := lib.FunctionCallExpression{
				FunctionName: funcName,
				Arguments:    argumentList,
			}

			return funcCallExpressionToken, index + 1

		} else if token.Type == "COMMA" {
			// parse the expression from the lastSepIndex
			expr, err := ParseExpressionTokens(funcTokenList)

			if err != nil {
				panic("Error while parsing argument expression!")
			}

			// append it to the argument list
			argumentList = append(argumentList, expr)
			funcTokenList = make([]lib.Token, 0)

			// update the lastSepIndex value
			// lastSepIndex = index + 1
		} else {
			// ignore

			if token.Type == "IDENTIFIER" && tokens[index+1].Type == "LEFT_PARANTHESES" {
				index++
				continue
			}

			funcTokenList = append(funcTokenList, token)
		}

		index++

	}

	return lib.FunctionCallExpression{
		FunctionName: funcName,
		Arguments:    argumentList,
	}, closeParan

}

func HandleParseFunctionCallExpressionWrapper(tokens []lib.Token, index int) (lib.FunctionCallExpression, int) {
	close := GetFunctionCallerMatchingParan(tokens, index)
	outFuncExpr, newIndex := HandlePreProcessFunctionCallExpression(tokens, index+1, close)
	return outFuncExpr, newIndex
}

func GetFunctionCallerMatchingParan(tokens []lib.Token, index int) int {

	stack := make([]lib.Token, 0)

	for index < len(tokens) {

		token := tokens[index]

		if token.Type == "LEFT_PARANTHESES" {
			stack = append(stack, token)
			index++
			continue
		} else if token.Type == "RIGHT_PARANTHESES" {

			stack = stack[0 : len(stack)-1]

			if len(stack) == 0 {
				return index
			}

		}

		index++
	}

	return -1

}

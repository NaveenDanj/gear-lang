package util

import (
	"gear-lang/pkg/lib"
)

func ParseArrayExpression(list []lib.Token, startIndex int, end int) (lib.ArrayExpressionElement, int) {

	arrList := make([]lib.ArrayExpressionElement, 0) // Changed to interface{} to accommodate both arrays and individual elements
	lastElemIndex := startIndex
	index := startIndex
	tokenList := make([]lib.Token, 0)

	for index <= end {
		token := list[index]

		if token.Type == "LEFT_BRACKET" {
			// Parse nested array
			terminalIndex := GetArrayIndexAccessMatchingBracket(list, index)
			elem, newIndex := ParseArrayExpression(list, index+1, terminalIndex)
			arrList = append(arrList, elem)
			lastElemIndex = newIndex + 1
			index = newIndex
			continue

		} else if token.Type == "RIGHT_BRACKET" {
			// End of the array, parse the last element before the bracket
			if lastElemIndex < index {
				expr, _ := ParseExpressionTokens(list[lastElemIndex:index])
				arrList = append(arrList, lib.ArrayExpressionElement{Elements: expr})
			}
			return lib.ArrayExpressionElement{Elements: arrList}, index + 1

			// else if token.Type == "IDENTIFIER" &&

		} else if token.Type == "COMMA" {
			// Parse element between commas
			if lastElemIndex < index {
				expr, _ := ParseExpressionTokens(list[lastElemIndex:index])
				arrList = append(arrList, lib.ArrayExpressionElement{Elements: expr})
			}
			lastElemIndex = index + 1
		} else {
			tokenList = append(tokenList, token)
		}
		index++
	}

	return arrList[0], index
}

func ParseArrayExpressionWrapper(list []lib.Token, startIndex int) (lib.ArrayExpressionElement, int) {
	terminalIndex := GetArrayIndexAccessMatchingBracket(list, startIndex)
	ParseArrayExpression, _ := ParseArrayExpression(list, startIndex, terminalIndex+1)
	return ParseArrayExpression, terminalIndex + 1
}

func HandleAccessArrayIndexExpression(tokens []lib.Token, index int, closeBracket int) (lib.ArrayIndexAccessExpression, int) {

	// find the last couple of brackets of the array index access expression
	exprTokenList := make([]lib.Token, 0)
	arrayName := ""
	if len(tokens) >= 3 {
		arrayName = tokens[index-2].Value
	}
	bracketExprList := make([]*lib.Expression, 0)

	for index <= closeBracket {

		token := tokens[index]

		if token.Type == "LEFT_BRACKET" {

			// get the last closing bracket index
			_closeBracket := GetArrayIndexAccessMatchingBracket(tokens, index)
			// recursively call the handle access array index expression function
			outElem, _ := HandleAccessArrayIndexExpression(tokens, index+1, _closeBracket)

			newToken := lib.Token{
				Type:  "ArrayIndexAccessExpression",
				Other: outElem,
			}

			exprTokenList = append(exprTokenList, newToken)

			index = _closeBracket + 1
			continue

		} else if token.Type == "RIGHT_BRACKET" && index == closeBracket {

			// parse the token into expressions
			expr, err := ParseExpressionTokens(exprTokenList)
			bracketExprList = append(bracketExprList, expr)

			if err != nil {
				panic("Error while parsing array index expression!")
			}

			// create new token
			newArrayExpression := lib.ArrayIndexAccessExpression{
				ArrayName:       arrayName,
				IndexExpression: bracketExprList,
			}

			newToken := lib.Token{
				Type:  "ArrayIndexAccessExpression",
				Other: newArrayExpression,
			}

			exprTokenList = append(exprTokenList, newToken)

			if tokens[closeBracket+1].Type == "LEFT_BRACKET" {
				break
			}

			return newArrayExpression, index

		} else {

			if token.Type == "IDENTIFIER" && tokens[index+1].Type == "LEFT_BRACKET" {
				index++
				continue
			}

			exprTokenList = append(exprTokenList, tokens[index])
		}

		index++

	}

	newArrayExpression := lib.ArrayIndexAccessExpression{
		ArrayName:       arrayName,
		IndexExpression: bracketExprList,
	}

	for tokens[closeBracket+1].Type == "LEFT_BRACKET" {
		newCloseBracket := GetArrayIndexAccessMatchingBracket(tokens, closeBracket+1)
		arrExpr, _ := HandleAccessArrayIndexExpression(tokens, closeBracket+2, newCloseBracket)
		newArrayExpression.IndexExpression = append(newArrayExpression.IndexExpression, arrExpr.IndexExpression[0])
		closeBracket = newCloseBracket
	}

	return newArrayExpression, closeBracket + 1

}

func HandleParseArrayIndexAccessExpressionWrapper(tokens []lib.Token, index int) (lib.ArrayIndexAccessExpression, int) {
	closeBracket := GetArrayIndexAccessMatchingBracket(tokens, index)
	outArrayExpression, newIndex := HandleAccessArrayIndexExpression(tokens, index+1, closeBracket)
	return outArrayExpression, newIndex
}

func GetArrayIndexAccessMatchingBracket(tokens []lib.Token, index int) int {

	stack := make([]lib.Token, 0)

	for index < len(tokens) {

		if tokens[index].Type == "LEFT_BRACKET" {
			stack = append(stack, tokens[index])
			index++
			continue
		} else if tokens[index].Type == "RIGHT_BRACKET" {

			stack = stack[0 : len(stack)-1]

			if len(stack) == 0 {
				return index
			}

		}

		index++
	}

	return index

}

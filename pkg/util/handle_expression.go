package util

import (
	"fmt"
	"gear-lang/pkg/lib"
)

func ParseExpression(index int, tokens []lib.Token, rlist []lib.Token) (*lib.Expression, int, error) {
	if index >= len(tokens) {
		return nil, index, fmt.Errorf("unexpected end of tokens")
	}

	var left *lib.Expression
	var err error

	// Parse the primary expression (number, identifier, or parenthesized expression)
	left, index, err = parsePrimaryExpression(tokens, index)
	if err != nil {
		return nil, index, err
	}

	for index < len(tokens) {
		token := tokens[index]

		if token.Type != "PLUS_OPERATOR" && token.Type != "MINUS_OPERATOR" && token.Type != "MULTIPLY_OPERATOR" && token.Type != "DIVIDE_OPERATOR" {
			break
		}

		index++
		right, newIndex, err := parsePrimaryExpression(tokens, index)
		if err != nil {
			return nil, index, err
		}

		left = &lib.Expression{
			Left:     left,
			Operator: token.Value,
			Right:    right,
		}
		index = newIndex
	}

	return left, index, nil
}

// Helper function to parse a primary expression (number, identifier, or parenthesized expression)
func parsePrimaryExpression(tokens []lib.Token, index int) (*lib.Expression, int, error) {
	if index >= len(tokens) {
		return nil, index, fmt.Errorf("unexpected end of tokens")
	}

	tokens = BaseExpressionParser(tokens)
	fmt.Println("Expression tokens are ----> ", tokens)
	token := tokens[index]

	switch token.Type {
	case "NUMERIC_LITERAL", "IDENTIFIER", "STRING_LITERAL", "FunctionCallExpressionToken", "ArrayIndexAccessExpression":

		if token.Type == "IDENTIFIER" {
			if IsPropertyExpressions(token.Value) {
				struct_expr := HandleParsePropertyExpressions(token.Value, index, "")
				expr_val := lib.Expression{Value: struct_expr.PropertyName.Value}
				return &expr_val, index + 1, nil
			}
		} else if token.Type == "FunctionCallExpressionToken" {
			expr_val := lib.Expression{Value: token}
			return &expr_val, index + 1, nil
		} else if token.Type == "ArrayIndexAccessExpression" {
			expr_val := lib.Expression{Value: token}
			return &expr_val, index + 1, nil
		}

		value := token.Value
		index++
		return &lib.Expression{Value: value}, index, nil

	case "LEFT_PARANTHESES":

		index++
		expr, newIndex, err := ParseExpression(index, tokens, nil)
		if err != nil {
			return nil, index, err
		}
		if tokens[newIndex].Type != "RIGHT_PARANTHESES" {
			return nil, index, fmt.Errorf("expected closing parenthesis")
		}
		newIndex++
		return expr, newIndex, nil

	default:
		return nil, index, fmt.Errorf("unexpected token: %s", token.Value)
	}
}

// Helper function to start parsing from the beginning
func ParseExpressionTokens(tokens []lib.Token) (*lib.Expression, error) {
	result, _, err := ParseExpression(0, tokens, nil)
	return result, err
}

func IsPropertyExpressions(str string) bool {
	for i := 0; i < len(str); i++ {
		if str[i] == ':' {
			return true
		}
	}
	return false
}

func HandleParsePropertyExpressions(str string, index int, prevString string) *lib.ObjectPropertyAccessExpression {

	if index == len(str)-1 {
		prevString += string(str[index])
		return &lib.ObjectPropertyAccessExpression{
			ObjectName:   "",
			PropertyName: nil,
			Value:        prevString,
		}
	}

	if str[index] == ':' {
		index++
		temp := prevString
		prevString = ""
		return &lib.ObjectPropertyAccessExpression{
			ObjectName:   temp,
			PropertyName: HandleParsePropertyExpressions(str, index, prevString),
		}
	}

	prevString += string(str[index])
	return HandleParsePropertyExpressions(str, index+1, prevString)
}

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

// Handle access array index function
func HandleAccessArrayIndexExpression(tokens []lib.Token, index int, closeBracket int) (lib.ArrayIndexAccessExpression, int) {

	// find the last couple of brackets of the array index access expression
	exprTokenList := make([]lib.Token, 0)
	arrayName := tokens[index-2].Value
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

// base function for picking correct parsing function for the expression
func BaseExpressionParser(tokens []lib.Token) []lib.Token {

	outList := make([]lib.Token, 0)
	index := 0

	// terminator can be any terminal lexeme such as SEMICOLON, RIGHT_PARANTHESES or RIGHT_BRACKET or RIGHT_BRACE
	for index < len(tokens) {
		// try to guess the type of expression in a switch

		if tokens[index].Type == "LEFT_BRACKET" && tokens[index-1].Type == "IDENTIFIER" {
			expr, newIndex := HandleParseArrayIndexAccessExpressionWrapper(tokens, index)

			newToken := lib.Token{
				Type:  "ArrayIndexAccessExpression",
				Other: expr,
			}

			outList = append(outList, newToken)
			index = newIndex
			continue
		} else if tokens[index].Type == "LEFT_PARANTHESES" && tokens[index-1].Type == "IDENTIFIER" {
			expr, newIndex := HandleParseFunctionCallExpressionWrapper(tokens, index)

			newToken := lib.Token{
				Type:  "FunctionCallExpressionToken",
				Other: expr,
			}

			outList = append(outList, newToken)

			index = newIndex
			continue
		} else {

			// remove array indentifier after parsing the array
			// remove the function identifier after parsing the function
			// // remove if have object referencing expression identifier after parsing it

			if index+1 < len(tokens) {
				if tokens[index].Type == "IDENTIFIER" && tokens[index+1].Type != "LEFT_PARANTHESES" {
					index++
					continue
				} else if tokens[index].Type == "IDENTIFIER" && tokens[index+1].Type != "LEFT_BRACKET" {
					index++
					continue
				}
			}

			outList = append(outList, tokens[index])
			// }

			// fmt.Println("Index data and other details -------> ", index, len(tokens)-2)

		}

		index++

	}

	// handle parse the tokens after preprocessing the token list

	return outList

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

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

	token := tokens[index]

	switch token.Type {
	case "NUMERIC_LITERAL", "IDENTIFIER", "STRING_LITERAL", "FunctionCallExpressionToken":

		fmt.Println("Elem => ", token.Type)

		if token.Type == "IDENTIFIER" {
			if IsPropertyExpressions(token.Value) {
				struct_expr := HandleParsePropertyExpressions(token.Value, index, "")
				expr_val := lib.Expression{Value: struct_expr.PropertyName.Value}
				return &expr_val, index + 1, nil
			}
		} else if token.Type == "FunctionCallExpressionToken" {
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

// TODO: function call expression
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

package util

import (
	"fmt"
	"gear-lang/pkg/lib"
)

// Main function to parse binary expressions
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

	fmt.Println("Expression tokens are ----> ", tokens)
	tokens = BaseExpressionParser(tokens)
	token := tokens[index]
	fmt.Println("Expression tokens are out ----> ", tokens, token)

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

// Base function for picking correct parsing function for the expression
func BaseExpressionParser(tokens []lib.Token) []lib.Token {
	outList, _ := CommonLoopFunction(tokens, 0)
	return outList
}

// this helper function can be used to parse expression inside another expression
func IntermediateBaseExpressionParser(tokens []lib.Token) ([]lib.Token, int) {
	outList, index := CommonLoopFunction(tokens, 0)
	fmt.Println("Common loop : ", outList)
	return outList, index
}

// TODO:common for loop func : probably should give meaningful name for this later
func CommonLoopFunction(tokens []lib.Token, index int) ([]lib.Token, int) {

	outList := make([]lib.Token, 0)

	for index < len(tokens) {
		// try to guess the type of expression in a switch
		if index >= 1 && tokens[index].Type == "LEFT_BRACKET" && tokens[index-1].Type == "IDENTIFIER" {
			expr, newIndex := HandleParseArrayIndexAccessExpressionWrapper(tokens, index)

			newToken := lib.Token{
				Type:  "ArrayIndexAccessExpression",
				Other: expr,
			}

			outList = append(outList, newToken)
			index = newIndex
			continue

		} else if index >= 1 && tokens[index].Type == "LEFT_PARANTHESES" && tokens[index-1].Type == "IDENTIFIER" {
			expr, newIndex := HandleParseFunctionCallExpressionWrapper(tokens, index)

			newToken := lib.Token{
				Type:  "FunctionCallExpressionToken",
				Other: expr,
			}

			outList = append(outList, newToken)
			index = newIndex
			continue

		} else if tokens[index].Type == "LEFT_BRACKET" {
			expr, newIndex := ParseArrayExpressionWrapper(tokens, index)

			newToken := lib.Token{
				Type:  "ArrayIndexAccessExpression",
				Other: expr,
			}

			outList = append(outList, newToken)
			index = newIndex
			continue
		} else {

			if index+1 < len(tokens) {
				if tokens[index].Type == "IDENTIFIER" && tokens[index+1].Type != "LEFT_PARANTHESES" {
					outList = append(outList, tokens[index])
					index++
					continue
				} else if tokens[index].Type == "IDENTIFIER" && tokens[index+1].Type != "LEFT_BRACKET" {
					outList = append(outList, tokens[index])
					index++
					continue
				}
			}

			outList = append(outList, tokens[index])
		}

		index++

	}

	return outList, index

}

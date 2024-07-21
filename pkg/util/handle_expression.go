package util

import (
	"fmt"
	"gear-lang/pkg/lib"
	"strings"
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
	case "NUMERIC_LITERAL", "IDENTIFIER", "STRING_LITERAL":

		// if token.Type == "IDENTIFIER" {
		// 	if IsPropertyExpressions() {

		// 	}
		// }

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

func ParseObjectPropertyExpressions(str string) *lib.ObjectPropertyAccessExpression {

	// split the string by : operator
	parts := strings.Split(str, ":")

	if len(parts) != 2 {
		return nil
	}

	newObject := lib.ObjectPropertyAccessExpression{
		ObjectName:   parts[0],
		PropertyName: parts[0],
	}
	return &newObject
}

func IsPropertyExpressions(str string) bool {
	for i := 0; i < len(str); i++ {
		if str[i] == ':' {
			return true
		}
	}
	return false
}

func HandleParsePropertyExpressions(index int, str string, prevObject string, prevProperty string, prevType string) *lib.ObjectPropertyAccessExpression {

	if index == len(str) {
		return &lib.ObjectPropertyAccessExpression{
			ObjectName:   prevObject,
			PropertyName: prevProperty,
		}
	}

	if str[index] == ':' && prevType == "Object" {
		index++ // skip the :
		return HandleParsePropertyExpressions(index, str[index:], prevObject, str[:index-1], "Property")
	}

	if str[index] == ':' && prevType == "Property" {
		index++ // skip the :
		return HandleParsePropertyExpressions(index, str[index:], prevObject, prevType, "Object")
	}

	return HandleParsePropertyExpressions(index, str[index:], prevObject, prevType, "Object")

}

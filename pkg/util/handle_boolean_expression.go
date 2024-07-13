package util

import (
	"fmt"
	"gear-lang/pkg/lib"
)

// ParseExpression recursively parses boolean expressions and builds the AST
func ParseBooleanExpression(index int, tokens []lib.Token) (*lib.Expression, int, error) {
	if index >= len(tokens) {
		return nil, index, fmt.Errorf("unexpected end of tokens")
	}

	var left *lib.Expression
	var err error

	// Parse the primary expression (boolean literal, identifier, or parenthesized expression)
	left, index, err = parsePrimaryBooleanExpression(tokens, index)
	if err != nil {
		return nil, index, err
	}

	for index < len(tokens) {
		token := tokens[index]

		if token.Type != "AND_OPERATOR" && token.Type != "OR_OPERATOR" &&
			token.Type != "EQUAL_OPERATOR" && token.Type != "NOT_EQUAL_OPERATOR" &&
			token.Type != "LESS_OPERATOR" && token.Type != "GREATER_OPERATOR" &&
			token.Type != "LESS_EQUAL_OPERATOR" && token.Type != "GREATER_EQUAL_OPERATOR" && token.Type != "DOUBLE_EQUALS_OPERATOR" {
			break
		}

		index++
		right, newIndex, err := parsePrimaryBooleanExpression(tokens, index)
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

// Helper function to parse a primary expression (boolean literal, identifier, or parenthesized expression)
func parsePrimaryBooleanExpression(tokens []lib.Token, index int) (*lib.Expression, int, error) {
	if index >= len(tokens) {
		return nil, index, fmt.Errorf("unexpected end of tokens")
	}

	token := tokens[index]

	switch token.Type {
	case "BOOLEAN_LITERAL", "IDENTIFIER", "STRING_LITERAL":
		value := token.Value
		index++
		return &lib.Expression{Value: value}, index, nil

	case "LEFT_PARANTHESES":
		index++
		expr, newIndex, err := ParseBooleanExpression(index, tokens)
		if err != nil {
			return nil, index, err
		}
		if newIndex >= len(tokens) || tokens[newIndex].Type != "RIGHT_PARANTHESES" {
			return nil, index, fmt.Errorf("expected closing parenthesis")
		}
		newIndex++
		return expr, newIndex, nil

	case "NOT_OPERATOR":
		index++
		expr, newIndex, err := parsePrimaryExpression(tokens, index)
		if err != nil {
			return nil, index, err
		}
		return &lib.Expression{
			Operator: "NOT",
			Right:    expr,
		}, newIndex, nil

	default:
		return nil, index, fmt.Errorf("unexpected token: %s", token.Value)
	}
}

func ParseBooleanExpressionTokens(tokens []lib.Token) (*lib.Expression, error) {
	result, _, err := ParseBooleanExpression(0, tokens)
	return result, err
}

// func (e *Expression) PrintExpression(indent string, last bool) {
// 	if e == nil {
// 		return
// 	}

// 	// Print the current node
// 	fmt.Print(indent)
// 	if last {
// 		fmt.Print("└─")
// 		indent += "  "
// 	} else {
// 		fmt.Print("├─")
// 		indent += "| "
// 	}

// 	if e.Operator != "" {
// 		fmt.Println(e.Operator)
// 	} else {
// 		fmt.Println(e.Value)
// 	}

// 	// Print the left and right children
// 	if e.Left != nil {
// 		e.Left.PrintExpression(indent, false)
// 	}
// 	if e.Right != nil {
// 		e.Right.PrintExpression(indent, true)
// 	}
// }

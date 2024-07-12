package util

import (
	"fmt"
	"gear-lang/pkg/lib"
)

type Expression struct {
	Left     *Expression
	Operator string
	Right    *Expression
	Value    string
}

func ParseExpression(index int, tokens []lib.Token, rlist []lib.Token) (*Expression, int, error) {
	if index >= len(tokens) {
		return nil, index, fmt.Errorf("unexpected end of tokens")
	}

	var left *Expression
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

		left = &Expression{
			Left:     left,
			Operator: token.Value,
			Right:    right,
		}
		index = newIndex
	}

	return left, index, nil
}

// Helper function to parse a primary expression (number, identifier, or parenthesized expression)
func parsePrimaryExpression(tokens []lib.Token, index int) (*Expression, int, error) {
	if index >= len(tokens) {
		return nil, index, fmt.Errorf("unexpected end of tokens")
	}

	token := tokens[index]

	switch token.Type {
	case "NUMERIC_LITERAL", "IDENTIFIER", "STRING_LITERAL":

		value := token.Value
		index++
		return &Expression{Value: value}, index, nil

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
func ParseExpressionTokens(tokens []lib.Token) (*Expression, error) {
	result, _, err := ParseExpression(0, tokens, nil)
	return result, err
}

func (e *Expression) PrintExpression(indent string, last bool) {
	if e == nil {
		return
	}

	// Print the current node
	fmt.Print(indent)
	if last {
		fmt.Print("└─")
		indent += "  "
	} else {
		fmt.Print("├─")
		indent += "| "
	}

	if e.Operator != "" {
		fmt.Println(e.Operator)
	} else {
		fmt.Println(e.Value)
	}

	// Print the left and right children
	if e.Left != nil {
		e.Left.PrintExpression(indent, false)
	}
	if e.Right != nil {
		e.Right.PrintExpression(indent, true)
	}
}

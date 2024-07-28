package util

import (
	"errors"
	"fmt"
	"gear-lang/pkg/lib"
)

func ParseArrayInitExpression(list []lib.Token, index int, currentList []*lib.Expression, prevExpIndex int) ([]*lib.Expression, int, error) {

	// var outList = make(map[int][]*lib.Expression)

	if index >= len(list) {
		return currentList, index, nil
	}

	switch list[index].Type {
	case "LEFT_BRACKET":
		var newList []*lib.Expression
		res, outIndex, err := ParseArrayInitExpression(list, index+1, newList, index+1)
		if err != nil {
			return nil, 0, err
		}

		// for i := 0; i < len(res); i++ {
		// 	currentList = append(currentList, res[i])
		// }

		fmt.Println("---------------------------------------- ->")
		for i := 0; i < len(res); i++ {
			fmt.Printf("%#v\n", res[i])
			// outList = append(outList, res[i])
		}
		fmt.Println("---------------------------------------- -<")

		return ParseArrayInitExpression(list, outIndex, currentList, outIndex)

	case "RIGHT_BRACKET":
		exprTokenList := list[prevExpIndex:index]
		if len(exprTokenList) > 0 {
			exp, err := ParseExpressionTokens(exprTokenList)
			if err != nil {
				return nil, 0, errors.New("error while parsing array expression: " + err.Error())
			}
			currentList = append(currentList, exp)
		}
		return currentList, index + 1, nil

	case "COMMA":
		exprTokenList := list[prevExpIndex:index]
		if len(exprTokenList) > 0 {
			exp, err := ParseExpressionTokens(exprTokenList)
			if err != nil {
				return nil, 0, errors.New("error while parsing array expression: " + err.Error())
			}
			currentList = append(currentList, exp)
		}
		return ParseArrayInitExpression(list, index+1, currentList, index+1)

	default:
		return ParseArrayInitExpression(list, index+1, currentList, prevExpIndex)
	}
}

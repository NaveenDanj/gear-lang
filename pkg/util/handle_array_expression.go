package util

import (
	"gear-lang/pkg/lib"
)

func ParseArrayExpression(list []lib.Token, startIndex int) (lib.ArrayExpressionElement, int) {
	arrList := make([]lib.ArrayExpressionElement, 0) // Changed to interface{} to accommodate both arrays and individual elements
	lastElemIndex := startIndex
	index := startIndex

	for index < len(list) && list[index].Type != "SEMICOLON" {
		token := list[index]

		if token.Type == "LEFT_BRACKET" {
			// Parse nested array
			elem, newIndex := ParseArrayExpression(list, index+1)
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

		} else if token.Type == "COMMA" {
			// Parse element between commas
			if lastElemIndex < index {
				expr, _ := ParseExpressionTokens(list[lastElemIndex:index])
				arrList = append(arrList, lib.ArrayExpressionElement{Elements: expr})
			}
			lastElemIndex = index + 1
		}
		index++
	}

	return arrList[0], index
}

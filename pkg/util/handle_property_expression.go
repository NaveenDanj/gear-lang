package util

import "gear-lang/pkg/lib"

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

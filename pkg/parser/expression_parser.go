package pkg

import "gear-lang/pkg/lib"

func ParseExpressionBlock(index int, exp_type string, tokens []lib.Token) {

	stack := 1
	counter := index + 1
	endBlockIndex := counter + 1

	// expressionStatement := lib.IBlockStatement{
	// 	Statements: make([]lib.IStatement[any], 0),
	// }

	for stack != 0 && counter < len(tokens) {
		if tokens[counter].Type == "LEFT_PARANTHESES" {
			stack -= 1
			if stack == 0 {
				endBlockIndex = counter
				break
			}
			counter++
			continue
		} else if tokens[counter].Type == "RIGHT_PARANTHESES" {
			stack += 1
		}

		counter++
	}

	i := index + 1

	for {
		if i > endBlockIndex {
			break
		}
	}

}

func ParseBooleanExpression() {

}

func ParseStringExpression() {

}

func ParseExpression() {

}

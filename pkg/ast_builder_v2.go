package pkg

import (
	"fmt"
	"gear-lang/pkg/lib"
)

type ASTBuilder struct {
	CurrentStatementIndex int
	Program               lib.Program
	TokenList             []lib.Token
}

// main parsing function to start parsing
func (ast *ASTBuilder) Parse() {
	ast.CurrentStatementIndex = 0

	for {

		if ast.CurrentStatementIndex > len(ast.TokenList)-1 {
			break
		}

		ast.CurrentStatementIndex = ast.ParseStatement(ast.CurrentStatementIndex)
	}
}

// to parse each and every statement in the program
func (ast *ASTBuilder) ParseStatement(index int) int {

	if ast.TokenList[index].Type == "LEFT_BRACE" {
		fmt.Print("Left brace \n")
		return ast.ParseBlock(index)
	}

	return index + 1
}

// to parse block statements essentially we'll have to use ParseStatement function inside of this function
func (ast *ASTBuilder) ParseBlock(index int) int {
	var stack = make([]int, 0)
	stack = append(stack, index)
	counter := index

	for len(stack) != 0 && counter < len(ast.TokenList) {

		if ast.TokenList[counter].Type == "RIGHT_BRACE" {
			// parse it
			startIndex := stack[len(stack)-1] + 1
			endIndex := counter

			str := ast.TokenList[startIndex:endIndex]
			fmt.Println(str)
			fmt.Println("======================================")
			stack = stack[0 : len(stack)-1]
		}

		counter++

	}

	return counter
}

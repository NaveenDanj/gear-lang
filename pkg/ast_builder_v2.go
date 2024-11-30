package pkg

import (
	"fmt"
	"gear-lang/pkg/lib"
)

type ASTBuilder struct {
	CurrentStatementIndex int
	Program               lib.IProgram
	TokenList             []lib.Token
}

// main parsing function to start parsing
func (ast *ASTBuilder) Parse() {
	ast.CurrentStatementIndex = 0

	for {

		if ast.CurrentStatementIndex > len(ast.TokenList)-1 {
			break
		}

		newStatement, index := ast.ParseStatement(ast.CurrentStatementIndex)
		ast.CurrentStatementIndex = index
		ast.Program.Statements = append(ast.Program.Statements, newStatement)
		fmt.Println("index : ", newStatement)

	}
}

// to parse each and every statement in the program
func (ast *ASTBuilder) ParseStatement(index int) (lib.IStatement[any], int) {

	if ast.TokenList[index].Type == "LEFT_BRACE" {
		statement := lib.IStatement[any]{
			StatementType: "BlockStatement",
		}
		newStatement, newIndex := ast.ParseBlock(index)
		statement.Value = newStatement
		return statement, newIndex
	}

	return lib.IStatement[any]{}, index + 1
}

// to parse block statements essentially we'll have to use ParseStatement function inside of this function
func (ast *ASTBuilder) ParseBlock(index int) (lib.IBlockStatement, int) {
	stack := 1
	counter := index + 1
	endBlockIndex := counter + 1

	blockStatement := lib.IBlockStatement{
		Statements: make([]lib.IStatement[any], 0),
	}

	// find the end of block statement
	for stack != 0 && counter < len(ast.TokenList) {
		if ast.TokenList[counter].Type == "RIGHT_BRACE" {
			stack -= 1
			if stack == 0 {
				endBlockIndex = counter
				break
			}
			counter++
			continue
		} else if ast.TokenList[counter].Type == "LEFT_BRACE" {
			stack += 1
		}

		counter++
	}

	i := index + 1

	for {
		if i > endBlockIndex {
			break
		}
		newStatement, newIndex := ast.ParseStatement(i)
		blockStatement.Statements = append(blockStatement.Statements, newStatement)
		i = newIndex
	}

	return blockStatement, endBlockIndex

}

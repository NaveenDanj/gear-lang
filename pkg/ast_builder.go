package pkg

import (
	"fmt"
	"gear-lang/pkg/nodes"
)

type ASTBuilder struct {
	CurrentStatementIndex int
	Program               nodes.Program
	TokenList             []Token
}

func (ast *ASTBuilder) Parse(index int) {

	ast.CurrentStatementIndex = index

	if len(ast.TokenList) == 0 {
		return
	}

	if ast.TokenList[ast.CurrentStatementIndex].Type == "KEYWORD" {
		ast.handleKeyword(ast.TokenList[ast.CurrentStatementIndex].Value)
	}

}

func (ast *ASTBuilder) handleKeyword(keyword string) {
	switch keyword {

	case "let":
		// ast.Program.Statements = append(ast.Program.Statements, nodes.NewVariableDeclarationStatement(ast.TokenList[ast.CurrentStatementIndex+1].Value))
		// ast.CurrentStatementIndex += 2 // Skip the type and variable name
	default:
		fmt.Println("x is not 1, 2, or 3")
	}

}

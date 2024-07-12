package pkg

import (
	"fmt"
	"gear-lang/pkg/lib"
	"gear-lang/pkg/nodes"
)

type ASTBuilder struct {
	CurrentStatementIndex int
	Program               lib.Program
	TokenList             []lib.Token
}

func (ast *ASTBuilder) Parse(index int) {

	ast.CurrentStatementIndex = index

	if len(ast.TokenList) == 0 || ast.CurrentStatementIndex == len(ast.TokenList)-1 {

		for _, item := range ast.Program.Statements {
			fmt.Printf("%#v\n", item)
		}

		return
	}

	if ast.TokenList[ast.CurrentStatementIndex].Type == "KEYWORD" {
		ast.handleKeyword(ast.TokenList[ast.CurrentStatementIndex].Value)
	} else {
		ast.CurrentStatementIndex += 1
	}

	ast.Parse(ast.CurrentStatementIndex)

}

func (ast *ASTBuilder) handleKeyword(keyword string) {
	switch keyword {

	case "let":
		index, newStatement := nodes.HandleVariableDeclarationStatement(ast.TokenList, ast.CurrentStatementIndex)
		ast.Program.Statements = append(ast.Program.Statements, newStatement)
		ast.CurrentStatementIndex = index
	case "print":
		index, newStatement := nodes.HandlePrintStatement(ast.TokenList, ast.CurrentStatementIndex)
		ast.Program.Statements = append(ast.Program.Statements, newStatement)
		ast.CurrentStatementIndex = index
	case "import":
		index, newStatement := nodes.HandleImportStatement(ast.TokenList, ast.CurrentStatementIndex)
		ast.Program.Statements = append(ast.Program.Statements, newStatement)
		ast.CurrentStatementIndex = index
	case "if":
		index, newStatement := nodes.HandleIfStatement(ast.TokenList, ast.CurrentStatementIndex)
		ast.Program.Statements = append(ast.Program.Statements, newStatement)
		ast.CurrentStatementIndex = index
	default:
		fmt.Printf("Unhandled keyword: %s\n", keyword)
		// panic("Error: Unhandled keyword")
		ast.CurrentStatementIndex += 1
	}

}

func (ast *ASTBuilder) handleParseStatementBlock(tokenList []lib.Token, index int) {
	return
}

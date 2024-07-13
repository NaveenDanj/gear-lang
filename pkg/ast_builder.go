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
		return
	}

	if ast.TokenList[ast.CurrentStatementIndex].Type == "KEYWORD" {
		index, newStatement := ast.handleKeyword(ast.TokenList[ast.CurrentStatementIndex].Value)

		if newStatement.StatementType == "Unhandled" {
			fmt.Println("unhandled keyword")
		}

		ast.Program.Statements = append(ast.Program.Statements, newStatement)
		ast.CurrentStatementIndex = index
	} else {
		ast.CurrentStatementIndex += 1
	}

	ast.Parse(ast.CurrentStatementIndex)

}

func (ast *ASTBuilder) handleKeyword(keyword string) (int, lib.Statement) {
	switch keyword {

	case "let":
		index, newStatement := nodes.HandleVariableDeclarationStatement(ast.TokenList, ast.CurrentStatementIndex)
		return index, newStatement

	case "print":
		index, newStatement := nodes.HandlePrintStatement(ast.TokenList, ast.CurrentStatementIndex)
		return index, newStatement
	case "import":
		index, newStatement := nodes.HandleImportStatement(ast.TokenList, ast.CurrentStatementIndex)
		return index, newStatement
	case "if":
		index, newStatement := nodes.HandleIfStatementCondition(ast.TokenList, ast.CurrentStatementIndex)
		fmt.Print("if statement : ")
		fmt.Printf("%#v\n", newStatement)
		return index, newStatement
	default:
		fmt.Printf("Unhandled keyword: %s\n", keyword)
		index := ast.CurrentStatementIndex + 1
		return index, lib.Statement{StatementType: "Unhandled"}
	}

}

func ParseBlockStatement(tokenList []lib.Token, index int) (int, lib.Statement) {

	if tokenList[index].Type == "LEFT_BRACE" {
		i := index + 1
		i, stmt := ParseBlockStatement(tokenList, &i)
		return i, stmt
	} else if tokenList[index].Type == "RIGHT_BRACE" {

		newBlock := lib.StatementBlock{
			Statements: []lib.Statement{},
		}

		newStmt := lib.Statement{
			StatementType: "StatementBlock",
			Value:         newBlock,
		}

	}

}

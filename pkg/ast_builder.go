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
		index, newStatement := ast.handleKeyword(ast.TokenList[ast.CurrentStatementIndex].Value, ast.CurrentStatementIndex)

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

func (ast *ASTBuilder) handleKeyword(keyword string, i int) (int, lib.Statement) {
	switch keyword {

	case "let":
		index, newStatement := nodes.HandleVariableDeclarationStatement(ast.TokenList, i)
		return index, newStatement
	case "print":
		index, newStatement := nodes.HandlePrintStatement(ast.TokenList, i)
		return index, newStatement
	case "import":
		index, newStatement := nodes.HandleImportStatement(ast.TokenList, i)
		return index, newStatement
	case "if":
		index, ifStatement := nodes.HandleIfStatementCondition(ast.TokenList, i)
		var l []lib.Statement
		newIndex, block := ast.ParseBlockStatement(ast.TokenList, l, index)
		block.StatementType = "BlockStatement"
		ifStatement.ThenBlock = block
		newSt := lib.Statement{StatementType: "IFStatement", Value: ifStatement}
		return newIndex, newSt
	default:
		fmt.Printf("Unhandled keyword: %s\n", keyword)
		index := i + 1
		return index, lib.Statement{StatementType: "Unhandled"}
	}

}

func (ast *ASTBuilder) ParseBlockStatement(tokenList []lib.Token, stmtList []lib.Statement, index int) (int, lib.Statement) {

	if tokenList[index].Type == "LEFT_BRACE" {
		i := index + 1
		var l []lib.Statement
		i, stmt := ast.ParseBlockStatement(tokenList, l, i)
		return i, stmt
	} else if tokenList[index].Type == "RIGHT_BRACE" {

		newBlock := lib.StatementBlock{
			Statements: stmtList,
		}

		newStmt := lib.Statement{
			StatementType: "StatementBlock",
			Value:         newBlock,
		}

		return index, newStmt

	} else {
		index_out, stmt := ast.handleKeyword(tokenList[index].Value, index)

		if stmt.StatementType == "Unhandled" {
			fmt.Println("unhandled keyword in block statement")
		} else {
			stmtList = append(stmtList, stmt)
		}

		i, stmt := ast.ParseBlockStatement(tokenList, stmtList, index_out)
		return i, stmt
	}

}

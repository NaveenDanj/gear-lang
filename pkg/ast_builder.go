package pkg

import (
	"fmt"
	"gear-lang/pkg/lib"
)

type ASTBuilder struct {
	CurrentStatementIndex int
	Program               lib.Program
	TokenList             []Token
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
		index, newStatement := handleVariableDeclarationStatement(ast.TokenList, ast.CurrentStatementIndex)
		ast.Program.Statements = append(ast.Program.Statements, newStatement)
		ast.CurrentStatementIndex = index
	case "print":
		index, newStatement := handlePrintStatement(ast.TokenList, ast.CurrentStatementIndex)
		ast.Program.Statements = append(ast.Program.Statements, newStatement)
		ast.CurrentStatementIndex = index
	default:
		fmt.Printf("Unhandled keyword: %s\n", keyword)
		// panic("Error: Unhandled keyword")
		ast.CurrentStatementIndex += 1
	}

}

func handleVariableDeclarationStatement(tokenList []Token, index int) (int, lib.Statement) {

	dataType := tokenList[index+1].Value
	varName := tokenList[index+2].Value
	expressionStr := ""
	counter := index + 4

	for i := index + 4; i < len(tokenList) && tokenList[i].Type != "SEMICOLON"; i++ {
		expressionStr += tokenList[i].Value
		counter += 1
	}

	// TODO: have to handle expression strings

	newLetStatement := lib.LetStatement{
		VariableName: varName,
		DataType:     dataType,
		Expression:   nil,
	}

	newStatement := lib.Statement{
		StatementType: "VARIABLE_DECLARATION",
		Value:         newLetStatement,
	}

	return counter, newStatement

}

func handlePrintStatement(tokenList []Token, index int) (int, lib.Statement) {
	expressionStr := ""
	counter := index + 1

	for i := index + 1; i < len(tokenList) && tokenList[i].Type != "SEMICOLON"; i++ {
		expressionStr += tokenList[i].Value
		counter += 1
	}

	newPrintStatement := lib.PrintStatement{
		Expression: expressionStr,
	}

	newStatement := lib.Statement{
		StatementType: "PRINT",
		Value:         newPrintStatement,
	}

	return counter, newStatement
}

// func handleImportStatement(tokenList []Token, index int) (int, nodes.Statement) {
// 	return 1
// }

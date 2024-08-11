package pkg

import (
	"fmt"
	"gear-lang/pkg/lib"
	"gear-lang/pkg/nodes"
	"gear-lang/pkg/util"
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
			// val, ok := item.Value.(lib.LetStatement)
			fmt.Printf("%#v\n", item)
			fmt.Println("----------------------------------------")
		}

		return
	}

	if ast.TokenList[ast.CurrentStatementIndex].Type == "KEYWORD" {
		index, newStatement := ast.handleKeyword(ast.TokenList[ast.CurrentStatementIndex].Value, ast.CurrentStatementIndex)

		if newStatement.StatementType == "Unhandled" {
			// fmt.Println("unhandled keyword")
		}

		ast.Program.Statements = append(ast.Program.Statements, newStatement)
		ast.CurrentStatementIndex = index
	} else if ast.TokenList[ast.CurrentStatementIndex].Type == "EQUAL_OPERATOR" {
		fmt.Println("Found variable assignment operation")
		index, stmt := nodes.HandleVariableAssignmentStatement(ast.TokenList, ast.CurrentStatementIndex)
		ast.Program.Statements = append(ast.Program.Statements, stmt)
		ast.CurrentStatementIndex = index
	} else {
		ast.CurrentStatementIndex += 1

		// check whether the token is a possible function calling expression
		if ast.TokenList[ast.CurrentStatementIndex].Type == "IDENTIFIER" && ast.TokenList[ast.CurrentStatementIndex+1].Type == "LEFT_PARANTHESES" {
			closeParan := util.GetFunctionCallerMatchingParan(ast.TokenList, ast.CurrentStatementIndex+1)
			funcExpr, _ := util.HandlePreProcessFunctionCallExpression(ast.TokenList, ast.CurrentStatementIndex+2, closeParan)

			if ast.TokenList[closeParan+1].Type == "SEMICOLON" {
				ast.Program.Statements = append(ast.Program.Statements, lib.Statement{
					StatementType: "FunctionCallExpression",
					Value:         funcExpr,
				})
			}

			ast.CurrentStatementIndex = closeParan

		} else if ast.TokenList[ast.CurrentStatementIndex].Type == "IDENTIFIER" && ast.TokenList[ast.CurrentStatementIndex+1].Type == "LEFT_BRACKET" {
			arrExpr, newIndex := util.HandleParseArrayIndexAccessExpressionWrapper(ast.TokenList, ast.CurrentStatementIndex+1)
			fmt.Printf("%#v\n", arrExpr.IndexExpression[0])
			ast.CurrentStatementIndex = newIndex
		}

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
		block.StatementType = "BLOCK_STATEMENT"
		ifStatement.ThenBlock = block

		if newIndex+1 < len(ast.TokenList) && ast.TokenList[newIndex+1].Type == "KEYWORD" && ast.TokenList[newIndex+1].Value == "else" {
			var l []lib.Statement
			newElseIndex, else_block := ast.ParseBlockStatement(ast.TokenList, l, newIndex+2)
			else_block.StatementType = "BLOCK_STATEMENT"
			ifStatement.ElseBlock = else_block
			newIndex = newElseIndex
		}

		newSt := lib.Statement{StatementType: "IF_STATEMENT", Value: ifStatement}
		return newIndex, newSt
	case "while":
		index, whileStmt := nodes.HandleWhileStatementCondition(ast.TokenList, i)
		var l []lib.Statement
		newIndex, block := ast.ParseBlockStatement(ast.TokenList, l, index)
		block.StatementType = "BLOCK_STATEMENT"
		whileStmt.Body = block
		newSt := lib.Statement{StatementType: "WHILE_STATEMENT", Value: whileStmt}
		return newIndex, newSt
	case "function":
		index, funcStmt := nodes.HandleFunctionDeclarationStatement(ast.TokenList, i, true)
		var l []lib.Statement
		newIndex, block := ast.ParseBlockStatement(ast.TokenList, l, index)
		funcStmt.Body = block
		newSt := lib.Statement{StatementType: "FUNCTION_DECLARATION", Value: funcStmt}
		return newIndex, newSt
	case "return":
		index, newStatement := nodes.HandleReturnStatement(ast.TokenList, i)
		return index, newStatement
	case "struct":
		std := lib.StructDeclarationStatement{
			Name: ast.TokenList[i+1].Value,
		}
		var l []lib.Statement
		newIndex, block := ast.ParseStructBlockStatement(ast.TokenList, l, i+2)
		std.Fields = block
		newSt := lib.Statement{StatementType: "STRUCT_DECLARATION", Value: std}
		return newIndex, newSt
	default:
		// fmt.Printf("Unhandled keyword: %s\n", keyword)
		index := i + 1
		return index, lib.Statement{StatementType: "Unhandled"}
	}

}

func (ast *ASTBuilder) ParseStructBlockStatement(tokenList []lib.Token, stmtList []lib.Statement, index int) (int, lib.Statement) {

	if tokenList[index].Type == "LEFT_BRACE" {
		i := index + 1
		var l []lib.Statement
		i, stmt := ast.ParseStructBlockStatement(tokenList, l, i)
		return i, stmt
	} else if tokenList[index].Type == "RIGHT_BRACE" {

		newBlock := lib.StatementBlock{
			Type:       "StatementBlock",
			Statements: stmtList,
		}

		newStmt := lib.Statement{
			StatementType: "StatementBlock",
			Value:         newBlock,
		}

		return index, newStmt

	} else {
		identifier := tokenList[index].Value
		typeChecker := tokenList[index+1].Value
		structField := lib.StructField{}

		if typeChecker == "function" {
			i, funcStmt := nodes.HandleFunctionDeclarationStatement(tokenList, index, false)
			var l []lib.Statement
			newIndex, block := ast.ParseBlockStatement(ast.TokenList, l, i)
			funcStmt.Body = block

			newFuncSt := lib.Statement{StatementType: "FUNCTION_DECLARATION", Value: funcStmt}

			structField.Name = identifier
			structField.DataType = "function"
			structField.Body = newFuncSt

			newSt := lib.Statement{StatementType: "STRUCT_FIELD", Value: structField}
			stmtList = append(stmtList, newSt)
			i, stmt := ast.ParseStructBlockStatement(tokenList, stmtList, newIndex+1)
			return i, stmt
		} else {
			structField.Name = identifier
			structField.DataType = tokenList[index+1].Value
			structField.Body = lib.Statement{StatementType: "STRUCT_FIELD", Value: nil}
			newSt := lib.Statement{StatementType: "STRUCT_FIELD", Value: structField}
			stmtList = append(stmtList, newSt)
			i, stmt := ast.ParseStructBlockStatement(tokenList, stmtList, index+3)
			return i, stmt
		}

	}

}

func (ast *ASTBuilder) ParseBlockStatement(tokenList []lib.Token, stmtList []lib.Statement, index int) (int, lib.Statement) {

	for index < len(tokenList) {
		token := tokenList[index]
		switch token.Type {
		case "LEFT_BRACE":
			// Recursively parse the nested block
			newIndex, block := ast.ParseBlockStatement(tokenList, []lib.Statement{}, index+1)
			stmtList = append(stmtList, block)
			index = newIndex
			continue
		case "RIGHT_BRACE":
			// Return the current block when encountering a closing brace
			newBlock := lib.StatementBlock{
				Type:       "StatementBlock",
				Statements: stmtList,
			}

			// fmt.Println("new stmt : ", newBlock)

			return index, lib.Statement{
				StatementType: "BLOCK_STATEMENT",
				Value:         newBlock,
			}
		case "KEYWORD":
			// Handle various keywords
			newIndex, stmt := ast.handleKeyword(token.Value, index)
			// fmt.Println("new stmt : KEYWORD", stmt)
			if stmt.StatementType != "Unhandled" {
				stmtList = append(stmtList, stmt)
			}
			index = newIndex + 1
			continue
		default:
			// Handle other statements
			// newIndex, stmt := ast.parseStatement(tokenList, index)
			// if stmt.StatementType != "Unhandled" {
			// 	stmtList = append(stmtList, stmt)
			// }
			// index++
		}
		index++
	}

	// Handle case where there is no matching RIGHT_BRACE
	newBlock := lib.StatementBlock{
		Type:       "StatementBlock",
		Statements: stmtList,
	}
	return index, lib.Statement{
		StatementType: "BLOCK_STATEMENT",
		Value:         newBlock,
	}

}

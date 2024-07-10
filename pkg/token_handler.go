package pkg

import (
	"fmt"
)

type Token struct {
	Type  string
	Value string
}

type TokenDriver struct {
	TokenList   []Token
	KeyWordList [30]string
	Operators   map[string]int
	Numbers     map[byte]int
}

func (t *TokenDriver) Init() {
	t.KeyWordList = [30]string{
		"let",
		"int",
		"string",
		"boolean",
		"start",
		"end",
		"real",
		"print",
		"function",
		"export",
		"import",
		"if",
		"else",
	}

	t.Operators = make(map[string]int)

	t.Operators["EQUAL_OPERATOR"] = 1
	t.Operators["NOT_EQUALS_OPERATOR"] = 1
	t.Operators["PLUS_OPERATOR"] = 1
	t.Operators["MINUS_OPERATOR"] = 1
	t.Operators["MULTIPLY_OPERATOR"] = 1
	t.Operators["DIVIDE_OPERATOR"] = 1
	t.Operators["DOT_OPERATOR"] = 1

	t.Numbers = make(map[byte]int)

	t.Numbers['0'] = 1
	t.Numbers['1'] = 1
	t.Numbers['2'] = 1
	t.Numbers['3'] = 1
	t.Numbers['4'] = 1
	t.Numbers['5'] = 1
	t.Numbers['6'] = 1
	t.Numbers['7'] = 1
	t.Numbers['8'] = 1
	t.Numbers['9'] = 1

}

func (t *TokenDriver) ParseTokens(lexemeList []Lexeme) {

	// var prev_lexemes []Lexeme
	i := 0
	str := ""

	for i < len(lexemeList) {
		lex := lexemeList[i]
		str += lex.Value

		if lex.LexType == "NEWLINE" ||
			lex.LexType == "WHITESPACE" {
			isKey := checkAndParseKeyword(str, t)

			if isKey {
				str = ""
				i += 1
				continue
			}

			isID := checkAndParseIdentifier(str, t)

			if isID {
				str = ""
				i += 1
				continue
			}

		} else if IsDigit(lex.Value) {
			isNumeric, out_index := checkAndParseNumericLiteral(lexemeList, i, t)

			if isNumeric {
				str = ""
				i = out_index
				continue
			}

		} else if lex.LexType == "QOUTE" {
			isString, index := CheckAndParseStringLiteral(lexemeList, i, t)

			if isString {
				str = ""
				i = index + 1
				continue
			}

		} else if str == "true" || str == "false" {
			new_token := Token{
				Type:  "BOOLEAN_LITERAL",
				Value: str,
			}
			t.TokenList = append(t.TokenList, new_token)
			str = ""
			i += 1
			continue
		} else if lex.LexType == "RIGHT_BRACE" ||
			lex.LexType == "LEFT_BRACE" ||
			lex.LexType == "COMMA" ||
			lex.LexType == "RIGHT_PARANTHESES" ||
			lex.LexType == "LEFT_PARANTHESES" ||
			lex.LexType == "PIPE" ||
			lex.LexType == "WHITESPACE" ||
			t.Operators[lex.LexType] != 0 ||
			lex.LexType == "SEMICOLON" {

			if CheckIsIdentifier(str) {
				checkAndParseIdentifier(str, t)
			}

			new_token := Token{
				Type:  lex.LexType,
				Value: lex.Value,
			}
			t.TokenList = append(t.TokenList, new_token)
			str = ""
			i += 1
			continue
		}

		fmt.Println("Out => " + str)
		i += 1
	}

	removeEmptyTokens(t)

	// for _, t := range lexemeList {
	// 	fmt.Printf("Lexeme Type : %s , Lexeme Value : %s \n", t.LexType, t.Value)
	// }

	for _, t := range t.TokenList {
		fmt.Printf("Token Type : %s , Token Value : %s \n", t.Type, t.Value)
	}

}

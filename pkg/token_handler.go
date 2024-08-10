package pkg

import (
	"fmt"
	"gear-lang/pkg/lib"
)

type TokenDriver struct {
	TokenList   []lib.Token
	KeyWordList [30]string
	Operators   map[string]int
	Numbers     map[byte]int
	Validator   map[byte]int
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
		"while",
		"export",
		"void",
		"return",
		"struct",
	}

	t.Operators = make(map[string]int)

	t.Operators["EQUAL_OPERATOR"] = 1
	t.Operators["NOT_EQUALS_OPERATOR"] = 1
	t.Operators["PLUS_OPERATOR"] = 1
	t.Operators["MINUS_OPERATOR"] = 1
	t.Operators["MULTIPLY_OPERATOR"] = 1
	t.Operators["DIVIDE_OPERATOR"] = 1
	t.Operators["DOT_OPERATOR"] = 1
	t.Operators["DOT_OPERATOR"] = 1
	t.Operators["PIPE_OPERATOR"] = 1
	t.Operators["AND_OPERATOR"] = 1

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

	t.Validator = make(map[byte]int)
	t.Validator['}'] = 1
	t.Validator['{'] = 1
	t.Validator[')'] = 1
	t.Validator['('] = 1
	t.Validator[']'] = 1
	t.Validator['['] = 1
	t.Validator[';'] = 1
	t.Validator[':'] = 1
	t.Validator[','] = 1
	t.Validator['+'] = 1
	t.Validator['-'] = 1
	t.Validator['*'] = 1
	t.Validator['/'] = 1
	t.Validator[' '] = 1
	t.Validator['.'] = 1

}

func (t *TokenDriver) Tokenizer(lexemeList []Lexeme) {

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
			new_token := lib.Token{
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
			lex.LexType == "WHITESPACE" ||
			t.Operators[lex.LexType] != 0 ||
			lex.LexType == "RIGHT_BRACKET" ||
			lex.LexType == "LEFT_BRACKET" ||
			lex.LexType == "SEMICOLON" {

			isKeyword := checkAndParseKeyword(str, t)

			fmt.Println("is left bracket => ", lex.LexType == "LEFT_BRACKET", str, isKeyword)

			if !isKeyword {
				if CheckIsIdentifier(str) {
					fmt.Println("it is a identifier", str, str[0] == ' ')
					checkAndParseIdentifier(str, t)
				}
			}

			if t.Operators[lex.LexType] != 0 {
				index := ParseOperators(str, i, lexemeList, t)
				str = ""
				i = index
				i += 1
				continue
			}

			new_token := lib.Token{
				Type:  lex.LexType,
				Value: lex.Value,
			}
			t.TokenList = append(t.TokenList, new_token)
			str = ""
			i += 1
			continue
		}

		i += 1
	}

	removeEmptyTokens(t)

	for _, t := range t.TokenList {
		fmt.Printf("Token Type : %s , Token Value : %s \n", t.Type, t.Value)
	}

}

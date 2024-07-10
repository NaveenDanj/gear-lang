package pkg

import "regexp"

type Lexeme struct {
	LexType string
	Value   string
}

type LexemeDriver struct {
	LexList []Lexeme
}

func (d *LexemeDriver) CheckLexeme(char string) {
	if char == " " {
		new_lex := Lexeme{
			LexType: "WHITESPACE",
			Value:   " ",
		}
		d.LexList = append(d.LexList, new_lex)
	} else if char == "\n" {
		new_lex := Lexeme{
			LexType: "NEWLINE",
			Value:   "\n",
		}
		d.LexList = append(d.LexList, new_lex)

	} else if char == "'" {

		new_lex := Lexeme{
			LexType: "QOUTE",
			Value:   "'",
		}

		d.LexList = append(d.LexList, new_lex)

	} else if char == "+" {

		new_lex := Lexeme{
			LexType: "PLUS_OPERATOR",
			Value:   "+",
		}

		d.LexList = append(d.LexList, new_lex)

	} else if char == "=" {

		new_lex := Lexeme{
			LexType: "EQUAL_OPERATOR",
			Value:   "=",
		}

		d.LexList = append(d.LexList, new_lex)

	} else if char == "(" {

		new_lex := Lexeme{
			LexType: "LEFT_PARANTHESES",
			Value:   "(",
		}

		d.LexList = append(d.LexList, new_lex)

	} else if char == ")" {

		new_lex := Lexeme{
			LexType: "RIGHT_PARANTHESES",
			Value:   ")",
		}

		d.LexList = append(d.LexList, new_lex)

	} else if char == "[" {

		new_lex := Lexeme{
			LexType: "LEFT_BRACKET",
			Value:   "[",
		}

		d.LexList = append(d.LexList, new_lex)

	} else if char == "]" {

		new_lex := Lexeme{
			LexType: "RIGHT_BRACKET",
			Value:   "]",
		}

		d.LexList = append(d.LexList, new_lex)

	} else if char == ";" {

		new_lex := Lexeme{
			LexType: "SEMICOLON",
			Value:   ";",
		}

		d.LexList = append(d.LexList, new_lex)

	} else if char == "}" {

		new_lex := Lexeme{
			LexType: "RIGHT_BRACE",
			Value:   "}",
		}

		d.LexList = append(d.LexList, new_lex)

	} else if char == "{" {

		new_lex := Lexeme{
			LexType: "LEFT_BRACE",
			Value:   "{",
		}

		d.LexList = append(d.LexList, new_lex)

	} else if char == "\\" {

		new_lex := Lexeme{
			LexType: "BACKSLASH",
			Value:   "\\",
		}

		d.LexList = append(d.LexList, new_lex)

	} else if char == "," {

		new_lex := Lexeme{
			LexType: "COMMA",
			Value:   ",",
		}

		d.LexList = append(d.LexList, new_lex)

	} else if char == "." {

		new_lex := Lexeme{
			LexType: "DOT",
			Value:   ".",
		}

		d.LexList = append(d.LexList, new_lex)

	} else if char == "|" {

		new_lex := Lexeme{
			LexType: "PIPE_OPERATOR",
			Value:   "|",
		}

		d.LexList = append(d.LexList, new_lex)

	} else if char == "&" {

		new_lex := Lexeme{
			LexType: "AND_OPERATOR",
			Value:   "&",
		}

		d.LexList = append(d.LexList, new_lex)

	} else if IsLetter(char) {
		new_lex := Lexeme{
			LexType: "LETTER",
			Value:   char,
		}
		d.LexList = append(d.LexList, new_lex)

	} else if IsDigit(char) {
		new_lex := Lexeme{
			LexType: "LETTER",
			Value:   char,
		}
		d.LexList = append(d.LexList, new_lex)
	}

}

func isLetter(char string) bool {
	pattern := "[a-zA-Z]"
	matched, err := regexp.MatchString(pattern, char)
	if err != nil {
		return false
	}
	return matched
}

package pkg

import (
	"regexp"
)

func CheckPrevLexemesKeyword(str string, keywordList [30]string) bool {
	for _, keyword := range keywordList {
		if str == keyword {
			return true
		}
	}
	return false
}

func CheckStringLiteral(i int, prev_lexemes []Lexeme, lexemeList []Lexeme) (bool, string, string, int) {
	str := ""

	for _, lex := range prev_lexemes {
		str += lex.Value
	}

	j := 0

	for j = i + 1; j < len(lexemeList); j++ {
		next_lex := lexemeList[j]
		if next_lex.LexType == "QOUTE" {
			break
		} else {
			str += next_lex.Value
		}
	}

	return true, "STRING_LITERAL", str, j + 1

}

func CheckBooleanLiteral(prev_lexemes []Lexeme) (bool, string, string) {
	str := ""

	for _, lex := range prev_lexemes {
		str += lex.Value
	}

	if str == "true" || str == "false" {
		return true, "BOOLEAN_LITERAL", str
	}

	return false, "", ""

}

func IsDigit(char string) bool {
	pattern := "[0-9]"
	matched, err := regexp.MatchString(pattern, char)
	if err != nil {
		return false
	}
	return matched
}

func IsLetter(char string) bool {
	pattern := "[a-zA-Z]"
	matched, err := regexp.MatchString(pattern, char)
	if err != nil {
		return false
	}
	return matched
}

func CheckIsIdentifier(str string) bool {

	valid := make(map[byte]int)
	valid['}'] = 1
	valid[')'] = 1
	valid[']'] = 1
	valid[';'] = 1
	valid[','] = 1
	valid['+'] = 1
	valid['-'] = 1
	valid['*'] = 1
	valid['/'] = 1
	valid[' '] = 1
	valid[':'] = 1

	if len(str) == 0 {
		return false
	}

	if valid[str[len(str)-1]] == 0 {
		return false
	}

	return true

}

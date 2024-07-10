package pkg

import (
	"fmt"
	"strings"
)

func CheckAndParseStringLiteral(lexemeList []Lexeme, index int, t *TokenDriver) (bool, int) {

	if lexemeList[index].LexType != "QOUTE" {
		return false, index
	}

	i := index + 1
	lastIndex := index

	for i < len(lexemeList) && lexemeList[i].LexType != "QOUTE" {
		if lexemeList[i].LexType == "BACKSLASH" {
			lastIndex = i
			i++
			continue
		}

		lastIndex = i
		i++

	}

	str := ""

	for i := index + 1; i <= lastIndex; i++ {
		if lexemeList[i].LexType == "BACKSLASH" {
			i++
			continue
		}
		str += lexemeList[i].Value
	}

	new_token := Token{
		Type:  "STRING_LITERAL",
		Value: str,
	}

	t.TokenList = append(t.TokenList, new_token)

	return true, lastIndex + 1

}

func checkAndParseNumericLiteral(lexemeList []Lexeme, index int, t *TokenDriver) (bool, int) {

	str := ""

	if index != 0 && (!IsDigit(lexemeList[index-1].Value) && (t.Operators[lexemeList[index-1].LexType] == 0 && lexemeList[index-1].LexType != "WHITESPACE" && lexemeList[index-1].LexType != "LEFT_PARANTHESES")) {
		return false, index
	}

	if !IsDigit(lexemeList[index].Value) {
		return false, index
	} else {
		str += lexemeList[index].Value
	}

	i := index + 1
	for i < len(lexemeList) && (IsDigit(lexemeList[i].Value) || lexemeList[i].Value == ".") {
		str += lexemeList[i].Value
		i++

	}

	if IsDigit(str) {

		new_token := Token{
			Type:  "NUMERIC_LITERAL",
			Value: str,
		}

		t.TokenList = append(t.TokenList, new_token)
		return true, i
	}

	return false, index

}

func checkAndParseKeyword(str string, t *TokenDriver) bool {
	str = strings.ReplaceAll(str, " ", "")
	str = strings.ReplaceAll(str, "\n", "")
	isKeyword := CheckPrevLexemesKeyword(str, t.KeyWordList)
	if isKeyword {
		new_token := Token{
			Type:  "KEYWORD",
			Value: str,
		}
		t.TokenList = append(t.TokenList, new_token)
		return true
	}

	return false

}

func checkAndParseIdentifier(str string, t *TokenDriver) bool {
	fmt.Println("Possobile identifier : " + str)

	if len(str) == 0 {
		return false
	}

	if str[0] == ' ' || t.Numbers[str[0]] != 0 {
		return false
	}

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

	if valid[str[len(str)-1]] == 0 {
		return false
	} else {
		// remove the last character
		str = str[:len(str)-1]
	}

	new_token := Token{
		Type:  "IDENTIFIER",
		Value: str,
	}
	t.TokenList = append(t.TokenList, new_token)
	return true
}

func removeEmptyTokens(t *TokenDriver) {
	var newList []Token
	for _, token := range t.TokenList {
		if token.Value != "" {
			newList = append(newList, token)
		}
	}
	t.TokenList = newList
}

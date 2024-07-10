package pkg

import "strings"

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

	if len(str) != 0 && t.Validator[str[len(str)-1]] != 0 {
		str = str[:len(str)-1]
	}

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

	if len(str) == 0 {
		return false
	}

	if str[0] == ' ' || t.Numbers[str[0]] != 0 {
		return false
	}

	if t.Validator[str[len(str)-1]] == 0 {
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

func ParseOperators(str string, index int, lexemeList []Lexeme, t *TokenDriver) int {

	ops := make(map[string]string)
	ops["=="] = "DOUBLE_EQUALS_OPERATOR"
	ops["!="] = "NOT_EQUALS_OPERATOR"
	ops["<="] = "LESS_THAN_OR_EQUALS_OPERATOR"
	ops[">="] = "GREATER_THAN_OR_EQUALS_OPERATOR"
	ops["&&"] = "AND_OPERATOR"
	ops["&"] = "REFERECE_OPERATOR"
	ops["||"] = "OR_OPERATOR"
	ops["="] = "EQUAL_OPERATOR"
	ops["+"] = "PLUS_OPERATOR"
	ops["-"] = "MINUS_OPERATOR"
	ops["*"] = "MULTIPLY_OPERATOR"
	ops["/"] = "DIVIDE_OPERATOR"

	if t.Operators[lexemeList[index+1].LexType] != 0 {
		str += lexemeList[index+1].Value
		index += 1
	}

	new_token := Token{
		Type:  ops[str],
		Value: str,
	}
	t.TokenList = append(t.TokenList, new_token)

	return index
}

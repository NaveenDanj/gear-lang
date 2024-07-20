package main

import (
	"gear-lang/pkg"
	"log"
	"os"
)

func main() {
	content, err := os.ReadFile("test/test3.ger")
	if err != nil {
		log.Fatal("Failed to parse the file: File not found!")
		return
	}

	content_string := string(content)

	var lexList []pkg.Lexeme
	driver := pkg.LexemeDriver{
		LexList: lexList,
	}

	for i := 0; i < len(content_string); i++ {
		driver.CheckLexeme(string(content_string[i]))
	}

	tokenDriver := pkg.TokenDriver{}
	astBuilder := pkg.ASTBuilder{}
	tokenDriver.Init()
	tokenDriver.Tokenizer(driver.LexList)
	astBuilder.TokenList = tokenDriver.TokenList
	astBuilder.Parse(0)
}

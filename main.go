package main

import (
	"gear-lang/pkg"
	"log"
	"os"
)

func main() {
	content, err := os.ReadFile("test/test1.ger")
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

	// do the token parsing
	tokenDriver := pkg.TokenDriver{}
	tokenDriver.Init()
	tokenDriver.ParseTokens(driver.LexList)

}

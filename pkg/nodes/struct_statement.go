package nodes

import (
	"gear-lang/pkg/lib"
)

func HandleStructDeclarationStatement(tokenList []lib.Token, index int) (int, lib.StructDeclarationStatement) {
	structName := tokenList[index+1].Value
	return index, lib.StructDeclarationStatement{Name: structName}
}

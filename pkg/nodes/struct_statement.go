package nodes

import "gear-lang/pkg/lib"

func HandleStructDeclarationStatement(tokenList []lib.Token, index int) (int, lib.StructDeclarationStatement) {
	// TODO: handle struct declaration
	structName := tokenList[index+1].Value

}

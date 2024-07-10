package nodes

import "gear-lang/pkg/lib"

func HandleImportStatement(tokenList []lib.Token, index int) (int, lib.Statement) {

	st := lib.ImportStatement{
		ImportPath: tokenList[index+1].Value,
	}

	newStatement := lib.Statement{
		StatementType: "IMPORT",
		Value:         st,
	}

	return index + 2, newStatement

}

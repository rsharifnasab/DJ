package ts

import (
	"context"
	"fmt"

	sitter "github.com/smacker/go-tree-sitter"
	"github.com/smacker/go-tree-sitter/javascript"
)

func Tree() {
	parser := sitter.NewParser()
	parser.SetLanguage(javascript.GetLanguage())
	sourceCode := []byte("let a = 1; let b = [1,2,3,4]; b[0]++;")
	tree, err := parser.ParseCtx(context.Background(), nil, sourceCode)
	if err != nil {
		panic(err)
	}

	fmt.Println(tree)
	n := tree.RootNode()

	fmt.Println(n) // (program (lexical_declaration (variable_declarator (identifier) (number))))

	child := n.NamedChild(0)
	fmt.Println(child.Type())      // lexical_declaration
	fmt.Println(child.StartByte()) // 0
	fmt.Println(child.EndByte())   // 9
}

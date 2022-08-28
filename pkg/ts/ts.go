package ts

import (
	"context"
	"fmt"
	"os"

	sitter "github.com/smacker/go-tree-sitter"
	"github.com/smacker/go-tree-sitter/c"
	"github.com/spf13/cobra"
	"gopkg.in/yaml.v3"
)

type Language string
type TSConfig map[Language]map[string][]string

func loadConfig() *TSConfig {
	data, err := os.ReadFile("./examples/question-add/ts.yaml")
	if err != nil {
		panic(err.Error())
	}

	conf := make(TSConfig)
	err = yaml.Unmarshal(data, &conf)
	if err != nil {
		panic(err.Error())
	}

	fmt.Printf("config : %+v\n", conf)

	return &conf
}

func Tree() {
	conf := loadConfig()
	_ = conf
	parser := sitter.NewParser()
	parser.SetLanguage(c.GetLanguage())
	sourcePath := "./examples/sol-add-c/main.c"
	sourceCode, err := os.ReadFile(sourcePath)
	cobra.CheckErr(err)
	tree, err := parser.ParseCtx(context.Background(), nil, sourceCode)
	if err != nil {
		panic(err.Error())
	}
	fmt.Printf("tree : %+v\n", tree)
	for i := 0; true; i++ {
		node := tree.RootNode().Child(i)
		if node == nil {
			break
		}
		fmt.Printf("-------node[%d]-----\n%v\n", i, node)
	}
	Query()
}

func Query() {
	parser := sitter.NewParser()
	parser.SetLanguage(c.GetLanguage())
	sourcePath := "./examples/sol-add-c/main.c"
	sourceCode, err := os.ReadFile(sourcePath)
	cobra.CheckErr(err)
	tree, err := parser.ParseCtx(context.Background(), nil, sourceCode)
	if err != nil {
		panic(err.Error())
	}
	query, err := sitter.NewQuery([]byte("(array_declarator)"), c.GetLanguage())
	if err != nil {
		panic(err.Error())
	}
	qc := sitter.NewQueryCursor()
	qc.Exec(query, tree.RootNode())

	var funcs []*sitter.Node
	for {
		m, ok := qc.NextMatch()
		if !ok {
			break
		}

		for _, c := range m.Captures {
			funcs = append(funcs, c.Node)
			fmt.Println(c.Node)
		}
	}
	_ = funcs
	query.Close()
}

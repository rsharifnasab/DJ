package ts

import (
	"context"
	"fmt"
	"os"

	"github.com/rsharifnasab/DJ/pkg/util"
	sitter "github.com/smacker/go-tree-sitter"
	"github.com/smacker/go-tree-sitter/bash"
	"github.com/smacker/go-tree-sitter/c"
	"github.com/smacker/go-tree-sitter/cpp"
	"github.com/smacker/go-tree-sitter/csharp"
	"github.com/smacker/go-tree-sitter/css"
	"github.com/smacker/go-tree-sitter/elm"
	"github.com/smacker/go-tree-sitter/golang"
	"github.com/smacker/go-tree-sitter/html"
	"github.com/smacker/go-tree-sitter/java"
	"github.com/smacker/go-tree-sitter/javascript"
	"github.com/smacker/go-tree-sitter/ocaml"
	"github.com/smacker/go-tree-sitter/php"
	"github.com/smacker/go-tree-sitter/python"
	"github.com/smacker/go-tree-sitter/ruby"
	"github.com/smacker/go-tree-sitter/rust"
	"github.com/smacker/go-tree-sitter/scala"
	"github.com/smacker/go-tree-sitter/typescript/typescript"
	"gopkg.in/yaml.v3"
)

var tsLanguages map[string]*sitter.Language

func init() {
	tsLanguages = make(map[string]*sitter.Language)

	tsLanguages["bash"] = bash.GetLanguage()
	tsLanguages["c"] = c.GetLanguage()
	tsLanguages["cpp"] = cpp.GetLanguage()
	tsLanguages["csharp"] = csharp.GetLanguage()
	tsLanguages["css"] = css.GetLanguage()
	tsLanguages["elm"] = elm.GetLanguage()
	tsLanguages["golang"] = golang.GetLanguage()
	tsLanguages["html"] = html.GetLanguage()
	tsLanguages["java"] = java.GetLanguage()
	tsLanguages["javascript"] = javascript.GetLanguage()
	tsLanguages["ocaml"] = ocaml.GetLanguage()
	tsLanguages["php"] = php.GetLanguage()
	tsLanguages["python"] = python.GetLanguage()
	tsLanguages["ruby"] = ruby.GetLanguage()
	tsLanguages["rust"] = rust.GetLanguage()
	tsLanguages["scala"] = scala.GetLanguage()
	tsLanguages["typescript"] = typescript.GetLanguage()
}

type languageConfig map[string]([]string)
type tsConfig map[string]languageConfig

func loadConfig(questionDir string) (tsConfig, error) {
	data, err := os.ReadFile(questionDir + "/ts.yaml")
	if err != nil {
		return nil, err
	}

	conf := make(tsConfig)
	err = yaml.Unmarshal(data, &conf)
	if err != nil {
		return nil, err
	}

	fmt.Printf("config : %+v\n", conf)
	return conf, nil
}

/*
func executeQuery(queryStr string, tree *sitter.Tree, lang *sitter.Language) []*sitter.Node {
	query, err := sitter.NewQuery([]byte(queryStr), lang)
	if err != nil {
		panic(err.Error())
	}
	qc := sitter.NewQueryCursor()
	qc.Exec(query, tree.RootNode())

	var results []*sitter.Node
	for {
		m, ok := qc.NextMatch()
		if !ok {
			break
		}

		for _, c := range m.Captures {
			results = append(results, c.Node)
			fmt.Println(c.Node)
		}
	}
	return results
}
*/

func AllNodes(tree *sitter.Node) []*sitter.Node {
	res := make([]*sitter.Node, 0)
	res = append(res, tree)

	for i := 0; i < int(tree.ChildCount()); i++ {
		res = append(res, AllNodes(tree.Child(i))...)
	}

	return res
}

func getAllImports(tree *sitter.Tree, source []byte) []string {
	res := make([]string, 0)

	allNodes := AllNodes(tree.RootNode())
	for _, node := range allNodes {
		fmt.Printf("%d : (type %v) %v\n\n", 0, node.Type(), node)
		switch node.Type() {
		case "preproc_include": // c, cpp
			start := node.ChildByFieldName("path").StartByte()
			end := node.ChildByFieldName("path").EndByte()
			res = append(res, string(source[start+1:end-1]))
		case "scoped_identifier", "scoped_type_identifier": // java
			fmt.Println("import dcl")
			start := node.StartByte()
			end := node.EndByte()
			res = append(res, string(source[start:end]))
		}

	}

	return res
}

func getAllStructures(tree *sitter.Tree, source []byte) []string {
	res := make([]string, 0)
	return res
}

func (config tsConfig) oneSource(sourceFile string, lang string) error {
	parser := sitter.NewParser()
	parser.SetLanguage(tsLanguages[lang])

	sourceCode, err := os.ReadFile(sourceFile)
	if err != nil {
		return err
	}
	tree, err := parser.ParseCtx(context.Background(), nil, sourceCode)
	if err != nil {
		panic(err.Error())
	}
	//fmt.Printf("tree : %+v\n", tree)

	actualImports := getAllImports(tree, sourceCode)
	//fmt.Printf("actualImports: %v\n", actualImports)
	actualStructures := getAllStructures(tree, sourceCode)

	rules := config[lang]
	whiteImports, prs := rules["import-white"]
	if prs {
	actual:
		for _, actualImport := range actualImports {
			// searching for actualImport in wImport
			for _, wImport := range whiteImports {
				if wImport == actualImport {
					continue actual
				}
			}
			return fmt.Errorf("importing [%s] isn't allowed because it is not in the white list", actualImport)

		}

	}

	blackImports, prs := rules["import-black"]
	if prs {
		for _, bImport := range blackImports {
			for _, actualImport := range actualImports {
				if actualImport == bImport {
					return fmt.Errorf("importing %s isn't allowed because it is in the black list", actualImport)
				}
			}
		}
	}

	blackStructures, prs := rules["structures"]
	if prs {
		for _, bStructure := range blackStructures {
			for _, actualStructure := range actualStructures {
				if actualStructure == bStructure {
					return fmt.Errorf("using %s isn't allowed because it is in the black list", actualStructure)
				}
			}
		}
	}

	return nil
}

func CheckSource(questionDir, submisionDir, lang string) error {
	conf, err := loadConfig(questionDir)
	if err != nil {
		if os.IsNotExist(err) { // no ts.yaml
			return nil
		}
		panic(err)
	}

	if lang == "" {
		lang, err = util.AutoDetectLanguage(submisionDir)
		if err != nil {
			return err
		}
	}

	sources, err := util.FilterSrcsByLang(submisionDir, lang)
	if err != nil {
		return err
	}
	fmt.Printf("sources : %v\n", sources)
	for _, sourceFile := range sources {
		err := conf.oneSource(sourceFile, lang)
		if err != nil {
			return fmt.Errorf("%s: %s", sourceFile, err.Error())
		}
	}
	return nil
}

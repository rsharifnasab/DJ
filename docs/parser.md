# parser
we need a general purpose parser which parse most programming languages. 
and preferably written in go


1. A lexer written in go (but not parser)
   + ease of use
   + no documentation about lexing (it is a syntax highlighter!)
   + pure go
   + link [+](https://github.com/alecthomas/chroma)
   + low level which isn't enough for our task

Sample code:
```go
func main() {
	theFile := `Main.java`
	lexer := lexers.Match(theFile)
	if lexer == nil {
		log.Fatal("cannot find any suitable parser")
	}
	contents, err := ioutil.ReadFile(theFile)
	if err != nil {
		panic(err)
	}

	iterator, err := lexer.Tokenise(nil, string(contents))
	if err != nil {
		panic(err)
	}
	for _, token := range iterator.Tokens() {
		if token.Type.Category() == chroma.EOFType {
			break
		} else if token.Type.InCategory(chroma.Text) {
			continue
		}
		fmt.Printf("--------\nv: %s \nt:  %v \n", token.GoString(), token)
	}
}

```


2. Treesitter
   + client should have .so files for each language
   + doesn't have go binding (at least by default) but python seems stable
   + can parse practically all languages
   + provide good information.

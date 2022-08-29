package util

import (
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestPLMap1(t *testing.T) {
	lang, err := ExtensionToLanguge("c")
	assert.NoError(t, err)
	assert.Equal(t, "c", lang)
}

func TestPLMap2(t *testing.T) {
	lang, err := ExtensionToLanguge("ex")
	assert.NoError(t, err)
	assert.Equal(t, "elixir", lang)
}

func TestAutoLanguage(t *testing.T) {
	lang, err := AutoDetectLanguage(".")
	assert.NoError(t, err)
	assert.Equal(t, "go", lang)
}

func TestFilterSrcs(t *testing.T) {
	tmp := MakeTempfolder()
	defer os.RemoveAll(tmp)

	os.WriteFile(tmp+"/a.go", []byte("package main"), 0777)
	os.WriteFile(tmp+"/b.go", []byte("package main"), 0777)
	os.WriteFile(tmp+"/a.c", []byte("int main(){return 0}"), 0777)
	srcs, err := FilterSrcsByLang(tmp, "go")
	assert.NoError(t, err)
	assert.Len(t, srcs, 2)

	srcs, err = FilterSrcsByLang(tmp, "c")
	assert.NoError(t, err)
	assert.Len(t, srcs, 1)
}

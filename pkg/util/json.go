package util

import (
	"encoding/json"
	"fmt"
)

func PrintStruct(structt interface{}) {
	// this is a good option too:
	// https://github.com/k0kubun/pp
	if jsonM, err := json.MarshalIndent(structt, " ", "\t"); err != nil {
		panic(err)
	} else {
		fmt.Println(string(jsonM))
	}
}

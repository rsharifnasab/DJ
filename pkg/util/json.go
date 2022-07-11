package util

import (
	"encoding/json"
	"fmt"

	"github.com/spf13/cobra"
)

func PrintStruct(object interface{}) {
	jsonM, err := json.MarshalIndent(object, " ", "\t")
	cobra.CheckErr(err)
	fmt.Println(string(jsonM))
	// alternative:
	// https://github.com/k0kubun/pp
}

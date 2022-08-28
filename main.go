/*
Copyright © 2022 Roozbeh Sharifnasab rsharifnasab@gmail.com
*/
package main

import (
	"math/rand"
	"time"

	"github.com/rsharifnasab/DJ/cmd"
)

func main() {
	//viper.Set("debug", true)
	rand.Seed(time.Now().UnixNano())
	//ts.Tree()
	cmd.Execute()
}

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
	rand.Seed(time.Now().UnixNano())
	cmd.Execute()
}

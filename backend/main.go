package main

import (
	"fmt"

	"github.com/hiruthikj/tictactoe-go-wasm/game"
)

func main() {
	fmt.Println("Welcome to tictactoe game!")

	game := game.NewGame()
	bytes, _ := game.Serialize()

	fmt.Println(string(bytes))
}

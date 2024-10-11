package main

import (
	"fmt"
	datareader "game/DataReader"
	game "game/Game"
)

func main() {
	fmt.Println("Hello world!")
	game.Game()
	datareader.Reader("test")
}

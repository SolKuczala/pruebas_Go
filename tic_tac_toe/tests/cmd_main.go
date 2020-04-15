package main

import (
	"fmt"

	T "../project/tictactoe"
)

var gameStruct T.Game

func main() {
	gameStruct = T.NewGame(3)
	plays := []string{
		"o", "x", "x",
		"o", "x", "x",
		"o", "x", "x",
	}
	for i, play := range plays {
		coor := T.Coord{X: uint(i % 3), Y: uint(i / 3)}
		_, err := T.Play(play, coor, &gameStruct)
		printIfError(err)
		T.PrintBoard(&gameStruct)
	}
	T.PrintBoard(&gameStruct)
}

//funcion que haga la siguiente linea:
func printIfError(err error) {
	if err != nil {
		fmt.Println(err.Error())
	}
}

//go:build !js || !wasm

package main

import (
	"fmt"
)

func main() {
	fmt.Println("GAME START")
	board := initNewBoard()
	board.Print()
	for {
		var input string
		fmt.Scan(&input)

		switch input {
		case "w":
			board.Move(Up)
		case "a":
			board.Move(Left)
		case "s":
			board.Move(Down)
		case "d":
			board.Move(Right)
		}
		board.Print()
	}
}

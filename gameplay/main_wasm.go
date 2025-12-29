//go:build js && wasm

package main

import (
	"encoding/json"
	"fmt"
	"syscall/js"
)

var currentBoard GameBoard

func main() {
	c := make(chan struct{}, 0)
	fmt.Println("WASM 2048 Initialized")
	registerCallbacks()
	<-c
}

func registerCallbacks() {
	js.Global().Set("newGame", js.FuncOf(newGame))
	js.Global().Set("move", js.FuncOf(moveDir))
	js.Global().Set("getState", js.FuncOf(getState))
}

func newGame(this js.Value, args []js.Value) interface{} {
	currentBoard = initNewBoard()
	return getBoardState()
}

func moveDir(this js.Value, args []js.Value) interface{} {
	if len(args) < 1 {
		return nil
	}

	if currentBoard.GameOver {
		return getBoardState()
	}

	dirStr := args[0].String()
	var dir Direction
	switch dirStr {
	case "up":
		dir = Up
	case "down":
		dir = Down
	case "left":
		dir = Left
	case "right":
		dir = Right
	default:
		return getBoardState()
	}

	currentBoard.Move(dir)
	return getBoardState()
}

func getState(this js.Value, args []js.Value) interface{} {
	return getBoardState()
}

func getBoardState() string {
	b, err := json.Marshal(currentBoard)
	if err != nil {
		fmt.Println("Error marshaling board:", err)
		return "{}"
	}
	return string(b)
}

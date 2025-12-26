package main

import (
	"fmt"
	"slices"
)

const Size = 4

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

func initNewBoard() GameBoard {
	newBoard := GameBoard{
		Score:    0,
		Tiles:    make([]int, Size*Size),
		Player:   "An",
		GameOver: false,
		Size:     Size,
	}
	randomTiles := []RandomTile{{Value: 1, Weight: 90}, {Value: 2, Weight: 10}}
	for i := 0; i < 2; i++ {
		newBoard.SpawnTile(randomTiles)
	}
	return newBoard
}

func move(board *GameBoard, move Direction) []int {
	newTiles := make([]int, len(board.Tiles))
	switch move {
	case Up:
		{
			for i := int8(0); i < board.Size; i++ {
				tileArr := board.GetTilesInCol(i)
				merged := compressAndMerge(tileArr)
				for j, tile := range merged {
					newTiles[i+(int8(j)*board.Size)] = tile
				}
			}
			break
		}
	case Down:
		{
			for i := int8(0); i < board.Size; i++ {
				tileArr := board.GetTilesInCol(i)
				slices.Reverse(tileArr)
				merged := compressAndMerge(tileArr)
				slices.Reverse(merged)
				for j, tile := range merged {
					newTiles[i+(int8(j)*board.Size)] = tile
				}
			}
			break
		}
	case Left:
		{
			var tmpTiles []int
			for i := int8(0); i < board.Size; i++ {
				tileArr := board.GetTilesInRow(i)
				merged := compressAndMerge(tileArr)
				tmpTiles = append(tmpTiles, merged...)
			}
			newTiles = tmpTiles
			break
		}
	case Right:
		{
			var tmpTiles []int
			for i := int8(0); i < board.Size; i++ {
				tileArr := board.GetTilesInRow(i)
				slices.Reverse(tileArr)
				merged := compressAndMerge(tileArr)
				slices.Reverse(merged)
				tmpTiles = append(tmpTiles, merged...)
			}
			newTiles = tmpTiles
			break
		}
	}
	return newTiles
}

func compressAndMerge(tiles []int) []int {
	var result []int
	size := len(tiles)
	compressed := false
	for _, tile := range tiles {
		if tile == 0 {
			continue
		}
		index := len(result) - 1
		if compressed || index < 0 || result[index] != tile {
			result = append(result, tile)
			compressed = false
		} else {
			result[index] += 1
			compressed = true
		}
	}
	for i := len(result); i < size; i++ {
		result = append(result, 0)
	}
	return result
}

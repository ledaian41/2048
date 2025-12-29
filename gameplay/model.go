package main

import (
	"fmt"
	"math/rand/v2"
	"slices"
)

type Direction int

const (
	Left Direction = iota
	Right
	Up
	Down
)

type GameBoard struct {
	Score    int
	Tiles    []int
	Player   string
	GameOver bool
	Size     int8
}

func (board *GameBoard) GetTilesInRow(rowIndex int8) []int {
	row := make([]int, board.Size)
	startIndex := rowIndex * board.Size
	for j := int8(0); j < board.Size; j++ {
		row[j] = board.Tiles[startIndex+j]
	}
	return row
}

func (board *GameBoard) GetTilesInCol(colIndex int8) []int {
	col := make([]int, board.Size)
	for j := int8(0); j < board.Size; j++ {
		col[j] = board.Tiles[colIndex+j*board.Size]
	}
	return col
}

func (board *GameBoard) CalculateScore() {
	score := 0
	for _, tile := range board.Tiles {
		if tile > 0 {
			score += 1 << tile
		}
	}
	board.Score = score
}

func (board *GameBoard) Move(direction Direction) bool {
	newTiles := move(board, direction)
	if slices.Equal(newTiles, board.Tiles) {
		return false
	}
	board.Tiles = newTiles
	board.SpawnTile([]RandomTile{{Value: 1, Weight: 90}, {Value: 2, Weight: 10}})
	board.CalculateScore()
	return true
}

func (board *GameBoard) SpawnTile(randomTiles []RandomTile) {
	var empty []int
	for i, v := range board.Tiles {
		if v == 0 {
			empty = append(empty, i)
		}
	}
	if len(empty) == 0 {
		return
	}
	idx := empty[rand.IntN(len(empty))]
	board.Tiles[idx] = randomNewTile(randomTiles).Value
}

func randomNewTile(tiles []RandomTile) RandomTile {
	random := int8(rand.IntN(100))
	var result RandomTile
	var cursor int8 = 0
	for _, tile := range tiles {
		cursor += tile.Weight
		if random < cursor {
			return tile
		}
	}
	return result
}

func (board *GameBoard) Print() {
	for i := int8(0); i < board.Size; i++ {
		row := board.GetTilesInRow(i)
		fmt.Println(row)
	}
	fmt.Println("---------")
}

type RandomTile struct {
	Value  int
	Weight int8
}

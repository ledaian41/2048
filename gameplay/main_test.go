package main

import (
	"slices"
	"testing"
)

// Helper to create a board with specific tiles for testing
func newTestBoard(tiles []int) *GameBoard {
	if len(tiles) != 16 {
		// Fill with 0 if not enough provided, or truncate
		padded := make([]int, 16)
		copy(padded, tiles)
		tiles = padded
	}
	return &GameBoard{
		Size:  4,
		Tiles: tiles,
		Score: 0,
	}
}

func TestCompressAndMerge(t *testing.T) {
	tests := []struct {
		name     string
		input    []int
		expected []int
	}{
		{"Empty", []int{0, 0, 0, 0}, []int{0, 0, 0, 0}},
		{"Single Tile", []int{2, 0, 0, 0}, []int{2, 0, 0, 0}},
		{"Simple Merge", []int{2, 2, 0, 0}, []int{3, 0, 0, 0}}, // 2^1 + 2^1 = 2^2
		{"No Merge Different", []int{2, 3, 0, 0}, []int{2, 3, 0, 0}},
		{"Merge With Gap", []int{2, 0, 2, 0}, []int{3, 0, 0, 0}},
		{"Multi Merge", []int{2, 2, 2, 2}, []int{3, 3, 0, 0}},
		{"Merge Priority Left", []int{2, 2, 2, 0}, []int{3, 2, 0, 0}},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := compressAndMerge(tt.input)
			if !slices.Equal(got, tt.expected) {
				t.Errorf("compressAndMerge(%v) = %v, want %v", tt.input, got, tt.expected)
			}
		})
	}
}

func TestGetTilesInRow(t *testing.T) {
	// 0 1 2 3
	// 4 5 6 7
	// ...
	tiles := make([]int, 16)
	for i := 0; i < 16; i++ {
		tiles[i] = i
	}
	board := newTestBoard(tiles)

	expectedRow1 := []int{4, 5, 6, 7}
	got := board.GetTilesInRow(1)
	if !slices.Equal(got, expectedRow1) {
		t.Errorf("GetTilesInRow(1) = %v, want %v", got, expectedRow1)
	}
}

func TestGetTilesInCol(t *testing.T) {
	// 0 1 2 3
	// 4 5 6 7
	// 8 9 10 11
	// 12 13 14 15
	tiles := make([]int, 16)
	for i := 0; i < 16; i++ {
		tiles[i] = i
	}
	board := newTestBoard(tiles)

	expectedCol1 := []int{1, 5, 9, 13}
	got := board.GetTilesInCol(1)
	if !slices.Equal(got, expectedCol1) {
		t.Errorf("GetTilesInCol(1) = %v, want %v", got, expectedCol1)
	}
}

func TestGameBoard_Move_Left(t *testing.T) {
	// [1, 1, 0, 0] -> [2, 0, 0, 0]
	initial := make([]int, 16)
	initial[0] = 1
	initial[1] = 1
	board := newTestBoard(initial)

	moved := board.Move(Left)
	if !moved {
		t.Error("Expected move Left to be valid")
	}
	if board.Tiles[0] != 2 {
		t.Errorf("Expected tile at [0] to be 2, got %d", board.Tiles[0])
	}
}

func TestGameBoard_Move_Right(t *testing.T) {
	// [0, 0, 1, 1] -> [0, 0, 0, 2]
	initial := make([]int, 16)
	initial[2] = 1
	initial[3] = 1
	board := newTestBoard(initial)

	moved := board.Move(Right)
	if !moved {
		t.Error("Expected move Right to be valid")
	}
	if board.Tiles[3] != 2 {
		t.Errorf("Expected tile at [3] to be 2, got %d", board.Tiles[3])
	}
	if board.Tiles[2] != 0 {
		t.Errorf("Expected tile at [2] to be 0, got %d", board.Tiles[2])
	}
}

func TestGameBoard_Move_Up(t *testing.T) {
	// [1] ...
	// [1] ...
	// Merge upwards -> [2] at top
	initial := make([]int, 16)
	initial[0] = 1
	initial[4] = 1
	board := newTestBoard(initial)

	moved := board.Move(Up)
	if !moved {
		t.Error("Expected move Up to be valid")
	}
	if board.Tiles[0] != 2 {
		t.Errorf("Expected tile at [0] to be 2, got %d", board.Tiles[0])
	}
}

func TestGameBoard_Move_Down(t *testing.T) {
	// ...
	// [1] (index 12)
	// [1] (index 8) -> wait, index are row major.
	// Row 2: index 8. Row 3: index 12.
	// We want to merge downwards to row 3.
	initial := make([]int, 16)
	initial[8] = 1
	initial[12] = 1
	board := newTestBoard(initial)

	moved := board.Move(Down)
	if !moved {
		t.Error("Expected move Down to be valid")
	}
	if board.Tiles[12] != 2 {
		t.Errorf("Expected tile at [12] to be 2, got %d", board.Tiles[12])
	}
}

func TestSpawnTile_FullBoard(t *testing.T) {
	// Fill board with unique items so no spawn possible
	initial := make([]int, 16)
	for i := 0; i < 16; i++ {
		initial[i] = 1
	}
	board := newTestBoard(initial)

	// Copy tiles to verify they don't change
	snapshot := make([]int, 16)
	copy(snapshot, board.Tiles)

	// Attempt access spawn (Move calls it, but we can call directly to test safe guard)
	// But Move() won't call SpawnTile if no move happened.
	// The function SpawnTile handles empty check.
	randomTiles := []RandomTile{{Value: 1, Weight: 100}}
	board.SpawnTile(randomTiles)

	if !slices.Equal(board.Tiles, snapshot) {
		t.Error("Expected board to remain unchanged when full")
	}
}

func TestIsGameOver(t *testing.T) {
	tests := []struct {
		name     string
		board    GameBoard
		expected bool
	}{
		{
			name: "Board has empty cells (not game over)",
			board: GameBoard{
				Size: 4,
				Tiles: []int{
					2, 4, 2, 4,
					4, 2, 4, 2,
					2, 4, 0, 4, // Empty cell here
					4, 2, 4, 2,
				},
			},
			expected: false,
		},
		{
			name: "Board full, horizontal merge possible (not game over)",
			board: GameBoard{
				Size: 4,
				Tiles: []int{
					2, 2, 4, 8, // Merge possible here (2, 2)
					4, 8, 2, 4,
					2, 4, 8, 2,
					8, 2, 4, 8,
				},
			},
			expected: false,
		},
		{
			name: "Board full, vertical merge possible (not game over)",
			board: GameBoard{
				Size: 4,
				Tiles: []int{
					2, 4, 2, 4,
					2, 8, 4, 2, // Vertical match at (0,0) and (1,0) -> value 2
					4, 2, 8, 4,
					2, 4, 2, 8,
				},
			},
			expected: false,
		},
		{
			name: "Board full, no moves possible (game over)",
			board: GameBoard{
				Size: 4,
				Tiles: []int{
					2, 4, 2, 4,
					4, 2, 4, 2,
					2, 4, 2, 4,
					4, 2, 4, 2,
				},
			},
			expected: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := isGameOver(&tt.board); got != tt.expected {
				t.Errorf("isGameOver() = %v, want %v", got, tt.expected)
			}
		})
	}
}

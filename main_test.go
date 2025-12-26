package main

import (
	"slices"
	"testing"
)

func TestCompressAndMerge(t *testing.T) {
	tests := []struct {
		name     string
		input    []int
		expected []int
	}{
		{
			name:     "Empty",
			input:    []int{0, 0, 0, 0},
			expected: []int{0, 0, 0, 0},
		},
		{
			name:     "Single Tile",
			input:    []int{2, 0, 0, 0},
			expected: []int{2, 0, 0, 0},
		},
		{
			name:     "Simple Merge",
			input:    []int{2, 2, 0, 0},
			expected: []int{3, 0, 0, 0}, // 2 is 2^1, result 3 is 2^2=4. Note: Logic uses +1 for merge
		},
		{
			name:     "No Merge Different",
			input:    []int{2, 3, 0, 0},
			expected: []int{2, 3, 0, 0},
		},
		{
			name:     "Merge With Gap",
			input:    []int{2, 0, 2, 0},
			expected: []int{3, 0, 0, 0}, // 2(2^1) + 2(2^1) -> 3(2^2)
		},
		{
			name:     "Multi Merge",
			input:    []int{2, 2, 2, 2},
			expected: []int{3, 3, 0, 0}, // 2+2->3, 2+2->3.
		},
		{
			name:     "Merge Priority Left",
			input:    []int{2, 2, 2, 0},
			expected: []int{3, 2, 0, 0}, // First two merge -> 3, remaining 2 stays.
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// Note: compressAndMerge implementation assumes input is just the non-zero tiles
			// or it filters internally. Let's check the implementation.
			// The implementation actively filters 0s.
			// "for _, tile := range tiles { if tile == 0 { continue } ... }"
			// So passing the full row including 0s is fine.

			got := compressAndMerge(tt.input)
			if !slices.Equal(got, tt.expected) {
				t.Errorf("compressAndMerge(%v) = %v, want %v", tt.input, got, tt.expected)
			}
		})
	}
}

func TestGameBoard_Move(t *testing.T) {
	// Helper to create a board with specific tiles
	newBoard := func(tiles []int) *GameBoard {
		return &GameBoard{
			Size:  4,
			Tiles: tiles,
			Score: 0,
		}
	}

	t.Run("Move Left", func(t *testing.T) {
		// [2, 2, 0, 0]
		// [0, 0, 0, 0]
		// [0, 0, 0, 0]
		// [0, 0, 0, 0]
		initial := make([]int, 16)
		initial[0] = 1 // Value 2^1
		initial[1] = 1 // Value 2^1

		board := newBoard(initial)
		// We need to mock random spawn or handle it.
		// Move() calls SpawnTile() internally.
		// However, SpawnTile uses the board state.
		// We can check if the move logic itself worked by inspecting the known positions.

		moved := board.Move(Left)
		if !moved {
			t.Error("Expected move to be valid")
		}

		// Expected after merge:
		// [4, 0, 0, 0] (plus one random tile)
		// Value 2 is 2^2 = 4 (represented as 2 in code logic? Wait, let's check logic)
		// Logic: result[index] += 1. So 1+1 -> 2.
		// Expected at [0]: 2.

		if board.Tiles[0] != 2 {
			t.Errorf("Expected merge result 2 at index 0, got %d", board.Tiles[0])
		}

		// Count non-zeros. Should be 2 (one merged tile + one spawned tile)
		count := 0
		for _, v := range board.Tiles {
			if v != 0 {
				count++
			}
		}
		if count != 2 {
			t.Errorf("Expected 2 non-zero tiles (1 merged + 1 spawned), got %d", count)
		}
	})

	t.Run("No Move Possible", func(t *testing.T) {
		// [2, 4, 2, 4] -> represented as [1, 2, 1, 2]
		// ... full board locked
		// For simplicity, just test a row that can't move left
		initial := make([]int, 16)
		// Set first row to 1, 2, 3, 4 (unique, no merges, packed)
		initial[0] = 1
		initial[1] = 2
		initial[2] = 3
		initial[3] = 4

		board := newBoard(initial)
		moved := board.Move(Left)
		if moved {
			t.Error("Expected no move possible for Left")
		}
	})
}

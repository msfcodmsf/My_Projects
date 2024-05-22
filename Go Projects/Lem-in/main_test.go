package main

import (
	"os"
	"testing"
)

func TestAntMovement(t *testing.T) {
	tests := []struct {
		filename      string
		expectedMoves [][]string
	}{
		{
			filename: "example00.txt",
			expectedMoves: [][]string{
				{"L1-2"},
				{"L1-3", "L2-2"},
				{"L1-1", "L2-3", "L3-2"},
				{"L2-1", "L3-3", "L4-2"},
				{"L3-1", "L4-3"},
				{"L4-1"},
			},
		},
		{
			filename: "example01.txt",
			expectedMoves: [][]string{
				{"L1-t", "L2-h", "L3-0"},
				{"L1-E", "L2-t", "L3-A", "L4-h", "L5-o", "L6-0"},
				{"L1-a", "L2-E", "L3-t", "L4-c", "L5-A", "L6-h", "L7-n", "L8-o", "L9-0"},
				{"L1-m", "L2-a", "L3-E", "L4-t", "L5-k", "L6-c", "L7-A", "L8-e", "L9-n", "L10-o"},
				{"L1-end", "L2-m", "L3-a", "L4-E", "L5-end", "L6-k", "L7-c", "L8-end", "L9-e", "L10-n"},
				{"L2-end", "L3-m", "L4-a", "L6-end", "L7-k", "L9-end", "L10-e"},
				{"L3-end", "L4-m", "L7-end", "L10-end"},
				{"L4-end"},
			},
		},
		{
			filename: "example02.txt",
			expectedMoves: [][]string{
				{"L1-t", "L2-h", "L3-0"},
				{"L1-E", "L2-t", "L3-A", "L4-h", "L5-o", "L6-0"},
				{"L1-a", "L2-E", "L3-t", "L4-c", "L5-A", "L6-h", "L7-n", "L8-o", "L9-0"},
				{"L1-m", "L2-a", "L3-E", "L4-t", "L5-k", "L6-c", "L7-A", "L8-e", "L9-n", "L10-o"},
				{"L1-end", "L2-m", "L3-a", "L4-E", "L5-end", "L6-k", "L7-c", "L8-end", "L9-e", "L10-n"},
				{"L2-end", "L3-m", "L4-a", "L6-end", "L7-k", "L9-end", "L10-e"},
				{"L3-end", "L4-m", "L7-end", "L10-end"},
				{"L4-end"},
			},
		},
		{
			filename: "example03.txt",
			expectedMoves: [][]string{
				{"L1-t", "L2-h", "L3-0"},
				{"L1-E", "L2-t", "L3-A", "L4-h", "L5-o", "L6-0"},
				{"L1-a", "L2-E", "L3-t", "L4-c", "L5-A", "L6-h", "L7-n", "L8-o", "L9-0"},
				{"L1-m", "L2-a", "L3-E", "L4-t", "L5-k", "L6-c", "L7-A", "L8-e", "L9-n", "L10-o"},
				{"L1-end", "L2-m", "L3-a", "L4-E", "L5-end", "L6-k", "L7-c", "L8-end", "L9-e", "L10-n"},
				{"L2-end", "L3-m", "L4-a", "L6-end", "L7-k", "L9-end", "L10-e"},
				{"L3-end", "L4-m", "L7-end", "L10-end"},
				{"L4-end"},
			},
		},
		{
			filename: "example04.txt",
			expectedMoves: [][]string{
				{"L1-t", "L2-h", "L3-0"},
				{"L1-E", "L2-t", "L3-A", "L4-h", "L5-o", "L6-0"},
				{"L1-a", "L2-E", "L3-t", "L4-c", "L5-A", "L6-h", "L7-n", "L8-o", "L9-0"},
				{"L1-m", "L2-a", "L3-E", "L4-t", "L5-k", "L6-c", "L7-A", "L8-e", "L9-n", "L10-o"},
				{"L1-end", "L2-m", "L3-a", "L4-E", "L5-end", "L6-k", "L7-c", "L8-end", "L9-e", "L10-n"},
				{"L2-end", "L3-m", "L4-a", "L6-end", "L7-k", "L9-end", "L10-e"},
				{"L3-end", "L4-m", "L7-end", "L10-end"},
				{"L4-end"},
			},
		},
		{
			filename: "example05.txt",
			expectedMoves: [][]string{
				{"L1-t", "L2-h", "L3-0"},
				{"L1-E", "L2-t", "L3-A", "L4-h", "L5-o", "L6-0"},
				{"L1-a", "L2-E", "L3-t", "L4-c", "L5-A", "L6-h", "L7-n", "L8-o", "L9-0"},
				{"L1-m", "L2-a", "L3-E", "L4-t", "L5-k", "L6-c", "L7-A", "L8-e", "L9-n", "L10-o"},
				{"L1-end", "L2-m", "L3-a", "L4-E", "L5-end", "L6-k", "L7-c", "L8-end", "L9-e", "L10-n"},
				{"L2-end", "L3-m", "L4-a", "L6-end", "L7-k", "L9-end", "L10-e"},
				{"L3-end", "L4-m", "L7-end", "L10-end"},
				{"L4-end"},
			},
		},
		// 2,3,4,5 testlerinin expextedMoves değiştirilmedi ama çıktıları veriyor.
	}

	for _, tt := range tests {
		t.Run(tt.filename, func(t *testing.T) {
			// Call main function with test filename
			os.Args = []string{"cmd", tt.filename}
			main()

			// The generated moves are printed to standard output,
			// so you can capture them and compare with expected moves
		})
	}
}

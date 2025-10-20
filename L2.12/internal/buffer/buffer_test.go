package buffer

import "testing"

func TestEnqueue(t *testing.T) {

	buf := NewBuffer(3)
	tests := []struct {
		line     string
		num      int
		expected int
	}{
		{"one", 1, 0},
		{"two", 2, 1},
		{"three", 3, 2},
		{"four", 4, 0},
	}

	for i, tt := range tests {
		idx := buf.Enqueue(tt.line, tt.num)
		if idx != tt.expected {
			t.Errorf("test %d: Enqueue(%q,%d) = %d, expected = %d", i, tt.line, tt.num, idx, tt.expected)
		}
	}

}

func TestMoveIndex(t *testing.T) {

	buf := NewBuffer(3)
	tests := []struct {
		input, expected int
	}{
		{0, 1},
		{1, 2},
		{2, 0},
	}

	for _, tt := range tests {
		result := buf.moveIndex(tt.input)
		if result != tt.expected {
			t.Errorf("(moveIndex %d) = %d, expected = %d", tt.input, result, tt.expected)
		}
	}

}

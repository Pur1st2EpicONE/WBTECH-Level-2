package main

import (
	"testing"
)

func TestUnpack(t *testing.T) {

	tests := []struct {
		input    string
		expected string
		wantErr  error
	}{
		{
			input:    "a4bc2d5e",
			expected: "aaaabccddddde",
			wantErr:  nil,
		},
		{
			input:    "abcd",
			expected: "abcd",
			wantErr:  nil,
		},
		{
			input:    "",
			expected: "",
			wantErr:  nil,
		},
		{
			input:    "45",
			expected: "",
			wantErr:  ErrNoCharacters,
		},
		{
			input:    "3abc",
			expected: "",
			wantErr:  ErrDigitWithoutChar,
		},
		{
			input:    "1",
			expected: "",
			wantErr:  ErrNoCharacters,
		},
		{
			input:    "0",
			expected: "",
			wantErr:  ErrNoCharacters,
		},
		{
			input:    "a0",
			expected: "",
			wantErr:  nil,
		},
		{
			input:    "a0b3",
			expected: "bbb",
			wantErr:  nil,
		},
		{
			input:    `qwe\4\5`,
			expected: "qwe45",
			wantErr:  nil,
		},
		{
			input:    `qwe\45`,
			expected: "qwe44444",
			wantErr:  nil,
		},
		{
			input:    `qwe\\4`,
			expected: `qwe\\\\`,
			wantErr:  nil,
		},
		{
			input:    `qwe\`,
			expected: `qwe\`,
			wantErr:  nil,
		},
		{
			input:    `\\\\\`,
			expected: `\\\`,
			wantErr:  nil,
		},
		{
			input:    `\`,
			expected: `\`,
			wantErr:  nil,
		},
		{
			input:    `!3^&2*6(`,
			expected: `!!!^&&******(`,
			wantErr:  nil,
		},
		{
			input:    `ф2ц1п0й\34ч\`,
			expected: `ффцй3333ч\`,
			wantErr:  nil,
		},
		{
			input:    `\00`,
			expected: "",
			wantErr:  nil,
		},
	}

	for _, tt := range tests {
		t.Run(tt.input, func(t *testing.T) {
			result, err := unpack(tt.input)
			if tt.wantErr != nil {
				if err == nil {
					t.Errorf("expected error %v, got nil", tt.wantErr)
					return
				}
				if err != tt.wantErr {
					t.Errorf("wrong error: got %v, expected %v", err, tt.wantErr)
					return
				}
				return
			}
			if err != nil {
				t.Errorf("unexpected error: %v", err)
				return
			}
			if result != tt.expected {
				t.Errorf("wrong result:\ngot: %q\nwant: %q", result, tt.expected)
			}
		})
	}

}

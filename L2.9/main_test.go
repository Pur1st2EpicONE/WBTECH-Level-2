package main

import (
	"testing"
)

func TestUnpack(t *testing.T) {
	tests := []struct {
		inputStr    string
		expectedStr string
		expectedErr bool
	}{
		{"a4bc2d5e", "aaaabccddddde", false},
		{"abcd", "abcd", false},
		{"45", "", true},
		{"", "", false},
		{`qwe\4\5`, "qwe45", false},
		{`qwe\45`, "qwe44444", false},
		{`qwe\\4`, `qwe\\\\`, false},
		{`qwe\`, `qwe\`, false},
		{`\\\\\`, `\\\`, false},
		{`\`, `\`, false},
		{`!3^&2*6(`, `!!!^&&******(`, false},
		{`ф2ц1п0й\34ч\`, `ффцй3333ч\`, false},
	}

	for _, test := range tests {
		result, err := unpack(test.inputStr)
		if (err != nil) != test.expectedErr {
			t.Errorf("unpack(%q) error = %v, expected error %v", test.inputStr, err, test.expectedErr)
			continue
		}
		if result != test.expectedStr {
			t.Errorf("unpack(%q) = %q, expected %q", test.inputStr, result, test.expectedStr)
		}
	}
}

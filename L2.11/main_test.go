package main

import (
	"fmt"
	"reflect"
	"testing"
)

func TestGroupAnagrams(t *testing.T) {

	tests := []struct {
		input    []string
		expected map[string][]string
	}{
		{
			input: []string{"пятак", "пятка", "тяпка", "листок", "слиток", "столик", "стол"},
			expected: map[string][]string{
				"пятак":  {"пятак", "пятка", "тяпка"},
				"листок": {"листок", "слиток", "столик"},
			},
		},
		{
			input:    []string{"здесь", "нет", "анаграмм"},
			expected: map[string][]string{},
		},
		{
			input: []string{"кот", "ток", "кот", "окт"},
			expected: map[string][]string{
				"кот": {"кот", "окт", "ток"},
			},
		},
		{
			input: []string{"Кулон", "кОЛуН", "уклон"},
			expected: map[string][]string{
				"колун": {"колун", "кулон", "уклон"},
			},
		},
		{
			input:    []string{},
			expected: map[string][]string{},
		},
	}

	for i, tt := range tests {

		t.Run(fmt.Sprintf("test %d", i), func(t *testing.T) {
			result := groupAnagrams(tt.input)
			if !compareAnagramMaps(result, tt.expected) {
				t.Errorf("result %v, expected %v", result, tt.expected)
			}
		})

	}

}

func compareAnagramMaps(expected, actual map[string][]string) bool {

	if len(expected) != len(actual) {
		return false
	}

	for key, expectedGroup := range expected {
		actualGroup, ok := actual[key]
		if !ok {
			return false
		}
		if !reflect.DeepEqual(expectedGroup, actualGroup) {
			return false
		}
	}

	return true

}

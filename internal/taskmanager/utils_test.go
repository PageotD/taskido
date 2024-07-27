package taskmanager

import (
	"testing"
)

func TestGetMatchValue(t *testing.T) {
	tests := []struct {
		name     string
		input    []string
		expected string
	}{
		{
			name:     "Empty input",
			input:    []string{},
			expected: "",
		},
		{
			name:     "Single element",
			input:    []string{"match"},
			expected: "",
		},
		{
			name:     "Two elements",
			input:    []string{"match1", "match2"},
			expected: "match2",
		},
		{
			name:     "More than two elements",
			input:    []string{"match1", "match2", "match3"},
			expected: "match2",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := GetMatchValue(tt.input)
			if result != tt.expected {
				t.Errorf("GetMatchValue(%v) = %v; want %v", tt.input, result, tt.expected)
			}
		})
	}
}

func TestExtractMatches(t *testing.T) {
	tests := []struct {
		name     string
		input    [][]string
		expected []string
	}{
		{
			name:     "Empty input",
			input:    [][]string{},
			expected: []string{},
		},
		{
			name:     "Single empty match",
			input:    [][]string{{}},
			expected: []string{""},
		},
		{
			name:     "Single match with value",
			input:    [][]string{{"match1", "match2"}},
			expected: []string{"match2"},
		},
		{
			name:     "Multiple matches",
			input:    [][]string{{"match1", "match2"}, {"match3", "match4"}},
			expected: []string{"match2", "match4"},
		},
		{
			name:     "Some matches are empty",
			input:    [][]string{{"match1", "match2"}, {"match3"}},
			expected: []string{"match2", ""},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := ExtractMatches(tt.input)
			if len(result) != len(tt.expected) {
				t.Errorf("ExtractMatches(%v) = %v; want %v", tt.input, result, tt.expected)
			}
			for i, v := range result {
				if v != tt.expected[i] {
					t.Errorf("ExtractMatches(%v)[%d] = %v; want %v", tt.input, i, v, tt.expected[i])
				}
			}
		})
	}
}

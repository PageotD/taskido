package taskmanager

import (
	"testing"
	"github.com/stretchr/testify/assert"
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

func TestExtractProjects(t *testing.T) {
    tests := []struct {
		name     string
        input    string
        expected []string
    }{
        {
			name: "One project",
            input: "+project1 taskido project ",
            expected: []string{"+project1"},
        },
        {
			name: "No project",
            input: "No projects here",
            expected: []string{},
        },
        {
			name: "+taskido +fileserver Multiple projects",
            input:    "Multiple +projects +project4",
            expected: []string{"+projects", "+project4"},
        },
    }

    for _, test := range tests {
        result := ExtractProjects(test.input)
        assert.ElementsMatch(t, test.expected, result, "For input %s", test.input)
    }
}

func TestExtractContexts(t *testing.T) {
    tests := []struct {
		name     string
        input    string
        expected []string
    }{
        {
			name: "Multiple contexts",
            input: "Task with @context1 and @context2",
            expected: []string{"@context1", "@context2"},
        },
        {
			name: "No context",
            input: "No contexts here",
            expected: []string{},
        },
        {
			name: "One context",
            input:    "Another task @context3",
            expected: []string{"@context3"},
        },
    }

    for _, test := range tests {
        result := ExtractContexts(test.input)
        assert.ElementsMatch(t, test.expected, result, "For input %s", test.input)
    }
}

func TestExtractDue(t *testing.T) {
    tests := []struct {
		name     string
        input    string
        expected string
    }{
        {
			name: "Valid date",
            input:    "Task due:2024-08-15",
            expected: "2024-08-15",
        },
        {
			name: "No date",
            input:    "No due date here",
            expected: "",
        },
        {
			name: "Invalid date",
            input:    "Invalid due date format due:2024-02-30",
            expected: "2024-02-30", // Even if invalid date, pattern match is correct
        },
    }

    for _, test := range tests {
        result := ExtractDue(test.input)
        assert.Equal(t, test.expected, result, "For input %s", test.input)
    }
}
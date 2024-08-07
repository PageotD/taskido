package libtaskido

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
			result := getMatchValue(tt.input)
			if result != tt.expected {
				t.Errorf("getMatchValue(%v) = %v; want %v", tt.input, result, tt.expected)
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
        result := extractProjects(test.input)
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
        result := extractContexts(test.input)
        assert.ElementsMatch(t, test.expected, result, "For input %s", test.input)
    }
}

func TestExtractDueDate(t *testing.T) {
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
        result := extractDueDate(test.input)
        assert.Equal(t, test.expected, result, "For input %s", test.input)
    }
}

func TestExtractPriority(t *testing.T) {
    tests := []struct {
		name     string
        input    string
        expected int
    }{
        {
			name: "Valid priority",
            input:    "Task priority:1",
            expected: 1,
        },
        {
			name: "No priority",
            input:    "No priority here",
            expected: 0,
        },
        {
			name: "Invalid priority",
            input:    "Invalid priority:9",
            expected: 0, // Even if invalid date, pattern match is correct
        },
    }

    for _, test := range tests {
        result := extractPriority(test.input)
        assert.Equal(t, test.expected, result, "For input %s", test.input)
    }
}
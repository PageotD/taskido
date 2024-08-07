package libtaskido

import (
	"strings"
	"testing"
	"time"
	"github.com/stretchr/testify/assert"
)

func TestFormatDate(t *testing.T) {
	tests := []struct {
		date        string
		expectation string
	}{
		// Two days after date
		{time.Now().Add(-48 * time.Hour).Format("2006-01-02"), "\033[0;31m⚑\033[0;0m "+time.Now().Add(-48*time.Hour).Format("2006-01-02")},
		// Yesterday date
		{time.Now().Add(-24 * time.Hour).Format("2006-01-02"), "\033[0;31m⚑\033[0;0m "+time.Now().Add(-24*time.Hour).Format("2006-01-02")},
		// Today date
		{time.Now().Truncate(24 * time.Hour).Format("2006-01-02"), "\033[0;31m⚑\033[0;0m "+ time.Now().Truncate(24*time.Hour).Format("2006-01-02")},
		// Tomorrow date
		{time.Now().Add(24 * time.Hour).Format("2006-01-02"), "\033[0;33m⚑\033[0;0m "+time.Now().Add(24*time.Hour).Format("2006-01-02")},
		// Two day before date
		{time.Now().Add(48 * time.Hour).Format("2006-01-02"), "  "+time.Now().Add(48*time.Hour).Format("2006-01-02")},
	}

	for _, tt := range tests {
		t.Run(tt.date, func(t *testing.T) {
			result := formatDate(tt.date)
			if result != tt.expectation {
				t.Errorf("formatDate(%s) = %s; want %s", tt.date, result, tt.expectation)
			}
		})
	}
}

func TestApplyColorToDescription(t *testing.T) {
	tests := []struct {
		description string
		expectation string
	}{
		{"@John", "\033[0;34m@John\033[0;0m"},                           // With @
		{"Meeting with @Jane", "Meeting with \033[0;34m@Jane\033[0;0m"}, // With @
		{"No context", "No context"},                                // No context
	}

	for _, tt := range tests {
		t.Run(tt.description, func(t *testing.T) {
			result := applyColorToDescription(tt.description)
			if result != tt.expectation {
				t.Errorf("applyColorToDescription(%s) = %s; want %s", tt.description, result, tt.expectation)
			}
		})
	}
}

func TestApplyColorToProject(t *testing.T) {
	tests := []struct {
		projects    []string
		expectation string
	}{
		{[]string{"+acc", "+work"}, "\033[0;35m+acc\033[0;0m \033[0;35m+work\033[0;0m"}, // Multiple projects
		{[]string{"+test"}, "\033[0;35m+test\033[0;0m"},                            // Single project
		{[]string{}, ""}, // No projects
	}

	for _, tt := range tests {
		t.Run(strings.Join(tt.projects, ", "), func(t *testing.T) {
			result := applyColorToProject(tt.projects)
			if result != tt.expectation {
				t.Errorf("applyColorToProject(%v) = %s; want %s", tt.projects, result, tt.expectation)
			}
		})
	}
}

func TestFormatTask(t *testing.T) {
	tests := []struct {
		task     Task
		expected []string
	}{
		{
			task: Task{
				ID:        1,
				Due:       "2024-07-24",
				Projects:  []string{"project1", "project2"},
				Description:   "@context1 Do something",
				Status:    "pending",
				Priority:  0,
			},
			expected: []string{"1", "\033[0;31m⚑\033[0;0m", "2024-07-24", "\033[0;35mproject1\033[0;0m", "\033[0;35mproject2\033[0;0m", "\033[0;34m@context1\033[0;0m", "Do", "something"},
		},
		{
			task: Task{
				ID:        2,
				Due:       "2024-07-23",
				Projects:  []string{"project3"},
				Description:   "@context2 Another task",
				Status: "completed",
				Priority:  3,
			},
			expected: []string{"2", "\033[0;31m\u278C\033[0;0m", "\033[0;31m⚑\033[0;0m", "2024-07-23", "\033[0;35mproject3\033[0;0m", "\033[0;34m@context2\033[0;0m", "Another", "task"},
		},
	}

	for _, tt := range tests {
		t.Run("", func(t *testing.T) {
			got := strings.Fields(formatTask(tt.task)) // Split by whitespace
			assert.ElementsMatch(t, tt.expected, got)
		})
	}
}

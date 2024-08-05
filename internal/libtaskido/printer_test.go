package libtaskido

import (
	"strings"
	"testing"
	"time"
)

func TestApplyColorToDate(t *testing.T) {
	tests := []struct {
		date        string
		expectation string
	}{
		// Two days after date
		{time.Now().Add(-48 * time.Hour).Format("2006-01-02"), "\033[31m" + time.Now().Add(-48*time.Hour).Format("2006-01-02") + "\033[0m"},
		// Yesterday date
		{time.Now().Add(-24 * time.Hour).Format("2006-01-02"), "\033[31m" + time.Now().Add(-24*time.Hour).Format("2006-01-02") + "\033[0m"},
		// Today date
		{time.Now().Truncate(24 * time.Hour).Format("2006-01-02"), "\033[31m" + time.Now().Truncate(24*time.Hour).Format("2006-01-02") + "\033[0m"},
		// Tomorrow date
		{time.Now().Add(24 * time.Hour).Format("2006-01-02"), "\033[33m" + time.Now().Add(24*time.Hour).Format("2006-01-02") + "\033[0m"},
		// Two day before date
		{time.Now().Add(48 * time.Hour).Format("2006-01-02"), "\033[32m" + time.Now().Add(48*time.Hour).Format("2006-01-02") + "\033[0m"},
	}

	for _, tt := range tests {
		t.Run(tt.date, func(t *testing.T) {
			result := applyColorToDate(tt.date)
			if result != tt.expectation {
				t.Errorf("applyColorToDate(%s) = %s; want %s", tt.date, result, tt.expectation)
			}
		})
	}
}

func TestApplyColorToSubject(t *testing.T) {
	tests := []struct {
		subject     string
		expectation string
	}{
		{"@John", "\033[34m@John\033[0m"},                           // With @
		{"Meeting with @Jane", "Meeting with \033[34m@Jane\033[0m"}, // With @
		{"No context", "No context"},                                // No context
	}

	for _, tt := range tests {
		t.Run(tt.subject, func(t *testing.T) {
			result := applyColorToSubject(tt.subject)
			if result != tt.expectation {
				t.Errorf("applyColorToSubject(%s) = %s; want %s", tt.subject, result, tt.expectation)
			}
		})
	}
}

func TestApplyColorToProject(t *testing.T) {
	tests := []struct {
		projects    []string
		expectation string
	}{
		{[]string{"+acc", "+work"}, "\033[35m+acc\033[0m \033[35m+work\033[0m"}, // Multiple projects
		{[]string{"+test"}, "\033[35m+test\033[0m"},                            // Single project
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
		expected string
	}{
		{
			task: Task{
				ID:        1,
				Due:       "2024-07-24",
				Projects:  []string{"project1", "project2"},
				Subject:   "@context1 Do something",
				Completed: false,
				Archived:  false,
				Priority:  0,
			},
			expected: "1    \x1b[31m2024-07-24\x1b[0m " + "0 " +
				"\x1b[35mproject1\x1b[0m \x1b[35mproject2\x1b[0m " +
				"\x1b[34m@context1\x1b[0m Do something\n",
		},
		{
			task: Task{
				ID:        2,
				Due:       "2024-07-23",
				Projects:  []string{"project3"},
				Subject:   "@context2 Another task",
				Completed: true,
				Archived:  false,
				Priority:  3,
			},
			expected: "2    \x1b[31m2024-07-23\x1b[0m " + "3 " +
				"\x1b[35mproject3\x1b[0m " +
				"\x1b[34m@context2\x1b[0m Another task\n",
		},
	}

	for _, tt := range tests {
		t.Run("", func(t *testing.T) {
			got := formatTask(tt.task)
			if got != tt.expected {
				t.Errorf("formatTask() = %v, want %v", got, tt.expected)
			}
		})
	}
}

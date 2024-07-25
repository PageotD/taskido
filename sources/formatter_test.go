package main

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
		{[]string{"acc", "work"}, "\033[35m+acc\033[0m \033[35m+work\033[0m"}, // Multiple projects
		{[]string{"test"}, "\033[35m+test\033[0m"},                            // Single project
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

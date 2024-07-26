package formatter

import (
	"regexp"
	"strings"
	"time"
)

// ANSI color codes
const (
	Red    = "\033[31m" // Red color code
	Orange = "\033[33m" // Orange color code
	Blue   = "\033[34m" // Blue color code
	Green  = "\033[32m" // Green color code
	Violet = "\033[35m" // Violet color code
	Reset  = "\033[0m"  // Reset color code
)

// applyColorToDate applies color based on the date's proximity to today
func ApplyColorToDate(dueDate string) string {
	// Parse the due date
	date, err := time.Parse("2006-01-02", dueDate)
	if err != nil {
		return dueDate // return the original date if parsing fails
	}

	// Get today's date
	now := time.Now()
	today := now.Truncate(24 * time.Hour)
	tomorrow := today.Add(24 * time.Hour)

	// Determine the color based on the date
	if date.Before(today) || date.Equal(today) {
		return Red + dueDate + Reset
	} else if date.Equal(tomorrow) {
		return Orange + dueDate + Reset
	} else {
		return Green + dueDate + Reset
	}
}

// Helper function to apply color to the contexts
//func applyColorToDate(subject string) string {
// Replace @ followed by any characters with blue color
//	return regexp.MustCompile(`@(\S+)`).ReplaceAllString(subject, Blue+"@$1"+Reset)
//}

// Helper function to apply color to the contexts
func ApplyColorToSubject(subject string) string {
	// Replace @ followed by any characters with blue color
	return regexp.MustCompile(`@(\S+)`).ReplaceAllString(subject, Blue+"@$1"+Reset)
}

// Helper function to apply color to the project names
func ApplyColorToProject(projectList []string) string {
	var coloredProjects []string
	for _, project := range projectList {
		coloredProjects = append(coloredProjects, Violet+"+"+project+Reset)
	}
	// Join colored projects with a space
	return strings.Join(coloredProjects, " ")
}

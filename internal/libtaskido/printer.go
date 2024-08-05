package libtaskido

import (
	"fmt"
	"strings"
	"time"
	"regexp"
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
func applyColorToDate(dueDate string) string {
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
func applyColorToSubject(subject string) string {
	// Replace @ followed by any characters with blue color
	return regexp.MustCompile(`@(\S+)`).ReplaceAllString(subject, Blue+"@$1"+Reset)
}

// Helper function to apply color to the project names
func applyColorToProject(projectList []string) string {
	var coloredProjects []string
	for _, project := range projectList {
		coloredProjects = append(coloredProjects, Violet+project+Reset)
	}
	// Join colored projects with a space
	return strings.Join(coloredProjects, " ")
}

// formatTask formats a single task into a string with a specific layout and color coding.
func formatTask(task Task) string {
	return fmt.Sprintf("%-4d %-12s %-1d %s %s\n", task.ID, applyColorToDate(task.Due), task.Priority, applyColorToProject(task.Projects), applyColorToSubject(task.Subject))
}

// HandleList lists all tasks
func PrintTaskList(taskList []Task) {

    fmt.Printf("\n\033[4mCurrent:\033[0m\n\n")
    for _, task := range taskList {
        if !task.Completed && !task.Archived {
            fmt.Printf(formatTask(task))
        }
    }

    fmt.Printf("\n\033[4mCompleted:\033[0m\n\n")
    for _, task := range taskList {
        if task.Completed && !task.Archived {
			fmt.Printf(formatTask(task))       
		}
    }

    fmt.Printf("\n\033[4mArchived:\033[0m\n\n")
    for _, task := range taskList {
        if task.Archived {
			fmt.Printf(formatTask(task))        
		}
    }

    fmt.Printf("\n")
}
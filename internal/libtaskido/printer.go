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
		return "          " //dueDate // return the original date if parsing fails
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

func formatPriority(priority int) string {

	if priority == 1 {
		return fmt.Sprintf("\033[32m\u278A\033[0m")
	} else if priority == 2 {
		return fmt.Sprintf("\033[33m\u278B\033[0m")
	} else if priority == 3 {
		return fmt.Sprintf("\033[31m\u278C\033[0m")
	} 

	return "\u0020"
}

// formatTask formats a single task into a string with a specific layout and color coding.
func formatTask(task Task) string {
	return fmt.Sprintf("%-4d\u0020\u0020 %s\u0020\u0020\u0020\u0020 %s\u0020\u0020\u0020\u0020 %s %s\n", task.ID, formatPriority(task.Priority), applyColorToDate(task.Due), applyColorToProject(task.Projects), applyColorToSubject(task.Subject))
}

// PrintTaskList lists all tasks grouped by their status (current, completed, archived)
func PrintTaskList(taskList []Task) {
	// Use map to collect projects by status
	statusTasks := make(map[string][]Task)

	// Iterate through tasks and collect projects
	for _, task := range taskList {
		if task.Archived {
			statusTasks["archived"] = append(statusTasks["archived"], task)
		} else if !task.Archived && task.Completed {
			statusTasks["completed"] = append(statusTasks["completed"], task)
		} else {
			statusTasks["current"] = append(statusTasks["current"], task)
		}
	}

	// Print tasks grouped by status
	fmt.Printf("\033[4mCurrent:\033[0m\n")
	for _, task := range statusTasks["current"] {
		fmt.Printf(formatTask(task)) 
	}
	fmt.Println()
	
	fmt.Printf("\033[4mCompleted:\033[0m\n")
	for _, task := range statusTasks["completed"] {
		fmt.Printf(formatTask(task)) 
	}
	fmt.Println()

	fmt.Printf("\033[4mArchived:\033[0m\n")
	for _, task := range statusTasks["archived"] {
		fmt.Printf(formatTask(task)) 
	}
	fmt.Println()
}

// PrintTaskListByProjects lists all tasks grouped by their associated projects.
func PrintTaskListByProjects(taskList []Task) {
	// Use a map to collect unique projects
	projectTasks := make(map[string][]Task)

	// Iterate through tasks and collect projects
	for _, task := range taskList {
		for _, project := range task.Projects {
			projectTasks[project] = append(projectTasks[project], task)
		}
	}

	// Print tasks grouped by projects
	for project, tasks := range projectTasks {
		fmt.Printf("\033[35m%s\033[0m\n", project)
		for _, task := range tasks {
			fmt.Printf(formatTask(task)) 
		}
		fmt.Println()
	}
}

// PrintTaskListByContexts lists all tasks grouped by their associated contexts.
func PrintTaskListByContexts(taskList []Task) {
	// Use a map to collect unique contexts
	contextTasks := make(map[string][]Task)

	// Iterate through tasks and collect projects
	for _, task := range taskList {
		for _, context := range task.Contexts {
			contextTasks[context] = append(contextTasks[context], task)
		}
	}

	// Print tasks grouped by projects
	for context, tasks := range contextTasks {
		fmt.Printf("\033[34m%s\033[0m\n", context)
		for _, task := range tasks {
			fmt.Printf(formatTask(task)) 
		}
		fmt.Println()
	}
}
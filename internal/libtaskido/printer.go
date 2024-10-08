package libtaskido

import (
	"fmt"
	"regexp"
	"strings"
	"time"
)

// ANSI color codes
const (
	Reset = "\033[0;0m"
	// Regular Colors
	Black  = "\033[0;30m"
	Red    = "\033[0;31m"
	Green  = "\033[0;32m"
	Yellow = "\033[0;33m"
	Blue   = "\033[0;34m"
	Purple = "\033[0;35m"
	Cyan   = "\033[0;36m"
	White  = "\033[0;37m"
)

// formatDate applies color based on the date's proximity to today
func formatDate(dueDate string) string {
	// Parse the due date
	date, err := time.Parse("2006-01-02", dueDate)
	if err != nil {
		return strings.Repeat(" ", 12) //dueDate // return the original date if parsing fails
	}

	// Get today's date
	now := time.Now()
	today := now.Truncate(24 * time.Hour)
	tomorrow := today.Add(24 * time.Hour)

	// Add Flag or not based on the date
	if date.Before(today) || date.Equal(today) {
		return fmt.Sprintf("%s\u2691%s %s", Red, Reset, dueDate)
	} else if date.Equal(tomorrow) {
		return fmt.Sprintf("%s\u2691%s %s", Yellow, Reset, dueDate)
	} else {
		return fmt.Sprintf("  %s", dueDate)
	}
}

// Helper function to apply color to the contexts
func applyColorToDescription(description string) string {
	// Replace @ followed by any characters with blue color
	return regexp.MustCompile(`@(\S+)`).ReplaceAllString(description, Blue+"@$1"+Reset)
}

// Helper function to apply color to the project names
func applyColorToProject(projectList []string) string {
	var coloredProjects []string
	for _, project := range projectList {
		coloredProjects = append(coloredProjects, Purple+project+Reset)
	}
	// Join colored projects with a space
	return strings.Join(coloredProjects, " ")
}

func formatPriority(priority int) string {

	if priority == 1 {
		return fmt.Sprintf("\033[0;32m\u278A\033[0;0m")
	} else if priority == 2 {
		return fmt.Sprintf("\033[0;33m\u278B\033[0;0m")
	} else if priority == 3 {
		return fmt.Sprintf("\033[0;31m\u278C\033[0;0m")
	}

	return "\u0020"
}

// formatTask formats a single task into a string with a specific layout and color coding.
func formatTask(task Task) string {
	return fmt.Sprintf("%-4d\u0020\u0020 %s\u0020\u0020\u0020\u0020 %s\u0020\u0020\u0020\u0020 %s %s\n", task.ID, formatPriority(task.Priority), formatDate(task.Due), applyColorToProject(task.Projects), applyColorToDescription(task.Description))
}

// PrintTaskList lists all tasks grouped by their status (current, completed, archived)
func PrintTaskList(taskList []Task) {
	// Use map to collect projects by status
	statusTasks := make(map[string][]Task)

	// Iterate through tasks and collect projects
	for _, task := range taskList {
		if task.Status == "archived" {
			statusTasks["archived"] = append(statusTasks["archived"], task)
		} else if task.Status == "completed" {
			statusTasks["completed"] = append(statusTasks["completed"], task)
		} else {
			statusTasks["pending"] = append(statusTasks["pending"], task)
		}
	}

	// Print tasks grouped by status
	fmt.Printf("\033[1;4mPending:\033[0;0m\n")
	for _, task := range statusTasks["pending"] {
		fmt.Printf(formatTask(task))
	}
	fmt.Println()

	fmt.Printf("\033[1;4mCompleted:\033[1;0m\n")
	for _, task := range statusTasks["completed"] {
		fmt.Printf(formatTask(task))
	}
	fmt.Println()

	fmt.Printf("\033[1;4mArchived:\033[1;0m\n")
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
		if len(task.Projects) > 0 {
			for _, project := range task.Projects {
				projectTasks[project] = append(projectTasks[project], task)
			}
		} else {
			projectTasks["no project"] = append(projectTasks["no project"], task)
		}
	}

	// Print "no project" first
	if len(projectTasks["no project"]) > 0 {
		fmt.Printf("\033[34mno project\033[0m\n")
		for _, task := range projectTasks["no project"] {
			fmt.Printf(formatTask(task))
		}
		fmt.Println()
	}

	// Print tasks grouped by projects except "no project"
	for project, tasks := range projectTasks {
		if project != "no project" {
			fmt.Printf("\033[34m%s\033[0m\n", project)
			for _, task := range tasks {
				fmt.Printf(formatTask(task))
			}
			fmt.Println()
		}
	}
}

// PrintTaskListByContexts lists all tasks grouped by their associated contexts.
func PrintTaskListByContexts(taskList []Task) {
	// Use a map to collect unique contexts
	contextTasks := make(map[string][]Task)

	// Iterate through tasks and collect projects
	for _, task := range taskList {
		if len(task.Contexts) > 0 {
			for _, context := range task.Contexts {
				contextTasks[context] = append(contextTasks[context], task)
			}
		} else {
			contextTasks["no context"] = append(contextTasks["no context"], task)
		}
	}

	// Print "no context" first
	if len(contextTasks["no context"]) > 0 {
		fmt.Printf("\033[34mno context\033[0m\n")
		for _, task := range contextTasks["no context"] {
			fmt.Printf(formatTask(task))
		}
		fmt.Println()
	}

	// Print tasks grouped by contexts except "no context"
	for context, tasks := range contextTasks {
		if context != "no context" {
			fmt.Printf("\033[34m%s\033[0m\n", context)
			for _, task := range tasks {
				fmt.Printf(formatTask(task))
			}
			fmt.Println()
		}
	}
}

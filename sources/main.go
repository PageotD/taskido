package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"os"
	"regexp"
	"strings"

	"github.com/google/uuid"
)

// ANSI color codes
const (
	Blue   = "\033[34m" // Blue color code
	Green  = "\033[32m" // Green color code
	Violet = "\033[35m" // Violet color code
	Reset  = "\033[0m"  // Reset color code
)

// Task structure corresponds to the JSON object
type Task struct {
	ID            int      `json:"id"`
	UUID          string   `json:"uuid"`
	Subject       string   `json:"subject"`
	Projects      []string `json:"projects"`
	Contexts      []string `json:"contexts"`
	Due           string   `json:"due"`
	Completed     bool     `json:"completed"`
	CompletedDate string   `json:"completedDate"`
	Archived      bool     `json:"archived"`
	IsPriority    bool     `json:"isPriority"`
	Notes         []string `json:"notes"`
}

func main() {
	// Define flags for the commands
	addFlag := flag.Bool("add", false, "Indicates that the following text should be processed")
	listFlag := flag.Bool("list", false, "List all tasks")

	// Parse the command-line flags
	flag.Parse()

	if *addFlag {
		// Handle the -add command
		handleAdd(flag.Args())
	} else if *listFlag {
		// Handle the -list command
		handleList()
	} else {
		fmt.Println("No valid flag provided. Use -add to add a task or -list to list tasks.")
	}
}

func handleAdd(args []string) {
	// Join all remaining arguments into a single string
	addText := strings.Join(args, " ")

	// Define regex patterns for different fields
	projectPattern := regexp.MustCompile(`\+(\S+)`)
	contextPattern := regexp.MustCompile(`@(\S+)`)
	duePattern := regexp.MustCompile(`due:(\d{4}-\d{2}-\d{2})`)

	// Extract fields using regex
	projectMatches := projectPattern.FindAllStringSubmatch(addText, -1)
	contextMatch := contextPattern.FindStringSubmatch(addText)
	dueMatch := duePattern.FindStringSubmatch(addText)

	// Extract the task description
	taskDescription := addText
	for _, match := range projectMatches {
		taskDescription = strings.Replace(taskDescription, match[0], "", 1)
	}
	if dueMatch != nil {
		taskDescription = strings.Replace(taskDescription, dueMatch[0], "", 1)
	}
	taskDescription = strings.TrimSpace(taskDescription)

	// Create the Task instance
	task := Task{
		ID:            0, // Will be updated later
		UUID:          uuid.NewString(),
		Subject:       taskDescription,
		Projects:      extractMatches(projectMatches),
		Contexts:      []string{getMatchValue(contextMatch)},
		Due:           getMatchValue(dueMatch),
		Completed:     false,
		CompletedDate: "",
		Archived:      false,
		IsPriority:    false,
		Notes:         nil,
	}

	// Load existing tasks from the file
	var tasks []Task
	filePath := "tasks.json"
	if _, err := os.Stat(filePath); err == nil {
		// File exists, read it
		file, err := os.ReadFile(filePath)
		if err != nil {
			fmt.Printf("Error reading file: %v\n", err)
			return
		}
		if err := json.Unmarshal(file, &tasks); err != nil {
			fmt.Printf("Error unmarshalling JSON: %v\n", err)
			return
		}
	}

	// Set the next ID
	task.ID = len(tasks) + 1

	// Append the new task to the list
	tasks = append(tasks, task)

	// Convert the tasks slice to JSON
	jsonData, err := json.MarshalIndent(tasks, "", "  ")
	if err != nil {
		fmt.Printf("Error marshalling JSON: %v\n", err)
		return
	}

	// Save JSON to a file
	err = os.WriteFile(filePath, jsonData, 0644)
	if err != nil {
		fmt.Printf("Error writing file: %v\n", err)
		return
	}

	fmt.Println("Task added to tasks.json")
}

func handleList() {
	// Load existing tasks from the file
	var tasks []Task
	filePath := "tasks.json"
	if _, err := os.Stat(filePath); err != nil {
		fmt.Println("No tasks found.")
		return
	}

	file, err := os.ReadFile(filePath)
	if err != nil {
		fmt.Printf("Error reading file: %v\n", err)
		return
	}
	if err := json.Unmarshal(file, &tasks); err != nil {
		fmt.Printf("Error unmarshalling JSON: %v\n", err)
		return
	}

	// Print the tasks
	for _, task := range tasks {
		// Apply color to parts of the project names
		subjectWithColor := applyColorToSubject(task.Subject)
		projectWithColor := applyColorToProject(task.Projects)
		fmt.Printf("%-4d \033[32m%-12s \033[35m%s \033[0m %s\n", task.ID, task.Due, projectWithColor, subjectWithColor)
	}
}

// Helper function to get the match value from regex capture groups
func getMatchValue(matches []string) string {
	if len(matches) > 1 {
		return matches[1]
	}
	return ""
}

// Helper function to extract projects from regex matches
func extractMatches(matches [][]string) []string {
	var result []string
	for _, match := range matches {
		result = append(result, getMatchValue(match))
	}
	return result
}

// Helper function to apply color to the subject
func applyColorToSubject(subject string) string {
	// Replace @ followed by any characters with blue color
	return regexp.MustCompile(`@(\S+)`).ReplaceAllString(subject, Blue+"@$1"+Reset)
}

// Helper function to apply color to the project names
func applyColorToProject(projectList []string) string {
	var coloredProjects []string
	for _, project := range projectList {
		coloredProjects = append(coloredProjects, Violet+"+"+project+Reset)
	}
	// Join colored projects with a space
	return strings.Join(coloredProjects, " ")
}

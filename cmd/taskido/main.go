package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"os"
	"regexp"
	"strings"
	"time"

	"github.com/google/uuid"
	"taskido/internal/formatter"
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
	cmpFlag := flag.Int("cmp", 0, "Mark a task as complete")

	// Parse the command-line flags
	flag.Parse()

	if *addFlag {
		// Handle the -add command
		handleAdd(flag.Args())
	} else if *listFlag {
		// Handle the -list command
		handleList()
	} else if *cmpFlag != 0 {
		// Handle the -cmp command
		handleComplete(*cmpFlag)
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
		subjectWithColor := formatter.ApplyColorToSubject(task.Subject)
		projectWithColor := formatter.ApplyColorToProject(task.Projects)
		dueDateWithColor := formatter.ApplyColorToDate(task.Due)
		fmt.Printf("%-4d %-12s %s %s\n", task.ID, dueDateWithColor, projectWithColor, subjectWithColor)
	}
}

func handleComplete(taskID int) {
	// Charger les tâches depuis le fichier JSON
	var tasks []Task
	filePath := "tasks.json"

	// Vérifier si le fichier existe
	if _, err := os.Stat(filePath); err == nil {
		file, err := os.ReadFile(filePath)
		if err != nil {
			fmt.Printf("Error reading file: %v\n", err)
			return
		}
		if err := json.Unmarshal(file, &tasks); err != nil {
			fmt.Printf("Error unmarshalling JSON: %v\n", err)
			return
		}
	} else {
		fmt.Println("No tasks found.")
		return
	}

	// Obtenir la date actuelle
	now := time.Now().Format("2006-01-02")

	// Mettre à jour la tâche spécifiée
	var updatedTasks []Task
	taskUpdated := false
	for _, task := range tasks {
		if task.ID == taskID {
			task.Completed = true
			task.CompletedDate = now
			taskUpdated = true
		}
		updatedTasks = append(updatedTasks, task)
	}

	if !taskUpdated {
		fmt.Printf("Task ID %d not found\n", taskID)
		return
	}

	// Écrire les modifications dans le fichier JSON
	jsonData, err := json.MarshalIndent(updatedTasks, "", "  ")
	if err != nil {
		fmt.Printf("Error marshalling JSON: %v\n", err)
		return
	}

	err = os.WriteFile(filePath, jsonData, 0644)
	if err != nil {
		fmt.Printf("Error writing file: %v\n", err)
		return
	}

	fmt.Println("Tasks updated successfully.")
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

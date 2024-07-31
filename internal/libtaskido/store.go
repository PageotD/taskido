package libtaskido

import (
	"encoding/json"
	"fmt"
	"os"
)

// Name of the JSON file where tasks are stored
const tasksJSONFile = "tasks.json"

// EnsureFileExists checks if the JSON file exists in the current directory.
// It returns true if the file exists, otherwise false.
func EnsureFileExists() bool {
	if _, err := os.Stat(tasksJSONFile); os.IsNotExist(err) {
		return false
	}
	return true
}

// InitializeFile creates a new JSON file with an empty array of tasks if the file does not exist.
// It returns an error if there is an issue creating or writing to the file.
func InitializeFile() error {
	file, err := os.Create(tasksJSONFile)
	if err != nil {
		return fmt.Errorf("error creating file: %v", err)
	}
	defer file.Close()

	// Write an empty JSON array to the file
	if err := json.NewEncoder(file).Encode([]Task{}); err != nil {
		return fmt.Errorf("error writing to file: %v", err)
	}

	return nil
}

// LoadTasks loads the list of tasks from the JSON file.
// It returns a slice of tasks and an error. If there is an issue opening or reading the file, 
// the function returns an error and an empty slice of tasks.
func LoadTasks() ([]Task, error) {
	file, err := os.Open(tasksJSONFile)
	if err != nil {
		return nil, fmt.Errorf("error opening file: %v", err)
	}
	defer file.Close()

	var tasks []Task
	if err := json.NewDecoder(file).Decode(&tasks); err != nil {
		return nil, fmt.Errorf("error reading from file: %v", err)
	}

	return tasks, nil
}

// SaveTasks saves the provided slice of tasks to the JSON file.
// It returns an error if there is an issue creating or writing to the file.
func SaveTasks(tasks []Task) error {
	file, err := os.Create(tasksJSONFile)
	if err != nil {
		return fmt.Errorf("error creating file: %v", err)
	}
	defer file.Close()

	if err := json.NewEncoder(file).Encode(tasks); err != nil {
		return fmt.Errorf("error writing to file: %v", err)
	}

	return nil
}

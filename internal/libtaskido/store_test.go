package libtaskido

import (
	"encoding/json"
	"os"
	"testing"
)

func TestEnsureFileExists(t *testing.T) {
	// Clean up the file if it exists before testing
	defer os.Remove(tasksJSONFile)

	// Test when the file does not exist
	if exists := EnsureFileExists(); exists {
		t.Errorf("Expected file to not exist, but it does.")
	}

	// Create the file and test again
	if err := InitializeFile(); err != nil {
		t.Fatalf("Error initializing file: %v", err)
	}

	if exists := EnsureFileExists(); !exists {
		t.Errorf("Expected file to exist, but it does not.")
	}
}

func TestInitializeFile(t *testing.T) {
	// Clean up the file if it exists before testing
	defer os.Remove(tasksJSONFile)

	// Test file initialization
	if err := InitializeFile(); err != nil {
		t.Fatalf("Error initializing file: %v", err)
	}

	// Verify that the file is created and is empty
	fileInfo, err := os.Stat(tasksJSONFile)
	if err != nil {
		t.Fatalf("Error statting file: %v", err)
	}

	if fileInfo.Size() == 0 {
		t.Errorf("File size should be greater than 0, but it is 0.")
	}

	// Verify that the file contains an empty JSON array
	file, err := os.Open(tasksJSONFile)
	if err != nil {
		t.Fatalf("Error opening file: %v", err)
	}
	defer file.Close()

	var tasks []Task
	if err := json.NewDecoder(file).Decode(&tasks); err != nil {
		t.Fatalf("Error decoding file content: %v", err)
	}

	if len(tasks) != 0 {
		t.Errorf("Expected an empty task list, but got %d tasks.", len(tasks))
	}
}

func TestLoadTasks(t *testing.T) {
	// Clean up the file if it exists before testing
	defer os.Remove(tasksJSONFile)

	// Prepare sample tasks
	tasks := []Task{
		{ID: 1, UUID: "uuid1", Description: "Task 1"},
		{ID: 2, UUID: "uuid2", Description: "Task 2"},
	}

	// Save tasks to the file
	if err := SaveTasks(tasks); err != nil {
		t.Fatalf("Error saving tasks: %v", err)
	}

	// Load tasks from the file
	loadedTasks, err := LoadTasks()
	if err != nil {
		t.Fatalf("Error loading tasks: %v", err)
	}

	if len(loadedTasks) != len(tasks) {
		t.Errorf("Expected %d tasks, but got %d.", len(tasks), len(loadedTasks))
	}

	for i, task := range tasks {
		if loadedTasks[i].ID != task.ID || loadedTasks[i].UUID != task.UUID || loadedTasks[i].Description != task.Description {
			t.Errorf("Loaded task %d does not match expected task. Got %+v, want %+v", i, loadedTasks[i], task)
		}
	}
}

func TestSaveTasks(t *testing.T) {
	// Clean up the file if it exists before testing
	defer os.Remove(tasksJSONFile)

	// Prepare sample tasks
	tasks := []Task{
		{ID: 1, UUID: "uuid1", Description: "Task 1"},
		{ID: 2, UUID: "uuid2", Description: "Task 2"},
	}

	// Save tasks to the file
	if err := SaveTasks(tasks); err != nil {
		t.Fatalf("Error saving tasks: %v", err)
	}

	// Verify that the file content is correct
	file, err := os.Open(tasksJSONFile)
	if err != nil {
		t.Fatalf("Error opening file: %v", err)
	}
	defer file.Close()

	var savedTasks []Task
	if err := json.NewDecoder(file).Decode(&savedTasks); err != nil {
		t.Fatalf("Error decoding file content: %v", err)
	}

	if len(savedTasks) != len(tasks) {
		t.Errorf("Expected %d tasks, but got %d.", len(tasks), len(savedTasks))
	}

	for i, task := range tasks {
		if savedTasks[i].ID != task.ID || savedTasks[i].UUID != task.UUID || savedTasks[i].Description != task.Description {
			t.Errorf("Saved task %d does not match expected task. Got %+v, want %+v", i, savedTasks[i], task)
		}
	}
}

package taskstorage

import (
	"encoding/json"
	"fmt"
	"os"
	"sync"
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

var (
	filePath = "tasks.json"
	mu       sync.Mutex
)

// ReadTasks reads the tasks from the JSON file
func ReadTasks() ([]Task, error) {
	mu.Lock()
	defer mu.Unlock()

	// Check if file exists
	if _, err := os.Stat(filePath); os.IsNotExist(err) {
		return nil, nil // No tasks found
	}

	file, err := os.ReadFile(filePath)
	if err != nil {
		return nil, fmt.Errorf("error reading file: %v", err)
	}

	var tasks []Task
	if err := json.Unmarshal(file, &tasks); err != nil {
		return nil, fmt.Errorf("error unmarshalling JSON: %v", err)
	}

	return tasks, nil
}

// WriteTasks writes the tasks to the JSON file
func WriteTasks(tasks []Task) error {
	mu.Lock()
	defer mu.Unlock()

	jsonData, err := json.MarshalIndent(tasks, "", "  ")
	if err != nil {
		return fmt.Errorf("error marshalling JSON: %v", err)
	}

	if err := os.WriteFile(filePath, jsonData, 0644); err != nil {
		return fmt.Errorf("error writing file: %v", err)
	}

	return nil
}

// AddTask adds a new task to the tasks file
func AddTask(newTask Task) error {
	tasks, err := ReadTasks()
	if err != nil {
		return err
	}

	newTask.ID = len(tasks) + 1
	tasks = append(tasks, newTask)

	return WriteTasks(tasks)
}

// UpdateTask updates a task in the tasks file
func UpdateTask(updatedTask Task) error {
	tasks, err := ReadTasks()
	if err != nil {
		return err
	}

	for i, task := range tasks {
		if task.ID == updatedTask.ID {
			tasks[i] = updatedTask
			return WriteTasks(tasks)
		}
	}

	return fmt.Errorf("task with ID %d not found", updatedTask.ID)
}

// DeleteTask deletes a task from the tasks file
func DeleteTask(taskID int) error {
	tasks, err := ReadTasks()
	if err != nil {
		return err
	}

	var updatedTasks []Task
	found := false
	for _, task := range tasks {
		if task.ID != taskID {
			updatedTasks = append(updatedTasks, task)
		} else {
			found = true
		}
	}

	if !found {
		return fmt.Errorf("task with ID %d not found", taskID)
	}

	return WriteTasks(updatedTasks)
}

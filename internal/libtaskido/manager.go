package libtaskido

import (
	"fmt"
	"time"
)

// searchByID searches for a task by ID and returns its index.
func searchByID(inputID int, taskList []Task) (int, error) {
	for i := range taskList {
		if taskList[i].ID == inputID {
			return i, nil
		}
	}
	return -1, fmt.Errorf("task ID %d not found", inputID)
}

// findNextID finds the smallest unused ID.
func findNextID(taskList []Task) int {
	idMap := make(map[int]bool)
	for _, task := range taskList {
		idMap[task.ID] = true
	}
	
	// Find the smallest ID not in idMap
	for i := 1; ; i++ {
		if !idMap[i] {
			return i
		}
	}
}

// AddTask adds a new task to the task list.
func AddTask(inputTask []string, taskList []Task) []Task {
	// Parse input task
	newTask, err := ParseNewTask(inputTask)
	if err != nil {
		fmt.Printf("error parsing task %v\n", err)
		return taskList
	}

	// Assign new ID
	newTask.ID = findNextID(taskList)
	taskList = append(taskList, newTask)

	return taskList
}

// MarkComplete marks a task as complete.
func MarkComplete(inputID int, taskList []Task) []Task {
	id, err := searchByID(inputID, taskList)
	if err != nil {
		fmt.Printf("error searching ID %v\n", err)
		return taskList
	}
	taskList[id].Completed = true
	taskList[id].CompletedDate = time.Now().Format("2006-01-02")

	return taskList
}

// MarkUncomplete marks a task as uncompleted.
func MarkUncomplete(inputID int, taskList []Task) []Task {
	id, err := searchByID(inputID, taskList)
	if err != nil {
		fmt.Printf("error searching ID %v\n", err)
		return taskList
	}
	taskList[id].Completed = false
	taskList[id].CompletedDate = ""

	return taskList
}

// MarkArchive marks a task as archived.
func MarkArchive(inputID int, taskList []Task) []Task {
	id, err := searchByID(inputID, taskList)
	if err != nil {
		fmt.Printf("error searching ID %v\n", err)
		return taskList
	}
	taskList[id].Archived = true

	return taskList
}

// MarkUnarchive marks a task as unarchived.
func MarkUnarchive(inputID int, taskList []Task) []Task {
	id, err := searchByID(inputID, taskList)
	if err != nil {
		fmt.Printf("error searching ID %v\n", err)
		return taskList
	}
	taskList[id].Archived = false

	return taskList
}

// DeleteTask deletes a task from the task list.
func DeleteTask(inputID int, taskList []Task) []Task {
	id, err := searchByID(inputID, taskList)
	if err != nil {
		fmt.Printf("error searching ID %v\n", err)
		return taskList
	}
	taskList = append(taskList[:id], taskList[id+1:]...)

	return taskList
}

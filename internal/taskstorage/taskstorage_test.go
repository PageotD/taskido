package taskstorage

import (
	"encoding/json"
	"os"
	"testing"
	"taskido/internal/taskmodel"
)

const testFilePath = "test_tasks.json"

func setupTestFile(t *testing.T) {
	t.Helper()
	// Crée un fichier avec des données spécifiques pour les tests
	initialData := []taskmodel.Task{
		{
			ID:            1,
			UUID:          "dd556cc5-ff74-4a4f-857f-2bcefac18e47",
			Subject:       "develop some tests for @taskstorage",
			Projects:      []string{"taskido"},
			Contexts:      []string{"taskstorage"},
			Due:           "2024-08-30",
			Completed:     true,
			CompletedDate: "2024-08-28",
			Archived:      false,
			Priority:    false,
			Notes:         nil,
		},
	}

	jsonData, err := json.MarshalIndent(initialData, "", "  ")
	if err != nil {
		t.Fatalf("Error marshalling initial data: %v", err)
	}

	if err := os.WriteFile(testFilePath, jsonData, 0644); err != nil {
		t.Fatalf("Error creating or resetting test file: %v", err)
	}
}

func teardownTestFile(t *testing.T) {
	t.Helper()
	// Supprime le fichier de test après les tests
	if err := os.Remove(testFilePath); err != nil && !os.IsNotExist(err) {
		t.Fatalf("Error removing test file: %v", err)
	}
}

func TestReadTasks(t *testing.T) {
	setupTestFile(t)
	defer teardownTestFile(t)

	tasks, err := ReadTasks()
	if err != nil {
		t.Fatalf("Error reading tasks: %v", err)
	}

	if len(tasks) != 0 {
		t.Errorf("Expected no tasks, got %d", len(tasks))
	}

	for _, task := range tasks {
		t.Run("task.ID", func(t *testing.T) {
			if task.ID != 1 {
				t.Errorf("task.ID %d want 1", task.ID)
			}
		})
		t.Run("task.UUID", func(t *testing.T) {
			if task.UUID != "dd556cc5-ff74-4a4f-857f-2bcefac18e47" {
				t.Errorf("task.UUID %s want 'dd556cc5-ff74-4a4f-857f-2bcefac18e47'", task.UUID)
			}
		})
		t.Run("task.Subject", func(t *testing.T) {
			if task.Subject != "develop some tests for @taskstorage" {
				t.Errorf("task.Subject %s want 'develop some tests for @taskstorage'", task.Subject)
			}
		})
		t.Run("task.Projects", func(t *testing.T) {
			if task.Projects[0] != "taskido" {
				t.Errorf("task.Projects %s want 'taskido'", task.Projects[0])
			}
		})
		t.Run("task.Contexts", func(t *testing.T) {
			if task.Contexts[0] != "taskstorage" {
				t.Errorf("task.Contexts %s want 'taskstorage'", task.Contexts[0])
			}
		})
		t.Run("task.Due", func(t *testing.T) {
			if task.Due != "2024-08-30" {
				t.Errorf("task.Due %s want '2024-08-30'", task.Due)
			}
		})
		t.Run("task.Completed", func(t *testing.T) {
			if task.Completed != true {
				t.Errorf("task.Completed %t want 'true'", task.Completed)
			}
		})
		t.Run("task.CompletedDate", func(t *testing.T) {
			if task.CompletedDate != "2024-08-28" {
				t.Errorf("task.Completed %s want '2024-08-28'", task.CompletedDate)
			}
		})
	}
}

func TestAddTask(t *testing.T) {
	setupTestFile(t)
	defer teardownTestFile(t)

	initialTask := taskmodel.Task{
		ID:      1,
		UUID:    "uuid1",
		Subject: "Initial Task",
	}

	if err := WriteTasks([]taskmodel.Task{initialTask}); err != nil {
		t.Fatalf("Error writing initial task: %v", err)
	}

	newTask := taskmodel.Task{
		UUID:    "uuid2",
		Subject: "New Task",
	}

	if err := AddTask(newTask); err != nil {
		t.Fatalf("Error adding task: %v", err)
	}

	tasks, err := ReadTasks()
	if err != nil {
		t.Fatalf("Error reading tasks: %v", err)
	}

	if len(tasks) != 2 || tasks[1].Subject != "New Task" {
		t.Errorf("Expected task with Subject 'New Task', got %v", tasks)
	}
}

func TestUpdateTask(t *testing.T) {
	setupTestFile(t)
	defer teardownTestFile(t)

	initialTask := taskmodel.Task{
		ID:      1,
		UUID:    "uuid1",
		Subject: "Initial Task",
	}

	if err := WriteTasks([]taskmodel.Task{initialTask}); err != nil {
		t.Fatalf("Error writing initial task: %v", err)
	}

	updatedTask := taskmodel.Task{
		ID:      1,
		UUID:    "uuid1",
		Subject: "Updated Task",
	}

	if err := UpdateTask(updatedTask); err != nil {
		t.Fatalf("Error updating task: %v", err)
	}

	tasks, err := ReadTasks()
	if err != nil {
		t.Fatalf("Error reading tasks: %v", err)
	}

	if len(tasks) != 1 || tasks[0].Subject != "Updated Task" {
		t.Errorf("Expected updated task with Subject 'Updated Task', got %v", tasks)
	}
}

func TestDeleteTask(t *testing.T) {
	setupTestFile(t)
	defer teardownTestFile(t)

	task := taskmodel.Task{
		ID:      1,
		UUID:    "uuid1",
		Subject: "Task to Delete",
	}

	if err := WriteTasks([]taskmodel.Task{task}); err != nil {
		t.Fatalf("Error writing task: %v", err)
	}

	if err := DeleteTask(1); err != nil {
		t.Fatalf("Error deleting task: %v", err)
	}

	tasks, err := ReadTasks()
	if err != nil {
		t.Fatalf("Error reading tasks: %v", err)
	}

	if len(tasks) != 0 {
		t.Errorf("Expected no tasks after deletion, got %d", len(tasks))
	}
}

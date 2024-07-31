package libtaskido

import (
	"testing"
)

func setUpTestTaskList(t *testing.T) []Task {
	t.Helper()
	initialData := []Task{
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
		{
			ID:            3,
			UUID:          "dd556cc5-ff74-4a4f-857f-2bcefac18e48",
			Subject:       "develop some @stuff",
			Projects:      []string{"taskido"},
			Contexts:      []string{"stuff"},
			Due:           "2024-08-30",
			Completed:     true,
			CompletedDate: "2024-08-28",
			Archived:      false,
			Priority:    false,
			Notes:         nil,
		},
		{
			ID:            4,
			UUID:          "dd556cc5-ff74-4a4f-857f-2bcefac18e58",
			Subject:       "test some @stuff",
			Projects:      []string{"taskido"},
			Contexts:      []string{"stuff"},
			Due:           "2024-08-30",
			Completed:     true,
			CompletedDate: "2024-08-28",
			Archived:      false,
			Priority:    false,
			Notes:         nil,
		},
	}
	return initialData
}

func TestSearchByIDValid (t *testing.T) {
	taskList := setUpTestTaskList(t)
	ID, _ := searchByID(3, taskList)
	if ID != 1 {
		t.Errorf("searchByID want 3 get %d", ID)
	}
}

func TestSearchByIDInvalid (t *testing.T) {
	taskList := setUpTestTaskList(t)
	ID, err := searchByID(2, taskList)
	if err == nil {
		t.Errorf("searchByID want ID nil get %d", ID)
	}
}

func TestFindNextID (t *testing.T) {
	taskList := setUpTestTaskList(t)
	ID := findNextID(taskList)
	if ID != 2 {
		t.Errorf("findNextID want 2 get %d", ID)
	}
}

func TestAddTask (t *testing.T) {
	taskList := setUpTestTaskList(t)
	inputTask := []string{"+taskido", "develop", "some", "@stuff", "due:2024-07-31"}
	newTaskList := AddTask(inputTask, taskList)
	if len(newTaskList) != 4 {
		t.Errorf("AddTask len: want 4 get %d", len(newTaskList))
	}
	if newTaskList[3].ID != 2 {
		t.Errorf("AddTask want taskList[3].ID = 2 get %d", newTaskList[3].ID)
	}
	if newTaskList[3].Projects[0] != "+taskido" {
		t.Errorf("AddTask want taskList[3].Projects[0] = '+taskido' get %s", newTaskList[3].Projects[0])
	}
	if newTaskList[3].Contexts[0] != "@stuff" {
		t.Errorf("AddTask want taskList[3].Contexts[0] = '@stuff' get %s", newTaskList[3].Contexts[0])
	}
}

func TestDeleteTask (t *testing.T) {
	taskList := setUpTestTaskList(t)
	inputID := 1
	newTaskList := DeleteTask(inputID, taskList)
	if len(newTaskList) != 2 {
		t.Errorf("DeleteTaskwant len(taskList)=2 get %d", len(newTaskList))
	}
	if newTaskList[0].ID != 3 {
		t.Errorf("DeleteTaskwant newTaskList[0].ID = 3 get %d", newTaskList[0].ID)
	}
}
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
			Description:   "develop some tests for @taskstorage",
			Projects:      []string{"taskido"},
			Contexts:      []string{"taskstorage"},
			Due:           "2024-08-30",
			Status:        "pending",
			Priority:      0,
			CreatedAt:     "2024-08-30 15:04:05",
		},
		{
			ID:            3,
			UUID:          "dd556cc5-ff74-4a4f-857f-2bcefac18e48",
			Description:   "develop some @stuff",
			Projects:      []string{"taskido"},
			Contexts:      []string{"stuff"},
			Due:           "2024-08-30",
			Status:        "completed",
			Priority:      0,
			CreatedAt:     "2024-08-30 15:04:05",
		},
		{
			ID:            4,
			UUID:          "dd556cc5-ff74-4a4f-857f-2bcefac18e58",
			Description:   "test some @stuff",
			Projects:      []string{"taskido"},
			Contexts:      []string{"stuff"},
			Due:           "2024-08-30",
			Status:        "archived",
			Priority:      0,
			CreatedAt:     "2024-08-30 15:04:05",
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

func TestMarkComplete (t *testing.T) {
	taskList := setUpTestTaskList(t)
	inputID := 1
	newTaskList := MarkComplete(inputID, taskList)
	if newTaskList[0].Status != "completed" {
		t.Errorf("MarkComplete want status == completed")
	}

}

func TestMarkUncomplete (t *testing.T) {
	taskList := setUpTestTaskList(t)
	inputID := 3
	newTaskList := MarkUncomplete(inputID, taskList)
	if newTaskList[1].Status != "pending" {
		t.Errorf("MarkUncomplete want status == pending")
	}

}

func TestMarkArchive (t *testing.T) {
	taskList := setUpTestTaskList(t)
	inputID := 3
	newTaskList := MarkArchive(inputID, taskList)
	if newTaskList[1].Status != "archived" {
		t.Errorf("MarkArchive want status == archived")
	}

}

func TestMarkUnarchive (t *testing.T) {
	taskList := setUpTestTaskList(t)
	inputID := 4
	newTaskList := MarkUnarchive(inputID, taskList)
	if newTaskList[2].Status != "pending" {
		t.Errorf("MarkUnarchive want status == pending")
	}

}
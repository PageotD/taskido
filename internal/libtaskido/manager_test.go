package libtaskido

import (
	"testing"
	"time"
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
			Completed:     false,
			CompletedDate: "",
			Archived:      false,
			Priority:      0,
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
			Priority:      0,
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
			Archived:      true,
			Priority:      0,
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

func TestMarkComplete (t *testing.T) {
	taskList := setUpTestTaskList(t)
	inputID := 1
	newTaskList := MarkComplete(inputID, taskList)
	if newTaskList[0].Completed != true {
		t.Errorf("MarkComplete want complete = true")
	}
	if newTaskList[0].CompletedDate != time.Now().Format("2006-01-02") {
		t.Errorf("MarkComplete completedDate want %s get %s", time.Now().Format("2006-01-02"), newTaskList[0].CompletedDate)
	}

}

func TestMarkUncomplete (t *testing.T) {
	taskList := setUpTestTaskList(t)
	inputID := 3
	newTaskList := MarkUncomplete(inputID, taskList)
	if newTaskList[1].Completed != false {
		t.Errorf("MarkUncomplete want completed = false")
	}
	if newTaskList[1].CompletedDate != "" {
		t.Errorf("MarkUncomplete completedDate want empty string get %s", newTaskList[0].CompletedDate)
	}

}

func TestMarkArchive (t *testing.T) {
	taskList := setUpTestTaskList(t)
	inputID := 3
	newTaskList := MarkArchive(inputID, taskList)
	if newTaskList[1].Archived != true {
		t.Errorf("MarkArchive want archived = true")
	}

}

func TestMarkUnarchive (t *testing.T) {
	taskList := setUpTestTaskList(t)
	inputID := 4
	newTaskList := MarkUnarchive(inputID, taskList)
	if newTaskList[2].Archived != false {
		t.Errorf("MarkUnarchive want archived = false")
	}

}
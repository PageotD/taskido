package main

import (
	"flag"
	"fmt"
	"strings"
	"time"
	"taskido/internal/formatter"
	"taskido/internal/taskstorage"
	"github.com/google/uuid"
	"regexp"
)

func main() {
	addFlag := flag.Bool("add", false, "Indicates that the following text should be processed")
	listFlag := flag.Bool("list", false, "List all tasks")
	completedFlag := flag.Int("complete", 0, "Mark a task as complete")
	uncompletedFlag := flag.Int("uncomplete", 0, "un-complete a tag")
	archivedFlag := flag.Int("archive", 0, "Archive a task")
	unarchivedFlag := flag.Int("unarchive", 0, "un-archive a task")

	flag.Parse()

	if *addFlag {
		handleAdd(flag.Args())
	} else if *listFlag {
		handleList()
	} else if *completedFlag != 0 {
		handleComplete(*completedFlag)
	} else if *uncompletedFlag != 0 {
		handleUncomplete(*uncompletedFlag)
	} else if *archivedFlag != 0 {
		handleArchived(*archivedFlag)
	} else if *unarchivedFlag != 0 {
		handleUnarchived(*unarchivedFlag)
	} else {
		fmt.Println("No valid flag provided. Use -a to add a task or -l to list tasks.")
	}
}

func handleAdd(args []string) {
	addText := strings.Join(args, " ")

	projectPattern := regexp.MustCompile(`\+(\S+)`)
	contextPattern := regexp.MustCompile(`@(\S+)`)
	duePattern := regexp.MustCompile(`due:(\d{4}-\d{2}-\d{2})`)

	projectMatches := projectPattern.FindAllStringSubmatch(addText, -1)
	contextMatch := contextPattern.FindStringSubmatch(addText)
	dueMatch := duePattern.FindStringSubmatch(addText)

	taskDescription := addText
	for _, match := range projectMatches {
		taskDescription = strings.Replace(taskDescription, match[0], "", 1)
	}
	if dueMatch != nil {
		taskDescription = strings.Replace(taskDescription, dueMatch[0], "", 1)
	}
	taskDescription = strings.TrimSpace(taskDescription)

	task := taskstorage.Task{
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

	if err := taskstorage.AddTask(task); err != nil {
		fmt.Printf("Error adding task: %v\n", err)
		return
	}

	fmt.Println("Task added to tasks.json")
}

func handleList() {
	tasks, err := taskstorage.ReadTasks()
	if err != nil {
		fmt.Printf("Error reading tasks: %v\n", err)
		return
	}

	// Current tasks
	fmt.Printf("Current:\n")
	for _, task := range tasks {
		if !task.Completed && !task.Archived {
			subjectWithColor := formatter.ApplyColorToSubject(task.Subject)
			projectWithColor := formatter.ApplyColorToProject(task.Projects)
			dueDateWithColor := formatter.ApplyColorToDate(task.Due)
			fmt.Printf("%-4d %-12s %s %s\n", task.ID, dueDateWithColor, projectWithColor, subjectWithColor)
		}
	}

	// Completed tasks
	fmt.Printf("\nCompleted:\n")
	for _, task := range tasks {
		if task.Completed && !task.Archived {
			subjectWithColor := formatter.ApplyColorToSubject(task.Subject)
			projectWithColor := formatter.ApplyColorToProject(task.Projects)
			dueDateWithColor := formatter.ApplyColorToDate(task.Due)
			fmt.Printf("%-4d %-12s %s %s\n", task.ID, dueDateWithColor, projectWithColor, subjectWithColor)
		}
	}

	// Completed tasks
	fmt.Printf("\nArchived:\n")
	for _, task := range tasks {
		if task.Archived {
			subjectWithColor := formatter.ApplyColorToSubject(task.Subject)
			projectWithColor := formatter.ApplyColorToProject(task.Projects)
			dueDateWithColor := formatter.ApplyColorToDate(task.Due)
			fmt.Printf("%-4d %-12s %s %s\n", task.ID, dueDateWithColor, projectWithColor, subjectWithColor)
		}
	}
}

func handleComplete(taskID int) {
	now := time.Now().Format("2006-01-02")
	tasks, err := taskstorage.ReadTasks()
	if err != nil {
		fmt.Printf("Error reading tasks: %v\n", err)
		return
	}

	var taskToUpdate *taskstorage.Task
	for i := range tasks {
		if tasks[i].ID == taskID {
			taskToUpdate = &tasks[i]
			break
		}
	}

	if taskToUpdate == nil {
		fmt.Printf("Task ID %d not found\n", taskID)
		return
	}

	taskToUpdate.Completed = true
	taskToUpdate.CompletedDate = now

	if err := taskstorage.UpdateTask(*taskToUpdate); err != nil {
		fmt.Printf("Error updating task: %v\n", err)
		return
	}

	fmt.Println("Task updated successfully.")
}

func handleUncomplete(taskID int) {
	tasks, err := taskstorage.ReadTasks()
	if err != nil {
		fmt.Printf("Error reading tasks: %v\n", err)
		return
	}

	var taskToUpdate *taskstorage.Task
	for i := range tasks {
		if tasks[i].ID == taskID {
			taskToUpdate = &tasks[i]
			break
		}
	}

	if taskToUpdate == nil {
		fmt.Printf("Task ID %d not found\n", taskID)
		return
	}

	taskToUpdate.Completed = false
	taskToUpdate.CompletedDate = ""

	if err := taskstorage.UpdateTask(*taskToUpdate); err != nil {
		fmt.Printf("Error updating task: %v\n", err)
		return
	}

	fmt.Println("Task updated successfully.")
}


func handleArchived(taskID int) {
	tasks, err := taskstorage.ReadTasks()
	if err != nil {
		fmt.Printf("Error reading tasks: %v\n", err)
		return
	}

	var taskToUpdate *taskstorage.Task
	for i := range tasks {
		if tasks[i].ID == taskID {
			taskToUpdate = &tasks[i]
			break
		}
	}

	if taskToUpdate == nil {
		fmt.Printf("Task ID %d not found\n", taskID)
		return
	}

	taskToUpdate.Archived = true

	if err := taskstorage.UpdateTask(*taskToUpdate); err != nil {
		fmt.Printf("Error updating task: %v\n", err)
		return
	}

	fmt.Println("Task archived successfully.")
}

func handleUnarchived(taskID int) {
	tasks, err := taskstorage.ReadTasks()
	if err != nil {
		fmt.Printf("Error reading tasks: %v\n", err)
		return
	}

	var taskToUpdate *taskstorage.Task
	for i := range tasks {
		if tasks[i].ID == taskID {
			taskToUpdate = &tasks[i]
			break
		}
	}

	if taskToUpdate == nil {
		fmt.Printf("Task ID %d not found\n", taskID)
		return
	}

	taskToUpdate.Archived = false

	if err := taskstorage.UpdateTask(*taskToUpdate); err != nil {
		fmt.Printf("Error updating task: %v\n", err)
		return
	}

	fmt.Println("Task archived successfully.")
}


func getMatchValue(matches []string) string {
	if len(matches) > 1 {
		return matches[1]
	}
	return ""
}

func extractMatches(matches [][]string) []string {
	var result []string
	for _, match := range matches {
		result = append(result, getMatchValue(match))
	}
	return result
}

// internal/taskmanager/taskmanager.go
package taskmanager

import (
    "fmt"
    "strings"
    "time"
    "taskido/internal/taskmodel"
    "taskido/internal/taskstorage"
	"taskido/internal/formatter"
    "github.com/google/uuid"
)

// HandleAdd adds a new task
func HandleAdd(args []string) (taskmodel.Task, error) {
    addText := strings.Join(args, " ")

    projectMatches := ExtractProjects(addText)
    contextMatches := ExtractContexts(addText)
    dueMatch := ExtractDue(addText)

    taskDescription := addText
    for _, match := range projectMatches {
        taskDescription = strings.Replace(taskDescription, match, "", 1)
    }
    if dueMatch != "" {
        taskDescription = strings.Replace(taskDescription, "due:"+dueMatch, "", 1)
    }
    taskDescription = strings.TrimSpace(taskDescription)

    task := taskmodel.Task{
        UUID:          uuid.NewString(),
        Subject:       taskDescription,
        Projects:      projectMatches, 
        Contexts:      contextMatches,
        Due:           dueMatch,
        Completed:     false,
        CompletedDate: "",
        Archived:      false,
        Priority:    false,
        Notes:         nil,
    }

    return task, nil
}

// HandleList lists all tasks
func HandleList() {
    tasks, err := taskstorage.ReadTasks()
    if err != nil {
        fmt.Printf("Error reading tasks: %v\n", err)
        return
    }

    fmt.Printf("\n\033[4mCurrent:\033[0m\n\n")
    for _, task := range tasks {
        if !task.Completed && !task.Archived {
            fmt.Printf("%-4d %-12s %s %s\n", task.ID, formatter.ApplyColorToDate(task.Due), formatter.ApplyColorToProject(task.Projects), formatter.ApplyColorToSubject(task.Subject))
        }
    }

    fmt.Printf("\n\033[4mCompleted:\033[0m\n\n")
    for _, task := range tasks {
        if task.Completed && !task.Archived {
            fmt.Printf("%-4d %-12s %s %s\n", task.ID, formatter.ApplyColorToDate(task.Due), formatter.ApplyColorToProject(task.Projects), formatter.ApplyColorToSubject(task.Subject))
        }
    }

    fmt.Printf("\n\033[4mArchived:\033[0m\n\n")
    for _, task := range tasks {
        if task.Archived {
            fmt.Printf("%-4d %-12s %s %s\n", task.ID, formatter.ApplyColorToDate(task.Due), formatter.ApplyColorToProject(task.Projects), formatter.ApplyColorToSubject(task.Subject))
        }
    }

    fmt.Printf("\n")
}

// HandleComplete marks a task as complete
func HandleComplete(taskID int) (*taskmodel.Task, error) {
    now := time.Now().Format("2006-01-02")
    tasks, err := taskstorage.ReadTasks()
    if err != nil {
        return nil, fmt.Errorf("error reading tasks: %w", err)
    }

    var taskToUpdate *taskmodel.Task

    for i := range tasks {
        if tasks[i].ID == taskID {
            taskToUpdate = &tasks[i]
            break
        }
    }

    if taskToUpdate == nil {
        return nil, fmt.Errorf("task ID %d not found", taskID)
    }

    taskToUpdate.Completed = true
    taskToUpdate.CompletedDate = now

    // Update the task in storage if needed
    if err := taskstorage.UpdateTask(*taskToUpdate); err != nil {
        return nil, fmt.Errorf("error updating task: %w", err)
    }

    return taskToUpdate, nil
}

// HandleUncomplete marks a task as uncompleted
func HandleUncomplete(taskID int) (*taskmodel.Task, error) {

    tasks, err := taskstorage.ReadTasks()

    if err != nil {
        return nil, fmt.Errorf("Error reading tasks: %v\n", err)
    }

    var taskToUpdate *taskmodel.Task

    for i := range tasks {
        if tasks[i].ID == taskID {
            taskToUpdate = &tasks[i]
            break
        }
    }

    if taskToUpdate == nil {
        return nil, fmt.Errorf("task ID %d not found", taskID)
    }

    taskToUpdate.Completed = false
    taskToUpdate.CompletedDate = ""

    // Update the task in storage if needed
    if err := taskstorage.UpdateTask(*taskToUpdate); err != nil {
        return nil, fmt.Errorf("error updating task: %w", err)
    }

    return taskToUpdate, nil

}

// HandleArchived archives a task
func HandleArchived(taskID int) (*taskmodel.Task, error)  {

    tasks, err := taskstorage.ReadTasks()

    if err != nil {
        return nil, fmt.Errorf("Error reading tasks: %v\n", err)
    }

    var taskToUpdate *taskmodel.Task

    for i := range tasks {
        if tasks[i].ID == taskID {
            taskToUpdate = &tasks[i]
            break
        }
    }

    if taskToUpdate == nil {
        return nil, fmt.Errorf("task ID %d not found", taskID)
    }

    taskToUpdate.Archived = true

    if err := taskstorage.UpdateTask(*taskToUpdate); err != nil {
        return nil, fmt.Errorf("Error updating task: %v\n", err)
    }

    // Update the task in storage if needed
    if err := taskstorage.UpdateTask(*taskToUpdate); err != nil {
        return nil, fmt.Errorf("error updating task: %w", err)
    }

    return taskToUpdate, nil
}

// HandleUnarchived unarchives a task
func HandleUnarchived(taskID int) (*taskmodel.Task, error) {

    tasks, err := taskstorage.ReadTasks()

    if err != nil {
        return nil, fmt.Errorf("Error reading tasks: %v\n", err)
    }

    var taskToUpdate *taskmodel.Task

    for i := range tasks {
        if tasks[i].ID == taskID {
            taskToUpdate = &tasks[i]
            break
        }
    }

    if taskToUpdate == nil {
        return nil, fmt.Errorf("task ID %d not found", taskID)
    }

    taskToUpdate.Archived = true

    if err := taskstorage.UpdateTask(*taskToUpdate); err != nil {
        fmt.Printf("Error updating task: %v\n", err)
        return nil, fmt.Errorf("Error updating task: %v\n", err)
    }

    // Update the task in storage if needed
    if err := taskstorage.UpdateTask(*taskToUpdate); err != nil {
        return nil, fmt.Errorf("error updating task: %w", err)
    }

    return taskToUpdate, nil
}

// HandleDelete deletes a task
func HandleDelete(taskID int) (*taskmodel.Task, error) {
    tasks, err := taskstorage.ReadTasks()
    if err != nil {
        fmt.Printf("Error reading tasks: %v\n", err)
        return
    }

    var taskToDelete *taskmodel.Task
    for i := range tasks {
        if tasks[i].ID == taskID {
            taskToDelete = &tasks[i]
            break
        }
    }

    if taskToDelete == nil {
        fmt.Printf("Task ID %d not found\n", taskID)
        return
    }

    if err := taskstorage.DeleteTask(taskID); err != nil {
        fmt.Printf("Error deleting task: %v\n", err)
        return
    }

    fmt.Println("Task deleted successfully.")
}
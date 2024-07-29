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
func HandleComplete(taskID int) {
    now := time.Now().Format("2006-01-02")
    tasks, err := taskstorage.ReadTasks()
    if err != nil {
        fmt.Printf("Error reading tasks: %v\n", err)
        return
    }

    var taskToUpdate *taskmodel.Task
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

// HandleUncomplete marks a task as uncompleted
func HandleUncomplete(taskID int) {
    tasks, err := taskstorage.ReadTasks()
    if err != nil {
        fmt.Printf("Error reading tasks: %v\n", err)
        return
    }

    var taskToUpdate *taskmodel.Task
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

// HandleArchived archives a task
func HandleArchived(taskID int) {
    tasks, err := taskstorage.ReadTasks()
    if err != nil {
        fmt.Printf("Error reading tasks: %v\n", err)
        return
    }

    var taskToUpdate *taskmodel.Task
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

// HandleUnarchived unarchives a task
func HandleUnarchived(taskID int) {
    tasks, err := taskstorage.ReadTasks()
    if err != nil {
        fmt.Printf("Error reading tasks: %v\n", err)
        return
    }

    var taskToUpdate *taskmodel.Task
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

    fmt.Println("Task unarchived successfully.")
}

// HandleDelete deletes a task
func HandleDelete(taskID int) {
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
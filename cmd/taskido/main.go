package main

import (
	"flag"
	"fmt"
	"taskido/internal/libtaskido"
)

// printHelp displays a help message that provides usage instructions for the command-line tool.
func printHelp() {
	fmt.Println("Usage:")
	fmt.Println("  -add       : Adds a new task with the following text")
	fmt.Println("  -list      : Lists all tasks")
	fmt.Println("  -complete  : Marks a task as complete, requires task ID")
	fmt.Println("  -uncomplete: Marks a task as uncompleted, requires task ID")
	fmt.Println("  -archive   : Archives a task, requires task ID")
	fmt.Println("  -unarchive : Unarchives a task, requires task ID")
	fmt.Println("  -delete    : Deletes a task, requires task ID")
	fmt.Println("  -help      : Displays this help message")
}

// initTaskido initializes the task management system by ensuring that the necessary file exists,
// and then loads the current list of tasks. If the file does not exist, it creates it.
func initTaskido () []libtaskido.Task {
	if !libtaskido.EnsureFileExists() {
		err := libtaskido.InitializeFile()
		if err != nil {
			fmt.Printf("error creating file")
		}
	}
	taskList, _ := libtaskido.LoadTasks()
	return taskList
}

// main parses command-line flags to determine which operation to perform (e.g., add, complete, delete),
// then executes the corresponding function to manipulate the task list. It also handles the loading and saving
// of tasks.
func main() {

	// Define command-line flags
	addFlag := flag.Bool("add", false, "Adds a new task")
	listFlag := flag.Bool("list", false, "Lists all tasks")
	completedFlag := flag.Int("complete", 0, "Marks a task as complete, requires task ID")
	uncompletedFlag := flag.Int("uncomplete", 0, "Marks a task as uncompleted, requires task ID")
	archivedFlag := flag.Int("archive", 0, "Archives a task, requires task ID")
	unarchivedFlag := flag.Int("unarchive", 0, "Unarchives a task, requires task ID")
	deleteFlag := flag.Int("delete", 0, "Deletes a task, requires task ID")
	helpFlag := flag.Bool("help", false, "Display help message")

	// Parse command-line flags
	flag.Parse()

	// Initialize task list
	taskList := initTaskido()
	taskModification := false

	// Execute command based on flags
    switch {
	// Add a new task
    case *addFlag:
		taskList = libtaskido.AddTask(flag.Args(), taskList)
		taskModification = true
	// Marks a task as complete
    case *completedFlag != 0:
        taskList = libtaskido.MarkComplete(*completedFlag, taskList)
		taskModification = true
	// Marks a task as uncompleted
    case *uncompletedFlag != 0:
        taskList = libtaskido.MarkUncomplete(*uncompletedFlag, taskList)
		taskModification = true
	// Marks a task as archived
    case *archivedFlag != 0:
        taskList = libtaskido.MarkArchive(*archivedFlag, taskList)
		taskModification = true
	// Marks a task as unarchived
    case *unarchivedFlag != 0:
        taskList = libtaskido.MarkUnarchive(*unarchivedFlag, taskList)
		taskModification = true
	// Delete a task
    case *deleteFlag != 0:
        taskList = libtaskido.DeleteTask(*deleteFlag, taskList)
		taskModification = true
	// List all tasks
	case *listFlag:
		if len(flag.Args()) > 0 && flag.Args()[0] == "projects"  {
			libtaskido.PrintTaskListByProjects(taskList)
		} else if len(flag.Args()) > 0 && flag.Args()[0] == "contexts"  {
			libtaskido.PrintTaskListByContexts(taskList)
		}else {
			libtaskido.PrintTaskList(taskList)
		}
	// Print help
	case *helpFlag:
		printHelp()
	// Default 
    default:
        fmt.Println("No valid flag provided. Use --help.")
    }

	if taskModification {
		err := libtaskido.SaveTasks(taskList)
		if err != nil {
			fmt.Printf("error during save %v", err)
			return
		}
		fmt.Printf("Tasklist updated succesfully.\n")
	}
}

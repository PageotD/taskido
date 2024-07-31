package main

import (
	"flag"
	"fmt"
	"taskido/internal/libtaskido"
)

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

func main() {

	addFlag := flag.Bool("add", false, "Adds a new task with the following text")
	listFlag := flag.Bool("list", false, "Lists all tasks")
	completedFlag := flag.Int("complete", 0, "Marks a task as complete, requires task ID")
	uncompletedFlag := flag.Int("uncomplete", 0, "Marks a task as uncompleted, requires task ID")
	archivedFlag := flag.Int("archive", 0, "Archives a task, requires task ID")
	unarchivedFlag := flag.Int("unarchive", 0, "Unarchives a task, requires task ID")
	deleteFlag := flag.Int("delete", 0, "Deletes a task, requires task ID")
	helpFlag := flag.Bool("help", false, "Display help message")

	flag.Parse()

	taskList := initTaskido()

    switch {
    case *addFlag:
		taskList = libtaskido.AddTask(flag.Args(), taskList)
		err := libtaskido.SaveTasks(taskList)
		if err != nil {
			fmt.Printf("error during save %v", err)
			return
		}
		fmt.Printf("Task added succesfully.\n")
    case *completedFlag != 0:
        taskList = libtaskido.MarkComplete(*completedFlag, taskList)
		err := libtaskido.SaveTasks(taskList)
		if err != nil {
			fmt.Printf("error during save %v", err)
			return
		}
		fmt.Printf("Task updated succesfully.\n")
    case *uncompletedFlag != 0:
        taskList = libtaskido.MarkUncomplete(*uncompletedFlag, taskList)
		err := libtaskido.SaveTasks(taskList)
		if err != nil {
			fmt.Printf("error during save %v", err)
			return
		}
		fmt.Printf("Task updated succesfully.\n")
    case *archivedFlag != 0:
        taskList = libtaskido.MarkArchive(*archivedFlag, taskList)
		err := libtaskido.SaveTasks(taskList)
		if err != nil {
			fmt.Printf("error during save %v", err)
			return
		}
		fmt.Printf("Task updated succesfully.\n")
    case *unarchivedFlag != 0:
        taskList = libtaskido.MarkUnarchive(*unarchivedFlag, taskList)
		err := libtaskido.SaveTasks(taskList)
		if err != nil {
			fmt.Printf("error during save %v", err)
			return
		}
		fmt.Printf("Task updated succesfully.\n")
    case *deleteFlag != 0:
        taskList = libtaskido.DeleteTask(*deleteFlag, taskList)
		err := libtaskido.SaveTasks(taskList)
		if err != nil {
			fmt.Printf("error during save %v", err)
			return
		}
		fmt.Printf("Task deleted succesfully.\n")
	case *listFlag:
		libtaskido.PrintTaskList(taskList)
	case *helpFlag:
		printHelp()
    default:
        fmt.Println("No valid flag provided. Use -a to add a task or -l to list tasks.")
    }
}

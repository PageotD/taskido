package main

import (
	"flag"
	"fmt"
	"taskido/internal/taskmanager"
	"taskido/internal/taskstorage"
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

	if *helpFlag {
		printHelp()
		return
	}

	if *addFlag {

		task, err := taskmanager.HandleAdd(flag.Args())
		if err != nil {
			fmt.Printf("Error adding task: %v\n", err)
			return
		}
		if err := taskstorage.AddTask(task); err != nil {
			fmt.Printf("Error adding task: %v\n", err)
			return
		}
		fmt.Println("Task added to tasks.json")

	} else if *completedFlag != 0 {

		tasks, _ := taskmanager.HandleComplete(*completedFlag)
		if tasks != nil {
			if err := taskstorage.UpdateTask(*tasks); err != nil {
				fmt.Printf("Error updating task: %v\n", err)
				return
			}
			fmt.Println("Task updated successfully.")
		}

	} else if *uncompletedFlag != 0 {

		tasks, _ := taskmanager.HandleUncomplete(*uncompletedFlag)
		if tasks != nil {
			if err := taskstorage.UpdateTask(*tasks); err != nil {
				fmt.Printf("Error updating task: %v\n", err)
				return
			}
			fmt.Println("Task updated successfully.")
		}

	} else if *archivedFlag != 0 {

		tasks, _ := taskmanager.HandleArchived(*archivedFlag)
		if tasks != nil {
			if err := taskstorage.UpdateTask(*tasks); err != nil {
				fmt.Printf("Error updating task: %v\n", err)
				return
			}
			fmt.Println("Task updated successfully.")
		}

	} else if *unarchivedFlag != 0 {

		tasks, _ := taskmanager.HandleUnarchived(*archivedFlag)
		if tasks != nil {
			if err := taskstorage.UpdateTask(*tasks); err != nil {
				fmt.Printf("Error updating task: %v\n", err)
				return
			}
			fmt.Println("Task updated successfully.")
		}

	} else if *deleteFlag != 0 {
		err := taskmanager.HandleDelete(*deleteFlag)
		if err != nil {
			fmt.Printf("%v", err)
		}
	} else if *listFlag {
		taskmanager.HandleList()
	} else {
		fmt.Println("No valid flag provided. Use -help to display the usage information.")
	}
}

package main

import (
	"flag"
	"fmt"
	"taskido/internal/taskmanager"
	"taskido/internal/taskstorage"
)

func main() {
	addFlag := flag.Bool("add", false, "Indicates that the following text should be processed")
	listFlag := flag.Bool("list", false, "List all tasks")
	completedFlag := flag.Int("complete", 0, "Mark a task as complete")
	uncompletedFlag := flag.Int("uncomplete", 0, "un-complete a tag")
	archivedFlag := flag.Int("archive", 0, "Archive a task")
	unarchivedFlag := flag.Int("unarchive", 0, "un-archive a task")
	deleteFlag := flag.Int("delete", 0, "delete a task")

	flag.Parse()

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
	} else if *listFlag {
		taskmanager.HandleList()
	} else if *completedFlag != 0 {
		taskmanager.HandleComplete(*completedFlag)
	} else if *uncompletedFlag != 0 {
		taskmanager.HandleUncomplete(*uncompletedFlag)
	} else if *archivedFlag != 0 {
		taskmanager.HandleArchived(*archivedFlag)
	} else if *unarchivedFlag != 0 {
		taskmanager.HandleUnarchived(*unarchivedFlag)
	} else if *deleteFlag != 0 {
		taskmanager.HandleDelete(*deleteFlag)
	} else {
		fmt.Println("No valid flag provided. Use -a to add a task or -l to list tasks.")
	}
}
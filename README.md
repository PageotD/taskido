# Taskido
Taskido is a lightweight, command-line task management tool written in Go. Designed for simplicity and efficiency, Taskido helps you keep track of your tasks and projects with ease.

## Managing your tasks

### Adding a Task

To add a new task, use the following command:

```
taskido -add [projects] <description> [due date]
```

- **`description`** _(required)_: This is the primary content of the task. The description may include contexts specified as `@context`. Multiple contexts can be included within the description.

- **`projects`** _(optional)_: Specifies one or more projects associated with the task. Projects must be listed at the beginning of the command and are denoted by `+project`. You can include multiple projects, separated by spaces (e.g., `+project1 +project2`).

- **`due date`** _(optional)_: Sets a due date for the task in the format `due:YYYY-MM-DD`. Each task can have a unique due date.

### Listing existing tasks

To list all your tasks, use the following command:
```
taskido -list
```

### Completing/uncompleting a task

To complete an existing task, use the following command:
```
taskido -complete [taskID]
```

To uncomplete an existing task, use the following command:
```
taskido -uncomplete [taskID]
```

### Archiving/unarchiving a task

To archive an existing task, use the following command:
```
taskido -archive [taskID]
```

To unarchive an existing task, use the following command:
```
taskido -unarchive [taskID]
```

### Deleting a task

To delete a task, use the following command:
```
taskido -delete [taskID]
```

package libtaskido

import (
	"fmt"
	"strings"
	"regexp"
	"github.com/google/uuid"
)

func ParseNewTask(input []string) (Task, error) {

	inputTask := strings.Join(input, " ")

	// extract projects
	projects := extractProjects(inputTask)
	// extract contexts
	contexts := extractContexts(inputTask)
	// extract due date
	dueDate := extractDueDate(inputTask)

	// Clean up description
	taskDescription := inputTask
	for _, match := range projects {
		taskDescription = strings.Replace(taskDescription, match, "", 1)
	}
	if dueDate != "" {
        taskDescription = strings.Replace(taskDescription, "due:"+dueDate, "", 1)
    }
    taskDescription = strings.TrimSpace(taskDescription)

	if taskDescription == "" {
		return Task{}, fmt.Errorf("a task description is needed")
	}

	task := Task{
        UUID:          uuid.NewString(),
        Subject:       taskDescription,
        Projects:      projects, 
        Contexts:      contexts,
        Due:           dueDate,
        Completed:     false,
        CompletedDate: "",
        Archived:      false,
        Priority:    false,
        Notes:         nil,
    }

    return task, nil
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

func extractProjects (input string) []string {

    pattern := regexp.MustCompile(`\+\S+`)
    matches := pattern.FindAllStringSubmatch(input, -1)
    
    return flattenString(matches)
}

func extractContexts (input string) []string {

    pattern := regexp.MustCompile(`\@\S+`)
    matches := pattern.FindAllStringSubmatch(input, -1)
    
    return flattenString(matches)
}

func extractDueDate (input string) string {

    pattern := regexp.MustCompile(`due:(\d{4}-\d{2}-\d{2})`)
    match := pattern.FindStringSubmatch(input)

    if len(match) > 1 {
        return match[1]
    }
    return ""
}

func flattenString (ss [][]string ) []string {
    var result []string
    for _, s := range ss {
        result = append(result, s...)
    }
    return result
}
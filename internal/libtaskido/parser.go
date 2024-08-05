package libtaskido

import (
	"fmt"
    "time"
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
    // extract priority level
    priority := extractPriority(inputTask)

	// Clean up description
	taskDescription := inputTask
	for _, match := range projects {
		taskDescription = strings.Replace(taskDescription, match, "", 1)
	}
	if dueDate != "" {
        taskDescription = strings.Replace(taskDescription, "due:"+dueDate, "", 1)
    }
    // remove priority from input task
    pattern := regexp.MustCompile(`\s*priority:\d+`)
    taskDescription = pattern.ReplaceAllString(inputTask, "")

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
        Priority:      priority,
        CreatedAt:     time.Now().Format("2006-01-02 15:04:05"),
        UpdatedAt:     "",
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

func extractProjects(input string) []string {

    pattern := regexp.MustCompile(`\+\S+`)
    matches := pattern.FindAllStringSubmatch(input, -1)
    
    return flattenString(matches)
}

func extractContexts(input string) []string {

    pattern := regexp.MustCompile(`\@\S+`)
    matches := pattern.FindAllStringSubmatch(input, -1)
    
    return flattenString(matches)
}

func extractDueDate(input string) string {

    pattern := regexp.MustCompile(`due:(\d{4}-\d{2}-\d{2})`)
    match := pattern.FindStringSubmatch(input)

    if len(match) > 1 {
        return match[1]
    }
    return ""
}

// extractPriority extracts and converts the priority level from a string input
func extractPriority(input string) int {
	// Define regular expressions for different priority formats
	patterns := []struct {
		Pattern string
		Value   int
	}{
		{Pattern: `\b1\b`, Value: 1},              // Matches '1'
		{Pattern: `\blow\b`, Value: 1},           // Matches 'low'
		{Pattern: `\b2\b`, Value: 2},              // Matches '2'
		{Pattern: `\bmedium\b`, Value: 2},        // Matches 'medium'
		{Pattern: `\b3\b`, Value: 3},              // Matches '3'
		{Pattern: `\bhigh\b`, Value: 3},           // Matches 'high'
	}

	// Iterate over the patterns and check for matches in the input string
	for _, p := range patterns {
		re := regexp.MustCompile(p.Pattern)
		if re.MatchString(input) {
			return p.Value
		}
	}

	// Return 0 if no match is found
	return 0
}

func flattenString(ss [][]string ) []string {
    var result []string
    for _, s := range ss {
        result = append(result, s...)
    }
    return result
}
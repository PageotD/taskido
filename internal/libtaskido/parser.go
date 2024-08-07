package libtaskido

import (
	"fmt"
	"regexp"
	"strings"
	"time"

	"github.com/google/uuid"
)

// ParseNewTask parses the input strings and creates a new Task.
func ParseNewTask(input []string) (Task, error) {
	inputTask := strings.Join(input, " ")

	// Extract details from the input
	projects := extractProjects(inputTask)
	contexts := extractContexts(inputTask)
	dueDate := extractDueDate(inputTask)
	priority := extractPriority(inputTask)

	// Clean up task description
	taskDescription := inputTask
	for _, match := range projects {
		taskDescription = strings.Replace(taskDescription, match, "", 1)
	}
	if dueDate != "" {
		taskDescription = strings.Replace(taskDescription, "due:"+dueDate, "", 1)
	}
	taskDescription = removePriority(taskDescription)

	taskDescription = strings.TrimSpace(taskDescription)

	if taskDescription == "" {
		return Task{}, fmt.Errorf("a task description is needed")
	}

	// Create and return a new Task
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

// removePriority removes the priority information from the task description.
func removePriority(description string) string {
	pattern := regexp.MustCompile(`\s*priority:(\d+|low|medium|high)\b`)
	return pattern.ReplaceAllString(description, "")
}

// extractProjects extracts project tags from the input string.
func extractProjects(input string) []string {
	pattern := regexp.MustCompile(`\+\S+`)
	matches := pattern.FindAllStringSubmatch(input, -1)
	return flattenString(matches)
}

// extractContexts extracts context tags from the input string.
func extractContexts(input string) []string {
	pattern := regexp.MustCompile(`\@\S+`)
	matches := pattern.FindAllStringSubmatch(input, -1)
	return flattenString(matches)
}

// extractDueDate extracts the due date from the input string.
func extractDueDate(input string) string {
	pattern := regexp.MustCompile(`due:(\d{4}-\d{2}-\d{2})`)
	match := pattern.FindStringSubmatch(input)
	return getMatchValue(match)
}

// extractPriority extracts and converts the priority level from a string input.
func extractPriority(input string) int {
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

// getMatchValue returns the first match value from a slice of matches.
func getMatchValue(matches []string) string {
	if len(matches) > 1 {
		return matches[1]
	}
	return ""
}

// flattenString flattens a slice of string slices into a single slice of strings.
func flattenString(ss [][]string) []string {
	var result []string
	for _, s := range ss {
		result = append(result, s...)
	}
	return result
}
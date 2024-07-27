package taskmanager

import "regexp"

func GetMatchValue(matches []string) string {
    if len(matches) > 1 {
        return matches[1]
    }
    return ""
}

func ExtractMatches(matches [][]string) []string {
    var result []string
    for _, match := range matches {
        result = append(result, GetMatchValue(match))
    }
    return result
}

func ExtractProjects (input string) []string {

    pattern := regexp.MustCompile(`\+\S+`)
    matches := pattern.FindAllStringSubmatch(input, -1)
    
    return flatten_string(matches)
}

func ExtractContexts (input string) []string {

    pattern := regexp.MustCompile(`\@\S+`)
    matches := pattern.FindAllStringSubmatch(input, -1)
    
    return flatten_string(matches)
}

func ExtractDue (input string) string {

    pattern := regexp.MustCompile(`due:(\d{4}-\d{2}-\d{2})`)
    match := pattern.FindStringSubmatch(input)

    if len(match) > 1 {
        return match[1]
    }
    return ""
}

func flatten_string (ss [][]string ) []string {
    var result []string
    for _, s := range ss {
        result = append(result, s...)
    }
    return result
}

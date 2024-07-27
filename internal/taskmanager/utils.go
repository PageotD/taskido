package taskmanager

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

package parsinglogfiles

import (
    "regexp"
    "fmt"
)

func IsValidLine(text string) bool {
    re := regexp.MustCompile(`^\[(TRC|DBG|INF|WRN|ERR|FTL)\]`)
    return re.MatchString(text)
}

func SplitLogLine(text string) []string {
    re := regexp.MustCompile(`<[-~*=]*>`)
    return re.Split(text, -1)
}

func CountQuotedPasswords(lines []string) int {
    count := 0
    re := regexp.MustCompile(`(?i)"[^"]*password[^"]*"`)
    for _, line := range lines {
        if re.MatchString(line) {
            count++
        }
    }
    return count
}

func RemoveEndOfLineText(text string) string {
    re := regexp.MustCompile(`end-of-line\d+`)
    return re.ReplaceAllString(text, "")
}

func TagWithUserName(lines []string) []string {
    re := regexp.MustCompile(`User\s+(\S+)`)
    for i, line := range lines {
        if parts := re.FindStringSubmatch(line); len(parts) > 1 {
            lines[i] = fmt.Sprintf("[USR] %s %s", parts[1], line)
        }
    }
    return lines
}

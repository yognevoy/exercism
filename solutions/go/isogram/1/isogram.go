package isogram

import "strings"

func IsIsogram(word string) bool {
    seen := make(map[rune]bool, len(word))
    runes := []rune(strings.ToLower(word))

    for _, ch := range runes {
        if ch == ' ' || ch == '-' {
            continue
        }
        if seen[ch] {
            return false
        } else {
            seen[ch] = true
        }
    }

    return true
}

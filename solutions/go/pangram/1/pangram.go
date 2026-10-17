package pangram

func IsPangram(input string) bool {
    seen := make(map[rune]bool, 26)
    
    for _, r := range input {
        if r >= 'A' && r <= 'Z' {
            r += 32
        }
        if r >= 'a' && r <= 'z' {
            seen[r] = true
        }
    }
    
    return len(seen) == 26
}

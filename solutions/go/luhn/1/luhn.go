package luhn

func Valid(id string) bool {
    shouldDouble := false
    sum := 0
    count := 0

    for i := len(id) - 1; i >= 0; i-- {
        ch := id[i]
        
        if ch == ' ' {
            continue
        }
        if ch < '0' || ch > '9' {
            return false
        }
        
        digit := int(ch - '0')
        
        if shouldDouble {
            digit *= 2
            if digit > 9 {
                digit -= 9
            }
        }

        sum += digit
        count += 1
        shouldDouble = !shouldDouble
    }

    return count > 1 && sum % 10 == 0
}

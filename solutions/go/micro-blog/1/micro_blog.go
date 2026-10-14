package microblog

func Truncate(phrase string) string {
    count := 0
    for i := range phrase {
        if count == 5 {
            return phrase[:i]
        }
        count++
    }
    return phrase
}

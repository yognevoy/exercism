package birdwatcher

func TotalBirdCount(birdsPerDay []int) int {
    result := 0
    for i := 0; i < len(birdsPerDay); i++ {
        result += birdsPerDay[i]
    }
    return result
}

func BirdsInWeek(birdsPerDay []int, week int) int {
    start := (week - 1) * 7
    end := week * 7
    return TotalBirdCount(birdsPerDay[start:end])
}

func FixBirdCountLog(birdsPerDay []int) []int {
    for i := 0; i < len(birdsPerDay); i++ {
        if i % 2 == 0 {
            birdsPerDay[i] += 1
        }
    }
    return birdsPerDay
}

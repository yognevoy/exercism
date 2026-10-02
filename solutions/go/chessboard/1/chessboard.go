package chessboard

type File []bool

type Chessboard map[string]File

// CountInFile returns how many squares are occupied in the chessboard,
// within the given file.
func CountInFile(cb Chessboard, file string) int {
    fileValue := cb[file]
    if fileValue == nil {
        return 0
    }
    result := 0
    for _, occupied := range fileValue {
        if occupied {
            result++
        }
    }
    return result
}

// CountInRank returns how many squares are occupied in the chessboard,
// within the given rank.
func CountInRank(cb Chessboard, rank int) int {
    if rank < 1 || rank > 8 {
        return 0
    }
    result := 0
    for _, file := range cb {
        if len(file) >= rank && file[rank-1] {
            result++
        }
    }
    return result
}

// CountAll should count how many squares are present in the chessboard.
func CountAll(cb Chessboard) int {
    result := 0
    for range cb {
        result += 8
    }
    return result
}

// CountOccupied returns how many squares are occupied in the chessboard.
func CountOccupied(cb Chessboard) int {
    result := 0
    for _, file := range cb {
        for _, occupied := range file {
            if occupied {
                result++
            }
        }
    }
    return result
}

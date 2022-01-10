package chessboard

// Declare a type named Rank which stores if a square is occupied by a piece - this will be a slice of bools
type Rank [8]bool

// Declare a type named Chessboard which contains a map of eight Ranks, accessed with keys from "A" to "H"
type Chessboard map[string]Rank

// CountInRank returns how many squares are occupied in the chessboard,
// within the given rank
func CountInRank(cb Chessboard, rank string) int {
	value_board, _ := cb[rank]
	ret := 0
	for _, value_rank := range value_board {
		if value_rank {
			ret++
		}
	}
	return ret
}

// CountInFile returns how many squares are occupied in the chessboard,
// within the given file
func CountInFile(cb Chessboard, file int) int {
	ret := 0
	if file >= 1 && file <= 8 {
		for _, v := range cb {
			if v[file-1] {
				ret++
			}
		}
	}
	return ret
}

// CountAll should count how many squares are present in the chessboard
func CountAll(cb Chessboard) int {
	ret := 0
	for range cb {
		ret += 8
	}
	return ret
}

// CountOccupied returns how many squares are occupied in the chessboard
func CountOccupied(cb Chessboard) int {
	ret := 0
	for k, _ := range cb {
		ret += CountInRank(cb, k)
	}
	return ret
}

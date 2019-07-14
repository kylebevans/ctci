package main

import (
	"fmt"
)

func Get8QueenSolns(currentQueens [][2]int) [][][2]int {
	currentRow := len(currentQueens)
	//fmt.Println(currentRow)
	var solutions [][][2]int
	for i := 0; i < 8; i++ {
		currentQueensCopy := currentQueens
		if IsLegalSpot([2]int{currentRow, i}, currentQueens) {
			currentQueensCopy = append(currentQueensCopy, [2]int{currentRow, i})
			if len(currentQueensCopy) == 8 {
				solutions =  append(solutions, currentQueensCopy)
				continue
			}
			spotSolutions := Get8QueenSolns(currentQueensCopy)
			//fmt.Println(spotSolutions)
			//fmt.Println(currentQueensCopy)
			//fmt.Println(i)
			if len(spotSolutions) != 0 {
				solutions = append(solutions, spotSolutions...)
			}
		}
	}
	return solutions
}

func IsLegalSpot(spot [2]int, currentQueens [][2]int) bool {
	for _, v := range currentQueens {
		if v[0] == spot[0] || v[1] == spot[1] {
			return false
		}
		rise := spot[1] - v[1]
		run := spot[0] - v[0]
		if (rise / run == 1 || rise / run == -1) && rise % run == 0 {
			return false
		}
	}
	return true
}

func main() {
	for _, v := range Get8QueenSolns([][2]int{}) {
		fmt.Println(v)
	}
	return
}
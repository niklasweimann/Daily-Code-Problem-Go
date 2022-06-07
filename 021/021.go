package main

import (
	"fmt"
	"math"
)

func main() {
	matrix := [][]int{
		{0, 3, 1, 1},
		{2, 0, 0, 4},
		{1, 5, 3, 1}}
	print(matrix)
	fmt.Println(solve(matrix))

}

func solve(matrix [][]int) int {
	for i := 1; i < len(matrix); i++ {
		matrix[i][0] += matrix[i-1][0]
		//print(matrix)
	}
	for c := 1; c < len(matrix[0]); c++ {
		matrix[0][c] += matrix[0][c-1]
		//print(matrix)
	}
	for r := 1; r < len(matrix); r++ {
		for c := 1; c < len(matrix[0]); c++ {
			matrix[r][c] += int(math.Max(float64(matrix[r-1][c]), float64(matrix[r][c-1])))
			//print(matrix)
		}
	}
	return matrix[len(matrix)-1][len(matrix[0])-1]
}
func print(matrix [][]int) {
	for i := 0; i < len(matrix); i++ {
		for j := 0; j < len(matrix[i]); j++ {
			fmt.Print(matrix[i][j])
		}
		fmt.Print("\n")
	}
	fmt.Println()
}

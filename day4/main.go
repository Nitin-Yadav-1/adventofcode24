package main

import (
	"fmt"
	"os"
	"strings"
)

func main() {
	data, err := os.ReadFile("day4_input.txt")
	if err != nil {
		panic(err)
	}

	var matrix [][]int = parseInput(data)
	fmt.Println("partone : ", partone(matrix))
	fmt.Println("partone : ", parttwo(matrix))
}

func parseInput(data []byte) [][]int {
	var matrix [][]int
	var row []int

	for _, ch := range strings.ReplaceAll(string(data), "\r\n", "\n") {
		if ch == '\n' {
			matrix = append(matrix, row)
			row = nil
		} else {
			row = append(row, int(ch))
		}
	}

	if row != nil {
		matrix = append(matrix, row)
	}

	return matrix
}

func partone(matrix [][]int) int {
	var count int = 0
	var word string = "XMAS"

	for i, row := range matrix {
		for j, elem := range row {
			if elem == int(word[0]) {
				// search horizontal
				if matchAt(i, j, matrix, 0, 1, word) {
					count += 1
				}
				if matchAt(i, j, matrix, 0, -1, word) {
					count += 1
				}
				// search vertical
				if matchAt(i, j, matrix, 1, 0, word) {
					count += 1
				}
				if matchAt(i, j, matrix, -1, 0, word) {
					count += 1
				}
				// search diagonal L to R
				if matchAt(i, j, matrix, 1, 1, word) {
					count += 1
				}
				if matchAt(i, j, matrix, -1, -1, word) {
					count += 1
				}
				// search diagonal R to L
				if matchAt(i, j, matrix, 1, -1, word) {
					count += 1
				}
				if matchAt(i, j, matrix, -1, 1, word) {
					count += 1
				}
			}
		}
	}

	return count
}

func parttwo(matrix [][]int) int {
	var count int = 0

	for i, row := range matrix {
		for j, elem := range row {
			if elem == 'A' {
				ltordiag := (matchAt(i, j, matrix, 1, 1, "AS") && matchAt(i, j, matrix, -1, -1, "AM")) || (matchAt(i, j, matrix, 1, 1, "AM") && matchAt(i, j, matrix, -1, -1, "AS"))
				rtoldiag := (matchAt(i, j, matrix, 1, -1, "AS") && matchAt(i, j, matrix, -1, 1, "AM")) || (matchAt(i, j, matrix, 1, -1, "AM") && matchAt(i, j, matrix, -1, 1, "AS"))
				isX := ltordiag && rtoldiag
				if isX {
					count += 1
				}
			}
		}
	}

	return count
}

func matchAt(i int, j int, matrix [][]int, di int, dj int, word string) bool {
	if len(matrix) == 0 {
		return len(word) == 0
	}

	if len(word) == 0 {
		return false
	}

	var rows int = len(matrix)
	var cols int = len(matrix[0])

	for _, ch := range word {
		if i < 0 || i >= rows {
			return false
		}
		if j < 0 || j >= cols {
			return false
		}
		if matrix[i][j] != int(ch) {
			return false
		}
		i += di
		j += dj
	}

	return true
}

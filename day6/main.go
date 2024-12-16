package main

import (
	"fmt"
	"os"
	"strings"
)


func main() {
	data, err := os.ReadFile("day6_input.txt")
	if err != nil {
		panic(err)
	}

	var matrix [][]int = parseInput(strings.ReplaceAll(string(data), "\r\n", "\n"))
	fmt.Println("partone : ", partone(matrix))
}


func parseInput(input string) [][]int {
	var matrix [][]int
	var row []int

	for _, ch := range input {
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
	i, j := findChar(matrix, int('^'))
	var isOnMap bool = true
	var di int = -1
	var dj int = 0
	var visited [][]bool = make([][]bool, len(matrix))
	for r, row := range matrix {
		visited[r] = make([]bool, len(row))
	}

	for isOnMap {
		newi, newj := canMove(i,j,di,dj,matrix)
		if newi == i && newj == j {
			newdi, newdj := rotateRight(di,dj)
			di = newdi
			dj = newdj
		} else if newi == -1 && newj == -1 {
			visited[i][j] = true
			isOnMap = false
		} else {
			visited[i][j] = true
			i += di
			j += dj
		}
	}

	for _, row := range visited {
		for _, val := range row {
			if val {
				count++
			}
		}
	}

	return count
}

func rotateRight(di int, dj int) (int, int) {
	var newdi int = 0
	var newdj int = 0
	if di == -1 && dj == 0 {
		newdi = 0
		newdj = 1
	} else if di == 0 && dj == 1 {
		newdi = 1
		newdj = 0
	} else if di == 1 && dj == 0 {
		newdi = 0
		newdj = -1
	} else if di == 0 && dj == -1 {
		newdi = -1
		newdj = 0
	}
	return newdi, newdj
}

func canMove(i int, j int, di int, dj int, matrix [][]int) (int,int) {
	if len(matrix) == 0 {
		return -1,-1
	}

	if i < 0 || i >= len(matrix) || j < 0 || j >= len(matrix[0]) {
		return -1,-1
	}
	
	var newi int = i + di
	var newj int = j + dj
	
	if newi < 0 || newi >= len(matrix) || newj < 0 || newj >= len(matrix[0]) {
		return -1,-1
	}

	if matrix[newi][newj] == '#' {
		return i,j
	}

	return newi, newj
}

func findChar(matrix [][]int, char int) (int,int) {
	for i, row := range matrix {
		for j, elem := range row {
			if elem == char{
				return i,j
			}
		}
	}

	return -1,-1
}
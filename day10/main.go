package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	data, err := os.ReadFile("input.txt")
	if err != nil {
		fmt.Println(err.Error())
	}

	var matrix [][]int = parseInput(data)
	fmt.Println("partone : ", partone(matrix))
	fmt.Println("parttwo : ", parttwo(matrix))
}


func parseInput(data []byte) [][]int {
	var matrix [][]int
	var row []int

	for _, ch := range strings.ReplaceAll(string(data), "\r\n", "\n") {
		if ch == '\n' {
			matrix = append(matrix, row)
			row = nil
		} else {
			row = append(row, int(ch) - int('0'))
		}
	}

	if row != nil {
		matrix = append(matrix, row)
		row = nil
	}

	return matrix
}


func partone(matrix [][]int) int {
	var count int = 0
	var visited [][]bool = makeVisited(matrix)

	for i, row := range matrix {
		for j, elem := range row {
			if elem == 0 {
				var uniquePostions []string
				traverseTrail(i, j, matrix, visited, &uniquePostions)
				count += len(uniquePostions)
			}
		}
	}

	return count
}


func parttwo(matrix [][]int) int {
	var count int = 0
	var visited [][]bool = makeVisited(matrix)

	for i, row := range matrix {
		for j, elem := range row {
			if elem == 0 {
				count += countTrail(i, j, matrix, visited)
			}
		}
	}

	return count
}


func traverseTrail(i int, j int, matrix [][]int, visited [][]bool, uniquePositions *[]string)  {
	if matrix[i][j] == 9 {
		pos := strconv.Itoa(i)+":"+strconv.Itoa(j)
		isPresent := false
		for _, unipos := range *uniquePositions {
			if pos == unipos {
				isPresent = true
			}
		}
		if !isPresent {
			*uniquePositions = append(*uniquePositions, pos)
		}
		return
	}

	visited[i][j] = true

	var di []int = []int{0,0,1,-1}
	var dj []int = []int{1,-1,0,0}
	for k := 0; k < 4; k++ {
		row := i + di[k]
		col := j + dj[k]
		if isValid(row, col, matrix) && (!visited[row][col]) && ((matrix[i][j] + 1) == matrix[row][col]) {
			traverseTrail(row, col, matrix, visited, uniquePositions)
		}
	}

	visited[i][j] = false
}


func countTrail(i int, j int, matrix [][]int, visited [][]bool) int {
	if matrix[i][j] == 9 {
		return 1
	}

	visited[i][j] = true

	var di []int = []int{0,0,1,-1}
	var dj []int = []int{1,-1,0,0}
	var count int = 0

	for k := 0; k < 4; k++ {
		row := i + di[k]
		col := j + dj[k]
		if isValid(row, col, matrix) && (!visited[row][col]) && ((matrix[i][j] + 1) == matrix[row][col]) {
			count += countTrail(row, col, matrix, visited)
		}
	}

	visited[i][j] = false

	return count
}

func isValid(i int, j int, matrix [][]int) bool {
	return (0 <= i && i < len(matrix)) && (0 <= j && j < len(matrix[0]))
}

func makeVisited(matrix [][]int) [][]bool {
	var visited [][]bool
	for _, row := range matrix{
		visited = append(visited, make([]bool, len(row)))
	}
	return visited
}


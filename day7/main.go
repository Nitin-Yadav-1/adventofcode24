package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

type fn func(int64,int64) int64

type Line struct {
	result int64
	nums []int
}

func add(a int64, b int64) int64 {
	return a + b
}

func mul(a int64, b int64) int64 {
	return a * b
}

func concat(a int64, b int64) int64{
	num, _ := strconv.ParseInt(strconv.FormatInt(a, 10) + strconv.FormatInt(b, 10), 10, 64)
	return num
}

func main() {
	data, err := os.ReadFile("input.txt")
	if err != nil {
		panic(err)
	}

	lines := parseInput(data)
	fmt.Println("partone : ", partone(lines))
	fmt.Println("parttwo : ", parttwo(lines))
}

func parseInput(data []byte) []Line {
	var lines []Line
	for _, line := range strings.Split(strings.ReplaceAll(string(data), "\r\n", "\n"), "\n") {
		var arr []int = []int{}
		num, _ := strconv.Atoi(strings.Split(line, ":")[0]) 
		nums := strings.Split(strings.Split(line, ":")[1], " ")
		for _, s := range nums {
			if s == "" {continue}
			elem, _ := strconv.Atoi(s)
			arr = append(arr, elem)
		}
		lines = append(lines, Line{int64(num), arr})
	}
	return lines
}

func partone(lines []Line) int64 {
	var sum int64 = 0
	var operators []fn = []fn{add, mul}
	for _, line := range lines {
		if evalRecursive(1, int64(line.nums[0]), line.nums, line.result, operators) > 0 {
			sum += line.result
		}
	}
	return sum
}


func parttwo(lines []Line) int64 {
	var sum int64 = 0
	var operators []fn = []fn{add, mul, concat}
	for _, line := range lines {
		if evalRecursive(1, int64(line.nums[0]), line.nums, line.result, operators) > 0 {
			sum += line.result
		}
	}
	return sum
}


func evalRecursive(currIndex int, currResult int64, nums []int, result int64, operators []fn) int {
	if currIndex >= len(nums) {
		if currResult == result{
			return 1
		} else {
			return 0
		}
	}

	var count int = 0
	for _, opr := range operators {
		res := opr(currResult, int64(nums[currIndex]))
		count += evalRecursive(currIndex+1, res, nums, result, operators)
	}

	return count
}
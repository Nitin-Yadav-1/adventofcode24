package main

import (
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
)

const LOW_SAFE_LIMIT int = 1
const HIGH_SAFE_LIMIT int = 3

func main() {
	data, err := os.ReadFile("./day2_input.txt")
	if err != nil {
		panic(err)
	}

	var lines []string = strings.Split(string(data), "\r\n")
	var reports [][]int

	for _, line := range lines {
		var report []int

		for _, num := range strings.Split(line, " ") {
			intnum, err := strconv.Atoi(num)
			if err != nil {
				panic(err)
			}
			report = append(report, intnum)
		}

		reports = append(reports, report)
	}

	fmt.Println("Part one : ", partone(reports))
	fmt.Println("Part two : ", parttwo(reports))
}

func partone(reports [][]int) int {
	var count int = 0

	for _, report := range reports {
		if isSafe(report) {
			count += 1
		}
	}

	return count
}

func parttwo(reports [][]int) int {
	var count int = 0

	for _, report := range reports {
		if isSafe(report) {
			count += 1
		} else if canBeSafe(report) {
			count += 1
		}
	}

	return count
}

func canBeSafe(report []int) bool {
	for i := 0; i < len(report); i++ {
		var newReport []int
		for j := 0; j < len(report); j++ {
			if i == j {
				continue
			}
			newReport = append(newReport, report[j])
		}
		if isSafe(newReport) {
			return true
		}
	}

	return false
}

func isSafe(report []int) bool {
	isStrictlyIncreasing, min, max := isStrictlyIncreasing(report[:])
	if isStrictlyIncreasing {
		return (LOW_SAFE_LIMIT <= min) && (max <= HIGH_SAFE_LIMIT)
	}

	isStrictlyDecreasing, min, max := isStrictlyDecreasing(report[:])
	if isStrictlyDecreasing {
		return (LOW_SAFE_LIMIT <= min) && (max <= HIGH_SAFE_LIMIT)
	}

	return false
}

func isStrictlyIncreasing(arr []int) (bool, int, int) {
	var minDiff int = math.MaxInt
	var maxDiff int = math.MinInt

	if len(arr) <= 1 {
		return true, 0, 0
	}

	for i := 0; i < len(arr)-1; i++ {
		diff := arr[i] - arr[i+1]
		if diff >= 0 {
			return false, 0, 0
		}
		minDiff = int(math.Min(float64(minDiff), math.Abs(float64(diff))))
		maxDiff = int(math.Max(float64(maxDiff), math.Abs(float64(diff))))
	}

	return true, minDiff, maxDiff
}

func isStrictlyDecreasing(arr []int) (bool, int, int) {
	var minDiff int = math.MaxInt
	var maxDiff int = math.MinInt

	if len(arr) <= 1 {
		return true, 0, 0
	}

	for i := 0; i < len(arr)-1; i++ {
		diff := arr[i] - arr[i+1]
		if diff <= 0 {
			return false, 0, 0
		}
		minDiff = int(math.Min(float64(minDiff), math.Abs(float64(diff))))
		maxDiff = int(math.Max(float64(maxDiff), math.Abs(float64(diff))))
	}

	return true, minDiff, maxDiff
}

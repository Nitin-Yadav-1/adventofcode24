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
		panic(err)
	}

	rules, updates := parseInput(data)
	fmt.Println("partone : ", partone(rules, updates))
	fmt.Println("parttwo : ", parttwo(rules, updates))
}

func parseInput(data []byte) (map[int][]int, [][]int) {
	var rules map[int][]int = map[int][]int{}
	var updates [][]int = [][]int{}

	parts := strings.Split(strings.ReplaceAll(string(data), "\r\n", "\n"), "\n\n")

	for _, line := range strings.Split(parts[0], "\n") {
		nums := strings.Split(line, "|")
		first, _ := strconv.Atoi(nums[0])
		second, _ := strconv.Atoi(nums[1])
		rules[first] = append(rules[first], second)
	}

	for _, line := range strings.Split(parts[1], "\n") {
		var row []int
		for _, element := range strings.Split(line, ",") {
			num, _ := strconv.Atoi(element)
			row = append(row, num)
		}
		updates = append(updates, row)
	}

	return rules, updates
}

func partone(rules map[int][]int, updates [][]int) int {
	var sum int = 0
	for _, update := range updates {
		if isValidUpdate(rules, update) {
			sum += update[len(update)/2]
		}
	}
	return sum
}

func parttwo(rules map[int][]int, updates [][]int) int {
	var sum int = 0
	for _, update := range updates {
		if !isValidUpdate(rules, update) {
			sort(update, rules)
			sum += update[len(update)/2]
		}
	}
	return sum
}

func isValidUpdate(rules map[int][]int, update []int) bool {
	for i := 0; i < len(update); i++ {
		for j := i + 1; j < len(update); j++ {
			if contains(update[i], rules[update[j]]) {
				return false
			}
		}
	}
	return true
}

func contains(num int, arr []int) bool {
	for _, element := range arr {
		if element == num {
			return true
		}
	}
	return false
}

func sort(update []int, rules map[int][]int) {
	for i := 0; i < len(update); i++ {
		for j := i + 1; j < len(update); j++ {
			if compare(update[i], update[j], rules) == 1 {
				temp := update[i]
				update[i] = update[j]
				update[j] = temp
			}
		}
	}
}

func compare(a int, b int, rules map[int][]int) int {
	if contains(b, rules[a]) {
		return -1
	}

	if contains(a, rules[b]) {
		return 1
	}

	return 0
}

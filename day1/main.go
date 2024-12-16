package main

import (
	"fmt"
	"math"
	"os"
	"sort"
	"strconv"
	"strings"
)

func main() {
	data, err := os.ReadFile("./day1_input.txt")
	if err != nil {
		panic(err)
	}

	var lines []string = strings.Split(strings.ReplaceAll(string(data), "\r\n", "\n"), "\n")
	var nums1 [1000]int
	var nums2 [1000]int

	for i := 0; i < len(lines); i++ {
		var nums []string = strings.Split(lines[i], "   ")
		num1, err1 := strconv.Atoi(strings.TrimSpace(nums[0]))
		num2, err2 := strconv.Atoi(strings.TrimSpace(nums[1]))
		if err1 != nil  || err2 != nil {
			panic("error")
		}
		nums1[i] = num1
		nums2[i] = num2
	}
	
	fmt.Println("Part 1 : ", partone(nums1[:], nums2[:]))
	fmt.Println("Part 2 : ", parttwo(nums1[:], nums2[:]))
}

func partone(left []int, right []int) int {
	sort.Ints(left[:])
	sort.Ints(right[:])

	var sum int = 0
	for i:= 0; i < 1000; i++ {
		diff := left[i] - right[i]
		sum += int(math.Abs(float64(diff)))
	}
	return sum
}

func parttwo(left []int, right []int) int {
	var score int = 0
	for i := 0; i < len(left); i++ {
		var count int = 0
		for j := 0; j < len(right); j++ {
			if right[j] == left[i]{ 
				count++
			}
		}
		score += left[i] * count
	}
	
	return score
}
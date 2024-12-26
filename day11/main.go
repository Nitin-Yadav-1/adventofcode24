package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)


type Key struct {
	num uint64
	blinkLeft int
}


func main() {
	data, err := os.ReadFile("input.txt")
	if err != nil {
		panic(err)
	}

	nums := parseInput(data)
	fmt.Println("partone : ", partone(nums))
	fmt.Println("parttwo : ", parttwo(nums))
}

func parseInput(data []byte) []uint64 {
	var nums []uint64 

	for _, ch := range strings.Split(string(data), " ") {
		num, _ := strconv.Atoi(ch)
		nums = append(nums, uint64(num))
	}

	return nums
}


func partone(nums []uint64) int {
	var count int = 0
	var memoizeMap map[Key]int = map[Key]int{}
	for _, num := range nums {
		count += recursiveBlink(num, 25, memoizeMap)
	}
	return count
}

func parttwo(nums []uint64) int {
	var count int = 0
	var memoizeMap map[Key]int = map[Key]int{}
	for _, num := range nums {
		count += recursiveBlink(num, 75, memoizeMap)
	}
	return count
}


func recursiveBlink(num uint64, blinkLeft int, memoizeMap map[Key]int) int {
	if blinkLeft == 0 {
		return 1
	}

	value, keyExists := memoizeMap[Key{num, blinkLeft}]
	if keyExists {
		return value
	}

	var count int = 0

	if num == 0 {
		count += recursiveBlink(1, blinkLeft-1, memoizeMap)
	} else if isEvenDigits, leftnum, rightnum := evenDigitRule(num); isEvenDigits {
		count += recursiveBlink(leftnum, blinkLeft-1, memoizeMap)
		count += recursiveBlink(rightnum, blinkLeft-1, memoizeMap)
	} else {
		count += recursiveBlink(num * 2024, blinkLeft-1, memoizeMap)
	}

	memoizeMap[Key{num, blinkLeft}] = count
	
	return count
}


func evenDigitRule(num uint64) (bool, uint64, uint64) {
	var digitCount int = 0
	var leftnum uint64 = 0
	var rightnum uint64 = 0

	var n uint64 = num
	for n > 0 {
		digitCount++
		n /= 10
	}

	if digitCount % 2 != 0 {
		return false, leftnum, rightnum
	}
	
	// get right half in reverse
	var temp uint64 = 1
	for i := 0; i < digitCount/2; i++ {
		temp = (temp * 10) + (num % 10)
		num /= 10
	}

	// reverse again right half
	var rev_n uint64 = 0
	for temp > 0 {
		rev_n = (rev_n * 10) + (temp % 10)
		temp /= 10
	}
	rev_n /= 10

	leftnum = num
	rightnum = rev_n

	return true, leftnum, rightnum
}

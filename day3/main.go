package main

import (
	"fmt"
	"os"
)

func main() {
	data, err := os.ReadFile("./day3_input.txt")
	if err != nil {
		panic(err)
	}

	memory := string(data)

	fmt.Println("Part one : ", partone(memory))
	fmt.Println("Part two : ", parttwo(memory))
}

func partone(memory string) uint64 {
	var result uint64 = 0

	var stack []int

	for _, ch := range memory {
		switch ch {
		case 'm':
			stack = nil
			stack = append(stack, int(ch))
		case 'u':
			if len(stack) > 0 && stack[len(stack)-1] == 'm' {
				stack = append(stack, int(ch))
			} 
		case 'l':
			if len(stack) > 0 && stack[len(stack)-1] == 'u' {
				stack = append(stack, int(ch))
			} 
		case '(':
			if len(stack) > 0 && stack[len(stack)-1] == 'l' {
				stack = append(stack, int(ch))
			} 
		case '0','1','2','3','4','5','6','7','8','9':
			if len(stack) > 0 && (stack[len(stack)-1] == '(' || stack[len(stack)-1] == ',' || isNumeric(stack[len(stack)-1])) {
				stack = append(stack, int(ch))
				} 
		case ',':
			if len(stack) > 0 && (1 <= numCountAfterChar(stack, '(') && numCountAfterChar(stack, '(') <= 3) && isNumeric(stack[len(stack)-1]) {
				stack = append(stack, int(ch))		
			}
		case ')':
			if len(stack) > 0 && (1 <= numCountAfterChar(stack, ',') && numCountAfterChar(stack, '(') <= 3) && isNumeric(stack[len(stack)-1]) && commaCount(stack) == 1 {
				stack = append(stack, int(ch))
				result += uint64(evaluateStack(stack))
				stack = nil
			}
		default:
			stack = nil
		}
	}

	return result
}

func parttwo(memory string) uint64 {
	var result uint64 = 0

	var stack []int
	var isEnabled bool = true

	for i, ch := range memory {
		if isEnabled {
			switch ch {
			case 'm':
				stack = nil
				stack = append(stack, int(ch))
			case 'u':
				if len(stack) > 0 && stack[len(stack)-1] == 'm' {
					stack = append(stack, int(ch))
				} 
			case 'l':
				if len(stack) > 0 && stack[len(stack)-1] == 'u' {
					stack = append(stack, int(ch))
				} 
			case '(':
				if len(stack) > 0 && stack[len(stack)-1] == 'l' {
					stack = append(stack, int(ch))
				} 
			case '0','1','2','3','4','5','6','7','8','9':
				if len(stack) > 0 && (stack[len(stack)-1] == '(' || stack[len(stack)-1] == ',' || isNumeric(stack[len(stack)-1])) {
					stack = append(stack, int(ch))
					} 
			case ',':
				if len(stack) > 0 && (1 <= numCountAfterChar(stack, '(') && numCountAfterChar(stack, '(') <= 3) && isNumeric(stack[len(stack)-1]) {
					stack = append(stack, int(ch))		
				}
			case ')':
				if len(stack) > 0 && (1 <= numCountAfterChar(stack, ',') && numCountAfterChar(stack, '(') <= 3) && isNumeric(stack[len(stack)-1]) && commaCount(stack) == 1 {
					stack = append(stack, int(ch))
					result += uint64(evaluateStack(stack))
					stack = nil
				}
			default:
				stack = nil
			}
		}

		if isDoInstruction(i, memory) {
			isEnabled = true
		}

		if isDontInstruction(i, memory) {
			isEnabled = false
		}

	}

	return result
}

func isDoInstruction(currIndex int, memory string) bool{
	var do string = "do()"
	var doIndex = 0
	
	for currIndex < len(memory) && doIndex < len(do) {
		if memory[currIndex] != do[doIndex] {
			return false
		}
		currIndex++
		doIndex++
	}

	return doIndex == len(do)
}

func isDontInstruction(currIndex int, memory string) bool{
	var do string = "don't()"
	var doIndex = 0

	for currIndex < len(memory) && doIndex < len(do) {
		if memory[currIndex] != do[doIndex] {
			return false
		}
		currIndex++
		doIndex++
	}

	return doIndex == len(do)
}

func printStack(stack []int) {
	for _, ch := range stack {
		fmt.Printf("%c", ch)
	}
	fmt.Println()
}

func isNumeric(char int) bool {
	return '0' <= char && char <= '9'
}

func commaCount(stack []int) int {
	var count int = 0
	for _, ch := range stack {
		if ch == ',' {
			count++
		}
	}
	return count
}

func numCountAfterChar(stack []int, afterChar int) int {
	var count int = 0
	var afterCharIndex int = findFirst(stack, afterChar)
	
	if afterCharIndex == -1 {
		return 0
	}

	var currIndex int = afterCharIndex + 1
	for currIndex < len(stack) && isNumeric(stack[currIndex]) {
		count++
		currIndex++
	}
	return count
}

func findFirst(stack []int, char int) int {
	var index int = -1
	for i, ch := range stack {
		if ch == char {
			index = i
			break
		}
	}
	return index
}

func sliceToInt(s []int) int {
    res := 0
    op := 1
    for i := len(s) - 1; i >= 0; i-- {
        res += charToInt(s[i]) * op
        op *= 10
    }
    return res
}

func charToInt(char int) int {
	return char - int('0')
}

func evaluateStack(stack []int) int {
	var commaIndex int = findFirst(stack, ',')
	a := sliceToInt(stack[4:commaIndex])
	b := sliceToInt(stack[commaIndex+1:len(stack)-1])
	return a * b
}
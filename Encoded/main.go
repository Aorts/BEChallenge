package main

import (
	"fmt"
)

func encode(input string) string {
	result := ""
	leftCount := 0
	rightCount := 0
	previousValue := 0
	for i := 0; i < len(input); i++ {
		if input[i] == 'L' {
			leftCount++
			if i > 0 && input[i-1] == 'L' {
				rightCount = 1
			}
			previousValue = leftCount
			if i == 0 {
				result += fmt.Sprintf("%d", leftCount+10)

			} else {
				result += fmt.Sprintf("%d", leftCount)
			}
		} else if input[i] == 'R' {
			rightCount++
			if i > 0 && input[i-1] == 'L' {
				leftCount = 0
			}
			if i == 0 {
				result += fmt.Sprintf("%d", rightCount+10)
			} else {
				result += fmt.Sprintf("%d", rightCount)
			}
			previousValue = rightCount
		} else if input[i] == '=' {
			if i == 0 {
				result += fmt.Sprintf("%d", previousValue+10)
			} else {
				result += fmt.Sprintf("%d", previousValue)
			}
		}
	}

	return result
}

func main() {
	//var userInput string
	//fmt.Print("Enter input string: ")
	//fmt.Scan(&userInput)

	result := encode("LLRR=")
	fmt.Println("Output:", result)
}

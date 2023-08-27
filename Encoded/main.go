package main

import (
	"fmt"
)

func main() {
	decoded := decode("LLRR=")
	fmt.Printf("Decoded number: %s\n", decoded)
}

func decode(encoded string) string {
	var result string
	var numbers []int
	var previousNum = 0
	for i := 0; i < len(encoded); i++ {
		symbol1 := encoded[i]
		if symbol1 == 'L' {
			if i == 0 {
				fmt.Println("i am here L 0", "------", i)

				numbers = append(numbers, 1)
				numbers = append(numbers, 0)
			} else {
				if encoded[i-1] == 'L' {
					fmt.Println("i am here L 1", "------", i)
					numbers = addAllNum(numbers)
					previousNum = numbers[i-1] - 1
					numbers = append(numbers, previousNum)
				} else {
					fmt.Println("i am here L 2", "------", i)
					previousNum = numbers[i-1]
					numbers = append(numbers, previousNum)
				}
			}
		} else if symbol1 == 'R' {
			if i == 0 {
				fmt.Println("i am here R 0", "------", i)
				numbers = append(numbers, 0)
				numbers = append(numbers, 1)
			} else {
				if encoded[i-1] == 'R' {
					fmt.Println("i am here R 1", "------", i)
					previousNum = numbers[i-1] + 1
					numbers = append(numbers, previousNum)
				} else {
					fmt.Println("i am here R 2", "------", i)
					previousNum = numbers[i-1] - 1
					fmt.Println(previousNum)
					numbers = append(numbers, previousNum)
				}
			}
		} else if symbol1 == '=' {
			if i == 0 {
				fmt.Println("i am here = 0", "------", i)
				numbers = append(numbers, 0)
			} else {
				fmt.Println("i am here = 1", "------", i)
				previousNum = numbers[i-1]
				numbers = append(numbers, previousNum)
			}
		}
	}
	fmt.Println(numbers)
	return result
}

func isLastSymbol(encoded string, i int) bool {
	return i+1 == len(encoded)
}

func addAllNum(numbers []int) []int {
	for i := 0; i < len(numbers); i++ {
		numbers[i] = numbers[i] + 1
	}
	return numbers
}

package main

import (
	"fmt"
	"log"
	"os"
	"os/signal"
	"strconv"
	"strings"
	"syscall"
)

func main() {
	signalChan := make(chan os.Signal, 1)
	signal.Notify(signalChan, os.Interrupt, syscall.SIGTERM)

	go func() {
		<-signalChan
		fmt.Println("\nReceived Ctrl+C. Exiting...")
		os.Exit(0)
	}()

	for {
		fmt.Print("Enter the encoded string: ")
		var encodedStr string
		_, err := fmt.Scan(&encodedStr)
		if err != nil {
			log.Fatal(err)
		}
		encodedStr = strings.ToUpper(encodedStr)
		decoded := decode(encodedStr)
		fmt.Printf("Decoded number: %s\n", decoded)
	}

}

func decode(encoded string) string {
	var result string
	var numbers []int
	var previousNum = 0
	for i := 0; i < len(encoded); i++ {
		symbol1 := encoded[i]
		if symbol1 == 'L' {
			if i == 0 {
				numbers = append(numbers, 1)
				numbers = append(numbers, 0)
			} else {
				if encoded[i-1] == 'L' {
					if numbers[i-1] == 0 {
						numbers[i] = numbers[i] - 1
						numbers = addAllNum(numbers, encoded)
					}
					previousNum = numbers[i-1] - 1
					numbers = addAllNum(numbers, encoded)
					numbers = append(numbers, previousNum)
				} else {
					previousNum = numbers[i-1]
					numbers = append(numbers, previousNum)
				}
			}
		} else if symbol1 == 'R' {
			if i == 0 {
				numbers = append(numbers, 0)
				numbers = append(numbers, 1)
			} else {
				if encoded[i-1] == 'R' {
					previousNum = numbers[i] + 1
					numbers = append(numbers, previousNum)
				} else {
					previousNum = numbers[i] + 1
					numbers = append(numbers, previousNum)
				}
			}
		} else if symbol1 == '=' {
			if i == 0 {
				numbers = append(numbers, 0)
				numbers = append(numbers, 0)
			} else {
				previousNum = numbers[i]
				numbers = append(numbers, previousNum)
			}
		}
	}
	result = sliceIntToStr(numbers)
	return result
}

func sliceIntToStr(numbers []int) string {
	var strSlice []string
	for _, num := range numbers {
		strSlice = append(strSlice, strconv.Itoa(num))
	}
	result := strings.Join(strSlice, "")
	return result
}

func addAllNum(numbers []int, encoded string) []int {
	for i := 0; i < len(numbers); i++ {
		if encoded[i] == 'L' && !lastCharIsEqual(encoded, i) {
			numbers[i]++
		} else if encoded[i] == 'L' && lastCharIsEqual(encoded, i) {
			for i := 0; i < len(numbers)-1; i++ {
				numbers[i]++
			}
		}
	}
	return numbers
}

func lastCharIsEqual(encoded string, i int) bool {
	if i == 0 {
		return false
	} else if encoded[i-1] == '=' {
		return true
	}
	return false
}

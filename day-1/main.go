package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
	"unicode"
)

func recursiveCheckNumWord(contentStr string, position int, numWord string, numWordPosition int) bool {
	if contentStr[position] == numWord[numWordPosition] {
		if numWordPosition == (len(numWord) - 1) {
			return true
		} else {
			return recursiveCheckNumWord(contentStr, position+1, numWord, numWordPosition+1)
		}
	} else {
		return false
	}
}

func reverseRecursiveCheckNumWord(contentStr string, position int, numWord string, numWordPosition int) bool {
	if contentStr[position] == numWord[numWordPosition] {
		if numWordPosition == 0 {
			return true
		} else {
			return reverseRecursiveCheckNumWord(contentStr, position-1, numWord, numWordPosition-1)
		}
	} else {
		return false
	}
}

func main() {
	content, err := os.ReadFile("values.txt")
	if err != nil {
		fmt.Println("unable to read file")
		return
	}
	numberWords := []string{"zero", "one", "two", "three", "four", "five", "six", "seven", "eight", "nine"}

	contentStr := string(content)
	contentArr := strings.Split(contentStr, "\n")
	total := 0
	numbers := []int{}
	for _, contentItem := range contentArr {
		firstDigit, secondDigit, firstDigitFound, secondDigitFound := 0, 0, false, false
		for i := 0; i < len(contentItem); i++ {
			firstChar := contentItem[i]
			corrLastChar := contentItem[len(contentItem)-1-i]
			if unicode.IsDigit(rune(firstChar)) && !firstDigitFound {
				firstDigit, _ = strconv.Atoi(string(firstChar))
				firstDigitFound = true
			} else if !firstDigitFound {
				for num, numWord := range numberWords {
					if rune(firstChar) == rune(numWord[0]) && recursiveCheckNumWord(contentItem, i, numWord, 0) {
						firstDigit = num
						firstDigitFound = true
						break
					}
				}
			}
			if unicode.IsDigit(rune(corrLastChar)) && !secondDigitFound {
				secondDigit, _ = strconv.Atoi(string(corrLastChar))
				secondDigitFound = true
			} else if !secondDigitFound {
				for num, numWord := range numberWords {
					if rune(corrLastChar) == rune(numWord[len(numWord)-1]) && reverseRecursiveCheckNumWord(contentItem, len(contentItem)-1-i, numWord, len(numWord)-1) {
						secondDigit = num
						secondDigitFound = true
						break
					}
				}
			}
			if firstDigitFound && secondDigitFound {
				break
			}
		}
		number := firstDigit*10 + secondDigit
		numbers = append(numbers, number)
		total += number
	}
	fmt.Println(numbers)
	fmt.Println(total)
}

package main

import (
	"fmt"
	"os"
	"strconv"
)

func main() {
	content, n, done := getInput()
	if done {
		return
	}

	var lines []string
	var line string

	for i := 0; i < n; i++ {
		if content[i] == '\n' {
			lines = append(lines, line)
			line = ""
		} else {
			line += string(content[i])
		}
	}

	var numberInLine []int

	for _, line := range lines {
		var d = getNumberInLinePartTwo(line)
		var number int
		if len(d) == 1 {
			number = 10*d[0] + d[0]
		} else {
			number = 10*d[0] + d[len(d)-1]
		}
		numberInLine = append(numberInLine, number)
	}

	var sum int
	for _, number := range numberInLine {
		sum += number
	}

	fmt.Println(sum)
}

func getNumberInLinePartTwo(numberInLineAsString string) []int {
	var numberInLine []int

	for n, number := range numberInLineAsString {
		switch number {
		case 'o':
			if len(numberInLineAsString)-n >= 3 {
				if checkForWord(numberInLineAsString, n, "one") {
					numberInLine = append(numberInLine, 1)
				}
			}
		case 't':
			if len(numberInLineAsString)-n >= 3 {
				if checkForWord(numberInLineAsString, n, "two") {
					numberInLine = append(numberInLine, 2)
				}
			}
			if len(numberInLineAsString)-n >= 5 {
				if checkForWord(numberInLineAsString, n, "three") {
					numberInLine = append(numberInLine, 3)
				}
			}
		case 'f':
			if len(numberInLineAsString)-n >= 4 {
				if checkForWord(numberInLineAsString, n, "four") {
					numberInLine = append(numberInLine, 4)
				}
			}
			if len(numberInLineAsString)-n >= 4 {
				if checkForWord(numberInLineAsString, n, "five") {
					numberInLine = append(numberInLine, 5)
				}
			}
		case 's':
			if len(numberInLineAsString)-n >= 3 {
				if checkForWord(numberInLineAsString, n, "six") {
					numberInLine = append(numberInLine, 6)
				}
			}
			if len(numberInLineAsString)-n >= 5 {
				if checkForWord(numberInLineAsString, n, "seven") {
					numberInLine = append(numberInLine, 7)
				}
			}
		case 'e':
			if len(numberInLineAsString)-n >= 5 {
				if checkForWord(numberInLineAsString, n, "eight") {
					numberInLine = append(numberInLine, 8)
				}
			}
		case 'n':
			if len(numberInLineAsString)-n >= 4 {
				if checkForWord(numberInLineAsString, n, "nine") {
					numberInLine = append(numberInLine, 9)
				}
			}
		default:
			numberAsInt, err := strconv.Atoi(string(number))
			if err == nil {
				numberInLine = append(numberInLine, numberAsInt)
			}
		}
	}

	return numberInLine
}

func checkForWord(asString string, n int, s string) bool {
	for i := 0; i < len(s); i++ {
		if asString[n+i] != s[i] {
			return false
		}
	}
	return true
}

func getInput() ([]byte, int, bool) {
	filePath := "input.txt"

	file, err := os.Open(filePath)
	if err != nil {
		fmt.Println("Error opening file:", err)
		return nil, 0, true
	}
	defer file.Close()

	fileInfo, err := file.Stat()
	if err != nil {
		fmt.Println("Error getting file info:", err)
		return nil, 0, true
	}
	fileSize := fileInfo.Size()

	bufferSize := int(fileSize)
	if bufferSize < 1024 {
		bufferSize = 1024
	}

	content := make([]byte, bufferSize)
	n, err := file.Read(content)
	if err != nil {
		fmt.Println("Error reading file:", err)
		return nil, 0, true
	}
	return content, n, false
}

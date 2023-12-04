package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
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

	var winningNumbers []int
	var lineResult int
	var finalResult int
	for _, line := range lines {
		line = line[strings.IndexByte(line, ':')+1:]
		winningNumbersString := line[:strings.IndexByte(line, '|')]
		line = line[strings.IndexByte(line, '|')+1:]
		for len(winningNumbersString) > 0 {
			breakingCharIndex := strings.IndexByte(winningNumbersString, ' ')
			numberAsString := winningNumbersString[:breakingCharIndex]
			numberAsString = cleanString(numberAsString)
			winningNumbersString = winningNumbersString[breakingCharIndex+1:]
			if numberAsString != "" {
				number, _ := strconv.Atoi(numberAsString)
				winningNumbers = append(winningNumbers, number)
			}
		}
		for len(line) > 0 {
			breakingCharIndex := strings.IndexByte(line, ' ')
			if breakingCharIndex == -1 {
				breakingCharIndex = len(line)
			}
			numberAsString := line[:breakingCharIndex]
			numberAsString = cleanString(numberAsString)
			if breakingCharIndex == len(line) {
				line = ""
			} else {
				line = line[breakingCharIndex+1:]
			}
			if numberAsString != "" {
				number, _ := strconv.Atoi(numberAsString)
				for _, winningNumber := range winningNumbers {
					if number == winningNumber {
						if lineResult == 0 {
							lineResult++
						} else {
							lineResult *= 2
						}
					}
				}
			}
		}
		finalResult += lineResult
		winningNumbers = nil
		lineResult = 0
	}
	fmt.Println(finalResult)
}

func cleanString(s string) string {
	var clean string
	for _, c := range s {
		if c != ' ' && c != '\n' && c != '\t' && c != ',' && c != ':' && c != ';' {
			clean += string(c)
		}
	}
	return clean
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

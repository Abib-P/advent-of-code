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

	var newNumbers string
	var finalSum int
	for actualLine, line := range lines {
		for actualChar, char := range line {
			if char >= '0' && char <= '9' {
				newNumbers += string(char)
			} else if newNumbers != "" {
				sum, done2 := funcName(actualChar, newNumbers, actualLine, lines, finalSum, line)
				if done2 {
					return
				}
				finalSum = sum
				newNumbers = ""
			}
			if actualChar == len(line)-1 && newNumbers != "" {
				sum, done2 := funcName(actualChar, newNumbers, actualLine, lines, finalSum, line)
				if done2 {
					return
				}
				finalSum = sum
				newNumbers = ""
			}
		}
	}

	fmt.Println(finalSum)
}

func funcName(actualChar int, newNumbers string, actualLine int, lines []string, finalSum int, line string) (int, bool) {
	for i := actualChar - len(newNumbers) - 1; i <= actualChar; i++ {
		if i >= 0 && actualLine-1 >= 0 && i < len(lines[actualLine-1]) {
			if lines[actualLine-1][i] < '0' || lines[actualLine-1][i] > '9' {
				if lines[actualLine-1][i] != '.' {
					i, err := strconv.Atoi(newNumbers)
					if err != nil {
						fmt.Println("Error converting string to int:", err)
						return 0, true
					}
					finalSum += i
					fmt.Println(newNumbers)
					break
				}
			}
		}
		if i >= 0 && actualLine+1 < len(lines) && i < len(lines[actualLine+1]) {
			if lines[actualLine+1][i] < '0' || lines[actualLine+1][i] > '9' {
				if lines[actualLine+1][i] != '.' {
					i, err := strconv.Atoi(newNumbers)
					if err != nil {
						fmt.Println("Error converting string to int:", err)
						return 0, true
					}
					finalSum += i
					fmt.Println(newNumbers)
					break
				}
			}
		}
	}
	// check left and right
	if actualChar-len(newNumbers)-1 >= 0 {
		if line[actualChar-len(newNumbers)-1] < '0' || line[actualChar-len(newNumbers)-1] > '9' {
			if line[actualChar-len(newNumbers)-1] != '.' {
				i, err := strconv.Atoi(newNumbers)
				if err != nil {
					fmt.Println("Error converting string to int:", err)
					return 0, true
				}
				finalSum += i
				fmt.Println(newNumbers)
			}
		}
	}
	if actualChar < len(line) {
		if line[actualChar] < '0' || line[actualChar] > '9' {
			if line[actualChar] != '.' {
				i, err := strconv.Atoi(newNumbers)
				if err != nil {
					fmt.Println("Error converting string to int:", err)
					return 0, true
				}
				finalSum += i
				fmt.Println(newNumbers)
			}
		}
	}
	return finalSum, false
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

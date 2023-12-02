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
		var d = getNumberInLine(line)
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

	fmt.Println("the result is : " + strconv.Itoa(sum))

}

func getNumberInLine(numberInLineAsString string) []int {
	var numberInLine []int

	for _, number := range numberInLineAsString {
		numberAsInt, err := strconv.Atoi(string(number))
		if err == nil {
			numberInLine = append(numberInLine, numberAsInt)
		}
	}
	return numberInLine
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

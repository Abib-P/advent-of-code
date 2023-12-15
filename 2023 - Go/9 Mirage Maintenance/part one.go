package main

import (
	"fmt"
	"os"
)

//file :
//0 3 6 9 12 15
//1 3 6 10 15 21
//10 13 16 21 30 45

//0   3   6   9  12  15 18
//  3   3   3   3   3  3
//    0   0   0   0  0

//1   3   6  10  15  21  28
//  2   3   4   5   6  7
//    1   1   1   1  1
//      0   0   0  0

//10  13  16  21  30  45  66
//  3   3   5   9  15  23
//    0   2   4   6  8
//      2   2   2  2
//        0   0  0

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

	var finalResult int

	for _, line := range lines {
		var lineAsInt []int
		var number int
		asNumber := false
		isNegative := false
		for i := 0; i < len(line); i++ {
			if (line[i] == ' ' || line[i] == '\t' || line[i] == '\n' || line[i] == '\r') && asNumber {
				lineAsInt = append(lineAsInt, number)
				number = 0
				asNumber = false
				isNegative = false
			} else if line[i] >= '0' && line[i] <= '9' && !isNegative {
				number = number*10 + int(line[i]-'0')
				asNumber = true
			} else if line[i] >= '0' && line[i] <= '9' && isNegative {
				number = number*10 - int(line[i]-'0')
				asNumber = true
			} else if line[i] == '-' {
				isNegative = true
			}
		}
		if number != 0 {
			lineAsInt = append(lineAsInt, number)
		}
		finalResult += solveLine(lineAsInt)
	}

	fmt.Println(finalResult)
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

func solveLine(line []int) int {
	if lineContainsOnlyZero(line) {
		return 0
	}
	if len(line) == 1 {
		return line[0]
	}

	var nextLine []int
	for i := 0; i < len(line)-1; i++ {
		nextLine = append(nextLine, line[i+1]-line[i])
	}
	return solveLine(nextLine) + line[len(line)-1]
}

func lineContainsOnlyZero(line []int) bool {
	for _, number := range line {
		if number != 0 {
			return false
		}
	}
	return true
}

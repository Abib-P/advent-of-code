package main

import (
	"fmt"
	"os"
)

type node struct {
	Left  string
	Right string
}

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

	var instructions string
	instructions = lines[0]
	lines = lines[2:]
	fmt.Println(instructions)

	var maps map[string]node
	maps = make(map[string]node)

	for _, line := range lines {
		var key string
		var right string
		var left string
		fmt.Sscanf(line, "%s = (%s %s)", &key, &left, &right)
		left = left[:len(left)-1]
		right = right[:len(right)-1]
		maps[key] = node{Left: left, Right: right}
	}

	var current string
	current = "AAA"
	i := 0
	numberOfInstructionsExecute := 0
	for current != "ZZZ" {
		if instructions[i] == 'R' {
			current = maps[current].Right
		} else {
			current = maps[current].Left
		}
		i++
		numberOfInstructionsExecute++
		if i == len(instructions) {
			i = 0
		}
	}

	fmt.Println(numberOfInstructionsExecute)
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

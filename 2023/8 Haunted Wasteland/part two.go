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

	var maps map[string]node
	maps = make(map[string]node)
	var currents []string

	for _, line := range lines {
		var key string
		var right string
		var left string
		fmt.Sscanf(line, "%s = (%s %s)", &key, &left, &right)
		left = left[:len(left)-1]
		right = right[:len(right)-1]
		if key[len(key)-1] == 'A' {
			currents = append(currents, key)
		}
		maps[key] = node{Left: left, Right: right}
	}

	currentInstructions := 0
	numberOfInstructionsExecute := 0
	numberOfGoodNodes := 0
	numberOfNodes := len(currents)
	for numberOfNodes != numberOfGoodNodes {
		numberOfGoodNodes = 0
		var newCurrents []string
		for _, current := range currents {
			if instructions[currentInstructions] == 'L' {
				newCurrents = append(newCurrents, maps[current].Left)
				if maps[current].Left[len(maps[current].Left)-1] == 'Z' {
					numberOfGoodNodes++
				}
			} else {
				newCurrents = append(newCurrents, maps[current].Right)
				if maps[current].Right[len(maps[current].Right)-1] == 'Z' {
					numberOfGoodNodes++
				}
			}
		}
		currents = newCurrents
		currentInstructions++
		if currentInstructions == len(instructions) {
			currentInstructions = 0
		}
		numberOfInstructionsExecute++
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

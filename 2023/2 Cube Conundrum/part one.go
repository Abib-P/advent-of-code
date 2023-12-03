package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

type Game struct {
	id    int
	blue  int
	red   int
	green int
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

	var games []Game

	for _, line := range lines {
		user := line[:strings.IndexByte(line, ':')]
		user = user[strings.LastIndexByte(user, ' ')+1:]

		var game Game

		game.id, _ = strconv.Atoi(user)

		line = line[strings.IndexByte(line, ':')+1:]

		for len(line) > 0 {

			breakingCharIndex := strings.IndexAny(line, "; :,\n")
			numberAsString := line[:breakingCharIndex]
			for cleanString(numberAsString) == "" && len(line) > 0 {
				line = line[breakingCharIndex+1:]
				breakingCharIndex = strings.IndexAny(line, "; :,\n")
				numberAsString = line[:breakingCharIndex]
			}

			numberAsString = cleanString(numberAsString)
			line = line[breakingCharIndex+1:]
			breakingCharIndex = strings.IndexAny(line, "; :,\n")
			color := line
			if breakingCharIndex != -1 {
				color = line[:breakingCharIndex]
			}
			for cleanString(color) == "" && len(line) > 0 {
				line = line[breakingCharIndex+1:]
				breakingCharIndex = strings.IndexAny(line, "; :,\n")
				color = line[:breakingCharIndex]
			}

			if breakingCharIndex == -1 {
				line = ""
			} else {
				line = line[breakingCharIndex+1:]
			}

			number, _ := strconv.Atoi(numberAsString)
			switch color {
			case "blue":
				if number > game.blue {
					game.blue = number
				}
			case "red":
				if number > game.red {
					game.red = number
				}
			case "green":
				if number > game.green {
					game.green = number
				}
			}
		}

		if gameIsPossible(game) {
			games = append(games, game)
		}
	}
	//sum of all ids
	var sum int
	for _, game := range games {
		sum += game.id
	}
	fmt.Println(sum)
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

func gameIsPossible(game Game) bool {
	if game.red > 12 || game.green > 13 || game.blue > 14 {
		return false
	}
	return true
}

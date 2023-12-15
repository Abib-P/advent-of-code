package main

import (
	"fmt"
	"os"
	"strconv"
)

type boat struct {
	Time     int
	Distance int
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

	var times []int
	lines[0] = lines[0][6:]
	actualNumber := ""
	for _, line := range lines[0] {
		if line >= '0' && line <= '9' {
			actualNumber += string(line)
		} else if actualNumber != "" {
			number, _ := strconv.Atoi(actualNumber)
			times = append(times, number)
			actualNumber = ""
		}
	}
	if actualNumber != "" {
		number, _ := strconv.Atoi(actualNumber)
		times = append(times, number)
		actualNumber = ""
	}

	var distances []int
	lines[1] = lines[1][9:]
	for _, line := range lines[1] {
		if line >= '0' && line <= '9' {
			actualNumber += string(line)
		} else if actualNumber != "" {
			number, _ := strconv.Atoi(actualNumber)
			distances = append(distances, number)
			actualNumber = ""
		}
	}
	if actualNumber != "" {
		number, _ := strconv.Atoi(actualNumber)
		distances = append(distances, number)
		actualNumber = ""
	}

	var boats []boat
	for i := 0; i < len(times); i++ {
		boats = append(boats, boat{times[i], distances[i]})
	}

	var numberOfPossibleBoats []int
	for _, boat := range boats {
		possibilities := calculateMaxDistance(boat)
		if possibilities != -1 {
			if !isEven(boat.Time) {
				numberOfPossibleBoats = append(numberOfPossibleBoats, (boat.Time+1)-(possibilities*2))
			} else {
				numberOfPossibleBoats = append(numberOfPossibleBoats, (boat.Time+1)-(possibilities*2))
			}
		}
	}

	finalResult := 1
	for _, numberOfPossibleBoat := range numberOfPossibleBoats {
		finalResult *= numberOfPossibleBoat
	}

	println(finalResult)
}

func isEven(possibilities int) bool {
	if possibilities%2 == 0 {
		return true
	}
	return false
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

func calculateMaxDistance(boat boat) int {
	var chargeTheBoatTime int
	chargeTheBoatTime = 0
	for chargeTheBoatTime < boat.Time {
		numberOfMilimetterPerMilisecond := chargeTheBoatTime
		distance := numberOfMilimetterPerMilisecond * (boat.Time - chargeTheBoatTime)
		if distance > boat.Distance {
			return chargeTheBoatTime
		}
		chargeTheBoatTime++
	}
	return -1
}

package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

type boat struct {
	Time     int
	Distance int
}

// type de fichier :
//32T3K 765
//T55J5 684
//KK677 28
//KTJJT 220
//QQQJA 483

type hand struct {
	Id   int
	Hand string
	Type int
}

// declare all possible card here
const listCard = "23456789TJQKA"

// declare const here
const fiveOfAKind = 7
const fourOfAKind = 6
const fullHouse = 5
const threeOfAKind = 4
const twoPairs = 3
const onePair = 2
const highCard = 1

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

	var hands []hand
	for _, line := range lines {
		var newHand hand
		newHand.Hand = line[:5]
		line = line[6:]
		newHand.Id, _ = strconv.Atoi(line)
		newHand.Type = calculateType(newHand.Hand)
		hands = append(hands, newHand)
	}

	//sort hands by type, if type is the same, compare first card then second card...
	/*for i := 0; i < len(hands); i++ {
		for j := i + 1; j < len(hands); j++ {
			if hands[i].Type < hands[j].Type {
				hands[i], hands[j] = hands[j], hands[i]
			} else if hands[i].Type == hands[j].Type {
				for k := 0; k < 5; k++ {
					if strings.IndexByte(listCard, hands[i].Hand[k]) < strings.IndexByte(listCard, hands[j].Hand[k]) {
						hands[i], hands[j] = hands[j], hands[i]
						break
					} else if strings.IndexByte(listCard, hands[i].Hand[k]) > strings.IndexByte(listCard, hands[j].Hand[k]) {
						break
					}
				}
			}
		}
	}*/
	//this but not from 7 to 1 but from 1 to 7
	for i := 0; i < len(hands); i++ {
		for j := i + 1; j < len(hands); j++ {
			if hands[i].Type > hands[j].Type {
				hands[i], hands[j] = hands[j], hands[i]
			} else if hands[i].Type == hands[j].Type {
				for k := 0; k < 5; k++ {
					if strings.IndexByte(listCard, hands[i].Hand[k]) > strings.IndexByte(listCard, hands[j].Hand[k]) {
						hands[i], hands[j] = hands[j], hands[i]
						break
					} else if strings.IndexByte(listCard, hands[i].Hand[k]) < strings.IndexByte(listCard, hands[j].Hand[k]) {
						break
					}
				}
			}
		}
	}

	fmt.Println(hands)

	var finalResult = 0

	for i := 0; i < len(hands); i++ {
		finalResult += hands[i].Id * (i + 1)
	}

	fmt.Println(finalResult)
}

func calculateType(hand string) int {
	//calculate type that depend on the number of pairs
	var numberOfEachCard [13]int
	for _, card := range hand {
		//use strings.IndexByte
		numberOfEachCard[strings.IndexByte(listCard, byte(card))]++
	}
	var numberOfPairs int
	var numberOfThree int
	var numberOfFour int
	for _, number := range numberOfEachCard {
		switch number {
		case 2:
			numberOfPairs++
		case 3:
			numberOfThree++
		case 4:
			numberOfFour++
		case 5:
			return fiveOfAKind
		}
	}
	if numberOfFour == 1 {
		return fourOfAKind
	}
	if numberOfThree == 1 && numberOfPairs == 1 {
		return fullHouse
	}
	if numberOfThree == 1 {
		return threeOfAKind
	}
	if numberOfPairs == 2 {
		return twoPairs
	}
	if numberOfPairs == 1 {
		return onePair
	}
	return highCard
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

package day4

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type Card struct {
	numbers []int
}

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func Day4(input string) {
	f, err := os.Open(input)
	check(err)
	s := bufio.NewScanner(f)
	fmt.Printf("Day 4 - Part I\n")
	wCards, sCards := parseCards(s)
	points := countScratchCardPoints(wCards, sCards)
	fmt.Printf("Points sum = %d\n", points)
}

func parseCards(s *bufio.Scanner) (winningCards []*Card, scratchCards []*Card) {
	var index int

	for s.Scan() {
		index++
		// Card 1: 41 48 83 86 17 | 83 86  6 31 17  9 48 53
		text := s.Text()
		fmt.Printf("%s\n", text)
		// [Card 1][41 48 83 86 17 | 83 86  6 31 17  9 48 53]
		s1 := strings.Split(text, ":")
		// [41 48 83 86 17][83 86  6 31 17  9 48 53]
		s2 := strings.Split(s1[1], "|")
		// [41][48][83][86][17]
		wNumbers := toIntSlice(strings.Fields(s2[0]))
		// [83][86][6][31][17][9][48][53]
		sNumbers := toIntSlice(strings.Fields(s2[1]))

		winningCards = append(winningCards, &Card{wNumbers})
		scratchCards = append(scratchCards, &Card{sNumbers})
	}

	return winningCards, scratchCards
}

func countScratchCardPoints(winningCards []*Card, scratchCards []*Card) int {
	var sumCardPoints int

	for i, sn := range scratchCards {
		var cardPoints int
		for _, sc := range sn.numbers {
			if contains(winningCards[i].numbers, sc) {
				if cardPoints != 0 {
					cardPoints = cardPoints * 2
				} else {
					cardPoints = 1
				}
			}
		}

		fmt.Printf("Card %d is worth %d point(s)\n", i+1, cardPoints)
		sumCardPoints += cardPoints
	}

	return sumCardPoints
}

func toIntSlice(input []string) (output []int) {
	for _, v := range input {
		value, err := strconv.Atoi(v)
		if err != nil {
			check(err)
		}
		output = append(output, value)
	}

	return
}

func contains(s []int, value int) bool {
	for _, v := range s {
		if v == value {
			return true
		}
	}

	return false
}

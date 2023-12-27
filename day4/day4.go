package day4

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

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
	points := parseCards(s)
	fmt.Printf("Points sum = %d\n", points)
}

func parseCards(s *bufio.Scanner) int {
	var points int
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

		var sumCardPoints int

		for _, sn := range sNumbers {
			if contains(wNumbers, sn) {
				if sumCardPoints != 0 {
					sumCardPoints = sumCardPoints * 2
				} else {
					sumCardPoints = 1
				}
			}
		}
		fmt.Printf("Card %d is worth %d point(s)\n", index, sumCardPoints)
		points += sumCardPoints
	}

	return points
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

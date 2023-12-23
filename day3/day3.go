package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

const input string = "day3/example-input.txt"

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func main() {
	fmt.Println("Day 3")
	f, err := os.Open(input)
	check(err)
	s := bufio.NewScanner(f)
	partOne(s)
}

func partOne(s *bufio.Scanner) {
	fmt.Println("Part I")

	for s.Scan() {
		text := s.Text()
		var current string

		for i := 0; i < len(text); i++ {
			str := string(text[i])
			_, err := strconv.Atoi(str)
			if err != nil {
				// check(err)
			}
			current += str
		}

		fmt.Printf("Number = %s", current)
	}
}

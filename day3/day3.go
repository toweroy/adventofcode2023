package main

import (
	"bufio"
	"fmt"
	"os"
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

	// for s.Scan() {
	// 	text := s.Text()
	// }

	// fmt.Printf("SUM = %d", sum)
}
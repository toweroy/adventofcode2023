package day2

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

const input string = "day2/input.txt"

func check(e error) {
	if e != nil {
		panic(e)
	}
}
// 12 red cubes, 13 green cubes, and 14 blue cubes
const RED int = 12
const GREEN int = 13
const BLUE int = 14

func main() {
	fmt.Println("Day 2")
	f, err := os.Open(input)
	check(err)
	s := bufio.NewScanner(f)
	// partOne(s)
	partTwo(s)
}

func partTwo(s *bufio.Scanner) {
	fmt.Println("Part II")
	var sum int

	for s.Scan() {
		text := s.Text()
		// [Game 1:] [7 blue, 9 red, 1 green; 8 green; 10 green, 5 blue, 3 red; 11 blue, 5 red, 1 green]
		s1 := strings.Split(text, ":")
		// Game 1
		gameId, err := strconv.Atoi(strings.Fields(s1[0])[1])
		if err != nil {
			check(err)
		}
		// 7 blue, 9 red, 1 green; 8 green; 10 green, 5 blue, 3 red; 11 blue, 5 red, 1 green
		s2 := s1[1]
		// [7 blue, 9 red, 1 green] [8 green] [10 green, 5 blue, 3 red] [11 blue, 5 red, 1 green]
		draws := strings.Split(s2, ";")
		
		sum += minimumCubes(draws, gameId)
	}

	fmt.Printf("SUM = %d", sum)
}

func minimumCubes(draws []string, gameId int) (int) {
	var lRed, lBlue, lGreen int

	for _, d := range draws {
		colors := strings.Split(d, ",")
		colorCount := make(map[string]int)

		for _, c := range colors {
			result := strings.TrimSpace(c)
			pair := strings.Fields(result)
			number, err := strconv.Atoi(pair[0])
			if err != nil {
				check(err)
			}
			color := strings.TrimSpace(pair[1])
			colorCount[color] = number
		}

		if colorCount["red"] > lRed {
			lRed = colorCount["red"]
		}
		if colorCount["blue"] > lBlue {
			lBlue = colorCount["blue"]
		}
		if colorCount["green"] > lGreen {
			lGreen = colorCount["green"]
		}
	}

	fmt.Printf("Game %d: %d red, %d blue, %d green\n", gameId, lRed, lBlue, lGreen)
	return lRed * lBlue * lGreen
}

func partOne(s *bufio.Scanner) {
	fmt.Println("Part I")
	var validGameIdSums int

	for s.Scan() {
		text := s.Text()
		// [Game 1:] [7 blue, 9 red, 1 green; 8 green; 10 green, 5 blue, 3 red; 11 blue, 5 red, 1 green]
		s1 := strings.Split(text, ":")
		// Game 1
		gameId, err := strconv.Atoi(strings.Fields(s1[0])[1])
		if err != nil {
			check(err)
		}
		// 7 blue, 9 red, 1 green; 8 green; 10 green, 5 blue, 3 red; 11 blue, 5 red, 1 green
		s2 := s1[1]
		// [7 blue, 9 red, 1 green] [8 green] [10 green, 5 blue, 3 red] [11 blue, 5 red, 1 green]
		draws := strings.Split(s2, ";")
		
		if  validGameId(draws, gameId) {
			validGameIdSums += gameId
		}
	}

	fmt.Printf("SUM = %d", validGameIdSums)
}

func validGameId(draws []string, gameId int) (bool) {
	for i, d := range draws {
		colors := strings.Split(d, ",")
		colorCount := make(map[string]int)

		for _, c := range colors {
			result := strings.TrimSpace(c)
			pair := strings.Fields(result)
			number, err := strconv.Atoi(pair[0])
			if err != nil {
				check(err)
			}
			color := strings.TrimSpace(pair[1])
			colorCount[color] = number
		}

		if colorCount["red"] <= RED && colorCount["blue"] <= BLUE && colorCount["green"] <= GREEN {
			fmt.Print("✅ ")
		} else {
			fmt.Print("❌ ")
			fmt.Printf("Game %d total draw %d = %d red, %d blue, %d green\n", gameId, i, colorCount["red"], colorCount["blue"], colorCount["green"])
			return false
		}
		fmt.Printf("Game %d total draw %d = %d red, %d blue, %d green\n", gameId, i, colorCount["red"], colorCount["blue"], colorCount["green"])
	}
	return true
}
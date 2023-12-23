package day3

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

const input string = "day3/example-input.txt"

type Location struct {
	row		int
	start	int
	end		int
}

type PartNumber struct {
	number 		int
	*Location
}

type Symbol struct {
	value		string
	*Location
}

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
	pn, sym := parsePartNumbers(s)
	fmt.Printf("%d numbers parsed\n", len(pn))
	fmt.Printf("%d symbols parsed\n", len(sym))
}

func parsePartNumbers(s *bufio.Scanner) ([]*PartNumber, []*Symbol) {
	fmt.Println("Part I")
	parts := []*PartNumber{}
	symbols := []*Symbol{}
	row := 0

	for s.Scan() {
		text := s.Text()
		var prev string
		var prevWasInt bool

		for i := 0; i < len(text); i++ {
			str := string(text[i])
			_, err := strconv.Atoi(str)

			if err == nil {
				fmt.Printf("Found number %s", str)
				
				if prev == "" || !prevWasInt {
					prev = str
				} else {
					fmt.Printf(" appending %s", str)
					prev += str
				}
				fmt.Printf("\n")
				prevWasInt = true
			} else if prevWasInt {
				partNumber, err := strconv.Atoi(prev)
				if err != nil {
					check(err)
				}
				parts = append(parts, &PartNumber{
					number: partNumber,
					Location: &Location{
						row,
						i,
						i + len(str),
					},
				})
				prevWasInt = false
				
				if str != "." {
					fmt.Printf("Found symbol %s\n", str)
					prev = str
					symbols = append(symbols, &Symbol{
						value:   str,
						Location: &Location{
							row,
							i,
							i,
						},
					})
				}
			} else if str != "." {
				prevWasInt = false
				fmt.Printf("Found symbol %s\n", str)
				prev = str
				symbols = append(symbols, &Symbol{
					value:   str,
					Location: &Location{
						row,
						i,
						i,
					},
				})
			}
		}
		
		prevWasInt = false
		prev = ""
		row++
	}

	return parts, symbols
}

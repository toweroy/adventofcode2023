package day3

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

type Location struct {
	start int
	end   int
}

type PartNumber struct {
	number int
	*Location
}

type Symbol struct {
	value 		string
	*Location
	gParts		[]*PartNumber
}

func check(e error) {
	if e != nil {
		panic(e)
	}
}

// Get all symbols [^\.|^(\d)] = 759
// Get all numbers (\d+) = 1211
func Day3(input string) {
	f, err := os.Open(input)
	check(err)
	s := bufio.NewScanner(f)
	pn, sym := parsePartNumbers(s)
	fmt.Printf("%d numbers parsed\n", len(pn))
	fmt.Printf("%d symbols parsed\n", len(sym))
	vPn := getValidPartNumbers(pn, sym)

	var sumPn int
	for _, v := range vPn {
		sumPn += v.number
	}
	fmt.Printf("Sum of part numbers = %d\n", sumPn)

	var sumGears int
	for _, v := range sym {
		for _, v := range v {
			if v.gParts != nil && len(v.gParts) == 2 {
				ratio := v.gParts[0].number * v.gParts[1].number
				fmt.Printf("Gear found: adding %d * %d = %d to sum\n", v.gParts[0].number, v.gParts[1].number, ratio)
				sumGears += ratio
			}
		}
	}
	fmt.Printf("Sum of gears = %d\n", sumGears)
}

func parsePartNumbers(s *bufio.Scanner) (map[int][]*PartNumber, map[int][]*Symbol) {
	fmt.Println("Part I")
	var totalParts int
	var totalSymbols int
	parts := map[int][]*PartNumber{}
	symbols := map[int][]*Symbol{}
	row := 0

	for s.Scan() {
		text := s.Text()
		var prev string
		var prevWasInt bool

		for i := 0; i < len(text); i++ {
			str := string(text[i])
			_, err := strconv.Atoi(str)

			if err == nil {
				// fmt.Printf("Found number %s", str)

				if prev == "" || !prevWasInt {
					prev = str
				} else {
					// fmt.Printf(" appending %s", str)
					prev += str
				}
				// fmt.Printf("\n")
				if i == (len(text) - 1) {
					partNumber, err := strconv.Atoi(prev)
					if err != nil {
						check(err)
					}
					newP := &PartNumber{
						number: partNumber,
						Location: &Location{
							i - len(prev),
							i - 1,
						}}
					if _, hasKey := parts[row]; !hasKey {
						parts[row] = []*PartNumber{newP}
					} else {
						parts[row] = append(parts[row], newP)
					}
					totalParts++
				}
				prevWasInt = true
			} else if prevWasInt {
				partNumber, err := strconv.Atoi(prev)
				if err != nil {
					check(err)
				}
				newP := &PartNumber{
					number: partNumber,
					Location: &Location{
						i - len(prev),
						i - 1,
					}}
				if _, hasKey := parts[row]; !hasKey {
					parts[row] = []*PartNumber{newP}
				} else {
					parts[row] = append(parts[row], newP)
				}
				
				if str != "." {
					// fmt.Printf("Found symbol %s\n", str)
					prev = str
					newSym := &Symbol{
						value: str,
						Location: &Location{
							i,
							i,
						},
					}
					if _, hasKey := symbols[row]; !hasKey {
						symbols[row] = []*Symbol{newSym}
						} else {
							symbols[row] = append(symbols[row], newSym)
						}
						totalSymbols++
					}

				totalParts++
				prevWasInt = false
			} else if str != "." {
				prevWasInt = false
				// fmt.Printf("Found symbol %s\n", str)
				prev = str
				newSym := &Symbol{
					value: str,
					Location: &Location{
						i,
						i,
					},
				}
				if _, hasKey := symbols[row]; !hasKey {
					symbols[row] = []*Symbol{newSym}
				} else {
					symbols[row] = append(symbols[row], newSym)
				}
				totalSymbols++
			}
		}

		prevWasInt = false
		prev = ""
		row++
	}

	fmt.Printf("Found %d numbers\n", totalParts)
	fmt.Printf("Found %d symbols\n", totalSymbols)
	return parts, symbols
}

func getValidPartNumbers(parts map[int][]*PartNumber, symbols map[int][]*Symbol) []*PartNumber {
	var validParts []*PartNumber
	for key, value := range parts {
		for _, v := range value {
			fmt.Printf("Part number: %d in row %d", v.number, key)
			sConnected := getSymbolsConnected(key, v, symbols)
			if len(sConnected) > 0 {
				fmt.Printf(" is connected ✅\n")
				validParts = append(validParts, v)
			} else {
				fmt.Printf(" not connected ❌\n")
			}
		}
	}
	fmt.Printf("Found %d valid part numbers\n", len(validParts))
	return validParts
}

func getSymbolsConnected(row int, part *PartNumber, symbols map[int][]*Symbol) ([]*Symbol) {
	r := []int{-1, 0, 1}
	var gSymbols []*Symbol

	if row == 0 {
		r = []int{0, 1}
	}

	for _, v := range r {
		if s, hasKey := symbols[row + v]; hasKey {
			for i := 0; i < len(s); i++ {
				if s[i].start >= part.start - 1 && s[i].start <= part.end + 1 {
					fmt.Printf("Found connnected symbol in row %d, index %d\n", row + v, i)

					if s[i].value == "*" {
						fmt.Printf("Symbol is a gear (*) adding part %d, row %d, index %d\n", part.number, row, part.start)
						if s[i].gParts != nil && len(s[i].gParts) > 0 {
							s[i].gParts = append(s[i].gParts, part)
						} else {
							s[i].gParts = []*PartNumber{part}
						}
					}
					gSymbols = append(gSymbols, s[i])
				}
			}
		}
	}

	return gSymbols
}

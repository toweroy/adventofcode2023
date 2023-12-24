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
	value string
	*Location
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

	// var allPn []int
	// for _, v := range pn {
	// 	for _, part := range v {
	// 		allPn = append(allPn, part.number)
	// 	}
	// }

	var sum int
	// for _, v := range allPn {
	// 	if slices.Contains(vPn, v) {
	// 		sum += v
	// 	}
	// }

	for _, v := range vPn {
		sum += v
	}
	fmt.Printf("SUM = %d\n", sum)
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

func getValidPartNumbers(parts map[int][]*PartNumber, symbols map[int][]*Symbol) []int {
	var validParts []int
	for key, value := range parts {
		for _, v := range value {
			fmt.Printf("Part number: %d in row %d", v.number, key)
			hasSymbol := hasSymbolAround(key, v, symbols)
			if hasSymbol {
				fmt.Printf(" is connected ✅\n")
				// fmt.Printf("%d\n", v.number)
				validParts = append(validParts, v.number)
			} else {
				fmt.Printf(" not connected ❌\n")
			}
		}
	}
	fmt.Printf("Found %d valid part numbers\n", len(validParts))
	return validParts
}

func hasSymbolAround(row int, part *PartNumber, symbols map[int][]*Symbol) bool {
	r := []int{-1, 0, 1}

	if row == 0 {
		r = []int{0, 1}
	}

	for _, v := range r {
		if s, hasKey := symbols[row + v]; hasKey {
			for i := 0; i < len(s); i++ {
				if s[i].start >= part.start - 1 && s[i].start <= part.end + 1 {
					return true
				}
			}
		}
	}

	return false
}

package day3

import (
	"bufio"
	"fmt"
	"os"
	"slices"
	"strconv"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

var _ = Describe("Parsing numbers", func() {
	var expectedParts = map[int][]*PartNumber{
		0: {
			&PartNumber{
				number: 467,
				Location: &Location{
					start: 0,
					end:   2,
				},
			},
			&PartNumber{
				number: 114,
				Location: &Location{
					start: 5,
					end:   7,
				},
			},
		},
		2: {
			&PartNumber{
				number: 35,
				Location: &Location{
					start: 2,
					end:   3,
				},
			},
			&PartNumber{
				number: 633,
				Location: &Location{
					start: 6,
					end:   8,
				},
			},
		},
		4: {
			&PartNumber{
				number: 617,
				Location: &Location{
					start: 0,
					end:   2,
				},
			},
		},
		5: {
			&PartNumber{
				number: 58,
				Location: &Location{
					start: 7,
					end:   8,
				},
			},
		},
		6: {
			&PartNumber{
				number: 592,
				Location: &Location{
					start: 2,
					end:   4,
				},
			},
		},
		7: {
			&PartNumber{
				number: 755,
				Location: &Location{
					start: 6,
					end:   8,
				},
			},
		},
		9: {
			&PartNumber{
				number: 664,
				Location: &Location{
					start: 1,
					end:   3,
				},
			},
			&PartNumber{
				number: 598,
				Location: &Location{
					start: 5,
					end:   7,
				},
			},
		},
	}
	var expectedSymbols = map[int][]*Symbol{
		1: {
			&Symbol{
				value: "*",
				Location: &Location{
					start: 3,
					end:   3,
				},
			},
		},
		3: {
			&Symbol{
				value: "#",
				Location: &Location{
					start: 6,
					end:   6,
				},
			},
		},
		4: {
			&Symbol{
				value: "*",
				Location: &Location{
					start: 3,
					end:   3,
				},
			},
		},
		5: {
			&Symbol{
				value: "+",
				Location: &Location{
					start: 5,
					end:   5,
				},
			},
		},
		8: {
			&Symbol{
				value: "$",
				Location: &Location{
					start: 3,
					end:   3,
				},
			},
			&Symbol{
				value: "*",
				Location: &Location{
					start: 5,
					end:   5,
				},
			},
		},
	}
	var symbols map[int][]*Symbol
	var parts map[int][]*PartNumber

	BeforeEach(func() {
		f, err := os.Open("example-input.txt")
		Expect(err).To(BeNil())
		s := bufio.NewScanner(f)
		parts, symbols = parsePartNumbers(s)
	})

	It("can get proper numbers", func() {
		Expect(parts).To(HaveLen(7))
		Expect(symbols).To(HaveLen(5))
		Expect(parts[0]).To(HaveExactElements(expectedParts[0]))
		Expect(parts[2]).To(HaveExactElements(expectedParts[2]))
		Expect(parts[4]).To(HaveExactElements(expectedParts[4]))
		Expect(parts[5]).To(HaveExactElements(expectedParts[5]))
		Expect(parts[6]).To(HaveExactElements(expectedParts[6]))
		Expect(parts[7]).To(HaveExactElements(expectedParts[7]))
		Expect(parts[9]).To(HaveExactElements(expectedParts[9]))
		Expect(symbols[1]).To(HaveExactElements(expectedSymbols[1]))
		Expect(symbols[3]).To(HaveExactElements(expectedSymbols[3]))
		Expect(symbols[4]).To(HaveExactElements(expectedSymbols[4]))
		Expect(symbols[5]).To(HaveExactElements(expectedSymbols[5]))
		Expect(symbols[8]).To(HaveExactElements(expectedSymbols[8]))
	})

	It("a part number with a symbol is detected", func() {
		// 467
		cSymbols := getSymbolsConnected(0, parts[0][0], expectedSymbols)
		Expect(cSymbols).To(HaveLen(1))
		// 114
		cSymbols = getSymbolsConnected(0, parts[0][1], expectedSymbols)
		Expect(cSymbols).To(HaveLen(0))
		// 35
		cSymbols = getSymbolsConnected(2, parts[2][0], expectedSymbols)
		Expect(cSymbols).To(HaveLen(1))
		// 633
		cSymbols = getSymbolsConnected(2, parts[2][1], expectedSymbols)
		Expect(cSymbols).To(HaveLen(1))
		// 617
		cSymbols = getSymbolsConnected(4, parts[4][0], expectedSymbols)
		Expect(cSymbols).To(HaveLen(1))
		// 58
		cSymbols = getSymbolsConnected(5, parts[5][0], expectedSymbols)
		Expect(cSymbols).To(HaveLen(0))
		// 592
		cSymbols = getSymbolsConnected(6, parts[6][0], expectedSymbols)
		Expect(cSymbols).To(HaveLen(1))
		// 755
		cSymbols = getSymbolsConnected(7, parts[7][0], expectedSymbols)
		Expect(cSymbols).To(HaveLen(1))
		// 664
		cSymbols = getSymbolsConnected(9, parts[9][0], expectedSymbols)
		Expect(cSymbols).To(HaveLen(1))
		// 598
		cSymbols = getSymbolsConnected(9, parts[9][1], expectedSymbols)
		Expect(cSymbols).To(HaveLen(1))
	})

	It("gets all part numbers", func() {
		f, err := os.Open("input.txt")
		Expect(err).To(BeNil())
		s := bufio.NewScanner(f)
		parts, symbols = parsePartNumbers(s)

		f, err = os.Open("all-parts.txt")
		Expect(err).To(BeNil())
		s = bufio.NewScanner(f)

		var expectedParts []int
		for s.Scan() {
			text := s.Text()
			number, err := strconv.Atoi(text)
			Expect(err).To(BeNil())
			expectedParts = append(expectedParts, number)
		}

		Expect(expectedParts).To(HaveLen(1211))
		fmt.Printf("Parsed %d number from all-parts.txt", len(expectedParts))

		var allPn []int
		for _, v := range parts {
			for _, part := range v {
				allPn = append(allPn, part.number)
			}
		}

		Expect(allPn).To(HaveLen(1211))
		for _, v := range expectedParts {
			if !slices.Contains(allPn, v) {
				fmt.Printf("Expect part not found %d\n", v)
			}
		}

		Expect(allPn).To(ContainElements(expectedParts))
	})
})

package day3

import (
	"bufio"
	"os"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

var _ = Describe("Parsing numbers", func() {
	var expectedParts = []*PartNumber{
		{
			number: 467,
			Location: &Location{
				row:   0,
				start: 0,
				end:   2,
			},
		},
		{
			number: 114,
			Location: &Location{
				row:   0,
				start: 5,
				end:   7,
			},
		},
		{
			number: 35,
			Location: &Location{
				row:   2,
				start: 2,
				end:   3,
			},
		},
		{
			number: 633,
			Location: &Location{
				row:   2,
				start: 6,
				end:   8,
			},
		},
		{
			number: 617,
			Location: &Location{
				row:   4,
				start: 0,
				end:   2,
			},
		},
		{
			number: 58,
			Location: &Location{
				row:   5,
				start: 7,
				end:   8,
			},
		},
		{
			number: 592,
			Location: &Location{
				row:   6,
				start: 2,
				end:   4,
			},
		},
		{
			number: 755,
			Location: &Location{
				row:   7,
				start: 6,
				end:   8,
			},
		},
		{
			number: 664,
			Location: &Location{
				row:   9,
				start: 1,
				end:   3,
			},
		},
		{
			number: 598,
			Location: &Location{
				row:   9,
				start: 5,
				end:   7,
			},
		},
	}
	var expectedSymbols = []*Symbol{
		{
			value: "*",
			Location: &Location{
				row:   1,
				start: 3,
				end:   3,
			},
		},
		{
			value: "#",
			Location: &Location{
				row:   3,
				start: 6,
				end:   6,
			},
		},
		{
			value: "*",
			Location: &Location{
				row:   4,
				start: 3,
				end:   3,
			},
		},
		{
			value: "+",
			Location: &Location{
				row:   5,
				start: 5,
				end:   5,
			},
		},
		{
			value: "$",
			Location: &Location{
				row:   8,
				start: 3,
				end:   3,
			},
		},
		{
			value: "*",
			Location: &Location{
				row:   8,
				start: 5,
				end:   5,
			},
		},
	}
	var symbols []*Symbol
	var parts []*PartNumber

	BeforeEach(func() {
		f, err := os.Open("example-input.txt")
		Expect(err).To(BeNil())
		s := bufio.NewScanner(f)
		parts, symbols = parsePartNumbers(s)
	})

	It("can get proper numbers", func() {
		Expect(parts).To(HaveLen(10))
		Expect(symbols).To(HaveLen(6))
		Expect(parts).To(HaveExactElements(expectedParts))
		Expect(symbols).To(HaveExactElements(expectedSymbols))
	})
})

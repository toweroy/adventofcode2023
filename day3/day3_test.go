package day3

import (
	"bufio"
	"os"

	. "github.com/onsi/ginkgo/v2"
    . "github.com/onsi/gomega"
)

var _ = Describe("Parsing numbers", func() {
	var parts []*PartNumber
	var symbols []*Symbol

	BeforeEach(func() {
		f, err := os.Open("example-input.txt")
		Expect(err).To(BeNil())
		s := bufio.NewScanner(f)
		parts, symbols = parsePartNumbers(s)
	})

	It("can get proper numbers", func() {
		Expect(parts).To(HaveLen(10))
		Expect(symbols).To(HaveLen(6))
	})
})
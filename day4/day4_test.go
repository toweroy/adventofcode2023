package day4

import (
	"bufio"
	"os"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

var _ = Describe("Parsing cards", func() {
	var points int

	BeforeEach(func() {
		f, err := os.Open("example-input.txt")
		Expect(err).To(BeNil())
		s := bufio.NewScanner(f)
		points = parseCards(s)
	})

	It("can get sum of winning numbers", func() {
		Expect(points).To(Equal(13))
	})
})

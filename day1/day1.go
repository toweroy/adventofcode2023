package day1

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

const input string = "day1/input.txt"

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func getLetter(s string) int {
	if strings.HasPrefix(s, "one") {
		return 1
	}
	if strings.HasPrefix(s, "two") {
		return 2
	}
	if strings.HasPrefix(s, "three") {
		return 3
	}
	if strings.HasPrefix(s, "four") {
		return 4
	}
	if strings.HasPrefix(s, "five") {
		return 5
	}
	if strings.HasPrefix(s, "six") {
		return 6
	}
	if strings.HasPrefix(s, "seven") {
		return 7
	}
	if strings.HasPrefix(s, "eight") {
		return 8
	}
	if strings.HasPrefix(s, "nine") {
		return 9
	}

	return 0
}

func main() {
	fmt.Println("Day 1")
	f, err := os.Open(input)
	check(err)
	s := bufio.NewScanner(f)
	var sum int

	for s.Scan() {
		text := s.Text()
		var nums []int

		for i := 0; i < len(text); i++ {
			sub := text[i:]
			v := getLetter(sub)

			if v > 0 {
				nums = append(nums, v)
			} else {
				c := string(text[i])
				v, err := strconv.Atoi(c)
				if err == nil {
					nums = append(nums, v)
				}
			}
		}
		var first int
		var last int
		if len(nums) > 1 {
			first = nums[0] * 10
			last = nums[len(nums)-1]
		} else {
			first = nums[0] * 10
			last = nums[0]
		}
		coord := first + last
		sum += coord
		fmt.Printf("first %d, last %d = %d \n", first, last, coord)
	}

	fmt.Printf("SUM = %d", sum)
}

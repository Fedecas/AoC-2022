package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"strconv"
	"strings"
)

const sep = "\r\n"

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func convertToInt(s string) (v int) {
	v, err := strconv.Atoi(s)
	check(err)
	return
}

func getInput() (res [][][]int) {
	bytes, err := ioutil.ReadAll(os.Stdin)
	check(err)
	content := string(bytes)

	lines := strings.Split(content, sep)
	for i := 0; i < len(lines); i++ {
		sections := strings.Split(lines[i], ",")
		if len(sections) > 1 {
			s1 := strings.Split(sections[0], "-")
			start1, end1 := convertToInt(s1[0]), convertToInt(s1[1])
			s2 := strings.Split(sections[1], "-")
			start2, end2 := convertToInt(s2[0]), convertToInt(s2[1])

			res = append(res, [][]int{{start1, end1}, {start2, end2}})
		}
	}
	return
}

func areFullyContained(start1, end1, start2, end2 int) bool {
	return (start1 >= start2 && end1 <= end2) || (start1 <= start2 && end1 >= end2)
}

func countFullyContained(sections [][][]int) (res int) {
	for i := 0; i < len(sections); i++ {
		s1, s2 := sections[i][0], sections[i][1]
		start1, end1 := s1[0], s1[1]
		start2, end2 := s2[0], s2[1]

		if areFullyContained(start1, end1, start2, end2) {
			res++
		}
	}
	return
}

func areOverlappingAtStart(start1, end1, start2, end2 int) bool {
	return (start1 <= start2 && end1 >= start2) || (start2 <= start1 && end2 >= start1)
}

func areOverlappingAtEnd(start1, end1, start2, end2 int) bool {
	return (start1 >= start2 && start1 <= end2) || (start2 >= start1 && start2 <= end1)
}

func areOverlapping(start1, end1, start2, end2 int) bool {
	return areFullyContained(start1, end1, start2, end2) ||
		areOverlappingAtStart(start1, end1, start2, end2) ||
		areOverlappingAtEnd(start1, end1, start2, end2)
}

func countOverlapping(sections [][][]int) (res int) {
	for i := 0; i < len(sections); i++ {
		s1, s2 := sections[i][0], sections[i][1]
		start1, end1 := s1[0], s1[1]
		start2, end2 := s2[0], s2[1]

		if areOverlapping(start1, end1, start2, end2) {
			res++
		}
	}
	return
}

func main() {
	sections := getInput()
	countContained := countFullyContained(sections)
	countOverlaps := countOverlapping(sections)

	fmt.Println("1:", countContained)
	fmt.Println("2:", countOverlaps)
}

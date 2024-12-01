package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
)

func main() {
	l0, l1 := readAndSortSpaceSeperatedLists("data")
	fmt.Println("Answer for Day1 Part1 is", sumOfDistance(l0, l1))
	fmt.Println("Answer for Day1 Part2 is", similarityScore(l0, l1))
}

// Ignoring error handling as its a controlled test cases
func readAndSortSpaceSeperatedLists(filename string) (l0, l1 []int) {
	f, _ := os.Open(filename)
	defer f.Close()

	in := bufio.NewScanner(f)
	for in.Scan() {
		line := strings.Fields(in.Text())
		d0, _ := strconv.Atoi(line[0])
		d1, _ := strconv.Atoi(line[1])

		l0 = append(l0, d0)
		l1 = append(l1, d1)
	}

	sort.Ints(l0)
	sort.Ints(l1)
	return l0, l1
}

func sumOfDistance(l0, l1 []int) int {
	var sum int
	for i := range l0 {
		sum += abs(l0[i] - l1[i])
	}
	return sum
}

func similarityScore(l0, l1 []int) int {
	occurenceInL1 := make(map[int]int)
	for _, v := range l1 {
		occurenceInL1[v]++
	}

	var score int
	for _, v := range l0 {
		score += v * occurenceInL1[v]
	}
	return score
}

func abs(a int) int {
	if a > 0 {
		return a
	}
	return -a
}

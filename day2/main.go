package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"
)

func main() {
	f, _ := os.Open("data")
	defer f.Close()
	fmt.Println("Answer for Day2 Part1 is", getSafeReports(f))
	f.Seek(0, 0)
	fmt.Println("Answer for Day2 Part2 is", getDampenSafeReports(f))
}

func getSafeReports(f io.Reader) int {
	var safeCount int

	s := bufio.NewScanner(f)
	for s.Scan() {
		lstr := strings.Fields(s.Text())
		var l []int

		for _, v := range lstr {
			num, _ := strconv.Atoi(v)
			l = append(l, num)
		}

		// controlled test data has more than 5 per list hence not checking Out of Bound
		// strictly increasing or strictly decreasing
		if isSafe(l, 1, 3) || isSafe(l, -3, -1) {
			safeCount++
		}

	}
	return safeCount
}

func getDampenSafeReports(f io.Reader) int {
	var safeCount int

	s := bufio.NewScanner(f)
	for s.Scan() {
		lstr := strings.Fields(s.Text())
		var l []int

		for _, v := range lstr {
			num, _ := strconv.Atoi(v)
			l = append(l, num)
		}

		if isDampenSafe(l, 1, 3) || isDampenSafe(l, -3, -1) {
			safeCount++
		}

	}
	return safeCount
}

func isSafe(list []int, llimit, ulimit int) bool {
	for i := 1; i < len(list); i++ {
		diff := list[i] - list[i-1]
		if diff > ulimit || diff < llimit {
			return false
		}
	}
	return true
}

func isDampenSafe(list []int, llimit, ulimit int) bool {
	var dampCount int
	i, j := 0, 1
	for j < len(list) {
		diff := list[j] - list[i]
		if diff > ulimit || diff < llimit {
			if dampCount > 0 {
				return false
			}
			dampCount++
			// --1. remove item at i  --2. remove item at j
			removeI := append(append([]int{}, list[:i]...), list[j:]...)
			removeJ := append(append([]int{}, list[:j]...), list[j+1:]...)
			return isSafe(removeI, llimit, ulimit) || isSafe(removeJ, llimit, ulimit)
		}
		i = j
		j++
	}
	return true
}

package main

import (
	"fmt"
	"os"
	"regexp"
	"strconv"
)

func main() {
	buf, _ := os.ReadFile("data")
	fmt.Println("Answer for Day3 Part1 is", getMultiplicationSum(buf))
	fmt.Println("Answer for Day3 Part2 is", getConditionalMultiplicationSum(buf))
}

func getMultiplicationSum(buf []byte) int {
	re := regexp.MustCompile(`mul\(\d+,\d+\)`)
	matches := re.FindAllString(string(buf), -1)
	if matches == nil {
		return 0
	}

	var res int
	for _, v := range matches {
		res += getDigitsAndMultiply(v)
	}

	return res
}

func getConditionalMultiplicationSum(buf []byte) int {
	enableStmt, disableStmt := "do()", "don't()"
	enabled := true

	re := regexp.MustCompile(`mul\(\d+,\d+\)|do\(\)|don't\(\)`)
	matches := re.FindAllString(string(buf), -1)
	if matches == nil {
		return 0
	}

	var res int
	for _, v := range matches {
		if v == enableStmt {
			enabled = true
			continue
		}

		if v == disableStmt {
			enabled = false
			continue
		}

		if enabled {
			res += getDigitsAndMultiply(v)
		}

	}
	return res
}

func getDigitsAndMultiply(v string) int {
	re := regexp.MustCompile(`\d+`)
	digits := re.FindAllString(v, -1)

	// due to regexp mul\(\d+,\d+\), guaranteed to have 2 numbers
	num0, _ := strconv.Atoi(digits[0])
	num1, _ := strconv.Atoi(digits[1])
	return num0 * num1
}

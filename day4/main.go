package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strings"
)

func main() {
	f, _ := os.Open("data")
	matrix := getInputMatrix(f)
	fmt.Println("Answer for Day4 Part1 is", getAllXMASOccurence(matrix))
	fmt.Println("Answer for Day4 Part2 is", getAllXShapeMASOccurence(matrix))
}

func getInputMatrix(f io.Reader) [][]rune {
	s := bufio.NewScanner(f)
	matrix := [][]rune{}
	for s.Scan() {
		lineArray := []rune{}
		lineArray = append(lineArray, []rune(s.Text())...)
		matrix = append(matrix, lineArray)
	}
	return matrix
}

func getAllXMASOccurence(matrix [][]rune) int {
	directions := [][]int{
		{0, 1},
		{1, 0},
		{1, 1},
		{1, -1},
		{0, -1},
		{-1, 0},
		{-1, 1},
		{-1, -1},
	}
	// just for readability :)
	x, y := 0, 1

	m, n := len(matrix), len(matrix[0])

	count := 0
	for i, line := range matrix {
		for j, c := range line {
			if c == 'X' {
				for _, dir := range directions {
					ni, nj := i+3*dir[x], j+3*dir[y]
					if ni >= 0 && ni < m && nj >= 0 && nj < n {
						word := string([]rune{
							matrix[i][j],
							matrix[i+dir[x]][j+dir[y]],
							matrix[i+2*dir[x]][j+2*dir[y]],
							matrix[i+3*dir[x]][j+3*dir[y]],
						})
						if word == "XMAS" {
							count++
						}
					}
				}
			}
		}
	}
	return count
}

func getAllXShapeMASOccurence(matrix [][]rune) int {
	m, n := len(matrix), len(matrix[0])
	count := 0
	for i, line := range matrix {
		for j, c := range line {
			if c == 'A' && i-1 >= 0 && i+1 < m && j-1 >= 0 && j+1 < n {

				diag1 := string([]rune{matrix[i-1][j-1], matrix[i][j], matrix[i+1][j+1]})
				diag2 := string([]rune{matrix[i-1][j+1], matrix[i][j], matrix[i+1][j-1]})

				if strings.Contains("MAS,SAM", diag1) && strings.Contains("MAS,SAM", diag2) {
					count++
				}
			}
		}
	}
	return count
}

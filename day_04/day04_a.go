package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	//Read input file
	input, _ := os.Open("input.txt")
	defer input.Close()
	sc := bufio.NewScanner(input)

	winningCardNums := make([][]int, 0)
	numbers := make([][]int, 0)
	line_num := 1
	totalPoints := 0
	for sc.Scan() {
		line := sc.Text()

		rhs := strings.Split(line, ":")[1]
		tmp := strings.Split(rhs, "|")
		winningNumsStr := tmp[0]
		numsStr := tmp[1]
		winningNums := make([]int, 0)
		for _, n := range strings.Fields(winningNumsStr) {
			num, _ := strconv.Atoi(n)
			winningNums = append(winningNums, num)
		}

		myNums := make([]int, 0)
		for _, n := range strings.Fields(numsStr) {
			num, _ := strconv.Atoi(n)
			myNums = append(myNums, num)
		}

		matches := intersect(winningNums, myNums)
		fmt.Printf("Line %d: matches: %v", line_num, matches)

		points := 0
		for _, _ = range matches {
			if points == 0 {
				points = 1
			} else {
				points = points * 2
			}
		}
		fmt.Printf(", points: %d\n", points)
		totalPoints += points

		winningCardNums = append(winningCardNums, winningNums)
		numbers = append(numbers, myNums)

		line_num++
	}
	fmt.Println(totalPoints)
}
func intersect(arr1, arr2 []int) []int {
	m := make(map[int]bool)
	var intersection []int

	// Store elements of arr1 in a map
	for _, v := range arr1 {
		m[v] = true
	}

	// Check if elements of arr2 are in the map
	for _, v := range arr2 {
		if _, ok := m[v]; ok {
			intersection = append(intersection, v)
			delete(m, v) // Ensure uniqueness
		}
	}

	return intersection
}

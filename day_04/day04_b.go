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
	input, _ := os.Open("input2.txt")
	defer input.Close()
	sc := bufio.NewScanner(input)

	lineNum := 1
	m := make(map[int]int)

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

		matches := intersectb(winningNums, myNums)
		m[lineNum] = len(matches)
		fmt.Printf("Line %d: matches: %v\n", lineNum, matches)

		lineNum++
	}

	cardCount := make(map[int]int)
	for i := 1; i < lineNum; i++ {
		cardCount[i] = 0
	}

	for i := 1; i < lineNum; i++ {
		cardCount[i]++
		additionalCards := m[i]
		for j := 1; j <= additionalCards; j++ {
			cardCount[i+j]++
		}
	}

	totalCards := 0
	for i := 1; i < lineNum; i++ {
		totalCards += cardCount[i]
	}
	fmt.Println(totalCards)
}

func intersectb(arr1, arr2 []int) []int {
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

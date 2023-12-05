package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
)

func main() {
	//Read input file
	input, _ := os.Open("input.txt")
	defer input.Close()
	sc := bufio.NewScanner(input)

	//var lines []string
	lines := make([]string, 139)
	symbol_positions := make([][]int, 139)

	/*
		Generate map of [line index](array of symbol index positions)

		Func for prev_line evaluation, current line, and next line

		For each line
		- Find numbers with start and end indices
		- For prev line (if 0th index, skip)
		    - Was symbol (non-digit and non ‘.’) present from start index-1 to end index+1
		- For cur line
		    - Is symbol present at start index-1 or end index+1
		- For next line
		    - Is symbol present from start index-1 to end index+1
	*/

	for sc.Scan() {
		line := sc.Text()
		lines = append(lines, line)

		positions := make([]int, 0)
		re, _ := regexp.Compile("[^a-zA-Z0-9 \\.]+")
		result := re.FindAllStringIndex(line, -1)
		for _, pos := range result {
			positions = append(positions, pos[0])
		}
	}

	fmt.Println("")
}

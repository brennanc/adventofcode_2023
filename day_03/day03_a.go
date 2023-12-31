package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strconv"
)

const NUM_INPUT_LINES = 140

func main() {
	//Read input file
	input, _ := os.Open("input.txt")
	defer input.Close()
	sc := bufio.NewScanner(input)

	//var lines []string
	lines := make([]string, 0)
	symbol_positions := make([][]int, 0)

	for sc.Scan() {
		line := sc.Text()
		lines = append(lines, line)

		positions := make([]int, 0)
		re, _ := regexp.Compile("[^a-zA-Z0-9 \\.]+")
		result := re.FindAllStringIndex(line, -1)
		for _, pos := range result {
			positions = append(positions, pos[0])
		}
		symbol_positions = append(symbol_positions, positions)
	}

	total_part_nums := 0
	part_num_sum := 0
	invalid_part_num_sum := 0
	for j := 0; j < NUM_INPUT_LINES; j++ {
		part_nums, invalid_part_nums := evaluateLine(lines, j, symbol_positions)
		for _, part_num := range part_nums {
			part_num_sum += part_num
			total_part_nums++
		}
		for _, inv_part_num := range invalid_part_nums {
			invalid_part_num_sum += inv_part_num
			total_part_nums++
		}
	}
	fmt.Printf("Valid sum: %d, Invalid sum: %d\n", part_num_sum, invalid_part_num_sum)
	fmt.Printf("Total part numbers (valid & invalid): %d\n", total_part_nums)
}

func evaluateLine(lines []string, index int, symbolPositions [][]int) ([]int, []int) {
	part_nums := make([]int, 0)
	invalid_part_nums := make([]int, 0)
	re, _ := regexp.Compile("(\\d+)")
	result := re.FindAllStringIndex(lines[index], -1)
	var isFirstLine bool = (index == 0)
	var isLastLine bool = (index == NUM_INPUT_LINES-1)
	for _, pos := range result { // for each potential part number
		start_idx := pos[0]
		end_idx := pos[1]
		isValidPartNum := false
		if (!isFirstLine) && (!isLastLine) {
			isValidPartNum = (evalAdjacentLine(symbolPositions[index-1], start_idx, end_idx) || evalCurLine(symbolPositions[index], start_idx, end_idx) || evalAdjacentLine(symbolPositions[index+1], start_idx, end_idx))
		} else if isFirstLine {
			isValidPartNum = evalCurLine(symbolPositions[index], start_idx, end_idx) || evalAdjacentLine(symbolPositions[index+1], start_idx, end_idx)
		} else if isLastLine {
			isValidPartNum = evalAdjacentLine(symbolPositions[index-1], start_idx, end_idx) || evalCurLine(symbolPositions[index], start_idx, end_idx)
		}
		if isValidPartNum {
			part_num, _ := strconv.Atoi(lines[index][pos[0]:pos[1]])
			part_nums = append(part_nums, part_num)
		} else {
			part_num, _ := strconv.Atoi(lines[index][pos[0]:pos[1]])
			invalid_part_nums = append(invalid_part_nums, part_num)
		}
	}
	fmt.Printf("LINE %d: VALID---%v\n", index, part_nums)
	fmt.Printf("LINE %d: INVALID-%v\n", index, invalid_part_nums)
	return part_nums, invalid_part_nums
}

func evalCurLine(symbolPositions []int, startIdx int, endIdx int) bool {
	for _, pos := range symbolPositions {
		if (pos == startIdx-1) || (pos == endIdx) {
			return true
		}
	}
	return false
}

func evalAdjacentLine(symbolPositions []int, startIdx int, endIdx int) bool {
	for _, pos := range symbolPositions {
		if (pos >= startIdx-1) && (pos <= endIdx) {
			return true
		}
	}
	return false
}

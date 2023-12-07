package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strconv"
)

const NUM_INPUT_LINES_PARTB = 140

func main() {
	input, _ := os.Open("input.txt")
	defer input.Close()
	sc := bufio.NewScanner(input)

	lines := make([]string, 0)
	symbol_positions := make([][]int, 0)
	num_positions := make([][]int, 0)
	part_nums := make([][]int, 0)

	for sc.Scan() {
		line := sc.Text()
		lines = append(lines, line)

		positions := make([]int, 0)
		num_pos := make([]int, 0)
		nums := make([]int, 0)
		re, _ := regexp.Compile("[\\*]+")
		result := re.FindAllStringIndex(line, -1)
		re2, _ := regexp.Compile("([\\d]+)")
		result2 := re2.FindAllStringIndex(line, -1)
		result3 := re2.FindAllStringSubmatch(line, -1)
		for _, pos := range result {
			positions = append(positions, pos[0])
		}
		for _, pos2 := range result2 {
			num_pos = append(num_pos, pos2[0])
			num_pos = append(num_pos, pos2[1])
		}
		for _, pos3 := range result3 {
			num, _ := strconv.Atoi(pos3[0])
			nums = append(nums, num)
		}
		symbol_positions = append(symbol_positions, positions)
		num_positions = append(num_positions, num_pos)
		part_nums = append(part_nums, nums)
	}

	product_sum := 0
	for j := 1; j < NUM_INPUT_LINES_PARTB-1; j++ { // start on second line and iterate through second-to-last line
		fmt.Printf("CURRENT LINE: %d\n", j+1)
		fmt.Printf("Part positions for prev line: %v\n", num_positions[j-1])
		fmt.Printf("Gear positions for curr line: %v\n", symbol_positions[j])
		fmt.Printf("Part positions for curr line: %v\n", num_positions[j])
		fmt.Printf("Part positions for next line: %v\n", num_positions[j+1])

		for _, gear := range symbol_positions[j] {
			result_set := make([]int, 0)
			nums_prev := findAdjacentLineNums(num_positions[j-1], part_nums[j-1], gear)
			fmt.Printf("Gear at position %d: prev line nums %v\n", gear, nums_prev)
			for _, num := range nums_prev {
				result_set = append(result_set, num)
			}

			nums_cur := findCurrentLineNums(num_positions[j], part_nums[j], gear)
			fmt.Printf("Gear at position %d: curr line nums %v\n", gear, nums_cur)
			for _, num := range nums_cur {
				result_set = append(result_set, num)
			}

			nums_next := findAdjacentLineNums(num_positions[j+1], part_nums[j+1], gear)
			fmt.Printf("Gear at position %d: next line nums %v\n", gear, nums_next)
			for _, num := range nums_next {
				result_set = append(result_set, num)
			}

			if len(result_set) < 2 {
				continue
			}

			product_sum += result_set[0] * result_set[1]
		}
		fmt.Println()
	}
	fmt.Println(product_sum)
}

func findCurrentLineNums(num_pos []int, part_nums []int, gear_idx int) []int {
	result_set := make([]int, 0)
	for i, j := 0, 0; i < len(num_pos); i, j = i+2, j+1 {
		start_idx := num_pos[i]
		end_idx := num_pos[i+1]
		if start_idx > gear_idx+1 {
			return result_set
		}
		if (start_idx-1 == gear_idx) || (end_idx == gear_idx) {
			result_set = append(result_set, part_nums[j])
		}
	}
	return result_set
}

func findAdjacentLineNums(num_pos []int, part_nums []int, gear_idx int) []int {
	result_set := make([]int, 0)
	for i, j := 0, 0; i < len(num_pos); i, j = i+2, j+1 {
		start_idx := num_pos[i]
		end_idx := num_pos[i+1]
		if start_idx > gear_idx+1 {
			return result_set
		}
		if (end_idx-1 == gear_idx) || (end_idx == gear_idx) ||
			(start_idx-1 == gear_idx) || (start_idx == gear_idx) || (start_idx+1 == gear_idx) {
			result_set = append(result_set, part_nums[j])
		}
	}
	return result_set
}

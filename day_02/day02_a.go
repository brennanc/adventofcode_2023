package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strconv"
	"strings"
)

func main() {
	//Read input file
	input, _ := os.Open("input.txt")
	defer input.Close()
	sc := bufio.NewScanner(input)

	max_red := 12
	max_green := 13
	max_blue := 14

	game_count := 0
	for sc.Scan() {
		line := sc.Text()
		game := strings.Split(line, ":")
		game_strid := strings.Split(game[0], " ")
		game_num, _ := strconv.Atoi(game_strid[1])
		game_data := strings.Split(game[1], ";")

		current_red_max := 0
		current_green_max := 0
		current_blue_max := 0
		for _, game_x := range game_data {
			colors := strings.Split(game_x, ",")
			re, _ := regexp.Compile("(\\d+)\\s(\\w)")
			count := 0
			color_name := ""
			for _, color := range colors {
				results := re.FindAllStringSubmatch(color, -1)
				count, _ = strconv.Atoi(results[0][1])
				color_name = results[0][2]
				switch {
				case color_name == "red":
					if count > current_red_max {
						current_red_max = count
					}
				case color_name == "green":
					if count > current_green_max {
						current_green_max = count
					}
				case color_name == "blue":
					if count > current_blue_max {
						current_blue_max = count
					}
				}
			}
		}
	}

	fmt.Println("")
}

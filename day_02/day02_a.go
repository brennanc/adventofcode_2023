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

	for sc.Scan() {
		line := sc.Text()
		game := strings.Split(line, ":")
		game_strid := strings.Split(game[0], " ")
		game_num, _ := strconv.Atoi(game_strid[1])
		game_data := strings.Split(game[1], ";")

		for _, game_x := range game_data {
			colors := strings.Split(game_x, ",")
			re, _ := regexp.Compile("(\\d+)\\s(\\w)")
			count := 0
			color_name := ""
			for _, color := range colors {
				results := re.FindAllStringSubmatch(color, -1)
				count = strconv.Atoi(results[0][1])
			}
		}
	}

	fmt.Println("")
}

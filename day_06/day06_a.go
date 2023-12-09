package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strconv"
)

func main() {
	//Read input file
	input, _ := os.Open("input.txt")
	defer input.Close()
	sc := bufio.NewScanner(input)

	times := make([]int, 0)
	distances := make([]int, 0)
	for sc.Scan() {
		line := sc.Text()

		re := regexp.MustCompile(`\d+`)
		matches := re.FindAllString(line, -1)
		for _, match := range matches {
			num, _ := strconv.Atoi(match)
			times = append(times, num)
		}

		sc.Scan()
		line = sc.Text()
		matches = re.FindAllString(line, -1)
		for _, match := range matches {
			num, _ := strconv.Atoi(match)
			distances = append(distances, num)
		}
	}

	waysToWinProduct := 0
	for i, time := range times {
		waysToWin := 0
		distance := distances[i]
		for speed := 1; speed <= time; speed++ {
			steps := distance / speed
			timeRemaining := time - speed
			distanceTraveled := speed * timeRemaining
			if steps > timeRemaining {
				continue
			}
			if distanceTraveled > distance {
				waysToWin++
			}
		}
		if waysToWin > 0 {
			if waysToWinProduct == 0 {
				waysToWinProduct = 1
			}
			waysToWinProduct *= waysToWin
		}
	}
	fmt.Println(waysToWinProduct)
}

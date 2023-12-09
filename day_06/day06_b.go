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

	time := 0
	distance := 0
	for sc.Scan() {
		line := sc.Text()

		numStr := ""
		re := regexp.MustCompile(`\d+`)
		matches := re.FindAllString(line, -1)
		for _, match := range matches {
			numStr = numStr + match
		}
		time, _ = strconv.Atoi(numStr)

		sc.Scan()
		line = sc.Text()
		distanceStr := ""
		matches = re.FindAllString(line, -1)
		for _, match := range matches {
			distanceStr = distanceStr + match
		}
		distance, _ = strconv.Atoi(distanceStr)
	}

	waysToWinProduct := 0
	waysToWin := 0
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
	fmt.Println(waysToWinProduct)
}

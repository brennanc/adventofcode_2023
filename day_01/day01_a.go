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

	re := regexp.MustCompile("^[^\\d]*([\\d]).*([\\d]+).*$")

	calibration_total := 0
	flag := "false"
	for sc.Scan() {
		line := sc.Text()

		results := re.FindAllStringSubmatch(line, -1)

		if len(results) == 0 {
			flag = "true"
			re2 := regexp.MustCompile("^[^\\d]*([\\d]).*([\\d]*).*$")
			results = re2.FindAllStringSubmatch(line, -1)
			fmt.Printf("%s-----%s%s\n", line, results[0][1], results[0][1])
			value, _ := strconv.Atoi(results[0][1] + results[0][1])
			calibration_total += value
			continue
		}

		for _, calibration_values := range results {
			value := 0
			if len(calibration_values) > 2 {
				fmt.Printf("%s-----%s%s\n", line, calibration_values[1], calibration_values[2])
				value, _ = strconv.Atoi(calibration_values[1] + calibration_values[2])
			}
			calibration_total += value
		}
	}

	fmt.Println(calibration_total)
	fmt.Println(flag)
}

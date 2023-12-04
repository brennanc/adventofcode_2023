package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strconv"
	"strings"
	"unicode"
)

func wordToDigit(word string) (string, bool) {
	digitMap := map[string]string{
		"zero":  "0",
		"one":   "1",
		"two":   "2",
		"three": "3",
		"four":  "4",
		"five":  "5",
		"six":   "6",
		"seven": "7",
		"eight": "8",
		"nine":  "9",
	}

	word = strings.ToLower(word)
	value, exists := digitMap[word]
	return value, exists
}
func getValue(word string, word2 string) int {
	value := 0
	if !unicode.IsDigit(rune(word[0])) {
		word, _ = wordToDigit(word)
	}
	if !unicode.IsDigit(rune(word2[0])) {
		word2, _ = wordToDigit(word2)
	}
	value, _ = strconv.Atoi(word + word2)
	return value
}

func main() {
	//Read input file
	input, _ := os.Open("input.txt")
	defer input.Close()
	sc := bufio.NewScanner(input)

	re := regexp.MustCompile("(one|1|two|2|three|3|four|4|five|5|six|6|seven|7|eight|8|nine|9).*(one|1|two|2|three|3|four|4|five|5|six|6|seven|7|eight|8|nine|9).*$")

	calibration_total := 0
	flag := "false"
	for sc.Scan() {
		line := sc.Text()

		results := re.FindAllStringSubmatch(line, -1)

		if len(results) == 0 {
			flag = "true"
			re2 := regexp.MustCompile("^[^\\d]*([\\d]).*([\\d]*).*$")
			results = re2.FindAllStringSubmatch(line, -1)
			value, _ := strconv.Atoi(results[0][1] + results[0][1])
			calibration_total += value
			fmt.Printf("================%s-----%s%s-----total: %d\n", line, results[0][1], results[0][1], calibration_total)
			continue
		}

		for _, calibration_values := range results {
			value := getValue(calibration_values[1], calibration_values[2])
			calibration_total += value
			fmt.Printf("%s-----%s %s-->%d-----total: %d\n", line, calibration_values[1], calibration_values[2], value, calibration_total)
		}
	}

	fmt.Println(calibration_total)
	fmt.Println(flag)
}

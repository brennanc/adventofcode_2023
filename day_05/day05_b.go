package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"regexp"
	"strconv"
)

func main() {
	//Read input file
	input, _ := os.Open("input.txt")
	defer input.Close()
	sc := bufio.NewScanner(input)

	seedToSoilMap := make([][]int64, 0)
	soilToFertilizerMap := make([][]int64, 0)
	fertilizerToWaterMap := make([][]int64, 0)
	waterToLightMap := make([][]int64, 0)
	lightToTempMap := make([][]int64, 0)
	tempToHumidityMap := make([][]int64, 0)
	humidityToLocationMap := make([][]int64, 0)
	seeds := make([]int64, 0)

	for sc.Scan() {
		line := sc.Text()
		re := regexp.MustCompile(`\d+`)
		matches := re.FindAllString(line, -1)
		for i := 0; i < len(matches); i += 2 {
			seedStart, _ := strconv.ParseInt(matches[i], 10, 64)
			numSeeds, _ := strconv.ParseInt(matches[i+1], 10, 64)
			for j := seedStart; j < seedStart+numSeeds; j++ {
				seeds = append(seeds, j)
			}
		}
		sc.Scan() // blank

		parseMapb(&seedToSoilMap, sc)
		parseMapb(&soilToFertilizerMap, sc)
		parseMapb(&fertilizerToWaterMap, sc)
		parseMapb(&waterToLightMap, sc)
		parseMapb(&lightToTempMap, sc)
		parseMapb(&tempToHumidityMap, sc)
		parseMapb(&humidityToLocationMap, sc)

		locations := make([]int64, 0)
		for i := 0; i < len(seeds); i++ {
			soil := sourceToDestb(seedToSoilMap, seeds[i])
			fert := sourceToDestb(soilToFertilizerMap, soil)
			water := sourceToDestb(fertilizerToWaterMap, fert)
			light := sourceToDestb(waterToLightMap, water)
			temp := sourceToDestb(lightToTempMap, light)
			humid := sourceToDestb(tempToHumidityMap, temp)
			loc := sourceToDestb(humidityToLocationMap, humid)
			locations = append(locations, loc)
			//fmt.Printf("Seed %d, soil %d, fertilizer %d, water %d, light %d, temp %d, humidity %d, location %d\n",
			//	seeds[i], soil, fert, water, light, temp, humid, loc)
		}
		fmt.Println(myminb(locations))
	}

}

func myminb(vals []int64) int64 {
	var mymin int64 = math.MaxInt64
	for i := 0; i < len(vals); i++ {
		if vals[i] < mymin {
			mymin = vals[i]
		}
	}
	return mymin
}

func parseMapb(mapping *[][]int64, sc *bufio.Scanner) {
	re := regexp.MustCompile(`\d+`)
	sc.Scan()
	sc.Scan()
	line := sc.Text()
	for len(line) > 0 {
		matches := re.FindAllString(line, -1)
		mapRow := make([]int64, 0)
		for _, match := range matches {
			num, _ := strconv.ParseInt(match, 10, 64)
			mapRow = append(mapRow, num)
		}
		*mapping = append(*mapping, mapRow)
		sc.Scan()
		line = sc.Text()
	}
}

func sourceToDestb(mapping [][]int64, sourceNum int64) int64 {
	var result int64 = -1
	for i := 0; i < len(mapping); i++ {
		sourceStart := mapping[i][1]
		destStart := mapping[i][0]
		offset := mapping[i][2]

		if (sourceNum >= sourceStart) && (sourceNum < sourceStart+offset) {
			result = destStart + (sourceNum - sourceStart)
			return result
		}
	}
	if result == -1 {
		result = sourceNum
	}
	return result
}

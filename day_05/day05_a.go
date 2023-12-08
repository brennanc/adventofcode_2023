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
		for _, match := range matches {
			num, _ := strconv.ParseInt(match, 10, 64)
			seeds = append(seeds, num)
		}
		sc.Scan() // blank

		parseMap(&seedToSoilMap, sc)
		parseMap(&soilToFertilizerMap, sc)
		parseMap(&fertilizerToWaterMap, sc)
		parseMap(&waterToLightMap, sc)
		parseMap(&lightToTempMap, sc)
		parseMap(&tempToHumidityMap, sc)
		parseMap(&humidityToLocationMap, sc)

		locations := make([]int64, 0)
		for i := 0; i < len(seeds); i++ {
			soil := sourceToDest(seedToSoilMap, seeds[i])
			fert := sourceToDest(soilToFertilizerMap, soil)
			water := sourceToDest(fertilizerToWaterMap, fert)
			light := sourceToDest(waterToLightMap, water)
			temp := sourceToDest(lightToTempMap, light)
			humid := sourceToDest(tempToHumidityMap, temp)
			loc := sourceToDest(humidityToLocationMap, humid)
			locations = append(locations, loc)
			fmt.Printf("Seed %d, soil %d, fertilizer %d, water %d, light %d, temp %d, humidity %d, location %d\n",
				seeds[i], soil, fert, water, light, temp, humid, loc)
		}
		fmt.Println(mymin(locations))
	}

}

func mymin(vals []int64) int64 {
	var mymin int64 = math.MaxInt64
	for i := 0; i < len(vals); i++ {
		if vals[i] < mymin {
			mymin = vals[i]
		}
	}
	return mymin
}

func parseMap(mapping *[][]int64, sc *bufio.Scanner) {
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

func sourceToDest(mapping [][]int64, sourceNum int64) int64 {
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

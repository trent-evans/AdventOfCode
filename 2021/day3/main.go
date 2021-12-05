package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
)

func filterList(numList []string, bitIdx int, mostCommonSearch bool) string {
	if len(numList) == 1 || len(numList[0]) == bitIdx {
		return numList[0]
	} else {
		/**
		- Fill in the recursion here
		*/
		numList = numList
		bitIdx++
		return filterList(numList, bitIdx, mostCommonSearch)
	}
}

func main() {
	file, err := os.Open("2021/day3/input.txt")

	if err != nil {
		log.Fatal(err)
	}

	defer file.Close()

	scanner := bufio.NewScanner(file)

	var data []string

	for scanner.Scan() {
		data = append(data, scanner.Text())
	}

	// Part 1
	var oneCounts [12]int

	for _, record := range data {
		for idx, bit := range record {
			if bit == '1' {
				oneCounts[idx]++
			}
		}
	}

	var threshold = len(data) / 2

	var gamma = ""
	var epsilon = ""

	for _, count := range oneCounts {
		if count > threshold {
			gamma += "1"
			epsilon += "0"
		} else {
			gamma += "0"
			epsilon += "1"
		}
	}

	fmt.Println(gamma)
	fmt.Println(epsilon)

	gammaVal, _ := strconv.ParseInt(gamma, 2, 32)
	epsilonVal, _ := strconv.ParseInt(epsilon, 2, 32)

	fmt.Println("Gamma Value:")
	fmt.Println(gammaVal)
	fmt.Println("Epsilon Value:")
	fmt.Println(epsilonVal)
	fmt.Println("Product Value:")
	fmt.Println(gammaVal * epsilonVal)

	// Part 2

	var oxygenData []string
	var co2Data []string

	for _, record := range data {
		if record[0] == '1' {
			oxygenData = append(oxygenData, record)
		} else {
			co2Data = append(co2Data, record)
		}
	}

	oxygenOut := filterList(oxygenData, 1, true)
	co2Out := filterList(co2Data, 1, false)

	oxygenVal, _ := strconv.ParseInt(oxygenOut, 2, 32)
	co2Val, _ := strconv.ParseInt(co2Out, 2, 32)

	fmt.Println("Oxygen Values:")
	fmt.Println(oxygenOut)
	fmt.Println(oxygenVal)
	fmt.Println("CO2 Values:")
	fmt.Println(co2Out)
	fmt.Println(co2Val)

	fmt.Println("Product:")
	fmt.Println(co2Val * oxygenVal)
}

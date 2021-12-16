package main

import (
	"bufio"
	"fmt"
	"log"
	"math"
	"os"
	"strconv"
	"strings"
)

func calculateSingleStepFuelUsage(startPositions []int, finalPosition int) int {

	fuelUsage := 0
	for i := 0; i < len(startPositions); i++ {
		fuelUsage += int(math.Abs(float64(startPositions[i] - finalPosition)))
	}
	return fuelUsage
}

func calculateMultiStepFuelUsage(startPositions []int, finalPosition int, stepCostMap map[int]int) int {
	fuelUse := 0
	for i := 0; i < len(startPositions); i++ {
		stepDiff := int(math.Abs(float64(startPositions[i] - finalPosition)))
		fuelCost, _ := stepCostMap[stepDiff]
		fuelUse += fuelCost
	}
	return fuelUse
}

func main() {
	fmt.Println("Hello world")

	filename := "2021/day7/input.txt"

	file, err := os.Open(filename)

	if err != nil {
		log.Fatal(err)
	}

	defer file.Close()

	scanner := bufio.NewScanner(file)

	var inputInts []int
	posMin := 1000
	posMax := 0
	for scanner.Scan() {
		inputText := scanner.Text() // Only one line in these inputs
		inputTextSplit := strings.Split(inputText, ",")

		for _, numStr := range inputTextSplit {
			intVal, _ := strconv.Atoi(numStr)
			inputInts = append(inputInts, intVal)

			if intVal < posMin {
				posMin = intVal
			}

			if intVal > posMax {
				posMax = intVal
			}
		}
	}

	fuelMinimum := posMax * len(inputInts)

	for i := posMin; i < posMax; i++ {
		fuelUse := calculateSingleStepFuelUsage(inputInts, i)
		if fuelUse < fuelMinimum {
			fuelMinimum = fuelUse
		}
	}

	fmt.Println(fmt.Sprintf("Minimum fuel usage for single steps: %d", fuelMinimum))

	maxDiff := posMax - posMin

	diffSums := make(map[int]int)

	currentSum := 0
	for i := 0; i <= maxDiff; i++ {
		currentSum += i
		diffSums[i] = currentSum
		//fmt.Println(fmt.Sprintf("Num: %d  SumToNum: %d", i, currentSum))
	}

	maxDiffFuel, _ := diffSums[maxDiff]
	multiStepFuelMin := posMax * maxDiffFuel * len(inputInts)

	for i := posMin; i < posMax; i++ {
		fuelUse := calculateMultiStepFuelUsage(inputInts, i, diffSums)
		if fuelUse < multiStepFuelMin {
			multiStepFuelMin = fuelUse
		}
	}

	fmt.Println(fmt.Sprintf("Minimum fuel usage for summed steps: %d", multiStepFuelMin))
}

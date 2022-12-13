package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"sort"
	"strconv"
)

func main() {
	file, err := os.Open("2022/day1/input.txt")

	if err != nil {
		log.Fatal(err)
	}

	defer file.Close()

	scanner := bufio.NewScanner(file)

	totalCalories := []int{}

	maxCals := 0
	currentCals := 0

	for scanner.Scan() {
		if scanner.Text() == "" {
			if currentCals > maxCals {
				maxCals = currentCals
			}
			totalCalories = append(totalCalories, currentCals)
			currentCals = 0
			continue
		}
		calVal, err := strconv.Atoi(scanner.Text())
		if err != nil {
			log.Fatal(err)
		}
		currentCals = currentCals + calVal
	}

	fmt.Printf("\n\nMax calories: %v\n\n", maxCals)

	sort.Ints(totalCalories)

	totalLen := len(totalCalories)
	top3Total := totalCalories[totalLen-1] + totalCalories[totalLen-2] + totalCalories[totalLen-3]

	fmt.Printf("Top 3 Total: %v\n\n", top3Total)
}

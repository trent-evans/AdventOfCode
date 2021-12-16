package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func calculateFishSum(fishMap map[int]int) int {
	sum := 0

	for i := 0; i < 9; i++ {
		sum += fishMap[i]
	}

	return sum
}

func main() {
	filename := "2021/day6/input.txt"

	file, err := os.Open(filename)

	if err != nil {
		log.Fatal(err)
	}

	defer file.Close()

	scanner := bufio.NewScanner(file)

	var initialListTest string

	for scanner.Scan() {
		initialListTest = scanner.Text() // There's only one line
	}

	initialListSplit := strings.Split(initialListTest, ",")

	var fishDays []int8

	fishDaysMap := map[int]int{
		0: 0,
		1: 0,
		2: 0,
		3: 0,
		4: 0,
		5: 0,
		6: 0,
		7: 0,
		8: 0,
	}

	for _, numStr := range initialListSplit {
		dayToInt, _ := strconv.Atoi(numStr)
		fishDays = append(fishDays, int8(dayToInt))

		fishDaysMap[dayToInt]++
	}

	mapDays := 256
	// Part 2 - This is also the better way to do part 1, but this is something I should have figured out earlier
	// Part 1 was initially a bunch of looping.  But I'm past that now
	for i := 1; i <= mapDays; i++ {

		numNewAndRolloverFish := fishDaysMap[0]

		for x := 0; x < 8; x++ { // Up to 7 to not exceed the bounds
			fishDaysMap[x] = fishDaysMap[x+1]
		}
		fishDaysMap[8] = numNewAndRolloverFish  // Pop in the new fish
		fishDaysMap[6] += numNewAndRolloverFish // Add the rollover fish

		if i == 80 {
			fishSum := calculateFishSum(fishDaysMap)
			if strings.Contains(filename, "example") {
				fmt.Println(fmt.Sprintf("Days: 80 -- Success: %t \nNum Fish: %d -- Num Expected: 5934", fishSum == 5934, fishSum))
			} else {
				fmt.Println(fmt.Sprintf("Days: 80 -- Success: %t \nNum Fish: %d -- Num Expected: 388739", fishSum == 388739, fishSum))
			}
		}
	}

	fmt.Println(fmt.Sprintf("# of Fish: %d", calculateFishSum(fishDaysMap)))
}

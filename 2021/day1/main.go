package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
)

func main() {
	file, err := os.Open("2021/day1/input.txt")

	if err != nil {
		log.Fatal(err)
	}

	defer file.Close()

	scanner := bufio.NewScanner(file)

	var data []int

	for scanner.Scan() {
		depth, err := strconv.Atoi(scanner.Text())
		if err != nil {
			log.Fatal(err)
		}
		data = append(data, depth)
	}

	count := 0

	for i := 0; i < len(data)-1; i++ {
		if data[i] < data[i+1] {

			count++
		}
	}

	fmt.Println(count)

	// Sliding window stuff

	var sliding_data []int
	for i := 0; i < len(data)-2; i++ {
		slidingValue := data[i] + data[i+1] + data[i+2]
		sliding_data = append(sliding_data, slidingValue)
	}

	sliding_count := 0
	for i := 0; i < len(sliding_data)-1; i++ {
		if sliding_data[i] < sliding_data[i+1] {

			sliding_count++
		}
	}

	fmt.Println(sliding_count)
}

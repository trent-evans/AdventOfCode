package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
)

type dataEntry struct {
	signals [10]string
	output  [4]string
}

func main() {
	fmt.Println("Hello world")

	filename := "2021/day7/input.txt"

	file, err := os.Open(filename)

	if err != nil {
		log.Fatal(err)
	}

	defer file.Close()

	var data []dataEntry
	var newEntry dataEntry

	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		scanText := scanner.Text()

		splitSignalAndOutput := strings.Split(scanText, "|")

		signalFull := splitSignalAndOutput[0]
		outputFull := splitSignalAndOutput[1]

	}

	// segmentValues := map[string]int{
	// 	"acbefg":  0,
	// 	"cf":      1,
	// 	"acdeg":   2,
	// 	"acdfg":   3,
	// 	"bcdf":    4,
	// 	"abdfg":   5,
	// 	"abdefg":  6,
	// 	"acf":     7,
	// 	"abcdefg": 8,
	// 	"abcdfg":  9,
	// }

}

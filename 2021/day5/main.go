package main

import (
	"bufio"
	"log"
	"os"
	"strconv"
	"strings"
)

type ventLine struct {
	x1 int
	y1 int
	x2 int
	y2 int
}

func main() {

	filename := "2021/day5/example.txt"

	file, err := os.Open("2021/day5/example.txt")

	if err != nil {
		log.Fatal(err)
	}

	defer file.Close()

	scanner := bufio.NewScanner(file)

	var ventList []ventLine

	for scanner.Scan() {
		entry := scanner.Text()
		entry = strings.Replace(entry, "->", " ", -1)
		entry = strings.Replace(entry, ",", " ", -1)

		valuesSplit := strings.Fields(entry)

		var newLine ventLine
		for idx, indVal := range valuesSplit {

			num, _ := strconv.Atoi(indVal)

			switch idx {

			case 0:
				newLine.x1 = num

			case 1:
				newLine.y1 = num

			case 2:
				newLine.x2 = num

			case 3:
				newLine.y2 = num
			}
		}

		ventList = append(ventList, newLine)
	}

	bound := 1000 // For the full size
	var thermalField [1000][1000]int

	if strings.Contains(filename, "example") {
		bound = 10
	}
}

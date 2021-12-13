package main

import (
	"bufio"
	"log"
	"os"
	"strings"
)

type indexPair struct {
	x1 int
	y1 int
	x2 int
	y2 int
}

func main() {

	filename := "2021/day5/example.txt"

	side := 1000 // For the full size

	if strings.Contains(filename, "example") {
		side = 10
	}

	var thermalField [side][side]int

	file, err := os.Open("2021/day5/example.txt")

	if err != nil {
		log.Fatal(err)
	}

	defer file.Close()

	scanner := bufio.NewScanner(file)

	for scanner.Scan() {

	}
}

package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

type instruction struct {
	direction string
	magnitude int
}

func main() {
	fmt.Println("Magic magic magics")

	file, err := os.Open("2021/day2/input.txt")

	if err != nil {
		log.Fatal(err)
	}

	defer file.Close()

	var instructionArray []instruction

	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		split := strings.Split(scanner.Text(), " ")

		convertedMag, err := strconv.Atoi(split[1])

		if err != nil {
			log.Fatal(err)
		}

		var newInstruction instruction
		newInstruction.direction = split[0]
		newInstruction.magnitude = convertedMag

		instructionArray = append(instructionArray, newInstruction)
	}

	var horizontalPos = 0
	var verticalPos = 0
	var aim = 0

	for _, step := range instructionArray {
		if step.direction == "up" {
			aim -= step.magnitude
		} else if step.direction == "down" {
			aim += step.magnitude
		} else {
			horizontalPos += step.magnitude
			verticalPos += step.magnitude * aim
		}
	}
	concatHor := fmt.Sprintf("Horizontal: %d", horizontalPos)
	concatVert := fmt.Sprintf("Vertical: %d", verticalPos)
	concatMult := fmt.Sprintf("Product: %d", horizontalPos*verticalPos)

	fmt.Println(concatHor)
	fmt.Println(concatVert)
	fmt.Println(concatMult)
}

package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func main() {

	filename := "2015/day2/input.txt"

	file, err := os.Open(filename)

	if err != nil {
		log.Fatal(err)
	}

	defer file.Close()

	scanner := bufio.NewScanner(file)

	total := 0
	ribbonTotal := 0

	for scanner.Scan() {
		scanText := scanner.Text()

		dimensions := strings.Split(scanText, "x")

		length, _ := strconv.Atoi(dimensions[0])
		width, _ := strconv.Atoi(dimensions[1])
		height, _ := strconv.Atoi(dimensions[2])

		topAndBottom := length * width
		frontAndBack := length * height
		sides := width * height

		topAndBottomPerim := length*2 + width*2
		frontAndBackPerim := length*2 + height*2
		sidesPerim := width*2 + height*2

		ribbonBow := length * width * height

		// Get the minimum side area
		minSide := topAndBottom
		if frontAndBack < minSide {
			minSide = frontAndBack
		}
		if sides < minSide {
			minSide = sides
		}

		minPerim := topAndBottomPerim
		if frontAndBackPerim < minPerim {
			minPerim = frontAndBackPerim
		}
		if sidesPerim < minPerim {
			minPerim = sidesPerim
		}

		subTotal := 2*topAndBottom + 2*frontAndBack + 2*sides + minSide

		total += subTotal

		ribbonSubTotal := ribbonBow + minPerim
		ribbonTotal += ribbonSubTotal
	}

	fmt.Println(total)
	fmt.Println(ribbonTotal)
}

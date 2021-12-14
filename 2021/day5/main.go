package main

import (
	"bufio"
	"fmt"
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

func countDangerZones(thermalField [1000][1000]int, bounds int) int {
	var dangerZones = 0

	for r := 0; r < bounds; r++ {
		for c := 0; c < bounds; c++ {
			if thermalField[r][c] > 1 {
				dangerZones++
			}
		}
	}

	return dangerZones
}

func printThermalField(thermalField [1000][1000]int, bounds int) {
	if bounds != 10 {
		fmt.Println("Not printing that field - waaaaay too big")
		return
	}

	for i := 0; i < 10; i++ {
		fmt.Println(thermalField[i][:10])
	}
}

func main() {

	filename := "2021/day5/example.txt"

	file, err := os.Open(filename)

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

		//fmt.Println(entry)

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

	fmt.Println(len(ventList))

	bound := 1000 // For the full size
	var thermalField [1000][1000]int
	diagnonalLines := 0

	if strings.Contains(filename, "example") {
		bound = 10
	}

	// Part 1
	for _, vent := range ventList {
		if vent.x1 == vent.x2 {

			y1 := vent.y1
			y2 := vent.y2

			if vent.y1 > vent.y2 {
				y1 = vent.y2
				y2 = vent.y1
			}

			for y := y1; y <= y2; y++ {
				thermalField[vent.x1][y]++
			}

		} else if vent.y1 == vent.y2 {

			x1 := vent.x1
			x2 := vent.x2

			if vent.x1 > vent.x2 {
				x1 = vent.x2
				x2 = vent.x1
			}

			for x := x1; x <= x2; x++ {
				thermalField[x][vent.y1]++
			}

		} else {
			diagnonalLines++
		}
	}

	dangerZones := countDangerZones(thermalField, bound)

	fmt.Println(fmt.Sprintf("Danger points: %d", dangerZones))
	fmt.Println(fmt.Sprintf("Diagonal vents: %d", diagnonalLines))

	// Part 2
	for _, vent := range ventList {
		if vent.x1 == vent.x2 || vent.y1 == vent.y2 { // Already considered in part 1
			continue
		}

		x1 := vent.x1
		x2 := vent.x2
		y1 := vent.y1
		y2 := vent.y2

		xDiff := x2 - x1
		yDiff := y2 - y1

		fmt.Println(fmt.Sprintf("xDiff: %d -- yDiff: %d", xDiff, yDiff))

		// Increment by 1s at a 45 degree - probably a better way to do this but this is easy
		if xDiff < 0 {

			var x = x2

			if yDiff < 0 {
				fmt.Println("xDiff and yDiff are negative")

				var y = y2
				fmt.Println(fmt.Sprintf("\tx: %d \ty: %d", x, y))

				for x >= x1 && y >= y1 {

					fmt.Println(fmt.Sprintf("\tx: %d \ty: %d", x, y))
					thermalField[x][y]++

					x--
					y--
				}

			} else {
				fmt.Println("xDiff is negative, yDiff is positive")

				var y = y1
				fmt.Println(fmt.Sprintf("\tx: %d \ty: %d", x, y))

				for x >= x1 && y <= y2 {
					fmt.Println(fmt.Sprintf("\tx: %d \ty: %d", x, y))
					thermalField[x][y]++

					x--
					y++
				}
			}

		} else {
			var x = x1

			if yDiff < 0 {
				fmt.Println("xDiff is positive, yDiff is negative")

				var y = y2
				fmt.Println(fmt.Sprintf("\tx: %d \ty: %d", x, y))

				for x <= x2 && y >= y1 {
					fmt.Println(fmt.Sprintf("\tx: %d \ty: %d", x, y))
					thermalField[x][y]++

					x++
					y--
				}

			} else {
				fmt.Println("xDiff and yDiff are both positive")

				var y = y1
				fmt.Println(fmt.Sprintf("\tx: %d \ty: %d", x, y))

				for x <= x2 && y <= y1 {
					fmt.Println(fmt.Sprintf("\tx: %d \ty: %d", x, y))
					thermalField[x][y]++

					x++
					y++
				}

			}
		}
	}

	dangerZones = countDangerZones(thermalField, bound)

	fmt.Println(fmt.Sprintf("Danger points with diagonals: %d", dangerZones))
	fmt.Println("")
	//printThermalField(thermalField, bound)
}

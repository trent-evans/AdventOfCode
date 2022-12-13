package main

import (
	"crypto/md5"
	"fmt"
	"strconv"
)

func main() {
	input := "ckczppom"

	fmt.Println(input)

	isCorrectValue := false
	tailValue := 0

	for isCorrectValue {
		testInput := input + strconv.Itoa(tailValue)

		hash := md5.Sum([]byte(testInput))

		if hash[0:5] == "00000" {
			fmt.Println(tailValue)
			break
		}

		tailValue++
	}
}

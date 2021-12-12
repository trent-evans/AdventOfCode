package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

type indexPair struct {
	row int
	col int
}

type bingoSpace struct {
	value    int
	isMarked bool
}

type bingoBoard struct {
	spaces       [5][5]bingoSpace
	spacesMarked int
	spaceMap     map[int]indexPair
	hasBingo     bool
}

func markSpace(value int, board *bingoBoard) bool {
	coords, isPresent := board.spaceMap[value]
	if isPresent {
		board.spaces[coords.row][coords.col].isMarked = true
		board.spacesMarked += 1
		fmt.Println(fmt.Sprintf("Spaces Marked: %d", board.spacesMarked))

		if board.spacesMarked >= 5 {
			board.hasBingo = doesBoardHaveBingo(coords, *board)
		}
	}
	return isPresent
}

func doesBoardHaveBingo(mostRecentSpace indexPair, board bingoBoard) bool {
	row := mostRecentSpace.row

	for c := 0; c < 5; c++ {
		if !board.spaces[row][c].isMarked {
			break
		}
		if c == 4 {
			return true
		}
	}

	col := mostRecentSpace.col

	for r := 0; r < 5; r++ {
		if !board.spaces[r][col].isMarked {
			break
		}
		if r == 4 {
			return true
		}
	}

	return false
}

func (board bingoBoard) addUnmarked() int {
	if !board.hasBingo {
		return -1
	}
	sum := 0
	for r := 0; r < 5; r++ {
		for c := 0; c < 5; c++ {
			if !board.spaces[r][c].isMarked {
				sum += board.spaces[r][c].value
			}
		}
	}
	return sum
}

func main() {
	file, err := os.Open("2021/day4/input.txt")

	if err != nil {
		log.Fatal(err)
	}

	defer file.Close()

	scanner := bufio.NewScanner(file)

	var callNumbers []int
	var boardStrings []string

	for scanner.Scan() {
		if len(callNumbers) == 0 {
			var callList = strings.Split(scanner.Text(), ",")
			for _, num := range callList {
				newCallNum, _ := strconv.Atoi(num)
				callNumbers = append(callNumbers, newCallNum)
			}
		} else {
			if scanner.Text() != "" {
				boardStrings = append(boardStrings, scanner.Text())
				// valuesAndLength := fmt.Sprintf("%s - Len: %d", scanner.Text(), len(scanner.Text()))
				// fmt.Println(valuesAndLength)
			}
		}
	}

	// lenCallNumbersOut := fmt.Sprintf("Length of call numbers: %d", len(callNumbers))
	// fmt.Println(lenCallNumbersOut)

	//var bingoBoardList []bingoBoard

	fmt.Println(len(boardStrings))

	var boardList []bingoBoard
	var newBoard = new(bingoBoard)
	newBoard.spaceMap = make(map[int]indexPair)
	newBoard.spacesMarked = 0
	newBoard.hasBingo = false

	for i := 0; i < len(boardStrings); i++ {
		//fmt.Println(boardStrings[i])

		valueSplit := strings.Fields(boardStrings[i]) // Splits on whitespace, not just a single character
		// Ran into a bug where single-digit numbers would split into 0 and # because there were two space characters

		// wholeRow := fmt.Sprintf("Whole row: %s", boardStrings[i])
		// fmt.Println(wholeRow)

		row := i % 5

		for col := 0; col < len(valueSplit); col++ {
			//fmt.Println(valueSplit[j])

			singleValue, _ := strconv.Atoi(valueSplit[col])

			// rowColData := fmt.Sprintf("Row: %d Col: %d -- Value: %d", row, col, singleValue)
			// fmt.Println(rowColData)

			newBoard.spaces[row][col] = bingoSpace{value: singleValue, isMarked: false}

			valueLoc := indexPair{row: row, col: col}
			newBoard.spaceMap[singleValue] = valueLoc
		}

		if row == 4 {
			boardList = append(boardList, *newBoard)

			newBoard = new(bingoBoard)
			newBoard.spaceMap = make(map[int]indexPair)
			newBoard.spacesMarked = 0
			newBoard.hasBingo = false
		}
	}

	//fmt.Println(len(boardList))

	for x, num := range callNumbers {
		fmt.Println(fmt.Sprintf("Idx: %d - CallNum: %d", x, num))

		for idx, board := range boardList {
			if markSpace(num, &board) {
				fmt.Println(fmt.Sprintf("Value added to Board ID: %d  - Num values: %d", idx, board.spacesMarked))
			}

			if board.hasBingo {
				unmarkedSum := board.addUnmarked()

				finalScore := unmarkedSum * num

				// scoreOut := fmt.Sprintf("Final Score: %d -- Sum: %d LastCalled: %d", finalScore, unmarkedSum, num)
				// fmt.Println(scoreOut)
				fmt.Println(fmt.Sprintf("Final Score: %d -- Sum: %d LastCalled: %d", finalScore, unmarkedSum, num))
				break
			}
		}
	}

}

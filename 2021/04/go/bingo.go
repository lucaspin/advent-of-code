package main

import (
	"fmt"
	"io/ioutil"
	"strconv"
	"strings"
)

type Bingo struct {
	Sequence []int
	Boards   []Board
}

func Init(lines []string) *Bingo {
	sequence := toInt(strings.Split(lines[0], ","))

	return &Bingo{
		Sequence: sequence,
		Boards:   parseBoards(lines[2:]),
	}
}

func parseBoards(lines []string) []Board {
	boards := []Board{}

	var currentBoard Board
	for _, line := range lines {
		if line == "" {
			boards = append(boards, currentBoard)
			currentBoard = Board{}
		} else {
			currentBoard.AddRow(toInt(trim(nonEmpty(strings.Split(line, " ")))))
		}
	}

	boards = append(boards, currentBoard)
	return boards
}

func toInt(values []string) []int {
	new := []int{}
	for _, value := range values {
		number, _ := strconv.Atoi(value)
		new = append(new, number)
	}

	return new
}

func trim(values []string) []string {
	newValues := []string{}
	for _, value := range values {
		newValues = append(newValues, strings.Trim(value, " "))
	}

	return newValues
}

func nonEmpty(values []string) []string {
	newValues := []string{}
	for _, value := range values {
		if value != "" {
			newValues = append(newValues, value)
		}
	}

	return newValues
}

func (b *Bingo) FindFirstWinner() (*Board, int) {
	for _, number := range b.Sequence {
		for _, board := range b.Boards {
			board.ProcessNumber(number)
			if board.IsWinner() {
				return &board, number
			}
		}
	}

	return nil, -1
}

func (b *Bingo) FindLastWinner() (*Board, int) {
	remainingBoards := b.Boards
	previousBoards := remainingBoards
	for _, number := range b.Sequence {
		for _, board := range remainingBoards {
			board.ProcessNumber(number)
		}

		previousBoards = remainingBoards
		remainingBoards = onlyNonWinners(remainingBoards)
		if len(remainingBoards) == 0 {
			return &previousBoards[0], number
		}
	}

	return nil, -1
}

func onlyNonWinners(boards []Board) []Board {
	onlyNonWinners := []Board{}
	for _, board := range boards {
		if !board.IsWinner() {
			onlyNonWinners = append(onlyNonWinners, board)
		}
	}

	return onlyNonWinners
}

type Board struct {
	Rows [][]*BoardCell
}

func (b *Board) PrintState() {
	for _, row := range b.Rows {
		for _, cell := range row {
			if cell.Marked {
				fmt.Print("1 ")
			} else {
				fmt.Print("0 ")
			}
		}

		fmt.Println()
	}
}

func (b *Board) Score(number int) int {
	return b.unmarkedSum() * number
}

func (b *Board) unmarkedSum() int {
	sum := 0
	for _, row := range b.Rows {
		for _, cell := range row {
			if !cell.Marked {
				sum += cell.Value
			}
		}
	}

	return sum
}

func (b *Board) AddRow(row []int) {
	newRow := []*BoardCell{}
	for _, value := range row {
		newRow = append(newRow, &BoardCell{Value: value, Marked: false})
	}

	b.Rows = append(b.Rows, newRow)
}

func (b *Board) ProcessNumber(number int) {
	for _, row := range b.Rows {
		for _, cell := range row {
			if cell.Value == number {
				cell.Marked = true
			}
		}
	}
}

func (b *Board) IsWinner() bool {
	return b.hasWinningRow() || b.hasWinningColumn()
}

func (b *Board) hasWinningRow() bool {
	for _, row := range b.Rows {
		if winningCells(row) {
			return true
		}
	}

	return false
}

func (b *Board) hasWinningColumn() bool {
	columns_number := len(b.Rows)

	for i := 0; i < columns_number; i++ {
		column := column(b.Rows, i)
		if winningCells(column) {
			return true
		}
	}

	return false
}

func column(rows [][]*BoardCell, columnIndex int) []*BoardCell {
	column := []*BoardCell{}
	for _, row := range rows {
		column = append(column, row[columnIndex])
	}

	return column
}

func winningCells(cells []*BoardCell) bool {
	for _, cell := range cells {
		if !cell.Marked {
			return false
		}
	}

	return true
}

type BoardCell struct {
	Value  int
	Marked bool
}

func main() {
	fmt.Printf("Part one...\n")
	partOne()
	fmt.Printf("Part two...\n")
	partTwo()
}

func partOne() {
	bingo := Init(readLines())
	winner, number := bingo.FindFirstWinner()
	fmt.Printf("Score: %d\n", winner.Score(number))
}

func partTwo() {
	bingo := Init(readLines())
	winner, number := bingo.FindLastWinner()
	fmt.Printf("Score: %d\n", winner.Score(number))
}

func readLines() []string {
	bytes, _ := ioutil.ReadFile("../input.txt")
	return strings.Split(string(bytes), "\n")
}

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
	if len(os.Args) > 3 || len(os.Args) < 2 {
		log.Fatalln("Invalid number of args passed")
		return
	}

	filename := os.Args[1]

	if len(os.Args) > 2 {
		part := os.Args[2]

		if part == "1" {
			part_1(filename)
		} else if part == "2" {
			part_2(filename)
		} else {
			log.Fatalln("Invalid part passed")
		}
	} else {
		part_1(filename)
		part_2(filename)
	}
}

func part_1(filename string) {
	array := read_file(filename)
	word := []rune("XMAS")
	count := 0
	for row := range array {
		for col := range array[row] {
			up := move(word, 0, array, row, col, "U")
			upright := move(word, 0, array, row, col, "UR")
			right := move(word, 0, array, row, col, "R")
			downright := move(word, 0, array, row, col, "DR")
			down := move(word, 0, array, row, col, "D")
			downleft := move(word, 0, array, row, col, "DL")
			left := move(word, 0, array, row, col, "L")
			upleft := move(word, 0, array, row, col, "UL")
			if up {
				count++
			}
			if upright {
				count++
			}
			if right {
				count++
			}
			if downright {
				count++
			}
			if down {
				count++
			}
			if downleft {
				count++
			}
			if left {
				count++
			}
			if upleft {
				count++
			}
		}
	}
	fmt.Println("part 1:", count)
}

func part_2(filename string) {
	array := read_file(filename)
	word := []rune("MAS")
	m := make(map[string]bool)
	count := 0
	for row := range array {
		for col := range array[row] {
			upright := move(word, 0, array, row, col, "UR")
			downright := move(word, 0, array, row, col, "DR")
			downleft := move(word, 0, array, row, col, "DL")
			upleft := move(word, 0, array, row, col, "UL")
			if upright && !already_used(row, col, "UR", m) {
				row_n := row - 2
				downright := move(word, 0, array, row_n, col, "DR")
				if downright && !already_used(row_n, col, "DR", m) {
					m[create_id(row_n, col, "DR")] = true
					m[create_id(row, col, "UR")] = true
					count++
				}
				col_n := col + 2
				upleft := move(word, 0, array, row, col_n, "UL")
				if upleft && !already_used(row, col_n, "UL", m) {
					m[create_id(row, col_n, "UL")] = true
					m[create_id(row, col, "UR")] = true
					count++
				}
			}
			if downright && !already_used(row, col, "DR", m) {
				row_n := row + 2
				upright := move(word, 0, array, row_n, col, "UR")
				if upright && !already_used(row_n, col, "UR", m) {
					m[create_id(row_n, col, "UR")] = true
					m[create_id(row, col, "DR")] = true
					count++
				}
				col_n := col + 2
				downleft := move(word, 0, array, row, col_n, "DL")
				if downleft && !already_used(row, col_n, "DL", m) {
					m[create_id(row, col_n, "DL")] = true
					m[create_id(row, col, "DR")] = true
					count++
				}
			}
			if downleft && !already_used(row, col, "DL", m) {
				row_n := row + 2
				upleft := move(word, 0, array, row_n, col, "UL")
				if upleft && !already_used(row_n, col, "UL", m) {
					m[create_id(row_n, col, "UL")] = true
					m[create_id(row, col, "DL")] = true
					count++
				}
				col_n := col - 2
				downright := move(word, 0, array, row, col_n, "DR")
				if downright && !already_used(row, col_n, "DR", m) {
					m[create_id(row, col_n, "DR")] = true
					m[create_id(row, col, "DL")] = true
					count++
				}
			}
			if upleft && !already_used(row, col, "UL", m) {
				row_n := row - 2
				downleft := move(word, 0, array, row_n, col, "DL")
				if downleft && !already_used(row_n, col, "DL", m) {
					m[create_id(row_n, col, "DL")] = true
					m[create_id(row, col, "UL")] = true
					count++
				}
				col_n := col - 2
				upright := move(word, 0, array, row, col_n, "UR")
				if upright && !already_used(row, col_n, "UR", m) {
					m[create_id(row, col_n, "UR")] = true
					m[create_id(row, col, "UL")] = true
					count++
				}
			}
		}
	}
	fmt.Println("part 2:", count)
}

func read_file(filename string) [][]rune {
	file, err := os.Open(filename)
	if err != nil {
		fmt.Printf("File %s does not exist", filename)
	}

	array := [][]rune{}
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		array = append(array, []rune(scanner.Text()))
	}

	return array
}

func move(word []rune, idx int, array [][]rune, row int, col int, direction string) bool {
	if idx == len(word) {
		return true
	}
	if row < 0 || row >= len(array) {
		return false
	}
	if col < 0 || col >= len(array[row]) {
		return false
	}
	if array[row][col] == word[idx] {
		if strings.Contains(direction, "U") {
			row = row - 1
		}
		if strings.Contains(direction, "R") {
			col = col + 1
		}
		if strings.Contains(direction, "D") {
			row = row + 1
		}
		if strings.Contains(direction, "L") {
			col = col - 1
		}
		return move(word, idx+1, array, row, col, direction)
	}
	return false
}

func already_used(row int, col int, direction string, m map[string]bool) bool {
	val := create_id(row, col, direction)
	_, prs := m[val]
	return prs
}

func create_id(row int, col int, direction string) string {
	return strconv.Itoa(row) + ", " + strconv.Itoa(col) + " " + direction
}

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
			if upright {
				ur := strconv.Itoa(row) + ", " + strconv.Itoa(col) + " UR"
				_, prs_ur := m[ur]
				row_n := row - 2
				downright := move(word, 0, array, row_n, col, "DR")
				dr := strconv.Itoa(row_n) + ", " + strconv.Itoa(col) + " DR"
				_, prs_dr := m[dr]
				if downright && !prs_dr && !prs_ur {
					m[dr] = true
					m[ur] = true
					count++
				}
				col_n := col + 2
				upleft := move(word, 0, array, row, col_n, "UL")
				ul := strconv.Itoa(row) + ", " + strconv.Itoa(col_n) + " UL"
				_, prs_ul := m[ul]
				if upleft && !prs_ul && !prs_ur {
					m[ul] = true
					m[ur] = true
					count++
				}
			}
			if downright {
				dr := strconv.Itoa(row) + ", " + strconv.Itoa(col) + " DR"
				_, prs_dr := m[dr]
				row_n := row + 2
				upright := move(word, 0, array, row_n, col, "UR")
				ur := strconv.Itoa(row_n) + ", " + strconv.Itoa(col) + " UR"
				_, prs_ur := m[ur]
				if upright && !prs_ur && !prs_dr {
					m[dr] = true
					m[ur] = true
					count++
				}
				col_n := col + 2
				downleft := move(word, 0, array, row, col_n, "DL")
				dl := strconv.Itoa(row) + ", " + strconv.Itoa(col_n) + " DL"
				_, prs_dl := m[dl]
				if downleft && !prs_dl && !prs_dr {
					m[dl] = true
					m[dr] = true
					count++
				}
			}
			if downleft {
				dl := strconv.Itoa(row) + ", " + strconv.Itoa(col) + " DL"
				_, prs_dl := m[dl]
				row_n := row + 2
				upleft := move(word, 0, array, row_n, col, "UL")
				ul := strconv.Itoa(row_n) + ", " + strconv.Itoa(col) + " UL"
				_, prs_ul := m[ul]
				if upleft && !prs_ul && !prs_dl {
					m[ul] = true
					m[dl] = true
					count++
				}
				col_n := col - 2
				downright := move(word, 0, array, row, col_n, "DR")
				dr := strconv.Itoa(row) + ", " + strconv.Itoa(col_n) + " DR"
				_, prs_dr := m[dr]
				if downright && !prs_dr && !prs_dl {
					m[dr] = true
					m[dl] = true
					count++
				}
			}
			if upleft {
				ul := strconv.Itoa(row) + ", " + strconv.Itoa(col) + " UL"
				_, prs_ul := m[ul]
				row_n := row - 2
				downleft := move(word, 0, array, row_n, col, "DL")
				dl := strconv.Itoa(row_n) + ", " + strconv.Itoa(col) + " DL"
				_, prs_dl := m[dl]
				if downleft && !prs_dl && !prs_ul {
					m[dl] = true
					m[ul] = true
					count++
				}
				col_n := col - 2
				upright := move(word, 0, array, row, col_n, "UR")
				ur := strconv.Itoa(row) + ", " + strconv.Itoa(col_n) + " UR"
				_, prs_ur := m[ur]
				if upright && !prs_ur && !prs_ul {
					m[ur] = true
					m[ul] = true
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

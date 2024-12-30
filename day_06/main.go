package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
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
	grid := read_file(filename)
	start_pos := get_start_pos(grid)
	dis_pos_count := find_distinct_positions(grid, start_pos, 0)
	fmt.Println("part 1:", dis_pos_count)
}

func part_2(filename string) {
	grid := read_file(filename)
	start_pos := get_start_pos(grid)
	inf_loop_count := find_infinite_loops(grid, start_pos, 0)
	fmt.Println("part 2:", inf_loop_count)
}

func find_infinite_loops(grid [][]rune, start_pos position, count int) int {
	for row := range grid {
		for col := range grid[row] {
			if grid[row][col] != '#' && !(row == start_pos.row && col == start_pos.col) {
				new_grid := make([][]rune, len(grid))
				new_grid = deep_copy(new_grid, grid)
				new_grid[row][col] = '#'
				if is_infinite_loop(new_grid, start_pos, make(map[position]bool)) {
					count++
				}
			}
		}
	}
	return count
}

func copy_map(new_map map[position]bool, m map[position]bool) map[position]bool {
	for k, v := range m {
		new_map[k] = v
	}
	return new_map
}

func deep_copy(new_grid [][]rune, grid [][]rune) [][]rune {
	for i := range grid {
		new_grid[i] = make([]rune, len(grid[i]))
		copy(new_grid[i], grid[i])
	}
	return new_grid
}

func is_infinite_loop(grid [][]rune, pos position, m map[position]bool) bool {
	_, prs := m[pos]
	if prs {
		return true
	}

	m[pos] = true

	new_row := pos.row + pos.row_move
	new_col := pos.col + pos.col_move

	if new_row < 0 ||
		new_col < 0 ||
		new_row >= len(grid) ||
		new_col >= len(grid[0]) {
		return false
	}

	var row_move int
	var col_move int
	if grid[new_row][new_col] == '#' {
		row_move, col_move = move_90(pos.row_move, pos.col_move)
		return is_infinite_loop(grid, position{pos.row, pos.col, row_move, col_move}, m)
	} 
	return is_infinite_loop(grid, position{pos.row + pos.row_move, pos.col + pos.col_move, pos.row_move, pos.col_move}, m)
}

func find_distinct_positions(grid [][]rune, pos position, count int) int {
	if grid[pos.row][pos.col] != 'X' {
		grid[pos.row][pos.col] = 'X'
		count++
	}

	new_row := pos.row + pos.row_move
	new_col := pos.col + pos.col_move

	if new_row < 0 ||
		new_col < 0 ||
		new_row >= len(grid) ||
		new_col >= len(grid[0]) {
		return count
	}

	var row_move int
	var col_move int
	if grid[new_row][new_col] == '#' {
		row_move, col_move = move_90(pos.row_move, pos.col_move)
	} else {
		row_move = pos.row_move
		col_move = pos.col_move
	}

	return find_distinct_positions(grid, position{pos.row + row_move, pos.col + col_move, row_move, col_move}, count)
}

func move_90(row, col int) (int, int) {
	if row == 1 && col == 0 {
		return 0, -1
	} else if row == -1 && col == 0 {
		return 0, 1
	} else if row == 0 && col == 1 {
		return 1, 0
	} else if row == 0 && col == -1 {
		return -1, 0
	}
	panic("Invalid position")
}

func get_start_pos(grid [][]rune) position {
	for row := range grid {
		for col := range grid[row] {
			switch grid[row][col] {
			case '^':
				return position{row, col, -1, 0}
			case '>':
				return position{row, col, 0, 1}
			case '<':
				return position{row, col, 0, -1}
			case 'v':
				return position{row, col, 1, 0}
			}
		}
	}
	panic("No starting position found")
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

type position struct {
	row      int
	col      int
	row_move int
	col_move int
}

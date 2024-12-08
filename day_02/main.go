package main

import (
	"bufio"
	"fmt"
	"log"
	"math"
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

	result := 0

	for row := range array {
		safe := evaluate_safety(array[row])
		if safe {
			result++
		}
	}

	fmt.Println("part 1:", result)
}

func part_2(filename string) {
	array := read_file(filename)

	result := 0

	for row := range array {
		safe := evaluate_safety(array[row])
		if safe {
			result++
		} else {
			row_safe := false
			for r := 0; r < len(array[row]); r++ {
				//remove value at index
				l := array[row][0:r]
				r := array[row][r+1:]
				new_array := []int{}
				new_array = append(new_array, l...)
				new_array = append(new_array, r...)

				safe := evaluate_safety(new_array)
				if safe {
					row_safe = true
					break
				}
			}
			if row_safe {
				result++
			}
		}
	}

	fmt.Println("part 2:", result)
}

func read_file(filename string) [][]int {
	file, err := os.Open(filename)
	if err != nil {
		fmt.Printf("File %d does not exist", filename)
	}

	array := [][]int{}
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		values := strings.Split(scanner.Text(), " ")
		integers := []int{}
		for _, v := range values {
			integer, _ := strconv.Atoi(v)
			integers = append(integers, integer)
		}
		array = append(array, integers)
	}

	return array
}

func evaluate_safety(array []int) bool {
	direction := 0
	l := 0
	r := 1
	for r < len(array) {
		//check if equal
		if array[l] == array[r] {
			break
		}

		//check if wrong direction
		if array[l] > array[r] {
			if direction == 1 {
				break
			}
			direction = -1
		}
		if array[l] < array[r] {
			if direction == -1 {
				break
			}
			direction = 1
		}

		//check if difference > 3
		difference := math.Abs(float64(array[l] - array[r]))
		if difference > 3 {
			break
		}
		l++
		r++
	}

	//row is safe
	return r == len(array)
}

func evaluateRow(array []int) bool {
	direction := 0
	success := false
	for i := 0; i < len(array)-1; i++ {
		j := i + 1
		success, direction = compareTwo(array, i, j, direction)
		if !success {
			return false
		}
	}
	return true
}

func compareTwo(array []int, i int, j int, direction int) (bool, int) {
	//is the next value decreasing/increasing
	if array[i] > array[j] {
		if direction > 0 {
			return false, direction
		}
		direction = -1
	}
	if array[i] < array[j] {
		if direction < 0 {
			return false, direction
		}
		direction = 1
	}
	if array[i] == array[j] {
		return false, direction
	}
	//is the difference larger than 3
	difference := array[i] - array[j]
	if difference < 0 {
		difference = difference * -1
	}
	if difference > 3 {
		return false, direction
	}
	return true, direction
}

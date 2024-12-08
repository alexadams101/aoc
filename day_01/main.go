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
	array1, array2 := read_file(filename)

	//sort them in asc order
	array1 = quicksort(array1)
	array2 = quicksort(array2)

	//sum the differences of each entry in the arrays
	sum := 0
	for i := 0; i < len(array1); i++ {
		diff := math.Abs(float64(array1[i] - array2[i]))
		sum = sum + int(diff)
	}

	fmt.Println("part 1:", sum)
}

func part_2(filename string) {
	array1, array2 := read_file(filename)

	//add count of each number in array2 to a map
	m := make(map[int]int)
	for _, v := range array2 {
		val, prs := m[v]
		if prs {
			m[v] = val + 1
		} else {
			m[v] = 1
		}
	}

	//multiple each element in array1 by the count in array2
	result := 0
	for _, v := range array1 {
		val, prs := m[v]
		if prs {
			result = result + v*val
		}
	}

	fmt.Println("part 2:", result)
}

func read_file(filename string) ([]int, []int) {
	file, err := os.Open(filename)
	if err != nil {
		fmt.Printf("File %s does not exist", filename)
	}

	array1 := []int{}
	array2 := []int{}
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		values := strings.Split(scanner.Text(), "   ")
		val1, _ := strconv.Atoi(values[0])
		val2, _ := strconv.Atoi(values[1])
		array1 = append(array1, val1)
		array2 = append(array2, val2)
	}
	return array1, array2
}

func quicksort(array []int) []int {
	if len(array) <= 1 {
		return array
	}

	idx := -1
	pivot := array[len(array)-1]

	for i := 0; i < len(array)-1; i++ {
		if array[i] <= pivot {
			idx++
			val := array[i]
			array[i] = array[idx]
			array[idx] = val
		}
	}
	idx++
	array[len(array)-1] = array[idx]
	array[idx] = pivot

	left := quicksort(array[0:idx])
	right := quicksort(array[idx+1:])

	left = append(left, array[idx])
	return append(left, right...)
}

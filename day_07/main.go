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
	m := read_file(filename)
	results := make(map[int]bool)
	result := 0
	for k, v := range m {
		goal := k
		combinations := get_combinations_p1(v, len(v)-1)
		for _, c := range combinations {
			if goal == c {
				_, prs := results[c]
				if !prs {
					results[c] = true
					result = result + c
				}
			}
		}
	}
	fmt.Println("part 1:", result)
}

func part_2(filename string) {
	m := read_file(filename)
	results := make(map[int]bool)
	result := 0
	for k, v := range m {
		goal := k
		combinations := get_combinations_p2(v, len(v)-1)
		for _, c := range combinations {
			if goal == c {
				_, prs := results[c]
				if !prs {
					results[c] = true
					result = result + c
				}
			}
		}
	}
	fmt.Println("part 2:", result)
}

func get_combinations_p2(numbers []int, pos int) []int {
	if pos == 0 {
		return []int{numbers[pos]}
	}

	combinations := get_combinations_p2(numbers, pos-1)

	list := []int{}
	for _, c := range combinations {
		pow := int(math.Floor(math.Log10(float64(numbers[pos])) + 1))
		list = append(list, c*int(math.Pow10(pow))+numbers[pos])
		list = append(list, numbers[pos]+c)
		list = append(list, numbers[pos]*c)
	}
	return list
}

func read_file(filename string) map[int][]int {
	file, err := os.Open(filename)
	if err != nil {
		fmt.Printf("File %s does not exist", filename)
	}

	m := make(map[int][]int)
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		strs := strings.Split(scanner.Text(), ": ")
		key, _ := strconv.Atoi(strs[0])
		val := []int{}
		for _, i := range strings.Split(strs[1], " ") {
			str, _ := strconv.Atoi(i)
			val = append(val, str)
		}
		m[key] = val
	}

	return m
}

func get_combinations_p1(numbers []int, pos int) []int {
	if pos == 0 {
		return []int{numbers[pos]}
	}

	combinations := get_combinations_p1(numbers, pos-1)

	list := []int{}
	for _, c := range combinations {
		list = append(list, numbers[pos]+c)
		list = append(list, numbers[pos]*c)
	}
	return list
}

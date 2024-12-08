package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"regexp"
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
	text := read_file(filename)
	regex, _ := regexp.Compile("mul\\((\\d{1,3}),(\\d{1,3})\\)")
	commands := regex.FindAllStringSubmatch(text, -1)

	result := 0
	for _, v := range commands {
		param_1, _ := strconv.Atoi(v[1])
		param_2, _ := strconv.Atoi(v[2])
		result = param_1*param_2 + result
	}
	fmt.Println("part 1:", result)
}

func part_2(filename string) {
	text := read_file(filename)
	regex, _ := regexp.Compile("don't\\(\\)|do\\(\\)|mul\\((\\d{1,3}),(\\d{1,3})\\)")
	commands := regex.FindAllStringSubmatch(text, -1)

	enabled := true
	result := 0
	for _, v := range commands {
		if strings.Contains(v[0], "do()") {
			enabled = true
		} else if strings.Contains(v[0], "don't()") {
			enabled = false
		} else if enabled {
			param_1, _ := strconv.Atoi(v[1])
			param_2, _ := strconv.Atoi(v[2])
			result = param_1*param_2 + result
		}
	}
	fmt.Println("part 2:", result)
}

func read_file(filename string) string {
	file, err := os.Open(filename)
	if err != nil {
		fmt.Printf("File %s does not exist", filename)
	}
	scanner := bufio.NewScanner(file)
	var sb strings.Builder
	for scanner.Scan() {
		sb.WriteString(scanner.Text())
	}
	return sb.String()
}

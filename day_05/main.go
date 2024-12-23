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
	rules, updates := read_file(filename)
	rules_map := create_rules_map(rules)
	result := 0
	for _, update := range updates {
		if valid_update(update, rules_map) {
			result = result + find_mid(update)
		}
	}
	fmt.Println("part 1:", result)
}

func part_2(filename string) {
	rules, updates := read_file(filename)
	rules_map := create_rules_map(rules)
	result := 0
	for _, update := range updates {
		if !valid_update(update, rules_map) {
			rules_map := create_rules_map_for_update(rules, update)
			rules_count := create_rules_count_for_update(rules, update)
			new_update := correct_order(update, rules_map, rules_count)
			result = result + find_mid(new_update)
		}
	}
	fmt.Println("part 2:", result)
}

func read_file(filename string) ([]Rule, [][]int) {
	file, err := os.Open(filename)
	if err != nil {
		fmt.Printf("File %s does not exist", filename)
	}

	rules := []Rule{}
	updates := [][]int{}
	scanner := bufio.NewScanner(file)
	x := true
	for scanner.Scan() {
		line := scanner.Text()
		if line == "" {
			x = false
		} else if x {
			ints := strings.Split(line, "|")
			before, _ := strconv.Atoi(ints[0])
			after, _ := strconv.Atoi(ints[1])
			var rule Rule
			rule.before = before
			rule.after = after
			rules = append(rules, rule)
		} else {
			update := strings.Split(line, ",")
			vals := []int{}
			for _, str := range update {
				val, _ := strconv.Atoi(str)
				vals = append(vals, val)
			}
			updates = append(updates, vals)
		}
	}
	return rules, updates
}

func create_rules_map(rules []Rule) map[int][]int {
	rules_map := make(map[int][]int)
	for _, rule := range rules {
		key := rule.before
		val := rule.after
		list, prs := rules_map[key]
		if !prs {
			rules_map[key] = []int{}
		}
		rules_map[key] = append(list, val)
	}
	return rules_map
}

func valid_update(update []int, rules map[int][]int) bool {
	for idx, val := range update {
		r := rules[val]
		remaining_values := update[0:idx]
		for _, i := range r {
			contains := false
			for _, v := range remaining_values {
				if v == i {
					contains = true
					break
				}
			}
			if contains {
				return false
			}
		}
	}
	return true
}

func find_mid(update []int) int {
	return update[len(update)/2]
}

func correct_order(update []int, rules_map map[int][]int, rules_count map[int]int) []int {
	list := []int{}
	for !counts_are_zero(rules_count) {
		for k, v := range rules_count {
			if v == 0 {
				if rules_count[k] == 0 {
					list = append(list, k)
					delete(rules_count, k)
				}
				dependants := rules_map[k]
				for _, d := range dependants {
					if rules_count[d] > 0 {
						rules_count[d] = rules_count[d] - 1
					}
				}
			}
		}
	}
	return list
}

func counts_are_zero(list map[int]int) bool {
	for _, v := range list {
		if v != 0 {
			return false
		}
	}
	return true
}

func create_rules_map_for_update(rules []Rule, update []int) map[int][]int {
	rules_map := make(map[int][]int)
	for _, rule := range rules {
		if contains(rule.before, rule.after, update) {
			key := rule.before
			val := rule.after
			list, prs := rules_map[key]
			if !prs {
				rules_map[key] = []int{}
			}
			rules_map[key] = append(list, val)
		}
	}
	return rules_map
}

func contains(x int, y int, list []int) bool {
	x_count := false
	y_count := false
	for _, i := range list {
		if i == x {
			x_count = true
		}
		if i == y {
			y_count = true
		}
	}
	return x_count && y_count
}

func create_rules_count_for_update(rules []Rule, update []int) map[int]int {
	rules_count := make(map[int]int)
	for _, u := range update {
		_, prs := rules_count[u]
		if !prs {
			rules_count[u] = 0
		}
		for _,r := range rules {
			if r.after == u && contains(r.before, r.after, update) {
                rules_count[u] = rules_count[u] + 1
			}
		}
	}
	return rules_count
}

type Rule struct {
	before int
	after  int
}

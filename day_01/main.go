package main

import (
	"fmt"
	"os"
	"bufio"
	"strings"
	"strconv"
)

func main() {
	if len(os.Args) != 2 {
		fmt.Println("Invalid number of args passed")
		return
	}

	data := os.Args[1]

	//read lists from file and and to array
	file,err := os.Open(data)
	if err != nil {
		fmt.Printf("File %d does not exist", data)
	}

	array1 := []int{}
	array2 := []int{}
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		values := strings.Split(scanner.Text(), "   ")
		val1,_ := strconv.Atoi(values[0])
		val2,_ := strconv.Atoi(values[1])
		array1 = append(array1, val1)
		array2 = append(array2, val2)
	}

	//add values in array2 to map
	m := make(map[int]int)
	for _,v := range array2 {
		val,prs := m[v]
		if prs {
			m[v] = val + 1
		} else {
			m[v] = 1
		}
	}

	//multiple each element by the count in array2
	result := 0
	for _,v := range array1 {
		val,prs := m[v]
		if prs {
			result = result + v * val
		}
	}
	fmt.Println(result)

//	//sort them in asc order
//	array1 = sort(array1)
//	array2 = sort(array2)
//
//
//	//for each entry in arrays sum the difference between the two
//	result := 0
//	for i:=0; i<len(array1); i++ {
//		num := array1[i] - array2[i]
//		if num < 0 {
//			num = num * -1
//		}
//		result = result + num
//	}
//	fmt.Println(result)
}

func sort(array []int) []int {
	if len(array) <= 1 {
		return array
	}

	idx := -1
	pivot := array[len(array)-1]

	for i:=0; i<len(array)-1; i++ {
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

	left := sort(array[0:idx])
	right := sort(array[idx+1:len(array)])

	left = append(left, array[idx])
	return append(left, right...)
}
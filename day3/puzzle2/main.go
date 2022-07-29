package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)


func maximumDigit(array [][]string, index int) []string {
	if index >= len(array[0]) || len(array) == 1 {
		return array[0]
	}
	var oneArray, zeroArray [][]string
	var counter1, counter2 int
	for _, val := range array {
		if val[index] == "1" {
			oneArray = append(oneArray, val)
			counter1++
		} else {
			zeroArray = append(zeroArray, val)
			counter2++
		}
	}
	if counter1 >= counter2 {
		return maximumDigit(oneArray, index+1)
	} else {
		return maximumDigit(zeroArray, index+1)
	}
}

func minimumDigit(array [][]string, index int) []string {
	if index >= len(array[0]) || len(array) == 1 {
		return array[0]
	}
	var oneArray, zeroArray [][]string
	var counter1, counter2 int
	for _, val := range array {
		if val[index] == "1" {
			oneArray = append(oneArray, val)
			counter1++
		} else {
			zeroArray = append(zeroArray, val)
			counter2++
		}
	}
	if counter1 >= counter2 {
		return minimumDigit(zeroArray, index+1)
	} else {
		return minimumDigit(oneArray, index+1)
	}
}

func main() {

	file, err := os.Open("input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()
	reader := bufio.NewScanner(file)

	var array [][]string
	for reader.Scan() {
		array = append(array, strings.Split(reader.Text(), ""))
	}
	num1 := maximumDigit(array, 0)
	num2 := minimumDigit(array, 0)

	var oxygen, co2 string
	for _, v := range num1 {
		oxygen += v
	}
	for _, v := range num2 {
		co2 += v
	}

	oxy, err := strconv.ParseInt(oxygen, 2, 13)
	if err != nil {
		log.Fatal(err)
	}
	c, err := strconv.ParseInt(co2, 2, 13)
	if err != nil {
		log.Fatal(err)
	}

	result := oxy * c
	fmt.Println("Результат:", result)
}

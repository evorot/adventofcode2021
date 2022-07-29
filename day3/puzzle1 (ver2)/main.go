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

	file, err := os.Open("input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()
	reader := bufio.NewScanner(file)

	var array [][]string
	var counter int
	var gamma string
	var epsilon string
	for reader.Scan() {
		array = append(array, strings.Split(reader.Text(), ""))

	}

	for i := 0; i < len(array[0]); i++ {
		for row := 0; row < len(array); row++ {
			if array[row][i] == "1" {
				counter++
			}
		}
		if counter > len(array)/2 {
			gamma += "1"
			epsilon += "0"
		} else {
			gamma += "0"
			epsilon += "1"
		}
		counter = 0
	}

	num1, err := strconv.ParseInt(gamma, 2, 13)
	if err != nil {
		log.Fatal(err)
	}
	num2, err := strconv.ParseInt(epsilon, 2, 13)
	if err != nil {
		log.Fatal(err)
	}
	result := num1 * num2
	fmt.Println("Результат:", result)
	fmt.Println(gamma, epsilon)
}

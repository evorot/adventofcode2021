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

	sum := make([]int, 12, 12)
	var counter int
	var gamma string
	var epsilon string
	for reader.Scan() {
		binary := strings.Split(reader.Text(), "")
		for i, val := range binary {
			if val == "1" {
				sum[i] += 1
			}
		}
		counter += 1
	}

	for _, val := range sum {
		if val > counter/2 {
			gamma += "1"
			epsilon += "0"
		} else {
			gamma += "0"
			epsilon += "1"
		}
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
	fmt.Println(sum)
	fmt.Println(gamma, epsilon)
}

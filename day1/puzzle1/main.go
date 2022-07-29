package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
)

func main() {
	file, err := os.Open("input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()
	reader := bufio.NewScanner(file)
	var array []int
	var increased int
	var decreased int
	for reader.Scan() {
		a, err := strconv.Atoi(reader.Text())
		if err != nil {
			log.Println(err)
		}
		array = append(array, a)
	}

	for i := 1; i < len(array); i++ {
		if array[i] > array[i-1] {
			increased++
		}
		if array[i] < array[i-1] {
			decreased++
		}
	}
	fmt.Println("Увеличилось:", increased)
	fmt.Println("Уменьшилось:", decreased)
}

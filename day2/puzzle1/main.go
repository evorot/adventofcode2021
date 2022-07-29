package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"regexp"
	"strconv"
)

func main() {
	file, err := os.Open("input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()
	reader := bufio.NewScanner(file)

	var (
		xsum, ysum int
	)

	for reader.Scan() {

		tPattern := regexp.MustCompile(`[a-z]+`)
		numPattern := regexp.MustCompile(`\d+`)
		t := tPattern.FindString(reader.Text())
		num, err := strconv.Atoi(numPattern.FindString(reader.Text()))
		if err != nil {
			log.Println(err)
		}
		switch t {
		case "forward":
			xsum += num
		case "up":
			ysum -= num
		case "down":
			ysum += num

		}
	}
	result := xsum * ysum

	fmt.Println("Результат:", result)

}

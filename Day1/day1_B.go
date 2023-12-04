package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"unicode"
)

func main() {

	digWords := []string{"one", "two", "three", "four", "five", "six", "seven", "eight", "nine"}

	f, _ := os.Open("./input.txt")
	defer f.Close()
	sc := bufio.NewScanner(f)

	num := 0
	for sc.Scan() {
		numStr := ""
		for i, c := range sc.Text() {
			if unicode.IsDigit(c) {
				numStr += string(c)
				continue
			}
			// Part 2
			for j, w := range digWords {
				if c == rune(w[0]) {
					if len(w)+i <= len(sc.Text()) && w == sc.Text()[i:len(w)+i] {
						numStr += strconv.Itoa(j + 1)
						break
					}

				}
			}
		}
		scNum, _ := strconv.Atoi(string(numStr[0]) + string(numStr[len(numStr)-1]))
		num += scNum
	}
	fmt.Println(num)
}

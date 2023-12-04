package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"unicode"
)

func main() {
	f, _ := os.Open("./input.txt")
	defer f.Close()
	sc := bufio.NewScanner(f)

	num := 0
	for sc.Scan() {
		numStr := ""
		for _, c := range sc.Text() {
			if unicode.IsDigit(c) {
				numStr += string(c)
				continue
			}

		}
		scNum, _ := strconv.Atoi(string(numStr[0]) + string(numStr[len(numStr)-1]))
		num += scNum
	}
	fmt.Println(num)
}

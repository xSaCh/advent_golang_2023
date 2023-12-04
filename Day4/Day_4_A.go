package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	f, _ := os.Open("./day4.txt")
	defer f.Close()
	sc := bufio.NewScanner(f)

	sum := 0
	for sc.Scan() {
		nums := strings.Split(sc.Text(), "|")
		winningNums := []int{}
		pt := 0
		for _, n := range strings.Split(strings.Split(nums[0], ":")[1], " ") {
			if n == "" {
				continue
			}
			in, _ := strconv.Atoi(n)
			winningNums = append(winningNums, in)
		}
		for _, n := range strings.Split(nums[1], " ") {
			if n == "" {
				continue
			}
			myNum, _ := strconv.Atoi(n)

			for _, wn := range winningNums {
				if myNum == wn {
					if pt == 0 {
						pt = 1
					} else {
						pt *= 2
					}
				}
			}
		}
		// fmt.Printf("pt: %v\n", pt)
		sum += pt
		// fmt.Printf("winningNums: %v\n", winningNums)
	}
	fmt.Printf("sum: %v\n", sum)
}

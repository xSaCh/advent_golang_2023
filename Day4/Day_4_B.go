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

	freq := map[int]int{}

	for sc.Scan() {
		winningNums := []int{}

		nums := strings.Split(sc.Text(), "|")
		sep := strings.Split(nums[0], ":")
		cardSep := strings.Split(sep[0], " ")

		cardNo, _ := strconv.Atoi(cardSep[len(cardSep)-1])
		freq[cardNo] = freq[cardNo] + 1

		for _, n := range strings.Split(sep[1], " ") {
			if n == "" {
				continue
			}
			in, _ := strconv.Atoi(n)
			winningNums = append(winningNums, in)
		}
		pt := 1
		for _, n := range strings.Split(nums[1], " ") {
			if n == "" {
				continue
			}
			myNum, _ := strconv.Atoi(n)

			for _, wn := range winningNums {
				if myNum == wn {
					freq[cardNo+pt] += freq[cardNo]
					pt += 1
					break
				}
			}
		}
	}
	sum := 0
	for _, v := range freq {
		sum += v
	}
	fmt.Printf("sum: %v\n", sum)
}

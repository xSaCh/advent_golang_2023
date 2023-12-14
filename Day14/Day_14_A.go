package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	f, _ := os.Open("./input.txt")
	defer f.Close()
	sc := bufio.NewScanner(f)

	dish := [][]rune{}

	for sc.Scan() {
		dish = append(dish, []rune(sc.Text()))
	}
	// dish = tilt(dish)
	tilt(dish)

	sum := 0
	for i, v := range dish {
		sum += strings.Count(string(v), "O") * (len(dish) - i)

	}
	fmt.Printf("sum: %v\n", sum)
}

func tilt(dish [][]rune) {
	for i := 0; i < len(dish); i++ {
		for j := 0; j < len(dish[i]); j++ {
			if dish[j][i] != 'O' {
				continue
			}

			newI := getEmptySpot(dish, j, i)
			if newI == -1 {
				continue
			}
			dish[newI][i], dish[j][i] = dish[j][i], dish[newI][i]
		}

	}
}

func getEmptySpot(d [][]rune, minH, cln int) int {
	lastEm := 0
	for i := minH - 1; i >= 0; i-- {
		if d[i][cln] == '.' {
			lastEm = i
		}
		if d[i][cln] == '#' || d[i][cln] == 'O' {
			return i + 1
		}
	}
	return lastEm
}

func printD(a [][]rune) {
	for _, v := range a {
		fmt.Println(string(v))
	}
}

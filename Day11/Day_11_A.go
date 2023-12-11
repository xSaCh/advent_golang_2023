package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"slices"
	"strings"
)

type Coord struct {
	x int
	y int
}

func main() {
	f, _ := os.Open("./input.txt")
	defer f.Close()
	sc := bufio.NewScanner(f)
	// sum := 0

	img := [][]rune{}
	for sc.Scan() {
		img = append(img, []rune(sc.Text()))
	}

	rows := []int{}
	cols := []int{}
	for i := 0; i < len(img); i++ {
		if !strings.ContainsRune(string(img[i]), '#') {
			rows = append(rows, i)
		}
		isEmpty := true
		for j := 0; j < len(img[0]); j++ {
			if img[j][i] == '#' {
				isEmpty = false
				break
			}
		}
		if isEmpty {
			cols = append(cols, i)
		}
	}
	for i, v := range rows {
		img = slices.Insert(img, v+i, []rune(strings.Repeat(".", len(img))))
	}
	for c, v := range cols {
		for i := 0; i < len(img); i++ {
			img[i] = slices.Insert(img[i], v+c, '.')
		}
	}

	galaxiesPos := []Coord{}
	for i, r := range img {
		for j, c := range r {
			if c == '#' {
				galaxiesPos = append(galaxiesPos, Coord{j, i})
			}
		}
	}

	sum := 0
	for i := 0; i < len(galaxiesPos); i++ {
		for j := i + 1; j < len(galaxiesPos); j++ {
			sum += manhanttenDis(galaxiesPos[i], galaxiesPos[j])
		}
	}
	fmt.Printf("sum: %v\n", sum)

}

func manhanttenDis(p1, p2 Coord) int {
	return int(math.Abs(float64(p1.x)-float64(p2.x)) + math.Abs(float64(p1.y)-float64(p2.y)))
}

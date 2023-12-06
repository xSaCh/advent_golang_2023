package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	f, _ := os.Open("./input.txt")
	defer f.Close()
	sc := bufio.NewScanner(f)

	sc.Scan()
	timesStr := strings.Fields(strings.Split(sc.Text(), ":")[1])
	sc.Scan()
	disStr := strings.Fields(strings.Split(sc.Text(), ":")[1])

	times, dis := []int{}, []int{}

	errMargin := 1
	for i := 0; i < len(timesStr); i++ {
		t, _ := strconv.Atoi(timesStr[i])
		times = append(times, t)
		d, _ := strconv.Atoi(disStr[i])
		dis = append(dis, d)

		ways := 0
		for j := 0; j <= t; j++ {
			curDis := (t - j) * j
			if curDis > d {
				ways++
			}
		}
		errMargin *= ways
	}
	fmt.Printf("errMargin: %v\n", errMargin)
}

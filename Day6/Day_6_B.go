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

	time, _ := strconv.Atoi(strings.Join(timesStr, ""))
	dis, _ := strconv.Atoi(strings.Join(disStr, ""))

	ways := 0
	for j := 0; j <= time; j++ {
		curDis := (time - j) * j
		if curDis > dis {
			ways++
		}
	}
	fmt.Printf("ways: %v\n", ways)
}

package main

import (
	"bufio"
	"os"
	"strconv"
	"strings"
)

const (
	NO_RED   = 12
	NO_GREEN = 13
	NO_BLUE  = 14
)

func main() {
	f, _ := os.Open("./input.txt")
	defer f.Close()
	sc := bufio.NewScanner(f)

	sum := 0
	for sc.Scan() {
		isPossible := true
		sp := strings.Split(sc.Text(), ":")
		id, _ := strconv.Atoi(strings.Split(sp[0], " ")[1])

	ballLoop:
		for _, set := range strings.Split(sp[1], ";") {
			for _, ball := range strings.Split(set, ",") {
				b := strings.Split(ball, " ")
				n, _ := strconv.Atoi(b[1])
				switch b[2] {
				case "red":
					if n > NO_RED {
						isPossible = false
					}
					break
				case "green":
					if n > NO_GREEN {
						isPossible = false
					}
					break
				case "blue":
					if n > NO_BLUE {
						isPossible = false
					}
					break
				}
				if !isPossible {
					break ballLoop
				}
			}
		}
		if isPossible {
			sum += id
		}
	}
	println(sum)
}

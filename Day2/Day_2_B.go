package main

import (
	"bufio"
	"os"
	"strconv"
	"strings"
)

func main() {
	f, _ := os.Open("./day2.txt")
	defer f.Close()
	sc := bufio.NewScanner(f)

	sum := 0
	for sc.Scan() {
		sp := strings.Split(sc.Text(), ":")
		maxR, maxG, maxB := 1, 1, 1

		for _, set := range strings.Split(sp[1], ";") {
			for _, ball := range strings.Split(set, ",") {
				b := strings.Split(ball, " ")
				n, _ := strconv.Atoi(b[1])
				switch b[2] {
				case "red":
					if n > maxR {
						maxR = n
					}
					break
				case "green":
					if n > maxG {
						maxG = n
					}
					break
				case "blue":
					if n > maxB {
						maxB = n
					}
					break
				}
			}
		}
		sum += (maxR * maxG * maxB)
	}
	println(sum)
}

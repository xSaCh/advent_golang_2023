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

	seqs := []string{}

	for sc.Scan() {
		str := sc.Text()
		seqs = strings.Split(str, ",")
	}

	sum := 0
	for _, s := range seqs {
		sum += hash(s)
	}
	fmt.Println(hash("qp"))
}

func hash(str string) int {
	n := 0
	for _, c := range str {
		n += int(c)
		n *= 17
		n %= 256
	}
	return n
}

package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

type Node struct {
	// cur   string
	left  string
	right string
}

func main() {
	f, _ := os.Open("./input.txt")
	defer f.Close()
	sc := bufio.NewScanner(f)

	sc.Scan()
	seq := sc.Text()
	sc.Scan()

	// network := []Coord{}
	network := map[string]Node{}
	for sc.Scan() {
		node := Node{}
		nodeS := strings.Split(sc.Text(), " = ")
		nxt := strings.Split(nodeS[1], ", ")

		node.left = nxt[0][1:len(nxt[0])]
		node.right = nxt[1][:len(nxt[1])-1]
		network[nodeS[0]] = node
	}

	curLoc := "AAA"
	curNode := network[curLoc]
	i := 0
	step := 0
	for curLoc != "ZZZ" {
		if seq[i] == 'R' {
			curLoc = curNode.right
		} else {
			curLoc = curNode.left
		}
		curNode = network[curLoc]
		i = (i + 1) % len(seq)
		step++
	}
	fmt.Printf("step: %v\n", step)
}

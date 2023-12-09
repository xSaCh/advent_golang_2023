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

	network := map[string]Node{}
	locWithA := []string{}
	for sc.Scan() {
		node := Node{}
		nodeS := strings.Split(sc.Text(), " = ")
		nxt := strings.Split(nodeS[1], ", ")

		node.left = nxt[0][1:len(nxt[0])]
		node.right = nxt[1][:len(nxt[1])-1]
		network[nodeS[0]] = node

		if nodeS[0][len(nodeS[0])-1] == 'A' {
			locWithA = append(locWithA, nodeS[0])
		}
	}
	fmt.Printf("locWithA: %v\n", locWithA)

	steps := []int{}
	for _, loc := range locWithA {
		curLoc := loc
		curNode := network[curLoc]
		i := 0
		step := 0
		for curLoc[len(curLoc)-1] != 'Z' {
			if seq[i] == 'R' {
				curLoc = curNode.right
			} else {
				curLoc = curNode.left
			}
			curNode = network[curLoc]
			i = (i + 1) % len(seq)
			step++
		}
		steps = append(steps, step)
	}

	fmt.Printf("steps: %v\n", lcm(steps))

}

func gcd(a, b int) int {
	if b == 0 {
		return a
	}
	return gcd(b, a%b)
}

func lcm(nums []int) int {
	result := nums[0]
	for i := 1; i < len(nums); i++ {
		result = (result * nums[i]) / gcd(result, nums[i])
	}
	return result
}

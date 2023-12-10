package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
	"time"
)

func main() {
	f, _ := os.Open("./input.txt")
	defer f.Close()
	sc := bufio.NewScanner(f)
	sum := 0

	s := time.Now()
	for sc.Scan() {
		hisStr := strings.Split(sc.Text(), " ")
		history := []int{}

		for i := len(hisStr) - 1; i >= 0; i-- {
			v, _ := strconv.Atoi(hisStr[i])
			history = append(history, v)
		}

		diffLists := [][]int{history}
		curDiff := getDiffList(history)
		diffLists = append(diffLists, curDiff)

		for !isZeroList(curDiff) {
			curDiff = getDiffList(curDiff)
			diffLists = append(diffLists, curDiff)
		}
		predictVal := 0
		for i := len(diffLists) - 2; i >= 0; i-- {

			his := diffLists[i]
			predictVal = his[len(his)-1] + predictVal
		}
		sum += predictVal
		// break
	}
	fmt.Printf("sum: %v\n", sum)

}

func isZeroList(list []int) bool {
	for _, v := range list {
		if v != 0 {
			return false
		}
	}
	return true
}

func getDiffList(list []int) []int {
	diffList := []int{}
	for i := 1; i < len(list); i++ {
		diffList = append(diffList, list[i]-list[i-1])
	}
	return diffList
}

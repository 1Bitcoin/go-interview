package main

import (
	"fmt"
	"math/rand"
)

func main() {
	fmt.Println(myRand(20))
}

func myRand(n int) []int {
	var ans []int
	duplicate := make(map[int]struct{})
	for len(ans) < n {
		value := rand.Intn(20)

		if _, exist := duplicate[value]; !exist {
			ans = append(ans, value)
			duplicate[value] = struct{}{}
		}
	}

	return ans
}

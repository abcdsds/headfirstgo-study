package main

import "fmt"

func Votes(votes []string) {
	counts := make(map[string]int)
	for _, vote := range votes {
		counts[vote]++
	}

	fmt.Println(counts)

	// for .. range 루프는 맵을 무작위 순위로 처리
	for key, value := range counts {
		fmt.Printf("%s vote %v \n", key, value)
	}
}

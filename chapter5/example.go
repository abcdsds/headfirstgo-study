package main

import (
	"fmt"
	"time"
)

func main() {
	var nodes [7]string
	nodes[0] = "do"
	nodes[1] = "re"
	nodes[2] = "mi"

	fmt.Println(nodes[0])
	fmt.Println(nodes[1])

	var primes [5]int
	primes[0] = 2
	primes[1] = 2

	fmt.Println(primes[0])

	var dates [4]time.Time
	dates[0] = time.Unix(1257894000, 0)
	dates[1] = time.Unix(1447920000, 0)
	dates[2] = time.Unix(1257894000, 0)

	fmt.Println(dates[1])
	fmt.Println(dates[3])

	//배열 리터럴
	rr := [3]int{9, 18, 27}
	fmt.Println(rr[0])
	fmt.Println(rr[1])

	rr[0] = 1

	fmt.Println(rr[0])

	nodes2 := [7]string{"A", "B", "C", "D", "E", "F", "G"}

	for i := 0; i < len(nodes2); i++ {
		fmt.Println(i, nodes2[i])
	}

	for index, element := range nodes2 {
		fmt.Println(index, element)
	}

	// _를 통한 return value 생략 가능
	for index, _ := range nodes2 {
		fmt.Println(index)
	}
}

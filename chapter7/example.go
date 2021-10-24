package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

func main() {
	lines, err := GetStrings("votes.txt")
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(lines)

	var names []string
	var counts []int
	for _, line := range lines {
		matched := false
		for i, name := range names {
			if name == line {
				counts[i]++
				matched = true
			}
		}
		if matched == false {
			names = append(names, line)
			counts = append(counts,1)
		}
	}

	for i, name := range names {
		fmt.Printf("%s: %d\n", name, counts[i])
	}

	// map 시작
	//var ranks map[string]int
	//ranks = make(map[string]int)

	ranks2 := make(map[string]int)

	ranks2["gold"] = 1
	ranks2["silver"] = 2
	ranks2["bronze"] = 3
	ranks2["iron"] = 4
	fmt.Println(ranks2["iron"])
	fmt.Println(ranks2["gold"])
	fmt.Println(ranks2["diamond"])

	// map 리터럴
	myMap := map[string]float64{"a": 1.2, "b": 5.6}
	fmt.Println(myMap["a"])
	fmt.Println(myMap["b"])
	fmt.Println(myMap["c"])

	value, ok := myMap["a"]
	fmt.Println(value, ok)

	value, ok = myMap["c"]
	fmt.Println(value, ok)


	myRanks := make(map[string]int)
	myRanks["bronze"] = 3
	rank, ok := myRanks["bronze"]
	fmt.Printf("rank: %d, ok: %v\n", rank, ok)
	delete(myRanks, "bronze")
	rank, ok = myRanks["bronze"]
	fmt.Printf("rank: %d, ok: %v\n", rank, ok)

	Votes(lines)

}

func GetStrings (fileName string) ([]string, error){
	var lines []string
	file, err := os.Open(fileName)
	if err != nil {
		return nil, err
	}

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		lines = append(lines, line)
	}

	err = file.Close()
	if err != nil {
		return nil, err
	}

	if scanner.Err() != nil {
		return nil, scanner.Err()
	}

	return lines, nil
}

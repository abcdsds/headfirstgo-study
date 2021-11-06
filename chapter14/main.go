package main

import "fmt"

func main() {
	phrases := []string{"my parents", "a rodeo clown"}
	fmt.Println("A photo of", JoinWithCommas(phrases))
}

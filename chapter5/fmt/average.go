package main

import (
	"fmt"
	"log"
	"main/file"
)

func main() {
	numbers := [3]float64{71.8, 56.2, 89.5}
	var sum float64 = 0
	for _, number := range numbers {
		sum += number
	}

	sampleCount := float64(len(numbers))
	fmt.Printf("Average: %0.2f\n", sum/sampleCount)

	myFloat, err := file.GetFloats("../file/data.txt")
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(myFloat)
	var sum2 float64 = 0
	for _, number2 := range myFloat {
		sum2 += number2
	}

	sampleCount2 := float64(len(myFloat))
	fmt.Printf("Average: %0.2f\n", sum2/sampleCount2)

}

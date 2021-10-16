package main

import (
	"fmt"
	"log"
	"math"
	"os"
	"strconv"
)


func main()  {
	var mySlice []string
	mySlice = make([]string, 7)

	mySlice[0] = "do"
	mySlice[1] = "re"
	mySlice[2] = "mi"

	fmt.Println(mySlice)

	notes := []string {"do", "re", "mi"}
	fmt.Println(notes)

	primes := []int {
		2,
		3,
		5,
	}

	fmt.Println(primes[0], primes[1], primes[2])

	testSlice := primes[2:3]
	fmt.Println(testSlice)

	testSlice2 := primes[:3]
	fmt.Println(testSlice2)

	testSlice3 := primes[3:]
	fmt.Println(testSlice3)

	fmt.Println(len(primes))

	testSlice3 = append(testSlice3, 6)
	fmt.Println(testSlice3)

	testSlice3 = append(testSlice3, 6,7,8)
	fmt.Println(testSlice3)

	testSlice4 := append(testSlice3, 6,7,8)
	fmt.Println(testSlice3, testSlice4)

	testSlice5 := append(testSlice4, 9)
	fmt.Println(testSlice3, testSlice4, testSlice5)

	testSlice4[1] = 1
	fmt.Println(testSlice3, testSlice4, testSlice5)

	s1 := []string{"s1", "s1"}
	s1 = append(s1, "s2", "s2")
	s3 := append(s1, "s3", "s3")
	s4 := append(s3, "s4", "s4")

	fmt.Println(s1, s1, s3, s4)
	s3[0] = "XX"
	fmt.Println(s1, s1, s3, s4)
	nums, err := GetFloats("data.txt")
	fmt.Println(err)
	fmt.Println(nums)

	var sum float64 = 0
	for _, number := range nums {
		sum += number
	}

	sampleCount := float64(len(nums))
	fmt.Printf("Average: %0.2f\n", sum / sampleCount)
	average2()

	// 가변 인자 함수
	fmt.Println(1)
	fmt.Println(1,2,3,4,5)
	myFunc("1","2","3","4","5")
	myFunc()

	max := maximum(1,2,3,4,5,6,7,8,9)
	fmt.Println(max)

	fmt.Println(inRange(1,100, -12.5, 3.2, 0, 50, 103.5))

	float64Slice := []float64{-12.5, 3.2, 0, 50, 103.5}
	fmt.Println(inRange(1,100, float64Slice...))


}

func average2() {
	arguments := os.Args[1:]
	var sum float64 = 0
	for _, argument := range arguments {
		num, err := strconv.ParseFloat(argument, 64)
		if err != nil {
			log.Fatal(err)
		}
		sum += num
	}

	sampleCount := float64(len(arguments))
	fmt.Printf("Average: %0.2f\n", sum / sampleCount)
}

func myFunc(param ...string) {
	fmt.Println(param)
}

func maximum(numbers ...float64) float64{
	max := math.Inf(-1)
	for _, number := range numbers {
		if number > max {
			max = number
		}
	}

	return max
}

func inRange(min float64, max float64, numbers ...float64) []float64 {
	var result []float64
	for _, number := range numbers {
		if number >= min && number <= max {
			result = append(result, number)
		}
	}
	return result
}
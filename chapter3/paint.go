package main

import (
	"fmt"
	"log"
)

func paint() {
	a := 4.2
	b := 3.0
	amount, err := paintNeeded(a, b) // 1. 함수로 인자를 전달할때
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(amount)

	paintNeeded(5.2, 3.5)
	paintNeeded(5.0, 3.3)

	am := 6
	newDivide(&am)
}

func paintNeeded(width float64, height float64) (float64, error) {
	area := width * height
	//fmt.Printf("%.2f liters needed\n", divide(area))
	return area / 10.0, nil
}

// 2. 매개변수는 인자 값을 복사한다.
func divide(number float64) float64 {
	number *= 2 // 3. 원래 값이 아닌 복사된 값을 변경한다.
	return number / 10.0
}

// 4. 인자의 값을 변경하기 위해서는 포인터를 사용해야 한다.
func newDivide(number *int) {
	*number *= 2
	*number /= 10.0
}

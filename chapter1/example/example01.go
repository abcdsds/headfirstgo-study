package main

import (
	"fmt"
	"math"
	"reflect"
	"strings"
)

type aa struct {
	v1 string
}

func main() {
	fmt.Println("Hello, Go!")

	fmt.Println(math.Floor(2.75))
	fmt.Println(strings.Title("head first go"))
	fmt.Println(reflect.TypeOf("Hello, Go!"))
	fmt.Println(reflect.TypeOf(true))
	fmt.Println(reflect.TypeOf(3.1415))
	fmt.Println(reflect.TypeOf(45))

	var quantity int
	var length, width float64
	var customerName string

	quantity = 2
	customerName = "daemon Cole"
	length, width = 1.2, 2.4

	printOrder(customerName, quantity, length, width)

	var quantity2 = 2
	var customerName2 = "daemon Cole"
	var length2, width2 = 1.2, 2.4

	printOrder(customerName2, quantity2, length2, width2)

	var myInt int
	var myFloat float64

	fmt.Println(myInt, myFloat)

	var myString string
	var myBool bool
	fmt.Println(myString, myBool)

	// 단축 변수 선언
	quantity3 := 4
	length3, width3 := 1.2, 2.4
	customerName3 := "Daemon Cole"

	printOrder(customerName3, quantity3, length3, width3)

	/** 네이밍 규칙
	 유효 length, stack2, sales.Total
	 잘못됨 2stack (숫자로 시작 불가능), sales.total(대문자로 시작하지 않는 변수는 타 패키지에서 접근할수 없다.) (중요한듯)
	 golang은 변수 네이밍을 카멜케이스를 권장

	 이름이 대문자로 시작하는 변수, 함수, 타입은 패키지 외부로 노출되고 외부의 다른 패키지에서 접근 가능
	**/

	// 타입 변환
	var myInt int = 2
	fmt.Println(reflect.TypeOf(myInt))
	fmt.Println(reflect.TypeOf(float64(myInt)))
}

func printOrder(customerName string, quantity int, length, width float64) {
	fmt.Println(customerName)
	fmt.Println("has ordered", quantity, "sheets")
	fmt.Println("each with an area of")
	fmt.Println(length*width, "square meters")
}

package main

import (
	"errors"
	"fmt"
	"reflect"
)

func main() {
	//var width, height, area float64
	reslutString := fmt.Sprintf("About one-thrid: %0.2f\n", 1.0/3.0)
	fmt.Printf(reslutString)

	//형식 동사 (formatting verb)
	fmt.Printf("The %s cost %d cents each.\n", "gumballs", 23)
	fmt.Printf("%v %v %v", "", "\t", "\n")
	fmt.Printf("%12s | %s\n", "Product", "Cost in cents")
	// %7.3 -> 7 은 전체 숫자의 최소 너비, 3은 소수점 이하 자릿수의 너비
	fmt.Printf("%%7.3f: %7.3f\n", 12.3456789)
	sayHi()
	repeatLine("hello", 3)
	paint()

	err := errors.New("height can't be negative")
	fmt.Println(err)
	//log.Fatal(err)

	myNum, myBool, myString := manyReturns()
	fmt.Println(myBool, myNum, myString)

	puzzle()

	// 포인터

	amount2 := 6
	fmt.Println(amount2)  // 변수의 값을 가져온다.
	fmt.Println(&amount2) // 변수의 주소를 가져온다.

	var myInt int
	fmt.Println(reflect.TypeOf(&myInt))
	var myFloat float64
	fmt.Println(reflect.TypeOf(&myFloat))
	var myBool2 bool
	fmt.Println(reflect.TypeOf(&myBool2))

	var myInt2 int
	var myIntPointer *int  // int 타입에 대한 포인터 값을 갖는 변수를 선언
	myIntPointer = &myInt2 // 변수에 포인터를 할당

	fmt.Println(myIntPointer)

	var myFloat2 float64
	var myFloatPointer *float64

	myFloatPointer = &myFloat2
	fmt.Println(myFloatPointer)

	var myBool3 bool
	myBoolPointer := &myBool3
	fmt.Println(myBoolPointer)

	// 포인터 값 가져오거나 변경하기

	mInt := 4
	mIntPointer := &mInt
	fmt.Println(&myInt)
	fmt.Println(mIntPointer)
	fmt.Println(*mIntPointer)

	// 포인터 주소는 그대로지만 값은 변경
	// 포인터 주소는 mInt의 주소이므로 mInt의 값도 변경된다.

	*mIntPointer = 8
	fmt.Println(mIntPointer)
	fmt.Println(*mIntPointer)

	fmt.Println(*createPointer())

}

/*
	카멜케이스를 사용
	다른 패키지에서 사용되는 경우에는 대문자로 시작 ( 대문자로 시작하지 않으면 외부에서 호출불가 )
*/

func sayHi() {
	fmt.Println("Hi!")
}

func repeatLine(line string, times int) {
	for i := 0; i < times; i++ {
		fmt.Println(line)
	}
}

func manyReturns() (int, bool, string) {
	return 1, true, "hello"
}

func createPointer() *float64 {
	myFloat := 98.5
	return &myFloat
}

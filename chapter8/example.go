package main

import (
	"fmt"
	maga "github.com/headfirstgo/magazine"
)



// 구조체를 외부 패키지에서 호출하려면 첫글자가 대문자여야 한다.
// 마찬가지로 필드도 첫글자가 대문자여야지 외부 패키지에서 사용가능하다.
type subscriber struct {
	name string
	rate float64
	active bool
	Rate float64
}

func main() {
	var myStruct struct {
		word string
		number int

	}

	fmt.Println(myStruct.number)
	fmt.Println(myStruct.word)

	myStruct.word = "abcde"
	myStruct.number = 111

	fmt.Println(myStruct.number)
	fmt.Println(myStruct.word)

	var subscriber1 subscriber
	subscriber1.name = "aman singh"
	subscriber1.rate = 10
	subscriber1.active = true

	showSubscriber(subscriber1)
	changeValue(subscriber1)
	showSubscriber(subscriber1)
	changeValue2(&subscriber1)
	showSubscriber(subscriber1)

	//포인터를 통해 구조체 필드 접근
	var value int = 2
	var pointer *int = &value
	fmt.Println(*pointer)

	var mySubscriber subscriber
	mySubscriber.name = "john"
	var pointerr *subscriber = &mySubscriber
	fmt.Println((*pointerr).name)
	//result: john
	fmt.Println(pointerr.name)
	//result: john
	changeValue3("john 2")
	showSubscriber(subscriber1)

	mySubscriber2 := subscriber{name: "d", active: true}
	fmt.Println(mySubscriber2)


	var address maga.Address
	address.State = "a"
	fmt.Println(address)

	/*
		외부 구조체의 익명 필드로 선언된 내부 구조체를 외부 구조체 안에 임베딩 되었다고 한다.
		임베딩된 구조체의 필드는 외부 구조체로 승격되는데 이는 내부 구조체의 필드를 마치 외부 구조체에 속해 있는 것처럼 접근할수 있다는 것을 의미
	*/
	var employee maga.Employee
	fmt.Println(employee.State)


}

func showSubscriber(s subscriber) {
	fmt.Println(s)
}

// 매개변수는 구조체의 복사본이기 때문에 변경해도 원본은 변경이 안된다.
// 자바는 객체는 주소를 가져와서 복사가 된다.
func changeValue(s subscriber) {
	s.active = false
}

//매개변수를 구조체가 아닌 구조체 포인터값으로 받으면 원본 변경가능
func changeValue2(s *subscriber) {
	s.active = false
}

func changeValue3(name string) *subscriber {
	var s subscriber
	s.name = name

	return &s
}

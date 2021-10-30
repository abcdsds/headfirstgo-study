package main

import (
	ddd "chapter10/date"
	ccc "chapter10/calendar"
	"errors"
	"fmt"
	"log"
)

type Date struct {
	Year  int
	Month int
	Day   int
}

type Date2 struct {
	year  int
	month int
	day   int
}

func main() {
	//임베딩된 Date의 필드에는 접근불가
	//필드와 같이 메소드도 승격됐기 때문에 메소드로 접근가능
	bb := ccc.Event{}
	fmt.Println(bb.Year())

	aa := ddd.Date2{}
	aa.SetYear(1998)
	fmt.Println(aa.Year())

	date1 := Date{Year: 1999, Month: 5, Day: 8}
	fmt.Println(date1)
	fmt.Println(date1.Year)
	err := date1.SetYear(2001)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(date1)

	date2 := Date2{year: 1888}
	fmt.Println(date2)
	//같은 패키지 내부라 필드를 대문자로 안해도 나옴
	fmt.Println(date2.year)

}

func (d *Date) SetYear(year int) error {
	if year < 1 {
		return errors.New("invalid year")
	}
	d.Year = year
	return nil
}

func (d *Date2) SetYear(year int) error {
	if year < 1 {
		return errors.New("invalid year")
	}
	d.year = year
	return nil
}

/* 접근자 메서드
go는 get접두사를 붙이는 컨벤션을 버리기로함. set접두사는 여전히 사용
*/

func (d *Date2) Year() int {
	return d.year
}

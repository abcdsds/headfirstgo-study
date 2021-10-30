package main

import "fmt"

type Liters float64
type Milliliters float64
type Gallons float64
type MyType string

func main() {
	var carFuel Gallons
	var busFuel Liters
	carFuel = Gallons(10.0)
	busFuel = Liters(240.0)
	fmt.Println(carFuel, busFuel)

	carFuel2 := ToGallons(busFuel)
	busFuel2 := ToLiters(carFuel)
	fmt.Println(carFuel2, busFuel2)
	mt := MyType("dd")
	mt.sayHi()
	mt.sayHi2(1, false)
	fmt.Println(mt.sayHi3())
	mt.changeString()
	fmt.Println(mt)
	mt.changeStringWithPointer()
	fmt.Println(mt)

	soda := Liters(2)
	fmt.Println(soda.ToGallons())
	water := Milliliters(10)
	fmt.Println(water.MillilitersToGallons())
}

func ToGallons(l Liters) Gallons {
	return Gallons(l * 0.264)
}

func ToLiters(g Gallons) Liters {
	return Liters(g * 3.785)
}

func MillilitersToGallons(m Milliliters) Gallons {
	return Gallons(m * 0.000264)
}

//리시버 매개변수
func (m MyType) sayHi() {
	fmt.Print(m)
}

func (m MyType) sayHi2(number int, flag bool) {
	fmt.Println(m)
	fmt.Println(number)
	fmt.Println(flag)
}

func (m MyType) sayHi3() bool{
	return true
}

//실제로 변경되지는 않음 값객체가 복사되기 떄문에 변경하려면 포인터를 전달해서 수정해야함
func (m MyType) changeString() {
	m = "nono"
}

//포인터 전달하여 수정
func (m *MyType) changeStringWithPointer() {
	*m = "nono2"
}

func (l Liters)ToGallons() Gallons {
	return Gallons(l * 0.264)
}

func (m Milliliters)MillilitersToGallons() Gallons {
	return Gallons(m * 0.000264)
}


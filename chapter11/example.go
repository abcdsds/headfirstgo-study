package main

import (
	"chapter11/gadget"
	"fmt"
)

type Whistle string
type Horn string
type Robot string
type NoiseMaker interface {
	MakeSound()
}
type AnyThing interface{}
type ComedyError string

func (w *Whistle) MakeSound() {
	fmt.Println("Tweet")
}
func (h *Horn) MakeSound() {
	fmt.Println("Honk")
}
func (r Robot) MakeSound() {
	fmt.Println("beep beep")
}
func (r Robot) Walk() {
	fmt.Println("Powering legs")
}
func (c ComedyError) Error() string {
	return string(c)
}

func main() {
	player := gadget.TapePlayer{}
	mixtape := []string{"first", "second", "third"}
	playList(player, mixtape)

	whistle := Whistle("Toyco Canary")
	horn := Horn("Toyco Blaster")
	var toy NoiseMaker = &whistle
	play(toy)
	toy = &horn
	play(toy)

	var noiseMaker NoiseMaker = Robot("Bot")
	noiseMaker.MakeSound()
	// 인터페이스를 구현한 메소드와 아닌 메소드가 같이 있는 구현체에서 아닌 메소드 사용시 따로 타입캐스팅하여 실행시켜야한다.
	var robot Robot = noiseMaker.(Robot)
	robot.Walk()

	//type check
	r, ok := noiseMaker.(Robot)
	if ok {
		r.Walk()
	}

	comedyError := ComedyError("beer Logger")
	fmt.Println(comedyError)

	AcceptAnything(1)
	AcceptAnything("1")

}

func playList(device gadget.Player, songs []string) {
	for _, song := range songs {
		device.Play(song)
	}
	device.Stop()
}

func play(n NoiseMaker) {
	n.MakeSound()
}

func AcceptAnything(thing AnyThing){
	fmt.Println(thing)
}
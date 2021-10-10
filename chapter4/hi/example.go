package main

import (
	"chapter4/greeting"
	"chapter4/keyboard"
	"fmt"

	keyboard2 "github.com/headfirstgo/keyboard"
)

func main() {
	greeting.Hi()
	greeting.Hello()
	keyboard.GetFloat()

	fmt.Print(keyboard.KeyboardName)
	fmt.Println(keyboard2.GetFloat())
	//fmt.Println(keyboard2)
}

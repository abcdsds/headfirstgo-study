package main

import (
	"fmt"
	"log"
	"time"
)

func test() error {
	return nil
}

func sendLetters(channel chan string) {
	time.Sleep(1 * time.Second)
	channel <- "a"
	time.Sleep(1 * time.Second)
	channel <- "b"
	time.Sleep(1 * time.Second)
	channel <- "c"
	time.Sleep(1 * time.Second)
	channel <- "d"
}

func main() {
	// err를 if문 안으로 넣어 err를 할당안해도 되게 
	if err := test(); err != nil {
		log.Fatal(err)
	}
	
	// go는 switch문에서 break 생략해도 됨
	
	/** go에는 버퍼있는 채널과 버퍼없는 채널이 존재
		한 고루틴이 버퍼없는 채널에 값을 전송하면 해당 고루틴은 다른 고루틴이 채널에서 값을 가져갈때 까지 블로킹 됩니다.
		반면에 버퍼있는 채널은 송신 고루틴을 블로킹하기 전까지 특정 개수의 값을 보유할 수 있습니다.
	*/

	
	// 버퍼개수를 넘어가면 데드락이 발생함
	channel := make(chan string, 3)
	channel <- "a"
	channel <- "b"
	channel <- "c"
	//channel <- "d"

	fmt.Println(<-channel)
	fmt.Println(<-channel)
	fmt.Println(<-channel)
	//fmt.Println(<-channel)

	// 채널이 없는경우 수신되기전까지 값을 가질수가 없다 그렇기 때문에 1초씩 기다렸다가 작동하지만 버퍼를 가지면
	// 추가로 값을 가질수 있기때문에 바로 저장한다.

	fmt.Println(time.Now())
	channel2 := make(chan string)
	go sendLetters(channel2)
	time.Sleep(5 * time.Second)
	fmt.Println(<-channel2, time.Now())
	fmt.Println(<-channel2, time.Now())
	fmt.Println(<-channel2, time.Now())
	fmt.Println(<-channel2, time.Now())
	fmt.Println(time.Now())

}

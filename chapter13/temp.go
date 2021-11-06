package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"time"
)

type Page struct {
	URL string
	Size int
}

func main() {

	/** 고루틴은 실행시점을 보장하지 않는다.
		고루틴은 return 값을 가진 함수를 사용할 수 없다.
		return 되지 않은 시점에 return 값을 받을 수 없기 때문
		fmt.Print로도 작동하지 않는다 아직 return 값이 없기 때문
	 */
	go responseSize("https://golang.org/") // 9964
	go responseSize("https://example.com/") // 1256
	go responseSize("https://golang.org/doc") // 18764

	// 고루틴은 채널을 통해 값을 반환 받을 수 있다.
	myChannel := make(chan string)
	go aa(myChannel)
	//fmt.Println(<-myChannel)

	//채널은 수신 고루틴이 값을 사용하기 전에 송신 고루틴이 값을 보냈음을 보장할수가 있다.
	// blocking으로 이를 보장한다 ( 고루틴의 모든 작업을 중지하는 역할 )
	//다른 고루틴이 해당 채널에 값을 보내기전까지 수신 고루틴을 블로킹하는 이런 방식을 통해 고루틴은 자기 자신의 행동을 동기화 할수 있다.
	channel := <-myChannel
	fmt.Println("hh " + channel)

	//main 고루틴을 실행중인 상태로 유지해야 위에 고루틴의 완료 응답을 받을 수 있다.
	time.Sleep(2 * time.Second)

	//strings := make(chan string)
	//c := make(chan string)
	//go abc(strings)
	//go def(c)
	//fmt.Println(<-strings)
	//fmt.Println(<-c)
	//fmt.Println(<-strings)
	//fmt.Println(<-c)
	//fmt.Println(<-strings)
	//fmt.Println(<-c)

	strings := make(chan string)
	go send(strings, "from main")
	reportNap("receiving goroutine", 5)
	fmt.Println(<-strings)
	fmt.Println(<-strings)
	time.Sleep(time.Second)
	fmt.Println("Time Over")
}

func responseSize(url string) {
	get, err := http.Get(url)
	if err != nil {
		log.Fatal(err)
	}

	defer get.Body.Close()
	all, err := ioutil.ReadAll(get.Body)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(len(all))
}

func aa(myChannel chan string) {
	myChannel <- "hi"
}


func abc(channel chan string) {
	channel <- "a"
	channel <- "b"
	channel <- "c"
}

func def(channel chan string) {
	channel <- "d"
	channel <- "e"
	channel <- "f"
}

func reportNap(name string, delay int) {
	for i := 0; i < delay; i++ {
		fmt.Println(name, "sleeping")
		time.Sleep(1 * time.Second)
	}
	fmt.Println(name, "wake up!")
}

func send(myChannel chan string, str string) {
	reportNap("sending goroutine", 2)
	fmt.Println("**** sending value ****")

	// main이 중지된 상태에서 이 전달은 블로킹 된다.
	myChannel <- str + " a"

	fmt.Println("**** sending value ****")
	myChannel <- str + " b"
}

// 구조체를 전달할수도 있음
func getSize(url string, channel chan Page) {
	channel <- Page {
		Size: 100,
		URL: url,
	}
}
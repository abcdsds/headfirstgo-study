package main

import (
	"fmt"
	"log"
	"net/http"
)

func viewHandler(writer http.ResponseWriter, message string) {
	//byte 타입은 문자열로 변환 가능
	//message := []byte("hello, web!")
	_, err := writer.Write([]byte(message))
	if err != nil {
		log.Fatal(err)
	}
}

func handler1(writer http.ResponseWriter, request *http.Request) {
	viewHandler(writer, "1")
}
func handler2(writer http.ResponseWriter, request *http.Request) {
	viewHandler(writer, "2")
}

func sayHi() {
	fmt.Println("hi")
}

func sayBye() {
	fmt.Println("bye")
}

func twice(function func()) {
	function()
	function()
}

func main() {
	hi := sayHi
	hi()

	bye := sayBye
	bye()

	twice(hi)
	twice(bye)

	http.HandleFunc("/hello", handler1)
	http.HandleFunc("/hello2", handler2)
	err := http.ListenAndServe("localhost:8080", nil)
	log.Fatal(err)
}



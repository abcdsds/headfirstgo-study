package main

import (
	"fmt"
	"log"
	"os"
)

func main() {
	/*
		err 이라고 변수명을 짓는 이유는 error 라는 타입의 이름을 가리기 떄문
		즉 기존에 이미 같은 이름으로 선언된것보다 우선시 된다.
	*/

	fileInfo, err := os.Stat("my.txt")
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(fileInfo.Size())
}

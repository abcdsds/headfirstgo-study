//pass_fail 프로그램은 성적의 합격 여부를 알려줍니다.
package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func pass() {
	// 시스템 위젯의 총 개수
	var TotalCount int // 정수만 가능합니다.
	TotalCount = 2
	fmt.Println(TotalCount)
	/*
		Package widfget에는 위젯을 다루는 모든 함수가 포함되어 있습니다.
	*/

	fmt.Print("Enter a grade : ")
	reader := bufio.NewReader(os.Stdin)
	// input, _ := reader.ReadString('\n') //반환값 무시
	// input, err := reader.ReadString('\n')
	// log.Fatal(err)
	input, err := reader.ReadString('\n')
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(input)

	// 단축 변수 선언은 한번만 가능
	s := "\t formerly surrounded by space \n"
	fmt.Println(strings.TrimSpace(s))

	input = strings.TrimSpace(input)
	grade, err := strconv.ParseFloat(input, 64)
	if err != nil {
		log.Fatal(err)
	}

	var status string
	if grade >= 60 {
		status = "passing"
	} else {
		status = "failing"
	}

	fmt.Println("A grade of", "is", status)

}

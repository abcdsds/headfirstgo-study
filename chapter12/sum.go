package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
)

func OpenFile(fileName string) (*os.File, error) {
	fmt.Println("Opening", fileName)
	return os.Open(fileName)
}

func CloseFile(file *os.File) {
	fmt.Println("Closing file")
	file.Close()
}

func GetFloats(fileName string) ([]float64, error) {
	var numbers []float64
	file, err := OpenFile(fileName)
	if err != nil {
		return nil, err
	}
	defer CloseFile(file)
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		number, err := strconv.ParseFloat(scanner.Text(), 64)
		if err != nil {
			return nil, err
		}
		numbers =append(numbers, number)
	}
	if scanner.Err() != nil {
		return nil, scanner.Err()
	}
	return numbers, nil
}

func main() {
	numbers, err := GetFloats(os.Args[1])
	if err != nil {
		log.Fatal(err)
	}
	var sum float64 = 0
	for _, number := range numbers {
		sum += number
	}
	fmt.Printf("Sum: %0.2f\n", sum)

	//err = Socialize()
	//if err != nil {
	//	log.Fatal(err)
	//}

	//강제로 패닉 일으키기
	//panic("oh no, going down")

	//panic이 없으면 nil return 함, 밑에 패닉이 있다면
	fmt.Println(recover())

	//defer는 크래시가 발생하기 전에 실행됨.
	one()

	fmt.Println("run")
}

func Socialize() error {
	// defer는 메소드 호출시에만 사용가능하다 for문과 변수선언과 같은 다른 명령문은 연기 불가능하다.
	defer fmt.Println("GoodBye!")
	fmt.Println("Hello")
	return fmt.Errorf("idw to talk")
	fmt.Println("Nice weather, eh?")
	return nil
}

func one() {
	defer fmt.Println("defeered in one()")
	two()
}

/**같은 스코프에 넣는건 의미가 없음. 패닉상태가 이어지기 떄문
	panic과 recover는 java의 exception처럼 사용하는걸 권장하지 않는다고 한다.
	go 개발자는 exception이 없는 코드가 더 나은 소프트웨어를 만들 수 있을거라고 함.
**/
func two() {
	defer fmt.Println("deffered in tow1()")
	defer fmt.Println("deffered in tow2()")
	defer calmDown()
	panic("deferred!")
	//panic 이후 코드는 실행되지 않음.
	fmt.Println("after panic")
}

//panic의 값은 recover를 통해 return 된다.
func calmDown() {
	fmt.Println("recorver", recover())
}


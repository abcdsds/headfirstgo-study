package main

import (
	"bufio"
	"fmt"
	"log"
	"math/rand"
	"os"
	"strconv"
	"strings"
	"time"
)

func main() {
	now := time.Now()
	year := now.Year()
	fmt.Println(year)

	broken := "G# r#cks!"
	replacer := strings.NewReplacer("#", "o")
	fixed := replacer.Replace(broken)
	fmt.Println(fixed)

	//pass()
	/* 조건이 true라 무한루프
	for x := 1; true; x++ {

	}
	*/
	seconds := time.Now().Unix()
	rand.Seed(seconds)
	target := rand.Intn(100)
	fmt.Println(target)

	success := false
	for guesses := 0; guesses < 10; guesses++ {
		fmt.Println("You have", 10-guesses, "guesses left.")

		reader := bufio.NewReader(os.Stdin)
		fmt.Print("Make a guess: ")
		input, err := reader.ReadString('\n')
		if err != nil {
			log.Fatal(err)
		}

		input = strings.TrimSpace(input)
		guess, err := strconv.Atoi(input)
		if err != nil {
			log.Fatal(err)
		}

		if guess < target {
			fmt.Println("you guess was Low")
		} else if guess > target {
			fmt.Println("you guess was High")
		} else {
			success = true
			fmt.Println("GJ you guessed it!")
			break
		}
	}
	if !success {
		fmt.Println("Sorry, you didn't guess my number. It was:", target)
	}

}

package main

import (
	"bufio"
	"fmt"
	"io"
	"math/rand"
	"os"
	"time"
)

var (
	questions = []string{
		"survey",
		"interview",
		"exhibition",
		"conference",
		"available",
		"annual",
		"equipment",
		"department",
		"refund",
	}
	denom = 0
	numer = 0
)

func main() {
	timeout := time.After(30 * time.Second)
	shuffled := shuffle(questions)
	ch := input(os.Stdin)
Game:
	for _, question := range shuffled {
		denom++
		fmt.Println(question)
		fmt.Print(">")
		select {
		case answer := <-ch:
			if answer == question {
				fmt.Println("ok!")
				numer++
			} else {
				fmt.Println("miss!")
			}
		case <-timeout:
			fmt.Println("Over the time limit!!")
			break Game
		}
	}
	fmt.Printf("Your score is %d/%d", numer, denom)
}

func shuffle(array []string) []string {
	// Copy array
	shuffled := array
	rand.Seed(time.Now().UnixNano())
	for i := range array {
		j := rand.Intn(i + 1)
		shuffled[i], shuffled[j] = shuffled[j], shuffled[i]
	}
	return shuffled
}

func input(r io.Reader) <-chan string {
	ch := make(chan string)
	go func() {
		s := bufio.NewScanner(r)
		for s.Scan() {
			ch <- s.Text()
		}
		close(ch)
	}()
	return ch
}

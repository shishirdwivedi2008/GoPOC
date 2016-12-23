package main

import (
	"fmt"
	"math/rand"
	"time"
)

func array() {
	var arr [100]int
	for x := 0; x < 100; x++ {
		rand.Seed(time.Now().UnixNano())
		y := rand.Intn(10000)
		fmt.Println("random number generated is", y)
		arr[x] = y
	}
}

func main() {
	rand.Seed(2) // Try changing this number!
	answers := []string{
		"It is certain",
		"It is decidedly so",
		"Without a doubt",
		"Yes definitely",
		"You may rely on it",
		"As I see it yes",
		"Most likely",
		"Outlook good",
		"Yes",
		"Signs point to yes",
		"Reply hazy try again",
		"Ask again later",
		"Better not tell you now",
		"Cannot predict now",
		"Concentrate and ask again",
		"Don't count on it",
		"My reply is no",
		"My sources say no",
		"Outlook not so good",
		"Very doubtful",
	}
	fmt.Println("Magic 8-Ball says:", answers[rand.Intn(len(answers))])
	array()
}

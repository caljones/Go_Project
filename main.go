package main

// I am going to write a simple Go program that has fun with Go, printing
// different things about what I have learned about Go so far.

import (
	"fmt"
	"math/rand/v2"
	"os"
	"time"
)

func main() {
	inputText := readFile("./input.txt")
	typePrint(inputText, 70)
}

func readFile(fileName string) string {
	file, err := os.Open(fileName)
	checkErr(err)
	defer file.Close()
	data, err := os.ReadFile(fileName)
	checkErr(err)
	return string(data)
}

func typePrint(input string, speed int) {
	r := rand.New(rand.NewPCG(0, 1))
	for _, c := range input {

		if string(c) == "%" {
			typePrint(readFile("./good_job.txt"), 2)
			continue
		} else if string(c) == "#" {
			typePrint(readFile("./notion.txt"), 2)
			continue
		}

		if string(c) == "@" {
			time.Sleep(time.Second / 2)
			continue
		} else if string(c) == "^" {
			typeDelete()
			time.Sleep(time.Duration(r.IntN(speed)) * time.Millisecond)
			continue
		}

		fmt.Print(string(c))
		time.Sleep(time.Duration(r.IntN(speed)) * time.Millisecond)
	}
}

func typeDelete() {
	fmt.Print("\b \b")
	time.Sleep(50 * time.Millisecond)
}

func checkErr(err error) {
	if err != nil {
		panic(err)
	}
}

package main

// I am going to write a simple Go program that has fun with Go, printing
// different things about what I have learned about Go so far.

import (
	"bufio"
	"fmt"
	"io"
	"math/rand/v2"
	"net/http"
	"os"
	"time"
)

const cali_url = "http://services.explorecalifornia.org/json/tours.php"

func main() {
	inputText := readFile("./input.txt")
	typePrint(inputText, 120)
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

		switch {
		// Print Good Job
		case string(c) == "$":
			typePrint(readFile("./good_job.txt"), 2)
		// Print Cali HTTP call
		case string(c) == "#":
			client := http.Client{}
			req, err := http.NewRequest("GET", cali_url, nil)
			checkErr(err)
			req.Header.Set("User-Agent", "")
			resp, err := client.Do(req)
			checkErr(err)
			bytes, err := io.ReadAll(resp.Body)
			checkErr(err)
			fmt.Println(string(bytes))
			resp.Body.Close()
		// Delay half a second
		case string(c) == "@":
			time.Sleep(time.Second / 2)
		// Backspace
		case string(c) == "^":
			typeDelete()
			time.Sleep(time.Duration(r.IntN(speed)) * time.Millisecond)
		// Dance animation
		case string(c) == "&":
			dance()
		default:
			fmt.Print(string(c))
		}
		time.Sleep(time.Duration(r.IntN(speed)) * time.Millisecond)
	}
}

func typeDelete() {
	fmt.Print("\b \b")
	time.Sleep(50 * time.Millisecond)
}

func dance() {
	file, err := os.Open("./dance.txt")
	checkErr(err)
	defer file.Close()
	scanner := bufio.NewScanner(file)
	counter := 0
	for true {
		for range 5 {
			scanner.Scan()
			s := scanner.Text()
			fmt.Println(s)
			if string(s[len(s)-1]) == "#" {
				goto end
			}
			counter++
		}
		time.Sleep(150 * time.Millisecond)
		for range 5 {
			fmt.Printf("\033[1A\033[K")
		}
	}
end:
}

func checkErr(err error) {
	if err != nil {
		panic(err)
	}
}

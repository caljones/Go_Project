// A program that reads a text file with encoded characters for special functions
package main

import (
	"bufio"
	"fmt"
	"io"
	"math/rand/v2"
	"net/http"
	"os"
	"time"
)

// Name of the file to be typed
const input_filename = "./input.txt"

// URI for an API call
const cali_uri = "http://services.explorecalifornia.org/json/tours.php"

// Millisecond delay between typed characters
const delay = 100

// Main function
func main() {
	inputText := readFile(input_filename)
	typePrint(inputText, delay)
}

// Reads a file from the name provided, outputs file as a string
func readFile(fileName string) string {
	file, err := os.Open(fileName)
	checkErr(err)
	defer file.Close()
	data, err := os.ReadFile(fileName)
	checkErr(err)
	return string(data)
}

// Prints the provided string with the provided millisecond delay between each
// printed character.
// Reads special character and respods accordingly:
//
//	$: Print an ascii art photo in a txt file called "good_job.txt"
//	#: Make an HTTP call to the URI constant
//	@: Sleeps for 0.5 seconds
func typePrint(input string, delay int) {
	r := rand.New(rand.NewPCG(0, 1))
	for _, c := range input {
		switch {
		// Print Good Job
		case string(c) == "$":
			typePrint(readFile("./good_job.txt"), 2)
		// Print Cali HTTP call
		case string(c) == "#":
			client := http.Client{}
			req, err := http.NewRequest("GET", cali_uri, nil)
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
		// Dance animation
		case string(c) == "&":
			dance()
		default:
			fmt.Print(string(c))
			time.Sleep(time.Duration(r.IntN(delay)) * time.Millisecond)
		}
	}
}

// Delete a character
func typeDelete() {
	fmt.Print("\b \b")
	time.Sleep(50 * time.Millisecond)
}

// Run the animation in "dance.txt"
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

// Check for an error and panic if there is one
func checkErr(err error) {
	if err != nil {
		panic(err)
	}
}

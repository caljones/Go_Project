// A program that reads a text file with encoded characters for special functions
package main

import (
	_ "embed"
	"io"
	"math/rand/v2"
	"net/http"
	"strings"
	"time"
)

//go:embed input.txt
var inputTxt string

//go:embed good_job.txt
var goodJobTxt string

//go:embed dance.txt
var danceTxt string

// URI for an API call (overridden in main_wasm.go to proxy through Next.js)
var cali_uri = "https://services.explorecalifornia.org/json/tours.php"

// Millisecond delay between typed characters
const delay = 100

// printFn is set by main_cli.go or main_wasm.go at init time.
var printFn func(string)

// Prints the provided string with the provided millisecond delay between each
// printed character.
// Reads special characters and responds accordingly:
//
//	$: Print an ascii art photo from good_job.txt
//	#: Make an HTTP call to the URI constant
//	@: Sleeps for 0.5 seconds
//	^: Backspace
//	&: Dance animation from dance.txt
func typePrint(input string, delay int) {
	r := rand.New(rand.NewPCG(0, 1))
	for _, c := range input {
		switch string(c) {
		// Print Good Job
		case "$":
			typePrint(goodJobTxt, 2)
		// Print Cali HTTP call
		case "#":
			client := http.Client{}
			req, err := http.NewRequest("GET", cali_uri, nil)
			checkErr(err)
			req.Header.Set("User-Agent", "")
			resp, err := client.Do(req)
			checkErr(err)
			bytes, err := io.ReadAll(resp.Body)
			checkErr(err)
			printFn(string(bytes) + "\n")
			resp.Body.Close()
		// Delay half a second
		case "@":
			time.Sleep(time.Second / 2)
		// Backspace
		case "^":
			typeDelete()
		// Dance animation
		case "&":
			dance()
		default:
			printFn(string(c))
			time.Sleep(time.Duration(r.IntN(delay)) * time.Millisecond)
		}
	}
}

// Delete a character
func typeDelete() {
	printFn("\b \b")
	time.Sleep(50 * time.Millisecond)
}

// Run the animation in dance.txt
func dance() {
	lines := strings.Split(strings.ReplaceAll(danceTxt, "\r\n", "\n"), "\n")
	i := 0
	for {
		for range 5 {
			if i >= len(lines) {
				return
			}
			s := lines[i]
			i++
			printFn(s + "\n")
			if len(s) > 0 && string(s[len(s)-1]) == "#" {
				return
			}
		}
		time.Sleep(150 * time.Millisecond)
		for range 5 {
			printFn("\033[1A\033[K")
		}
	}
}

// Check for an error and panic if there is one
func checkErr(err error) {
	if err != nil {
		panic(err)
	}
}

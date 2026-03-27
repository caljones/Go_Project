//go:build !js

package main

import "fmt"

func init() {
	printFn = func(s string) {
		fmt.Print(s)
	}
}

func main() {
	typePrint(inputTxt, delay)
}

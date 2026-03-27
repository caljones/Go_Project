//go:build js && wasm

package main

import "syscall/js"

func init() {
	printFn = func(s string) {
		js.Global().Call("__goOutput", s)
	}
	cali_uri = js.Global().Get("__toursUrl").String()
}

func main() {
	typePrint(inputTxt, delay)
	js.Global().Call("__goDone")
	select {} // block forever to keep WASM runtime alive
}

package main

import (
	"fmt"
	"syscall/js"
)

func main() {
	c := make(chan struct{}, 0)
	
	// Register our greeting function
	js.Global().Set("greet", js.FuncOf(greet))
	
	// Keep the program running
	<-c
}

func greet(this js.Value, args []js.Value) interface{} {
	message := "Hello from GoAF!"
	// Update DOM
	document := js.Global().Get("document")
	messageDiv := document.Call("getElementById", "message")
	messageDiv.Set("textContent", message)
	
	// Also log to console
	fmt.Println(message)
	return nil
}
package main

import (
	"fmt"
	"io"
	"log"
	"os"
)

func main() {
	r, w, err := os.Pipe()
	if err != nil {
		log.Panic(err)
	}
	os.Stdout = w

	fmt.Println("Hello, playground") // this gets captured

	w.Close()
	out, _ := io.ReadAll(r)

	// os.Stdin.Write() - to write

	fmt.Printf("Captured: %s", out) // prints: Captured: Hello, playground
}

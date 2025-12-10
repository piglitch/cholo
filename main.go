package main

import (
	"fmt"
	"io"
	"log"
	"os"
)
func main() {
	f, err := os.Open("messages.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()
	for {
		b := make([]byte, 8)
		n, err := f.Read(b)
		fmt.Printf("read: %s\n", string(b[:n]))
		if err != nil {
			os.Exit(0)
			return
		}
		if err == io.EOF {
			os.Exit(0)
		}
	}
}
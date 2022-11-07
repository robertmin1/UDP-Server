package main

import (
	"fmt"
	"log"
	"os"
)

func main() {
	file, err := os.Create("text.txt")
    if err != nil {
        log.Fatal(err)
    }
    fmt.Println("File created successfully")
    defer file.Close()
}

package main

import (
	"fmt"
	"os"
)

func getBalanaceFromFile() {
	data, _ := os.ReadFile("name.txt")
	fmt.Println("Hello,", string(data))
}

func writeToFile() {
	var name string
	fmt.Print("Enter your name: ")
	fmt.Scan(&name)
	os.WriteFile("name.txt", []byte(name), 0644)
}

func main() {
	writeToFile()
	getBalanaceFromFile()
}

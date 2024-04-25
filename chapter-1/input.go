package main

import "fmt"

func main() {
	fmt.Printf("Pleace give me your name: ")
	var name string
	_, err := fmt.Scanln(&name)
	if err != nil {
		return
	}
	fmt.Println("Your name is", name)
}

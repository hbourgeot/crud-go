package main

import "fmt"

func printTitle(title string) {
	for i := 0; i < len(title)+4; i++ {
		fmt.Print("*")
	}
	fmt.Print("\n")
	fmt.Printf("* %s *\n", title)
	for i := 0; i < len(title)+4; i++ {
		fmt.Print("*")
	}
}

func main() {
	printTitle("     Hola mundo!    ")
}

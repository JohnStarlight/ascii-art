package main

import (
	"fmt"
	"os"

	"ascii-art/internal"
)

func main() {
	if len(os.Args) != 2 {
		fmt.Println("The correct usage of this program is:")
		fmt.Println("go run . \"your-text-here\"")
		return
	}
	text := os.Args[1]
	fmt.Println("In which style would you like that?")
	fmt.Println("1 = Standard")
	fmt.Println("2 = Shadow")
	fmt.Println("3 = Thinkertoy")

	var choice int
	var filename string
	fmt.Scan(&choice)

	switch choice {
	case 1:
		filename = "standard.txt"
	case 2:
		filename = "shadow.txt"
	case 3:
		filename = "thinkertoy.txt"
	default:
		fmt.Println("The acceptable input is 1, 2, 3.")
		return
	}
	fmt.Printf("%s", internal.PrintAscii(text, filename))
}

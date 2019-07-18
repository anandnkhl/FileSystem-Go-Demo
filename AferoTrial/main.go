package main

import (
	"FileSystem-Go-Demo/AferoTrial/Handlers"
	"fmt"
	"github.com/gookit/color"
)

func main(){
	var input int
	fmt.Println("1.\t Make a Directory")
	fmt.Println("2.\t Make a File")
	fmt.Println("3.\t Write to a File")
	fmt.Println("4.\t Exit")
	_, _ = fmt.Scanf("%d", &input)
	switch input {
	case 1:
		Handlers.CreateDirectory()
		main()
	case 2:
		Handlers.CreateFile()
		main()
	case 3:
		Handlers.WriteFile()
		main()
	case 4:
		color.Red.Println("Bye bye")
		return
	default:
		color.Red.Println("Please enter valid option\n")
		main()
	}
}

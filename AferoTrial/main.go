package main

import (
	"fmt"
	"github.com/gookit/color"
	"github.com/spf13/afero"
)

var AppFs = afero.NewOsFs()

func CreateDirectory(){
	var dirName string
	fmt.Println("Enter the name of directory you want to create:")
	_, _ = fmt.Scanf("%s", &dirName)

	if status,_ := afero.DirExists(AppFs, "./"+dirName); !status{
		_ = AppFs.Mkdir("./"+dirName, 0777)
	}else{
		color.Red.Println("This folder already exists\n")
	}
}

func CreateFile(){
	var fileName string
	fmt.Println("Enter the name of file you want to create:")
	_,_ = fmt.Scanf("%s", &fileName)

	if status,_ := afero.Exists(AppFs, "./"+fileName); !status{
		_, _ = AppFs.Create(fileName)
	} else {
		color.Red.Println("File already exists\n")
	}
}

func main(){
	var input int
	fmt.Println("1.\t Make a Directory")
	fmt.Println("2.\t Make a File")
	fmt.Println("3.\t Exit")
	_, _ = fmt.Scanf("%d", &input)
	switch input {
	case 1:
		CreateDirectory()
		main()
	case 2:
		CreateFile()
		main()
	case 3:
		color.Red.Println("Bye bye")
		return
	default:
		color.Red.Println("Please enter valid option\n")
		main()
	}
}

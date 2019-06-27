package main

import (
	"fmt"
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
		fmt.Println("This folder already exists")
	}
}

func main(){
	var input int
	fmt.Println("1.\t Make a Directory")
	fmt.Scanf("%d", &input)
	switch input {
	case 1: CreateDirectory()
	case 2:
	default:
		fmt.Println("Please enter valid option")
	}
}

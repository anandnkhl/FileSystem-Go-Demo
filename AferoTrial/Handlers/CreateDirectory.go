package Handlers

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
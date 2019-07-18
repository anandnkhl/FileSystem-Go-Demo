package Handlers

import (
	"fmt"
	"github.com/gookit/color"
	"github.com/spf13/afero"
)

func CreateFile(){
	var fileName string
	fmt.Println("Enter the name of file you want to create:")
	_,_ = fmt.Scanf("%s", &fileName)

	if status,_ := afero.Exists(AppFs, "./"+fileName); !status{
		_, err := AppFs.Create(fileName)
		if err != nil {
			color.Red.Println(err,"\n")
		}
	} else {
		color.Red.Println("File already exists\n")
	}
}

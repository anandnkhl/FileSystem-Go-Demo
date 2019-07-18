package Handlers

import (
	"fmt"
	"github.com/gookit/color"
	"os"
)

func WriteFile(){
	var fileName string
	fmt.Println("Enter the file you want to edit:")
	_,_ = fmt.Scanf("%s", &fileName)
	file, err := AppFs.OpenFile("./"+fileName, os.O_APPEND | os.O_WRONLY | os.O_CREATE , 777)
	if err == nil {
		color.Green.Println("File opened successfully")
		fmt.Println("Enter the string to be written:")
		var toAppend string
		_, _ = fmt.Scanf("%s", &toAppend)
		_,err := file.WriteString(toAppend) //WriteFile(AppFs, "./"+fileName, data, 777)
		if err != nil {
			color.Red.Println(err)
		}
	} else{
		color.Red.Println("File did not open / Doesn't exist")
	}
}

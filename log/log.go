package log

import "fmt"

func Log(message string, param ...interface{}) {
	fmt.Print(message)
	//fmt.Println(param)
}

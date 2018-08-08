package models

import "fmt"

type TestConn struct {
	Conn    chan string
	Sclosed chan bool
}


func Txt(in chan string) {
	fmt.Println(<-in)
}

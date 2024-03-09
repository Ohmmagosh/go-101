package ohmmaagoch

import (
	"fmt"
)

func privateSayTest() {
	fmt.Println("hello private test")
}

func SayTest() {
	fmt.Println("HELLO test")
	privateSayTest()
}

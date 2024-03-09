package main

import (
	"fmt"

	"github.com/google/uuid"
)

func main() {
	fmt.Println("hello uuid")
	id := uuid.New().String()
	fmt.Println("id -> ", id)

}

package pointer

import "fmt"

func changeValue(ptr *int) {
	*ptr = 50
}

func Pointer() {
	x := 20
	fmt.Println("x -> before", x)
	changeValue(&x)
	fmt.Println("x -> after", x)
}

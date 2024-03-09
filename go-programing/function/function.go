package function

import (
	"fmt"
)

func add(num1 int, num2 int) int {
	fmt.Println("ADD")
	fmt.Println(num1, " + ", num2, " = ", num1+num2)
	return num1 + num2
}

func sayHello(name string, time int) {
	fmt.Println("time -> ", time)
	for i := 0; i < time; i++ {
		fmt.Println("name -> ", name)
	}
}

func Function() {
	fadd := add(1, 2)
	fmt.Println("value from ADD -> ", fadd)
	sayHello("ohm", 5)
}

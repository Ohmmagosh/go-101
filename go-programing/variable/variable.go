package variable

import "fmt"

func variable2() {
	fullname := "ohm"
	age := 12
	fmt.Printf("hello %s Yay! age = %d\n", fullname, age)
}

func Variable() {
	// BOOLEAN without value
	var bo bool
	// BOOLEAN with value
	var bo1 bool = true
	// INTEGER without value
	var i int
	var i8 int8
	var i16 int16
	var i32 int32
	var i64 int64
	// INTEGER with value
	var i1 int = 12345
	var i81 int8 = 123
	var i161 int16 = 1234
	var i321 int32 = 1234
	var i641 int64 = 1234
	// STRING without value
	var str string
	// STRING with value
	var str1 string = "HELLLO"

	fmt.Println(bo, bo1, i, i1, i8, i81, i16, i161, i32, i321, i64, i641, str, str1)
	variable2()
}

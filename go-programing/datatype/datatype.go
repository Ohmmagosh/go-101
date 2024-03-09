package datatype

import (
	"fmt"
)

func structure() {
	fmt.Println("structure")
	type Person struct {
		Name string
		Age  int
	}
	var a Person
	var b Person = Person{Name: "world", Age: 11}
	a.Name = "hello"
	a.Age = 20
	fmt.Println(a)
	fmt.Println(b)

}

func mymap() {
	myMap := make(map[string]int)

	myMap["apple"] = 5
	myMap["banana"] = 10
	myMap["orange"] = 15
	fmt.Println("myMap ->", myMap)

	for key, value := range myMap {
		fmt.Println("key -> ", key, "val -> ", value)
	}

	val, ok := myMap["apple"]
	if ok {
		fmt.Println("Apple is value : ", val)
	} else {
		fmt.Println("Apple not found in map")
	}
}

func slice() {
	fmt.Println("Slice")
	mySlice := []int{1, 2, 3, 4, 5, 6}

	fmt.Println(mySlice)
	mySlice = append(mySlice, 7)
	fmt.Println(mySlice)
	fmt.Println(len(mySlice))
	fmt.Println(cap(mySlice))

	fmt.Println("Slicing")
	subSlice := mySlice[2:3]
	fmt.Println(subSlice)
	fmt.Println(len(subSlice))
	fmt.Println(cap(subSlice))
}

func array() {
	fmt.Println("ARRAY")
	var arr [5]int
	for i := 0; i < len(arr); i++ {
		arr[i] = i
	}
	fmt.Println("arr len ->", len(arr))
	fmt.Printf("arr: %v\n", arr)
}

type Person struct {
	Name     string
	LastName string
}
type Speaker interface {
	Speak() string
}

func (p Person) fullname() {
	fmt.Println(p.Name+" ", p.LastName)
}
func (p Person) Speak() string {
	return p.Name + " say!! "
}

func oop() {
	fmt.Println("OOP")
	oop := Person{
		Name:     "dave",
		LastName: "world",
	}
	oop.fullname()
	fmt.Println(oop.Speak())
}

func Datatype() {
	array()
	slice()
	mymap()
	structure()
	oop()
}

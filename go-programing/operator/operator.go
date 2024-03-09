package operator

import (
	"fmt"
)

func forloop() {
	fmt.Println("For Loop")
	for i := 0; i < 100; i++ {
		ScoreSwitch(i)
	}
}

func doWhile() {
	fmt.Println("Do While")
	i := 0
	for {
		fmt.Println("i -> ", i)
		i++
		if i > 50 {
			break
		}
	}
}

func iteration() {
	forloop()
	doWhile()
}

func ScoreSwitch(score int) {
	var grade string
	switch {
	case score >= 80:
		grade = ("A")
	case score >= 70:
		grade = ("B")
	case score >= 60:
		grade = ("C")
	case score >= 50:
		grade = ("D")
	default:
		grade = ("F")
	}
	fmt.Println("score -> ", score, " grade -> ", grade)
}

func ScoreIfElse(score int) {
	if score >= 70 {
		fmt.Println("A")
	} else if score >= 60 {
		fmt.Println("B")
	} else if score >= 50 {
		fmt.Println("C")
	} else {
		fmt.Println("F")
	}
}

func stateMent() {
	a := true
	b := false

	fmt.Println("a -> ", a)
	fmt.Println("b -> ", b)
	fmt.Println("a && b -> ", a && b)
	// fmt.Println("a && a -> ", a && a)
	fmt.Println("!a -> ", !a)
	fmt.Println("!b -> ", !b)
}

func Operator() {
	num := 20
	str := "hello"
	num = num * 2

	num *= 2
	str += " world"

	fmt.Println("num -> ", num)
	fmt.Println("string -> ", str)

	if num == 20 {
		fmt.Println("num equal 20")
	}
	if num >= 0 && num <= 20 {
		fmt.Println("num is more than 0 and less then 20")
	}
	stateMent()
	iteration()
}

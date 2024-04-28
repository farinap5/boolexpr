package main

import (
	"fmt"

	b "boolexpr"
)

func main() {
	l := b.Init("a | b")

	data := "Qualquer dado"
	c, err := l.Eval(data)
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	fmt.Println(data," ",c)
}
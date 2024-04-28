package main

import (
	"fmt"

	b "boolsearch"
)

func main() {
	l := b.Init("a")

	data := "Qualquer dado"
	c, err := l.Eval(data)
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	fmt.Println(data," ",c)
}
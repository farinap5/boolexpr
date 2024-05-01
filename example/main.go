package main

import (
	"fmt"

	b "boolexpr"
)

func main() {
	l := b.Init("(((b&x|x)&x|x&q&q|x)&((x|(o&(o&x)))&q)&(q&b)&(((o&q)&a|l)))")

	data := "Qualquer dado"
	c, err := l.Eval(data)
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	fmt.Println(data," ",c)
}

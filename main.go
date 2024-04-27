package main

import "fmt"

func main() {
	l := Init("( x & y ) | Z")

	l.GetToken()
	for l.Token.Type != 0 {
		fmt.Println(l.Token.Type, l.Token.Data)
		l.GetToken()
	}
}
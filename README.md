# Boolexpr

### Overview

boolexpr is a lightweight library designed for validating strings against patterns using a simple, intuitive syntax. Unlike regular expressions, boolexpr uses straightforward logical operators to combine patterns, making it easier to write and understand complex validation rules.
Syntax

boolexpr patterns are constructed using the following symbols:

- AND: & (both conditions must be true)
- OR: | (at least one condition must be true)
- Grouping: () (used to control the precedence of operations)

### BNF Grammar

The syntax of boolexpr is defined by the following Backus-Naur Form (BNF) grammar:

```
<exp> ::= 
    <group> "&" <exp>
    | <group> "|" <exp>
    | <group>

<group> ::=
    "(" <exp> ")"
    | <data>

<data> ::=
    [a-z]+
```

### Example

The following code must return `false` from the `Eval()` since there is no `x` in the string validated.

```go
l := b.Init("(((b&x|x)&x|x&q&q|x)&((x|(o&(o&x)))&q)&(q&b)&(((o&q)&a|l)))")

data := "Qualquer dado"
c, err := l.Eval(data)
if err != nil {
	fmt.Panic(err.Error())
}
```

The following code must return `true` since the string validated has `a`.

```go
l := b.Init("a | z")

data := "Qualquer dado"
c, err := l.Eval(data)
if err != nil {
	fmt.Panic(err.Error())
}
```

The grammar must generate the following sentences.

```
((xy&g)|g|(vhs&tp)&e&h)&s|mb|b
(m&jua)
((wiaae|(r&dj&h)&(i&p))|((i&g|r)))

((o|b|b&b)&(l))|(a)|o
((b&b|l)&o|l&q&q|l)&((q|(o))&q)|(q&b)&(((o&q)&a|l))
```

[BNF Playground](http://bnfplayground.pauliankline.com/?bnf=%3Cexp%3E%20%3A%3A%3D%20%0A%09%3Cgroup%3E%20%22%26%22%20%3Cexp%3E%0A%20%20%20%20%7C%20%3Cgroup%3E%20%22%7C%22%20%3Cexp%3E%0A%20%20%20%20%7C%20%3Cgroup%3E%0A%0A%3Cgroup%3E%20%3A%3A%3D%0A%09%22(%22%20%3Cexp%3E%20%22)%22%0A%20%20%20%20%7C%20%3Cdata%3E%0A%0A%3Cdata%3E%20%3A%3A%3D%0A%09%5Ba-z%5D%2B&name=Bool%20Language)

https://www.inf.ufpr.br/lesoliveira/download/c-completo-total.pdf
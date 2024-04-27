The BNF grammar looks like the following pattern

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

It must accept the following sentences.

```
((xy&g)|g|(vhs&tp)&e&h)&s|mb|b
(m&jua)
```

[BNF Playground](http://bnfplayground.pauliankline.com/?bnf=%3Cexp%3E%20%3A%3A%3D%20%0A%09%3Cgroup%3E%20%22%26%22%20%3Cexp%3E%0A%20%20%20%20%7C%20%3Cgroup%3E%20%22%7C%22%20%3Cexp%3E%0A%20%20%20%20%7C%20%3Cgroup%3E%0A%0A%3Cgroup%3E%20%3A%3A%3D%0A%09%22(%22%20%3Cexp%3E%20%22)%22%0A%20%20%20%20%7C%20%3Cdata%3E%0A%0A%3Cdata%3E%20%3A%3A%3D%0A%09%5Ba-z%5D%2B&name=Bool%20Language)
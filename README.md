The BNF grammar for evaluating boolen expressions looks like the following pattern.

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

It must generate the following sentences.

```
((xy&g)|g|(vhs&tp)&e&h)&s|mb|b
(m&jua)
((wiaae|(r&dj&h)&(i&p))|((i&g|r)))

((o|b|b&b)&(l))|(a)|o
((b&b|l)&o|l&q&q|l)&((q|(o))&q)|(q&b)&(((o&q)&a|l))
```

[BNF Playground](http://bnfplayground.pauliankline.com/?bnf=%3Cexp%3E%20%3A%3A%3D%20%0A%09%3Cgroup%3E%20%22%26%22%20%3Cexp%3E%0A%20%20%20%20%7C%20%3Cgroup%3E%20%22%7C%22%20%3Cexp%3E%0A%20%20%20%20%7C%20%3Cgroup%3E%0A%0A%3Cgroup%3E%20%3A%3A%3D%0A%09%22(%22%20%3Cexp%3E%20%22)%22%0A%20%20%20%20%7C%20%3Cdata%3E%0A%0A%3Cdata%3E%20%3A%3A%3D%0A%09%5Ba-z%5D%2B&name=Bool%20Language)

https://www.inf.ufpr.br/lesoliveira/download/c-completo-total.pdf

```
<prog> ::= "vnrx" <n> <sttm>

<sttm> ::= <ass_sttm>
    | <while_sttm>
    | <if_sttm>
    | "print" <n> "(" <x_exp> ")"
	| <sttm> ";" <n> <sttm>

<while_sttm> ::= "while" <n> "(" <b_exp> ")" <n> "{" <n> <sttm> <n> "}"

<if_sttm> ::= "if" <n> "(" <b_exp> ")" <n>  "{" <n> <sttm> <n> "}" 
    | "else " <n> "if" <n> "(" <b_exp> ")" <n>  "{" <n> <sttm> <n> "}" 
    | "else" <n> "{" <sttm> "}" 

<ass_sttm> ::= <var> "=" <n> <x_exp>
<var> ::= "var " <n> <name> <n>
	| <name> <n>

<b_exp> ::= <bool>
    | <x_exp> <n> <c_op> <n> <x_exp>
    | "!" <b_exp>
    | "(" <n> <b_exp> <n> ")"
    | <b_exp> <n> <b_op> <n> <b_exp>

<x_exp> ::= <b_exp> | <m_exp>
<c_op> ::= "==" | "!=" | ">" | ">=" | "<" | "<="
<b_op> ::= "&&" | "||"
<bool> ::= "true" | "false"

<m_exp> ::= <factor> "+" <m_exp> 
	| <factor> "-" <m_exp> 
    | <factor>

<factor> ::= <num> "*" <expo> 
	| <num> "/" <expo>
	| <expo>

<expo> ::= <num> "^" <par> 
    | <par>

<par> ::= "(" <m_exp> ")"
    | <num>

<name> ::= [a-z] ([a-z] | [A-Z] | [0-9])*
<num> ::= [0-9]+ | [0-9]+ "." [0-9]+ | <name>
<n> ::= ("\n" | " ")*
```


```
<exp> ::= <sttm>
<sttm> ::= <name> <n> "{" <n> <sttm> <n> "}"
    | <name> <n> "=" <n> <val> <n> ";"
    | <sttm> <n> <sttm>

<val> ::= <num> | <bool> | <string> | <list>
<list> ::= "[" <n> <val> <n> ("," <n> <val>)* <n> "]"
 
<bool> ::= "true" | "false"
<num> ::= [0-9]+ | [0-9]+ "." [0-9]+
<string> ::= "\"" <name> (<n> <name>)* "\""
<name> ::= [a-z] ([a-z] | [A-Z] | [0-9])*
<n> ::= ("\n" | " ")*
```

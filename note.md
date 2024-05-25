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

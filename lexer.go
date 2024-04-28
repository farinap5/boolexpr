package boolsearch

const (
	INV = 0 // Invalid data or end of code
	AND = 1
	OR  = 2
	OPN = 3
	CLS = 4
	STR = 5 // A string
)

type Lex struct {
	Code 	string
	Index 	uint
	Token struct {
		Type uint8
		Data string
	}
}


/*
	Init will create a new lexer for a different sentence
*/
func Init(Data string) *Lex {
	lex := new(Lex)
	lex.Code = Data
	lex.Index = 0
	return lex
}

func (l *Lex)isOperator() bool {
	switch l.Code[l.Index] {
	case '&':
		l.Index++
		l.Token.Type = AND
		return true
	case '|':
		l.Index++
		l.Token.Type = OR
		return true
	default:
		return false
	}
}

/*
	isGroup Validates parenthesis nesting.
*/
func (l *Lex)isGroup() bool {
	switch l.Code[l.Index] {
	case '(':
		l.Index++
		l.Token.Type = OPN
		return true
	case ')':
		l.Index++
		l.Token.Type = CLS
		return true
	default:
		return false
	}
}

func Blank(char byte) bool {
	switch char {
	case '\n':
		return true
	case '\t':
		return true
	case ' ':
		return true
	default:
		return false
	}
}

func (l *Lex)isString() bool {
	var aux []byte
	ret := false
	for (int(l.Index) < len(l.Code)) && !Blank(l.Code[l.Index]) {
		if !ret {ret = true}
		aux = append(aux, l.Code[l.Index])
		l.Index++
	}
	l.Token.Type = STR
	l.Token.Data = string(aux)
	return ret
}


func (l *Lex)GetToken() {
	/* 
	starts nulled cuz if has data, ending blank returns as the last data saw.
	*/
	l.Token.Type = 0

	if int(l.Index) == len(l.Code) {
		return
	}

	for int(l.Index) < len(l.Code) {
		if !Blank(l.Code[l.Index]) {
			if int(l.Index) == len(l.Code) {
				l.Token.Type = 0
				return
			}

			if l.isOperator() {
				return
			}

			if l.isGroup() {
				return
			}

			if l.isString() {
				return
			}
		} else {
			l.Index++
		}
	}	
}
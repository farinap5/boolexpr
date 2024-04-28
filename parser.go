package boolexpr

import (
	"errors"
	"strings"
)

func (l *Lex) Eval(data string) (bool, error) {
	l.GetToken()
	if l.Token.Type == INV {
		return false, errors.New("invalid operation")
	}
	
	r := new(bool)
	_, err := eval_op(l, r, &data)
	return *r, err
}

func eval_op(l *Lex, r *bool, data *string) (string, error) {
	s, err := eval_grp(l, r, data)
	if err != nil {return "", err}

	if l.Token.Type == AND || l.Token.Type == OR {
		*r = strings.Contains(*data, s)
		auxt := l.Token.Type // aux type since it is replaced by the nxt func
		l.GetToken()
		s2, err := eval_op(l, r, data)
		if err != nil {return "", err}

		switch auxt {
		case AND:
			*r = *r && strings.Contains(*data, s2)
		case OR:
			*r = *r || strings.Contains(*data, s2)
		}
	}
	return s, nil
}

func eval_grp(l *Lex, r *bool, data *string) (string, error) {
	if l.Token.Type == OPN {
		l.GetToken()
		s, err := eval_op(l, r, data)
		if err != nil {return "", err}
		
		if l.Token.Type != CLS {
			return "", errors.New("no statement to close")
		}
		l.GetToken()
		return s, nil
	} else {
		return str(l)
	}
}

func str(l *Lex) (string, error) {
	if l.Token.Type == STR {
		aux := l.Token.Data
		l.GetToken()
		return aux, nil
		
	}
	return "", errors.New("not a valid token")
}
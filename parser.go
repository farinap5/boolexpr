package boolsearch

import (
	"errors"
	"strings"
)

func (l *Lex) Eval(data string) (bool, error) {
	l.GetToken()
	if l.Token.Type == INV {
		return false, errors.New("")
	}
	
	r := new(bool)
	_, err := eval_op(l, r, &data)
	return *r, err
}

func eval_op(l *Lex, r *bool, data *string) (string, error) {
	s1, err := eval_grp(l, r, data)
	if err != nil {return "", err}
	if l.Token.Type == AND || l.Token.Type == OR {
		l.GetToken()
		s2, err := eval_op(l, r, data)
		if err != nil {return "", err}
		switch l.Token.Type {
		case AND:
			*r = strings.Contains(*data, s1) && strings.Contains(*data, s2)
		case OR:
			*r = strings.Contains(*data, s1) || strings.Contains(*data, s2)
		}
	}
	return "",nil
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
		return l.Token.Data, nil
	}
	return "", errors.New("not a valid token")
}
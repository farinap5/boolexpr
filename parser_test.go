package boolexpr_test

import (
	"testing"
	b "boolexpr"
)

func TestEvalBasic(t *testing.T) {
	l := b.Init("(o&q)|a")
	data := "Qualquer dado"
	c, err := l.Eval(data)
	if err != nil {
		t.Error(err.Error())
	} else if c != true {
		t.Error("got false where must be true")
	}
}

func TestEvalExtesive(t *testing.T) {
	l := b.Init("(((b&b|l)&o|l&q&q|l)&((q|(o|o))&q)|(q&b)&(((o&q)&a|l)))&x")
	data := "Qualquer dado"
	c, err := l.Eval(data)
	if err != nil {
		t.Error(err.Error())
	} else if c == true {
		t.Error("got true where must be false")
	}

	l = b.Init("(((b&b|l)&o|l&q&q|l)&((q|(o|o))&q)|(q&b)&(((o&q)&a|l)))|x")
	c, err = l.Eval(data)
	if err != nil {
		t.Error(err.Error())
	} else if c == true {
		t.Error("got false where must be true")
	}

	l = b.Init("(((b&x|x)&x|x&q&q|x)&((x|(o&(o&x)))&q)&(q&b)&(((o&q)&a|l)))")
	c, err = l.Eval(data)
	if err != nil {
		t.Error(err.Error())
	} else if c == true {
		t.Error("got true where must be false")
	}
}

func TestEvalForceError(t *testing.T) {
	l := b.Init("(o&q)|&a")
	data := "Qualquer dado"
	_, err := l.Eval(data)
	if err == nil {
		t.Error("invalid sequence of token but no error")
	}

	l = b.Init("((o&q)|&a")
	_, err = l.Eval(data)
	if err == nil {
		t.Error("didt clone parentesis but no error")
	}
}
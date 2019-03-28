package lex

import (
	"bufio"
	"fmt"
	"github.com/keizo042/goecma/js/token"
)

type Lexer struct {
	src    *bufio.Reader
	state  stateFn
	tokens chan token.Value
}

type stateFn func(*lexer) stateFn

func New() (*Lexer, error) {
	return nil, fmt.Errorf("TBD")
}

func (l *Lexer) Lex() error {
	go l.start()
	return nil
}

func (l *Lexer) NextToken() chan token.Item {
	return l.tokens
}

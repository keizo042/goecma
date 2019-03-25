package lex

import (
	"fmt"
	"github.com/keizo042/goecma/js/token"
)

type Lexer struct {
	tokens chan token.Value
}

func New() (*Lexer, error) {
	return nil, fmt.Errorf("TBD")
}

func (l *Lexer) Lex() error {
	return fmt.Errorf("TBD")
}

func (l *Lexer) NextToken() chan token.Value {
	return l.tokens
}

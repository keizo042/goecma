package lex

import (
	"github.com/keizo042/goecma/js/token"
	"strconv"
	"strings"
)

const eof = -1

func (l *Lexer) peek() rune {
}

func (l *Lexer) next() rune {
}

func (l *Lexer) emit(typ token.ItemType) {
}

func (l *Lexer) start() {
	for l.state = lexText; l.state != nil; {
		l.state = l.state(l)
	}
}

func lexText(l *Lexer) statenFn {
	for {
		r := l.peek()
		switch r {
		case eof:
			return nil
		case '\n':
		case '\t':
		case '\r':
			l.next()
			continue
		case '<':
			// TODO(keizo042): emit left bracket
			continue
		case '>':
			// TODO(keizo042): emit right bracket
		case '"':
			l.next()
			return lexString
		default:
			return lexIdent
		}
	}
}

func lexString(l *Lexer) stateFn {
	r := l.peek()
	switch r {
	case eof:
		// TODO(keizo042): emit Lexical Error
		return nil
	case '"':
	default:
		l.next()
		continue
	}
}

func lexIdent(l *Lexer) stateFn {
	r := l.peek()
	switch r {
	case eof:
		// TODO(keizo042): emit Lexical Error
	default:
		// check is string
		continue
	}
}

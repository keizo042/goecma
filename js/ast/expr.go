package ast

import (
	"github.com/keizo042/goecma/js/token"
)

type Expr interface {
	exprNode()
}

type UniOpExpr struct {
	Op    token.Value
	Value token.Value
}

type BinOpExpr struct {
	Op     token.Value
	LValue token.Value
	RValue token.Value
}

func (u UniOpExpr) exprNode() {}
func (b BinOpExpr) exprNode() {}

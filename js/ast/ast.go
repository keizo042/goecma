package ast

type Node interface {
	Node() Expr
	Nodes() Node
}

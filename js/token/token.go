package token

type Item struct {
	Type  ItemType
	Value string
}

type int ItemType

const (
	ItemUndefined ItemType = iota
	ItemLeftBracket
	ItemRightBracket
	ItemString
	ItemIdent
	ItemEqual
)

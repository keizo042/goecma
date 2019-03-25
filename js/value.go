package js

type Value interface {
	ValueType() ValueType
}

type ValueType int

type StringValue struct {
	val string
}

type NumberValue struct {
	num float64
}

type ObjectValue struct {
}

const (
	ValueTypeInvalid ValueType = iota
	ValueTypeString
	ValueTypeNumber
	ValueTypeObject
)

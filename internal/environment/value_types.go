package environment

type ValueType string

const (
	VTIdentifier ValueType = "identifier"
	VTBuiltinFn  ValueType = "builtin"
	VTList       ValueType = "list"
	VTString     ValueType = "string"
	VTNumber     ValueType = "number"
)

type Value interface {
	GetValue() interface{}
	GetType() ValueType
}

type BuiltinFunc func(args []Value) (Value, error)

func (bf *BuiltinFunc) GetType() ValueType {
	return VTBuiltinFn
}

func (bf *BuiltinFunc) GetValue() interface{} {
	return bf
}

type Identifier string

func (i *Identifier) GetType() ValueType {
	return VTIdentifier
}

func (i *Identifier) GetValue() interface{} {
	return i
}

type List []Value

func (l *List) GetType() ValueType {
	return VTList
}

func (l *List) GetValue() interface{} {
	return l
}

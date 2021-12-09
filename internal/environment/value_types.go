package environment

import (
	"log"

	"github.com/nnnewb/minilang/pkg/ast"
)

type ValueType string

const (
	VTIdentifier ValueType = "identifier"
	VTBuiltinFn  ValueType = "builtin"
	VTList       ValueType = "list"
	VTString     ValueType = "string"
	VTNumber     ValueType = "number"
	VTBoolean    ValueType = "boolean"
	VTQuoted     ValueType = "quote"
)

type Value interface {
	GetValue() interface{}
	GetType() ValueType
}

func NewValueFromASTNode(node ast.Node) Value {
	switch n := node.(type) {
	case *ast.List:
		val := make(List, 0)
		elements := n.GetValue().([]ast.Node)
		for _, elem := range elements {
			val = append(val, NewValueFromASTNode(elem))
		}

		return &val
	case *ast.Quoted:
		return NewQuoted(NewValueFromASTNode(n.GetValue().(ast.Node)))
	case ast.Boolean:
		return Boolean(bool(n))
	case ast.Identifier:
		return Identifier(string(n))
	case ast.Number:
		return Number(float64(n))
	case ast.String:
		return String(string(n))
	default:
		log.Fatalf("unexpected ast node %v", node)
		return nil
	}
}

type BuiltinFunc func(env *ExecutionEnv, args []Value) (Value, error)

func (bf BuiltinFunc) GetType() ValueType {
	return VTBuiltinFn
}

func (bf BuiltinFunc) GetValue() interface{} {
	return bf
}

type Identifier string

func (i Identifier) GetType() ValueType {
	return VTIdentifier
}

func (i Identifier) GetValue() interface{} {
	return i
}

type List []Value

func (l *List) GetType() ValueType {
	return VTList
}

func (l *List) GetValue() interface{} {
	return l
}

type String string

func (s String) GetValue() interface{} {
	return s
}

func (s String) GetType() ValueType {
	return VTString
}

type Number float64

func (n Number) GetValue() interface{} {
	return n
}

func (n Number) GetType() ValueType {
	return VTNumber
}

type Boolean bool

func (n Boolean) GetValue() interface{} {
	return n
}

func (n Boolean) GetType() ValueType {
	return VTBoolean
}

type Quoted struct {
	value Value
}

func (q Quoted) GetValue() interface{} {
	return q.value
}

func (q Quoted) GetType() ValueType {
	return VTQuoted
}

func NewQuoted(val Value) *Quoted {
	return &Quoted{
		value: val,
	}
}

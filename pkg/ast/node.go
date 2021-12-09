package ast

import "strconv"

type NodeType string

const (
	NTList    = "list"
	NTIdent   = "ident"
	NTNumber  = "number"
	NTString  = "string"
	NTBoolean = "boolean"
)

type Node interface {
	GetValue() interface{}
	GetType() NodeType
}

type List struct {
	elements []Node
}

func (l *List) GetType() NodeType {
	return NTList
}

func (l *List) GetValue() interface{} {
	return l.elements
}

func (l *List) Append(node Node) *List {
	l.elements = append(l.elements, node)
	return l
}

func NewList() *List {
	return &List{
		elements: make([]Node, 0),
	}
}

func NewListWithInitial(initial Node) *List {
	l := NewList()
	return l.Append(initial)
}

type Identifier string

func (i Identifier) GetType() NodeType {
	return NTIdent
}

func (i Identifier) GetValue() interface{} {
	return i
}

type Number float64

func (n Number) GetType() NodeType {
	return NTNumber
}

func (n Number) GetValue() interface{} {
	return n
}

func NewNumber(text string) (Number, error) {
	if f, err := strconv.ParseFloat(text, 64); err != nil {
		return 0, err
	} else {
		return Number(f), nil
	}
}

type String string

func (n String) GetType() NodeType {
	return NTString
}

func (n String) GetValue() interface{} {
	return n
}

type Boolean bool

func (n Boolean) GetType() NodeType {
	return NTBoolean
}

func (n Boolean) GetValue() interface{} {
	return n
}

package ast

import "strconv"

type NodeType string

const (
	NTList    = "list"
	NTIdent   = "ident"
	NTNumber  = "number"
	NTString  = "string"
	NTBoolean = "boolean"
	NTQuote   = "quote"
)

type Node interface {
	GetValue() interface{}
}

type List struct {
	elements []Node
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

type Symbol string

func (s Symbol) GetValue() interface{} {
	return s
}

type Identifier string

func (i Identifier) GetValue() interface{} {
	return i
}

type UInt float64

func (n UInt) GetValue() interface{} {
	return n
}

func NewUInt(text string) (UInt, error) {
	if f, err := strconv.ParseUint(text, 10, 64); err != nil {
		return 0, err
	} else {
		return UInt(f), nil
	}
}

type Float float64

func (n Float) GetValue() interface{} {
	return n
}

func NewFloat(text string) (Float, error) {
	if f, err := strconv.ParseFloat(text, 64); err != nil {
		return 0, err
	} else {
		return Float(f), nil
	}
}

type String string

func (n String) GetValue() interface{} {
	return n
}

type Boolean bool

func (n Boolean) GetValue() interface{} {
	return n
}

type Quoted struct {
	elements []Node
}

func (q *Quoted) GetValue() interface{} {
	return q.elements
}

func NewQuoted(lst *List) *Quoted {
	if lst == nil {
		return &Quoted{
			elements: make([]Node, 0),
		}
	}
	elements := make([]Node, len(lst.elements))
	copy(elements, lst.elements)
	return &Quoted{
		elements: elements,
	}
}
